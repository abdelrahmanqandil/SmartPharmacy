-- ─────────────────────────────────────────────────────────────────────────────
-- Seed data for edge.db
-- Scenario: Al-Nour Community Branch (branch-1 / edge node)
--   Pharmacist : Dr. Khalil Mansour  (dr.khalil.mansour)
--   Period     : December 2025 – April 2026
--
-- The branch stocks 12 common medications — a focused subset of the central
-- formulary.  All 20 patients are present because the patient registry
-- replicates bidirectionally from central.
-- ─────────────────────────────────────────────────────────────────────────────

PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;

DELETE FROM replication_log;
DELETE FROM inventory_ledger;
DELETE FROM transactions;
DELETE FROM patients;
DELETE FROM medications;

-- ═════════════════════════════════════════════════════════════════════════════
--  MEDICATIONS  — branch formulary (12 common items)
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
  ('METR-500',  'Metronidazole 500mg',          'tablet',   50);

-- ═════════════════════════════════════════════════════════════════════════════
--  PATIENTS  — full registry (replicated from central)
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
--  BRANCH-1 — initial procurement (Dec 5, 2025)
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
--  BRANCH-1 — restocks (Feb & Mar 2026)
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
--  BRANCH-1 — dispenses (Jan–Apr 2026)
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
