#!/usr/bin/env bash
set -e

APP_CFG="config-edge.yaml"
MARMOT_CFG="marmot-edge.toml"

# ── dependency checks ────────────────────────────────────────────────────────
if [ ! -x "./marmot" ]; then
  echo "ERROR: marmot binary not found."
  echo "       Run 'make deps' to download it, or place it in this directory."
  exit 1
fi

trap 'echo; echo "[edge] shutting down..."; kill $(jobs -p) 2>/dev/null; wait' INT TERM

printf '\n'
printf '  ╔══════════════════════════════════════════════════════╗\n'
printf '  ║  Smart Pharmacy  ·  Edge Node (Branch)              ║\n'
printf '  ║  Replication engine: Marmot + NATS                  ║\n'
printf '  ╚══════════════════════════════════════════════════════╝\n'
printf '\n'

# ── 1. Marmot sidecar ────────────────────────────────────────────────────────
printf '[marmot] starting edge sidecar (node_id=2)...\n'
printf '[marmot] connecting to NATS at nats://127.0.0.1:4222\n'
printf '[marmot] seeding from central node at 127.0.0.1:8080\n'
./marmot -config "$MARMOT_CFG" &

sleep 1
printf '[marmot] sidecar running · bidirectional replication active\n'

# ── 2. Application server ────────────────────────────────────────────────────
printf '[app] starting edge node on :9091...\n'
printf '\n'

go run ./cmd/app -config "$APP_CFG"

wait
