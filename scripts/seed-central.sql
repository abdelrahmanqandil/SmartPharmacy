-- ─────────────────────────────────────────────────────────────────────────────
-- Seed data for central.db
-- Scenario: Al-Nour Central Hospital Pharmacy (central node)
--   Pharmacist : Dr. Sarah Hassan   (dr.sarah.hassan)
--   Branch-1   : Al-Nour Community Branch, operated by Dr. Khalil Mansour
--   Period     : December 2025 – April 2026
--
-- Central carries a full formulary of 25 medications.
-- Branch-1 stocks a subset of 12 common medications and its transactions
-- are replicated back here, so central holds the complete audit trail.
-- ─────────────────────────────────────────────────────────────────────────────

PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;

DELETE FROM replication_log;
DELETE FROM inventory_ledger;
DELETE FROM transactions;
DELETE FROM patients;
DELETE FROM medications;

-- ═════════════════════════════════════════════════════════════════════════════
--  MEDICATIONS  — full hospital formulary (25 items)
-- ═════════════════════════════════════════════════════════════════════════════
INSERT INTO medications (id, name, unit, min_stock) VALUES
  ('PARA-500',  'Paracetamol 500mg',           'tablet',  100),
  ('IBUP-400',  'Ibuprofen 400mg',              'tablet',   80),
  ('AMOX-500',  'Amoxicillin 500mg',            'capsule',  60),
  ('AZIT-250',  'Azithromycin 250mg',           'tablet',   40),
  ('OMEP-20',   'Omeprazole 20mg',              'capsule',  50),
  ('CETI-10',   'Cetirizine 10mg',              'tablet',   60),
  ('SALB-INH',  'Salbutamol Inhaler 100mcg',    'inhaler',  20),
  ('METF-500',  'Metformin 500mg',              'tablet',   80),
  ('ATOR-10',   'Atorvastatin 10mg',            'tablet',   60),
  ('VITC-500',  'Vitamin C 500mg',              'tablet',  100),
  ('IRON-FE',   'Ferrous Sulphate 200mg',       'tablet',   40),
  ('METR-500',  'Metronidazole 500mg',          'tablet',   50),
  ('DIAZ-5',    'Diazepam 5mg',                 'tablet',   20),
  ('AMLO-5',    'Amlodipine 5mg',               'tablet',   40),
  ('LISIN-10',  'Lisinopril 10mg',              'tablet',   40),
  ('VITD-1000', 'Vitamin D3 1000 IU',           'capsule',  60),
  ('DICL-50',   'Diclofenac Sodium 50mg',       'tablet',   50),
  ('LORA-10',   'Loratadine 10mg',              'tablet',   60),
  ('CIPR-500',  'Ciprofloxacin 500mg',          'tablet',   30),
  ('DOXY-100',  'Doxycycline 100mg',            'capsule',  30),
  ('PRDN-5',    'Prednisone 5mg',               'tablet',   20),
  ('WARF-5',    'Warfarin 5mg',                 'tablet',   15),
  ('INSL-R',    'Insulin Regular 100 IU/ml',    'vial',     10),
  ('CALC-500',  'Calcium Carbonate 500mg',      'tablet',   50),
  ('ASPR-75',   'Aspirin 75mg',                 'tablet',   80);

-- ═════════════════════════════════════════════════════════════════════════════
--  PATIENTS  — 20 registered patients (central + branch-1)
-- ═════════════════════════════════════════════════════════════════════════════
INSERT INTO patients (id, name, date_of_birth, contact, branch_id, last_updated) VALUES
  ('PAT-001', 'Ahmed Hassan',      '1985-03-12', '+20-100-123-4567', 'central',  '2026-01-10 09:15:00'),
  ('PAT-002', 'Fatima Al-Rashidi','1972-07-28', '+20-101-234-5678', 'central',  '2026-01-15 10:30:00'),
  ('PAT-003', 'Omar Khalil',      '1990-11-05', '+20-102-345-6789', 'branch-1', '2026-01-20 11:00:00'),
  ('PAT-004', 'Layla Ibrahim',    '1968-02-14', '+20-103-456-7890', 'branch-1', '2026-01-22 14:00:00'),
  ('PAT-005', 'Yusuf Mansour',    '1995-09-30', '+20-104-567-8901', 'central',  '2026-01-25 09:45:00'),
  ('PAT-006', 'Nour El-Din',      '1980-04-17', '+20-105-678-9012', 'central',  '2026-02-01 10:00:00'),
  ('PAT-007', 'Mariam Saad',      '2001-06-22', '+20-106-789-0123', 'branch-1', '2026-02-05 11:30:00'),
  ('PAT-008', 'Khaled Farouk',    '1963-12-08', '+20-107-890-1234', 'central',  '2026-02-10 08:30:00'),
  ('PAT-009', 'Samira Nasser',    '1978-08-19', '+20-108-901-2345', 'branch-1', '2026-02-14 13:00:00'),
  ('PAT-010', 'Ibrahim Taha',     '1955-01-30', '+20-109-012-3456', 'central',  '2026-02-20 09:00:00'),
  ('PAT-011', 'Hana Mahmoud',     '1992-05-11', '+20-110-123-4567', 'central',  '2026-03-01 10:15:00'),
  ('PAT-012', 'Tariq Al-Amin',    '1987-10-25', '+20-111-234-5678', 'branch-1', '2026-03-05 11:00:00'),
  ('PAT-013', 'Rania Zaki',       '1975-03-08', '+20-112-345-6789', 'central',  '2026-03-10 14:30:00'),
  ('PAT-014', 'Mostafa Sayed',    '1998-07-14', '+20-113-456-7890', 'branch-1', '2026-03-15 09:00:00'),
  ('PAT-015', 'Yasmin Adel',      '1983-11-28', '+20-114-567-8901', 'central',  '2026-03-20 10:00:00'),
  ('PAT-016', 'Adel Barakat',     '1960-04-03', '+20-115-678-9012', 'central',  '2026-03-25 11:30:00'),
  ('PAT-017', 'Dina Karim',       '2003-09-16', '+20-116-789-0123', 'branch-1', '2026-04-01 08:45:00'),
  ('PAT-018', 'Sherif Gamal',     '1970-02-22', '+20-117-890-1234', 'central',  '2026-04-05 10:30:00'),
  ('PAT-019', 'Noha Fathi',       '1988-06-07', '+20-118-901-2345', 'branch-1', '2026-04-10 13:15:00'),
  ('PAT-020', 'Walid Ashraf',     '1976-12-19', '+20-119-012-3456', 'central',  '2026-04-15 09:30:00');

