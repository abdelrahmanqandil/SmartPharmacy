# Smart Pharmacy

A distributed pharmacy management system demonstrating bidirectional SQLite replication between a central hospital node and edge branch nodes using **Marmot** and **NATS**.

## Architecture

Two nodes — a central hospital and a branch — each run a Go HTTP server backed by a local SQLite database. A Marmot sidecar process runs alongside each node and watches the SQLite WAL. Changes are published through a shared NATS server, so both databases stay in sync without any replication logic in the application code.

- Central: `:9090`, `central.db`, Marmot node 1
- Edge: `:9091`, `edge.db`, Marmot node 2
- NATS on `:4222` (started with the central node)

Replicated tables: `medications`, `patients`, `transactions`, `inventory_ledger`

## Prerequisites

- **Go 1.21+**
- **Python 3** (for seeding demo data)
- **make**
- `nc` / `netcat` (used by startup script to wait for NATS)

## Quick Start

### 1. Download dependencies

```bash
make deps
```

This downloads the `marmot` and `nats-server` binaries into the project directory.

### 2. Seed demo data (optional)

```bash
make seed-central
make seed-edge
```

Loads 5 months of realistic pharmacy data (25 medications, 20 patients, ~150 transactions).

### 3. Start the central node

```bash
make central   # terminal 1
```

Starts: NATS server → Marmot sidecar → Go API on :9090

### 4. Start the edge node

```bash
make edge      # terminal 2
```

Starts: Marmot sidecar → Go API on :9091

After a few seconds, both databases will be in sync.

## API Reference

All endpoints (except `/api/status`) require the header:
```
X-Pharmacy-Token: dev-pharmacist-token
```
or `Authorization: Bearer dev-pharmacist-token`.

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/status` | Node health and replication state |
| GET | `/api/medications` | List all medications |
| POST | `/api/medications/register` | Register a medication |
| GET | `/api/patients` | List all patients |
| POST | `/api/patients/register` | Register a patient |
| GET | `/api/inventory` | Global inventory |
| GET | `/api/inventory?branch=branch-1` | Branch-level inventory |
| GET | `/api/transactions` | Recent transactions (last 200) |
| POST | `/api/stock/add` | Add stock |
| POST | `/api/dispense` | Dispense medication |
| GET | `/api/sync/export` | Export full snapshot (JSON) |
| POST | `/api/sync/import` | Import snapshot |

### Example: Dispense

```bash
curl -s -X POST http://localhost:9091/api/dispense \
  -H "X-Pharmacy-Token: dev-pharmacist-token" \
  -H "Content-Type: application/json" \
  -d '{"medication_id":"PARA-500","quantity":10,"patient_id":"PAT-003","pharmacist_id":"dr.khalil.mansour"}'
```

## Configuration

| File | Purpose |
|------|---------|
| `config.yaml` | Central node config (port, DB path, auth token) |
| `config-edge.yaml` | Edge node config |
| `marmot-central.toml` | Marmot sidecar config for central (node_id=1, gRPC :8080) |
| `marmot-edge.toml` | Marmot sidecar config for edge (node_id=2, seeds from central :8080) |

Change the `auth.token` in `config.yaml` / `config-edge.yaml` for production deployments. The token can also be set via the `PHARMACY_AUTH_TOKEN` environment variable.

## Multi-Machine Deployment

To run central and edge on separate machines, update the IP addresses in:
- `marmot-central.toml` → `grpc_advertise_address` and `transport.address`
- `marmot-edge.toml` → `transport.address` and `cluster.seed_nodes`
- `scripts/start-central.sh` / `scripts/start-edge.sh` → NATS server host

## Makefile Targets

| Target | Description |
|--------|-------------|
| `make deps` | Download `marmot` and `nats-server` binaries |
| `make central` | Start NATS + Marmot + central app |
| `make edge` | Start Marmot + edge app |
| `make stop` | Kill all running processes |
| `make seed-central` | Load demo data into central.db |
| `make seed-edge` | Load demo data into edge.db |
| `make build` | Build the app binary |
| `make clean` | Remove databases and Marmot state |