-- ═════════════════════════════════════════════════════════════════════════════
--  CENTRAL — initial procurement (Dec 1, 2025)
-- ═════════════════════════════════════════════════════════════════════════════
INSERT INTO transactions (id, medication_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp) VALUES
  ('sc-si-01', 'PARA-500',  'dr.sarah.hassan', 1000, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:00:00'),
  ('sc-si-02', 'IBUP-400',  'dr.sarah.hassan',  800, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:05:00'),
  ('sc-si-03', 'AMOX-500',  'dr.sarah.hassan',  600, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:10:00'),
  ('sc-si-04', 'AZIT-250',  'dr.sarah.hassan',  400, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:15:00'),
  ('sc-si-05', 'OMEP-20',   'dr.sarah.hassan',  600, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:20:00'),
  ('sc-si-06', 'CETI-10',   'dr.sarah.hassan',  700, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:25:00'),
  ('sc-si-07', 'SALB-INH',  'dr.sarah.hassan',   60, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:30:00'),
  ('sc-si-08', 'METF-500',  'dr.sarah.hassan', 1200, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:35:00'),
  ('sc-si-09', 'ATOR-10',   'dr.sarah.hassan',  900, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:40:00'),
  ('sc-si-10', 'VITC-500',  'dr.sarah.hassan', 1200, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:45:00'),
  ('sc-si-11', 'IRON-FE',   'dr.sarah.hassan',  500, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:50:00'),
  ('sc-si-12', 'METR-500',  'dr.sarah.hassan',  600, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 08:55:00'),
  ('sc-si-13', 'DIAZ-5',    'dr.sarah.hassan',  200, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:00:00'),
  ('sc-si-14', 'AMLO-5',    'dr.sarah.hassan',  500, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:05:00'),
  ('sc-si-15', 'LISIN-10',  'dr.sarah.hassan',  500, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:10:00'),
  ('sc-si-16', 'VITD-1000', 'dr.sarah.hassan',  800, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:15:00'),
  ('sc-si-17', 'DICL-50',   'dr.sarah.hassan',  600, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:20:00'),
  ('sc-si-18', 'LORA-10',   'dr.sarah.hassan',  700, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:25:00'),
  ('sc-si-19', 'CIPR-500',  'dr.sarah.hassan',  400, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:30:00'),
  ('sc-si-20', 'DOXY-100',  'dr.sarah.hassan',  400, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:35:00'),
  ('sc-si-21', 'PRDN-5',    'dr.sarah.hassan',  300, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:40:00'),
  ('sc-si-22', 'WARF-5',    'dr.sarah.hassan',  200, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:45:00'),
  ('sc-si-23', 'INSL-R',    'dr.sarah.hassan',   80, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:50:00'),
  ('sc-si-24', 'CALC-500',  'dr.sarah.hassan',  700, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 09:55:00'),
  ('sc-si-25', 'ASPR-75',   'dr.sarah.hassan', 1000, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2025-12-01 10:00:00');

INSERT INTO inventory_ledger (id, medication_id, quantity_delta, entry_type, transaction_id, branch_id, node_id, recorded_by, timestamp) VALUES
  ('led-sc-si-01', 'PARA-500',  1000, 'STOCK_IN', 'sc-si-01', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:00:00'),
  ('led-sc-si-02', 'IBUP-400',   800, 'STOCK_IN', 'sc-si-02', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:05:00'),
  ('led-sc-si-03', 'AMOX-500',   600, 'STOCK_IN', 'sc-si-03', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:10:00'),
  ('led-sc-si-04', 'AZIT-250',   400, 'STOCK_IN', 'sc-si-04', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:15:00'),
  ('led-sc-si-05', 'OMEP-20',    600, 'STOCK_IN', 'sc-si-05', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:20:00'),
  ('led-sc-si-06', 'CETI-10',    700, 'STOCK_IN', 'sc-si-06', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:25:00'),
  ('led-sc-si-07', 'SALB-INH',    60, 'STOCK_IN', 'sc-si-07', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:30:00'),
  ('led-sc-si-08', 'METF-500',  1200, 'STOCK_IN', 'sc-si-08', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:35:00'),
  ('led-sc-si-09', 'ATOR-10',    900, 'STOCK_IN', 'sc-si-09', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:40:00'),
  ('led-sc-si-10', 'VITC-500',  1200, 'STOCK_IN', 'sc-si-10', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:45:00'),
  ('led-sc-si-11', 'IRON-FE',    500, 'STOCK_IN', 'sc-si-11', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:50:00'),
  ('led-sc-si-12', 'METR-500',   600, 'STOCK_IN', 'sc-si-12', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 08:55:00'),
  ('led-sc-si-13', 'DIAZ-5',     200, 'STOCK_IN', 'sc-si-13', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:00:00'),
  ('led-sc-si-14', 'AMLO-5',     500, 'STOCK_IN', 'sc-si-14', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:05:00'),
  ('led-sc-si-15', 'LISIN-10',   500, 'STOCK_IN', 'sc-si-15', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:10:00'),
  ('led-sc-si-16', 'VITD-1000',  800, 'STOCK_IN', 'sc-si-16', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:15:00'),
  ('led-sc-si-17', 'DICL-50',    600, 'STOCK_IN', 'sc-si-17', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:20:00'),
  ('led-sc-si-18', 'LORA-10',    700, 'STOCK_IN', 'sc-si-18', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:25:00'),
  ('led-sc-si-19', 'CIPR-500',   400, 'STOCK_IN', 'sc-si-19', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:30:00'),
  ('led-sc-si-20', 'DOXY-100',   400, 'STOCK_IN', 'sc-si-20', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:35:00'),
  ('led-sc-si-21', 'PRDN-5',     300, 'STOCK_IN', 'sc-si-21', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:40:00'),
  ('led-sc-si-22', 'WARF-5',     200, 'STOCK_IN', 'sc-si-22', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:45:00'),
  ('led-sc-si-23', 'INSL-R',      80, 'STOCK_IN', 'sc-si-23', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:50:00'),
  ('led-sc-si-24', 'CALC-500',   700, 'STOCK_IN', 'sc-si-24', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 09:55:00'),
  ('led-sc-si-25', 'ASPR-75',   1000, 'STOCK_IN', 'sc-si-25', 'central', 'central-1', 'dr.sarah.hassan', '2025-12-01 10:00:00');

-- ═════════════════════════════════════════════════════════════════════════════
--  BRANCH-1 — initial procurement (Dec 5, 2025) [replicated from edge]
-- ═════════════════════════════════════════════════════════════════════════════
INSERT INTO transactions (id, medication_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp) VALUES
  ('se-si-01', 'PARA-500',  'dr.khalil.mansour', 300, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:00:00'),
  ('se-si-02', 'IBUP-400',  'dr.khalil.mansour', 200, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:05:00'),
  ('se-si-03', 'AMOX-500',  'dr.khalil.mansour', 150, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:10:00'),
  ('se-si-04', 'AZIT-250',  'dr.khalil.mansour', 100, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:15:00'),
  ('se-si-05', 'OMEP-20',   'dr.khalil.mansour', 250, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:20:00'),
  ('se-si-06', 'CETI-10',   'dr.khalil.mansour', 200, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:25:00'),
  ('se-si-07', 'SALB-INH',  'dr.khalil.mansour',  25, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:30:00'),
  ('se-si-08', 'METF-500',  'dr.khalil.mansour', 300, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:35:00'),
  ('se-si-09', 'ATOR-10',   'dr.khalil.mansour', 200, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:40:00'),
  ('se-si-10', 'VITC-500',  'dr.khalil.mansour', 300, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:45:00'),
  ('se-si-11', 'IRON-FE',   'dr.khalil.mansour', 150, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:50:00'),
  ('se-si-12', 'METR-500',  'dr.khalil.mansour', 150, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2025-12-05 08:55:00');

INSERT INTO inventory_ledger (id, medication_id, quantity_delta, entry_type, transaction_id, branch_id, node_id, recorded_by, timestamp) VALUES
  ('led-se-si-01', 'PARA-500',  300, 'STOCK_IN', 'se-si-01', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:00:00'),
  ('led-se-si-02', 'IBUP-400',  200, 'STOCK_IN', 'se-si-02', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:05:00'),
  ('led-se-si-03', 'AMOX-500',  150, 'STOCK_IN', 'se-si-03', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:10:00'),
  ('led-se-si-04', 'AZIT-250',  100, 'STOCK_IN', 'se-si-04', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:15:00'),
  ('led-se-si-05', 'OMEP-20',   250, 'STOCK_IN', 'se-si-05', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:20:00'),
  ('led-se-si-06', 'CETI-10',   200, 'STOCK_IN', 'se-si-06', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:25:00'),
  ('led-se-si-07', 'SALB-INH',   25, 'STOCK_IN', 'se-si-07', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:30:00'),
  ('led-se-si-08', 'METF-500',  300, 'STOCK_IN', 'se-si-08', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:35:00'),
  ('led-se-si-09', 'ATOR-10',   200, 'STOCK_IN', 'se-si-09', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:40:00'),
  ('led-se-si-10', 'VITC-500',  300, 'STOCK_IN', 'se-si-10', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:45:00'),
  ('led-se-si-11', 'IRON-FE',   150, 'STOCK_IN', 'se-si-11', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:50:00'),
  ('led-se-si-12', 'METR-500',  150, 'STOCK_IN', 'se-si-12', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2025-12-05 08:55:00');

-- ═════════════════════════════════════════════════════════════════════════════
--  CENTRAL — restocks (Feb & Mar 2026)
-- ═════════════════════════════════════════════════════════════════════════════
INSERT INTO transactions (id, medication_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp) VALUES
  ('sc-rs-01', 'PARA-500',  'dr.sarah.hassan', 500, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2026-02-01 08:00:00'),
  ('sc-rs-02', 'METF-500',  'dr.sarah.hassan', 600, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2026-02-01 08:10:00'),
  ('sc-rs-03', 'ATOR-10',   'dr.sarah.hassan', 400, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2026-02-01 08:20:00'),
  ('sc-rs-04', 'VITC-500',  'dr.sarah.hassan', 600, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2026-03-01 08:00:00'),
  ('sc-rs-05', 'ASPR-75',   'dr.sarah.hassan', 500, 'STOCK_IN', 'central', 'central-1', 'COMPLETED', '2026-03-01 08:10:00');

INSERT INTO inventory_ledger (id, medication_id, quantity_delta, entry_type, transaction_id, branch_id, node_id, recorded_by, timestamp) VALUES
  ('led-sc-rs-01', 'PARA-500', 500, 'STOCK_IN', 'sc-rs-01', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-01 08:00:00'),
  ('led-sc-rs-02', 'METF-500', 600, 'STOCK_IN', 'sc-rs-02', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-01 08:10:00'),
  ('led-sc-rs-03', 'ATOR-10',  400, 'STOCK_IN', 'sc-rs-03', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-01 08:20:00'),
  ('led-sc-rs-04', 'VITC-500', 600, 'STOCK_IN', 'sc-rs-04', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-01 08:00:00'),
  ('led-sc-rs-05', 'ASPR-75',  500, 'STOCK_IN', 'sc-rs-05', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-01 08:10:00');

-- ═════════════════════════════════════════════════════════════════════════════
--  BRANCH-1 — restocks (Feb & Mar 2026) [replicated from edge]
-- ═════════════════════════════════════════════════════════════════════════════
INSERT INTO transactions (id, medication_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp) VALUES
  ('se-rs-01', 'PARA-500',  'dr.khalil.mansour', 150, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-01 09:00:00'),
  ('se-rs-02', 'METF-500',  'dr.khalil.mansour', 150, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-01 09:10:00'),
  ('se-rs-03', 'VITC-500',  'dr.khalil.mansour', 150, 'STOCK_IN', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-01 09:00:00');

INSERT INTO inventory_ledger (id, medication_id, quantity_delta, entry_type, transaction_id, branch_id, node_id, recorded_by, timestamp) VALUES
  ('led-se-rs-01', 'PARA-500', 150, 'STOCK_IN', 'se-rs-01', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-01 09:00:00'),
  ('led-se-rs-02', 'METF-500', 150, 'STOCK_IN', 'se-rs-02', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-01 09:10:00'),
  ('led-se-rs-03', 'VITC-500', 150, 'STOCK_IN', 'se-rs-03', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-01 09:00:00');

-- ═════════════════════════════════════════════════════════════════════════════
--  CENTRAL — dispenses (Jan–Apr 2026)
-- ═════════════════════════════════════════════════════════════════════════════
INSERT INTO transactions (id, medication_id, patient_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp) VALUES
  -- January 2026
  ('sc-ds-01', 'PARA-500',  'PAT-001', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-05 10:15:00'),
  ('sc-ds-02', 'METF-500',  'PAT-002', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-05 10:30:00'),
  ('sc-ds-03', 'ATOR-10',   'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-07 09:00:00'),
  ('sc-ds-04', 'WARF-5',    'PAT-010', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-07 09:15:00'),
  ('sc-ds-05', 'IBUP-400',  'PAT-001', 'dr.sarah.hassan',  20, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-09 11:00:00'),
  ('sc-ds-06', 'AMOX-500',  'PAT-005', 'dr.sarah.hassan',  21, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-10 09:30:00'),
  ('sc-ds-07', 'OMEP-20',   'PAT-013', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-12 10:00:00'),
  ('sc-ds-08', 'VITD-1000', 'PAT-002', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-14 11:30:00'),
  ('sc-ds-09', 'CETI-10',   'PAT-006', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-15 09:45:00'),
  ('sc-ds-10', 'ASPR-75',   'PAT-016', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-16 10:15:00'),
  ('sc-ds-11', 'AMLO-5',    'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-18 09:00:00'),
  ('sc-ds-12', 'LISIN-10',  'PAT-010', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-19 09:30:00'),
  ('sc-ds-13', 'AZIT-250',  'PAT-011', 'dr.sarah.hassan',   6, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-21 11:00:00'),
  ('sc-ds-14', 'DICL-50',   'PAT-015', 'dr.sarah.hassan',  20, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-22 10:30:00'),
  ('sc-ds-15', 'IRON-FE',   'PAT-018', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-25 09:15:00'),
  ('sc-ds-16', 'DIAZ-5',    'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-28 10:00:00'),
  ('sc-ds-17', 'INSL-R',    'PAT-002', 'dr.sarah.hassan',   2, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-10 11:00:00'),
  ('sc-ds-18', 'SALB-INH',  'PAT-006', 'dr.sarah.hassan',   2, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-01-15 10:00:00'),
  -- February 2026
  ('sc-ds-19', 'PARA-500',  'PAT-001', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-03 10:15:00'),
  ('sc-ds-20', 'METF-500',  'PAT-002', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-03 10:30:00'),
  ('sc-ds-21', 'ATOR-10',   'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-05 09:00:00'),
  ('sc-ds-22', 'WARF-5',    'PAT-010', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-05 09:15:00'),
  ('sc-ds-23', 'CIPR-500',  'PAT-005', 'dr.sarah.hassan',  14, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-07 10:00:00'),
  ('sc-ds-24', 'OMEP-20',   'PAT-013', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-08 10:00:00'),
  ('sc-ds-25', 'LORA-10',   'PAT-006', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-10 09:45:00'),
  ('sc-ds-26', 'ASPR-75',   'PAT-016', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-11 10:15:00'),
  ('sc-ds-27', 'AMLO-5',    'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-12 09:00:00'),
  ('sc-ds-28', 'LISIN-10',  'PAT-010', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-14 09:30:00'),
  ('sc-ds-29', 'DOXY-100',  'PAT-020', 'dr.sarah.hassan',  14, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-15 11:00:00'),
  ('sc-ds-30', 'VITC-500',  'PAT-011', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-17 10:30:00'),
  ('sc-ds-31', 'PRDN-5',    'PAT-015', 'dr.sarah.hassan',  10, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-18 10:00:00'),
  ('sc-ds-32', 'IRON-FE',   'PAT-018', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-20 09:15:00'),
  ('sc-ds-33', 'METR-500',  'PAT-013', 'dr.sarah.hassan',  14, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-22 11:00:00'),
  ('sc-ds-34', 'ASPR-75',   'PAT-016', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-25 10:15:00'),
  ('sc-ds-35', 'VITD-1000', 'PAT-002', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-26 11:30:00'),
  ('sc-ds-36', 'DIAZ-5',    'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-28 10:00:00'),
  ('sc-ds-37', 'INSL-R',    'PAT-002', 'dr.sarah.hassan',   2, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-02-10 11:00:00'),
  -- March 2026
  ('sc-ds-38', 'PARA-500',  'PAT-001', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-03 10:15:00'),
  ('sc-ds-39', 'METF-500',  'PAT-002', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-03 10:30:00'),
  ('sc-ds-40', 'ATOR-10',   'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-05 09:00:00'),
  ('sc-ds-41', 'WARF-5',    'PAT-010', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-07 09:15:00'),
  ('sc-ds-42', 'OMEP-20',   'PAT-013', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-08 10:00:00'),
  ('sc-ds-43', 'AMOX-500',  'PAT-005', 'dr.sarah.hassan',  21, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-10 09:30:00'),
  ('sc-ds-44', 'CETI-10',   'PAT-006', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-11 09:45:00'),
  ('sc-ds-45', 'AMLO-5',    'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-13 09:00:00'),
  ('sc-ds-46', 'LISIN-10',  'PAT-010', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-14 09:30:00'),
  ('sc-ds-47', 'VITC-500',  'PAT-011', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-16 10:30:00'),
  ('sc-ds-48', 'DICL-50',   'PAT-015', 'dr.sarah.hassan',  20, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-18 10:30:00'),
  ('sc-ds-49', 'ASPR-75',   'PAT-016', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-20 10:15:00'),
  ('sc-ds-50', 'IRON-FE',   'PAT-018', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-22 09:15:00'),
  ('sc-ds-51', 'CALC-500',  'PAT-020', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-25 11:00:00'),
  ('sc-ds-52', 'IBUP-400',  'PAT-001', 'dr.sarah.hassan',  20, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-28 11:00:00'),
  ('sc-ds-53', 'VITD-1000', 'PAT-002', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-30 11:30:00'),
  ('sc-ds-54', 'DIAZ-5',    'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-28 10:00:00'),
  ('sc-ds-55', 'INSL-R',    'PAT-002', 'dr.sarah.hassan',   2, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-10 11:00:00'),
  ('sc-ds-56', 'SALB-INH',  'PAT-006', 'dr.sarah.hassan',   2, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-03-15 10:00:00'),
  -- April 2026
  ('sc-ds-57', 'PARA-500',  'PAT-001', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-01 10:15:00'),
  ('sc-ds-58', 'METF-500',  'PAT-002', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-01 10:30:00'),
  ('sc-ds-59', 'ATOR-10',   'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-05 09:00:00'),
  ('sc-ds-60', 'WARF-5',    'PAT-010', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-07 09:15:00'),
  ('sc-ds-61', 'OMEP-20',   'PAT-013', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-08 10:00:00'),
  ('sc-ds-62', 'LORA-10',   'PAT-006', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-10 09:45:00'),
  ('sc-ds-63', 'ASPR-75',   'PAT-016', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-12 10:15:00'),
  ('sc-ds-64', 'AZIT-250',  'PAT-011', 'dr.sarah.hassan',   6, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-14 11:00:00'),
  ('sc-ds-65', 'PRDN-5',    'PAT-015', 'dr.sarah.hassan',  10, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-16 10:00:00'),
  ('sc-ds-66', 'CIPR-500',  'PAT-020', 'dr.sarah.hassan',  14, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-18 11:00:00'),
  ('sc-ds-67', 'IRON-FE',   'PAT-018', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-20 09:15:00'),
  ('sc-ds-68', 'ASPR-75',   'PAT-016', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-22 10:15:00'),
  ('sc-ds-69', 'VITD-1000', 'PAT-002', 'dr.sarah.hassan',  60, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-25 11:30:00'),
  ('sc-ds-70', 'PARA-500',  'PAT-001', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-28 10:15:00'),
  ('sc-ds-71', 'ATOR-10',   'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-30 09:00:00'),
  ('sc-ds-72', 'WARF-5',    'PAT-010', 'dr.sarah.hassan',  15, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-30 09:15:00'),
  ('sc-ds-73', 'DIAZ-5',    'PAT-008', 'dr.sarah.hassan',  30, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-25 10:00:00'),
  ('sc-ds-74', 'INSL-R',    'PAT-002', 'dr.sarah.hassan',   2, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-10 11:00:00'),
  ('sc-ds-75', 'DOXY-100',  'PAT-020', 'dr.sarah.hassan',  14, 'DISPENSE', 'central', 'central-1', 'COMPLETED', '2026-04-15 11:00:00');

INSERT INTO inventory_ledger (id, medication_id, quantity_delta, entry_type, transaction_id, branch_id, node_id, recorded_by, timestamp) VALUES
  ('led-sc-ds-01', 'PARA-500',    -30, 'DISPENSE', 'sc-ds-01', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-05 10:15:00'),
  ('led-sc-ds-02', 'METF-500',    -60, 'DISPENSE', 'sc-ds-02', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-05 10:30:00'),
  ('led-sc-ds-03', 'ATOR-10',     -30, 'DISPENSE', 'sc-ds-03', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-07 09:00:00'),
  ('led-sc-ds-04', 'WARF-5',      -30, 'DISPENSE', 'sc-ds-04', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-07 09:15:00'),
  ('led-sc-ds-05', 'IBUP-400',    -20, 'DISPENSE', 'sc-ds-05', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-09 11:00:00'),
  ('led-sc-ds-06', 'AMOX-500',    -21, 'DISPENSE', 'sc-ds-06', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-10 09:30:00'),
  ('led-sc-ds-07', 'OMEP-20',     -30, 'DISPENSE', 'sc-ds-07', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-12 10:00:00'),
  ('led-sc-ds-08', 'VITD-1000',   -60, 'DISPENSE', 'sc-ds-08', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-14 11:30:00'),
  ('led-sc-ds-09', 'CETI-10',     -30, 'DISPENSE', 'sc-ds-09', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-15 09:45:00'),
  ('led-sc-ds-10', 'ASPR-75',     -60, 'DISPENSE', 'sc-ds-10', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-16 10:15:00'),
  ('led-sc-ds-11', 'AMLO-5',      -30, 'DISPENSE', 'sc-ds-11', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-18 09:00:00'),
  ('led-sc-ds-12', 'LISIN-10',    -30, 'DISPENSE', 'sc-ds-12', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-19 09:30:00'),
  ('led-sc-ds-13', 'AZIT-250',     -6, 'DISPENSE', 'sc-ds-13', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-21 11:00:00'),
  ('led-sc-ds-14', 'DICL-50',     -20, 'DISPENSE', 'sc-ds-14', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-22 10:30:00'),
  ('led-sc-ds-15', 'IRON-FE',     -30, 'DISPENSE', 'sc-ds-15', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-25 09:15:00'),
  ('led-sc-ds-16', 'DIAZ-5',      -30, 'DISPENSE', 'sc-ds-16', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-28 10:00:00'),
  ('led-sc-ds-17', 'INSL-R',       -2, 'DISPENSE', 'sc-ds-17', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-10 11:00:00'),
  ('led-sc-ds-18', 'SALB-INH',     -2, 'DISPENSE', 'sc-ds-18', 'central', 'central-1', 'dr.sarah.hassan', '2026-01-15 10:00:00'),
  ('led-sc-ds-19', 'PARA-500',    -30, 'DISPENSE', 'sc-ds-19', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-03 10:15:00'),
  ('led-sc-ds-20', 'METF-500',    -60, 'DISPENSE', 'sc-ds-20', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-03 10:30:00'),
  ('led-sc-ds-21', 'ATOR-10',     -30, 'DISPENSE', 'sc-ds-21', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-05 09:00:00'),
  ('led-sc-ds-22', 'WARF-5',      -30, 'DISPENSE', 'sc-ds-22', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-05 09:15:00'),
  ('led-sc-ds-23', 'CIPR-500',    -14, 'DISPENSE', 'sc-ds-23', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-07 10:00:00'),
  ('led-sc-ds-24', 'OMEP-20',     -30, 'DISPENSE', 'sc-ds-24', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-08 10:00:00'),
  ('led-sc-ds-25', 'LORA-10',     -30, 'DISPENSE', 'sc-ds-25', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-10 09:45:00'),
  ('led-sc-ds-26', 'ASPR-75',     -60, 'DISPENSE', 'sc-ds-26', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-11 10:15:00'),
  ('led-sc-ds-27', 'AMLO-5',      -30, 'DISPENSE', 'sc-ds-27', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-12 09:00:00'),
  ('led-sc-ds-28', 'LISIN-10',    -30, 'DISPENSE', 'sc-ds-28', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-14 09:30:00'),
  ('led-sc-ds-29', 'DOXY-100',    -14, 'DISPENSE', 'sc-ds-29', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-15 11:00:00'),
  ('led-sc-ds-30', 'VITC-500',    -60, 'DISPENSE', 'sc-ds-30', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-17 10:30:00'),
  ('led-sc-ds-31', 'PRDN-5',      -10, 'DISPENSE', 'sc-ds-31', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-18 10:00:00'),
  ('led-sc-ds-32', 'IRON-FE',     -30, 'DISPENSE', 'sc-ds-32', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-20 09:15:00'),
  ('led-sc-ds-33', 'METR-500',    -14, 'DISPENSE', 'sc-ds-33', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-22 11:00:00'),
  ('led-sc-ds-34', 'ASPR-75',     -60, 'DISPENSE', 'sc-ds-34', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-25 10:15:00'),
  ('led-sc-ds-35', 'VITD-1000',   -60, 'DISPENSE', 'sc-ds-35', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-26 11:30:00'),
  ('led-sc-ds-36', 'DIAZ-5',      -30, 'DISPENSE', 'sc-ds-36', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-28 10:00:00'),
  ('led-sc-ds-37', 'INSL-R',       -2, 'DISPENSE', 'sc-ds-37', 'central', 'central-1', 'dr.sarah.hassan', '2026-02-10 11:00:00'),
  ('led-sc-ds-38', 'PARA-500',    -30, 'DISPENSE', 'sc-ds-38', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-03 10:15:00'),
  ('led-sc-ds-39', 'METF-500',    -60, 'DISPENSE', 'sc-ds-39', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-03 10:30:00'),
  ('led-sc-ds-40', 'ATOR-10',     -30, 'DISPENSE', 'sc-ds-40', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-05 09:00:00'),
  ('led-sc-ds-41', 'WARF-5',      -30, 'DISPENSE', 'sc-ds-41', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-07 09:15:00'),
  ('led-sc-ds-42', 'OMEP-20',     -30, 'DISPENSE', 'sc-ds-42', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-08 10:00:00'),
  ('led-sc-ds-43', 'AMOX-500',    -21, 'DISPENSE', 'sc-ds-43', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-10 09:30:00'),
  ('led-sc-ds-44', 'CETI-10',     -30, 'DISPENSE', 'sc-ds-44', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-11 09:45:00'),
  ('led-sc-ds-45', 'AMLO-5',      -30, 'DISPENSE', 'sc-ds-45', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-13 09:00:00'),
  ('led-sc-ds-46', 'LISIN-10',    -30, 'DISPENSE', 'sc-ds-46', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-14 09:30:00'),
  ('led-sc-ds-47', 'VITC-500',    -60, 'DISPENSE', 'sc-ds-47', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-16 10:30:00'),
  ('led-sc-ds-48', 'DICL-50',     -20, 'DISPENSE', 'sc-ds-48', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-18 10:30:00'),
  ('led-sc-ds-49', 'ASPR-75',     -60, 'DISPENSE', 'sc-ds-49', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-20 10:15:00'),
  ('led-sc-ds-50', 'IRON-FE',     -30, 'DISPENSE', 'sc-ds-50', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-22 09:15:00'),
  ('led-sc-ds-51', 'CALC-500',    -60, 'DISPENSE', 'sc-ds-51', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-25 11:00:00'),
  ('led-sc-ds-52', 'IBUP-400',    -20, 'DISPENSE', 'sc-ds-52', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-28 11:00:00'),
  ('led-sc-ds-53', 'VITD-1000',   -60, 'DISPENSE', 'sc-ds-53', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-30 11:30:00'),
  ('led-sc-ds-54', 'DIAZ-5',      -30, 'DISPENSE', 'sc-ds-54', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-28 10:00:00'),
  ('led-sc-ds-55', 'INSL-R',       -2, 'DISPENSE', 'sc-ds-55', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-10 11:00:00'),
  ('led-sc-ds-56', 'SALB-INH',     -2, 'DISPENSE', 'sc-ds-56', 'central', 'central-1', 'dr.sarah.hassan', '2026-03-15 10:00:00'),
  ('led-sc-ds-57', 'PARA-500',    -30, 'DISPENSE', 'sc-ds-57', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-01 10:15:00'),
  ('led-sc-ds-58', 'METF-500',    -60, 'DISPENSE', 'sc-ds-58', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-01 10:30:00'),
  ('led-sc-ds-59', 'ATOR-10',     -30, 'DISPENSE', 'sc-ds-59', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-05 09:00:00'),
  ('led-sc-ds-60', 'WARF-5',      -30, 'DISPENSE', 'sc-ds-60', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-07 09:15:00'),
  ('led-sc-ds-61', 'OMEP-20',     -30, 'DISPENSE', 'sc-ds-61', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-08 10:00:00'),
  ('led-sc-ds-62', 'LORA-10',     -30, 'DISPENSE', 'sc-ds-62', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-10 09:45:00'),
  ('led-sc-ds-63', 'ASPR-75',     -60, 'DISPENSE', 'sc-ds-63', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-12 10:15:00'),
  ('led-sc-ds-64', 'AZIT-250',     -6, 'DISPENSE', 'sc-ds-64', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-14 11:00:00'),
  ('led-sc-ds-65', 'PRDN-5',      -10, 'DISPENSE', 'sc-ds-65', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-16 10:00:00'),
  ('led-sc-ds-66', 'CIPR-500',    -14, 'DISPENSE', 'sc-ds-66', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-18 11:00:00'),
  ('led-sc-ds-67', 'IRON-FE',     -30, 'DISPENSE', 'sc-ds-67', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-20 09:15:00'),
  ('led-sc-ds-68', 'ASPR-75',     -60, 'DISPENSE', 'sc-ds-68', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-22 10:15:00'),
  ('led-sc-ds-69', 'VITD-1000',   -60, 'DISPENSE', 'sc-ds-69', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-25 11:30:00'),
  ('led-sc-ds-70', 'PARA-500',    -30, 'DISPENSE', 'sc-ds-70', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-28 10:15:00'),
  ('led-sc-ds-71', 'ATOR-10',     -30, 'DISPENSE', 'sc-ds-71', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-30 09:00:00'),
  ('led-sc-ds-72', 'WARF-5',      -15, 'DISPENSE', 'sc-ds-72', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-30 09:15:00'),
  ('led-sc-ds-73', 'DIAZ-5',      -30, 'DISPENSE', 'sc-ds-73', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-25 10:00:00'),
  ('led-sc-ds-74', 'INSL-R',       -2, 'DISPENSE', 'sc-ds-74', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-10 11:00:00'),
  ('led-sc-ds-75', 'DOXY-100',    -14, 'DISPENSE', 'sc-ds-75', 'central', 'central-1', 'dr.sarah.hassan', '2026-04-15 11:00:00');

-- ═════════════════════════════════════════════════════════════════════════════
--  BRANCH-1 — dispenses (Jan–Apr 2026) [replicated from edge]
-- ═════════════════════════════════════════════════════════════════════════════
INSERT INTO transactions (id, medication_id, patient_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp) VALUES
  -- January 2026
  ('se-ds-01', 'PARA-500',  'PAT-003', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-01-06 10:00:00'),
  ('se-ds-02', 'METF-500',  'PAT-004', 'dr.khalil.mansour',  60, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-01-06 10:15:00'),
  ('se-ds-03', 'AMOX-500',  'PAT-007', 'dr.khalil.mansour',  21, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-01-08 09:30:00'),
  ('se-ds-04', 'CETI-10',   'PAT-009', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-01-10 10:00:00'),
  ('se-ds-05', 'IBUP-400',  'PAT-003', 'dr.khalil.mansour',  20, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-01-12 09:45:00'),
  ('se-ds-06', 'ATOR-10',   'PAT-004', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-01-15 10:30:00'),
  ('se-ds-07', 'OMEP-20',   'PAT-012', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-01-18 09:00:00'),
  ('se-ds-08', 'VITC-500',  'PAT-014', 'dr.khalil.mansour',  60, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-01-22 11:00:00'),
  -- February 2026
  ('se-ds-09', 'PARA-500',  'PAT-003', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-04 10:00:00'),
  ('se-ds-10', 'METF-500',  'PAT-004', 'dr.khalil.mansour',  60, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-04 10:15:00'),
  ('se-ds-11', 'AZIT-250',  'PAT-007', 'dr.khalil.mansour',   6, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-06 09:30:00'),
  ('se-ds-12', 'CETI-10',   'PAT-009', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-10 10:00:00'),
  ('se-ds-13', 'ATOR-10',   'PAT-004', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-14 10:30:00'),
  ('se-ds-14', 'OMEP-20',   'PAT-012', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-16 09:00:00'),
  ('se-ds-15', 'IRON-FE',   'PAT-017', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-19 11:00:00'),
  ('se-ds-16', 'METR-500',  'PAT-019', 'dr.khalil.mansour',  14, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-02-22 10:00:00'),
  -- March 2026
  ('se-ds-17', 'PARA-500',  'PAT-003', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-04 10:00:00'),
  ('se-ds-18', 'METF-500',  'PAT-004', 'dr.khalil.mansour',  60, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-04 10:15:00'),
  ('se-ds-19', 'AMOX-500',  'PAT-007', 'dr.khalil.mansour',  21, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-06 09:30:00'),
  ('se-ds-20', 'CETI-10',   'PAT-009', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-10 10:00:00'),
  ('se-ds-21', 'ATOR-10',   'PAT-004', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-14 10:30:00'),
  ('se-ds-22', 'OMEP-20',   'PAT-012', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-16 09:00:00'),
  ('se-ds-23', 'VITC-500',  'PAT-014', 'dr.khalil.mansour',  60, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-20 11:00:00'),
  ('se-ds-24', 'IRON-FE',   'PAT-019', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-24 10:00:00'),
  ('se-ds-25', 'SALB-INH',  'PAT-003', 'dr.khalil.mansour',   1, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-03-28 09:00:00'),
  -- April 2026
  ('se-ds-26', 'PARA-500',  'PAT-003', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-04-02 10:00:00'),
  ('se-ds-27', 'METF-500',  'PAT-004', 'dr.khalil.mansour',  60, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-04-02 10:15:00'),
  ('se-ds-28', 'IBUP-400',  'PAT-009', 'dr.khalil.mansour',  20, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-04-08 09:45:00'),
  ('se-ds-29', 'OMEP-20',   'PAT-012', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-04-10 09:00:00'),
  ('se-ds-30', 'ATOR-10',   'PAT-004', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-04-14 10:30:00'),
  ('se-ds-31', 'IRON-FE',   'PAT-017', 'dr.khalil.mansour',  30, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-04-18 11:00:00'),
  ('se-ds-32', 'METR-500',  'PAT-019', 'dr.khalil.mansour',  14, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-04-22 10:00:00'),
  ('se-ds-33', 'VITC-500',  'PAT-014', 'dr.khalil.mansour',  60, 'DISPENSE', 'branch-1', 'edge-branch-1', 'COMPLETED', '2026-04-25 11:00:00');

INSERT INTO inventory_ledger (id, medication_id, quantity_delta, entry_type, transaction_id, branch_id, node_id, recorded_by, timestamp) VALUES
  ('led-se-ds-01', 'PARA-500',   -30, 'DISPENSE', 'se-ds-01', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-01-06 10:00:00'),
  ('led-se-ds-02', 'METF-500',   -60, 'DISPENSE', 'se-ds-02', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-01-06 10:15:00'),
  ('led-se-ds-03', 'AMOX-500',   -21, 'DISPENSE', 'se-ds-03', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-01-08 09:30:00'),
  ('led-se-ds-04', 'CETI-10',    -30, 'DISPENSE', 'se-ds-04', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-01-10 10:00:00'),
  ('led-se-ds-05', 'IBUP-400',   -20, 'DISPENSE', 'se-ds-05', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-01-12 09:45:00'),
  ('led-se-ds-06', 'ATOR-10',    -30, 'DISPENSE', 'se-ds-06', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-01-15 10:30:00'),
  ('led-se-ds-07', 'OMEP-20',    -30, 'DISPENSE', 'se-ds-07', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-01-18 09:00:00'),
  ('led-se-ds-08', 'VITC-500',   -60, 'DISPENSE', 'se-ds-08', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-01-22 11:00:00'),
  ('led-se-ds-09', 'PARA-500',   -30, 'DISPENSE', 'se-ds-09', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-04 10:00:00'),
  ('led-se-ds-10', 'METF-500',   -60, 'DISPENSE', 'se-ds-10', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-04 10:15:00'),
  ('led-se-ds-11', 'AZIT-250',    -6, 'DISPENSE', 'se-ds-11', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-06 09:30:00'),
  ('led-se-ds-12', 'CETI-10',    -30, 'DISPENSE', 'se-ds-12', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-10 10:00:00'),
  ('led-se-ds-13', 'ATOR-10',    -30, 'DISPENSE', 'se-ds-13', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-14 10:30:00'),
  ('led-se-ds-14', 'OMEP-20',    -30, 'DISPENSE', 'se-ds-14', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-16 09:00:00'),
  ('led-se-ds-15', 'IRON-FE',    -30, 'DISPENSE', 'se-ds-15', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-19 11:00:00'),
  ('led-se-ds-16', 'METR-500',   -14, 'DISPENSE', 'se-ds-16', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-02-22 10:00:00'),
  ('led-se-ds-17', 'PARA-500',   -30, 'DISPENSE', 'se-ds-17', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-04 10:00:00'),
  ('led-se-ds-18', 'METF-500',   -60, 'DISPENSE', 'se-ds-18', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-04 10:15:00'),
  ('led-se-ds-19', 'AMOX-500',   -21, 'DISPENSE', 'se-ds-19', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-06 09:30:00'),
  ('led-se-ds-20', 'CETI-10',    -30, 'DISPENSE', 'se-ds-20', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-10 10:00:00'),
  ('led-se-ds-21', 'ATOR-10',    -30, 'DISPENSE', 'se-ds-21', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-14 10:30:00'),
  ('led-se-ds-22', 'OMEP-20',    -30, 'DISPENSE', 'se-ds-22', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-16 09:00:00'),
  ('led-se-ds-23', 'VITC-500',   -60, 'DISPENSE', 'se-ds-23', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-20 11:00:00'),
  ('led-se-ds-24', 'IRON-FE',    -30, 'DISPENSE', 'se-ds-24', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-24 10:00:00'),
  ('led-se-ds-25', 'SALB-INH',    -1, 'DISPENSE', 'se-ds-25', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-03-28 09:00:00'),
  ('led-se-ds-26', 'PARA-500',   -30, 'DISPENSE', 'se-ds-26', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-04-02 10:00:00'),
  ('led-se-ds-27', 'METF-500',   -60, 'DISPENSE', 'se-ds-27', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-04-02 10:15:00'),
  ('led-se-ds-28', 'IBUP-400',   -20, 'DISPENSE', 'se-ds-28', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-04-08 09:45:00'),
  ('led-se-ds-29', 'OMEP-20',    -30, 'DISPENSE', 'se-ds-29', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-04-10 09:00:00'),
  ('led-se-ds-30', 'ATOR-10',    -30, 'DISPENSE', 'se-ds-30', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-04-14 10:30:00'),
  ('led-se-ds-31', 'IRON-FE',    -30, 'DISPENSE', 'se-ds-31', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-04-18 11:00:00'),
  ('led-se-ds-32', 'METR-500',   -14, 'DISPENSE', 'se-ds-32', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-04-22 10:00:00'),
  ('led-se-ds-33', 'VITC-500',   -60, 'DISPENSE', 'se-ds-33', 'branch-1', 'edge-branch-1', 'dr.khalil.mansour', '2026-04-25 11:00:00');

COMMIT;
PRAGMA foreign_keys=ON;
