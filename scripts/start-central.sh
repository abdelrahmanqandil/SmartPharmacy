#!/usr/bin/env bash
set -e

APP_CFG="config.yaml"
MARMOT_CFG="marmot-central.toml"

# ── dependency checks ────────────────────────────────────────────────────────
if [ ! -x "./nats-server" ]; then
  echo "ERROR: nats-server binary not found."
  echo "       Run 'make deps' to download it, or place it in this directory."
  exit 1
fi
if [ ! -x "./marmot" ]; then
  echo "ERROR: marmot binary not found."
  echo "       Run 'make deps' to download it, or place it in this directory."
  exit 1
fi

trap 'echo; echo "[central] shutting down..."; kill $(jobs -p) 2>/dev/null; wait' INT TERM

printf '\n'
printf '  ╔══════════════════════════════════════════════════════╗\n'
printf '  ║  Smart Pharmacy  ·  Central Node                    ║\n'
printf '  ║  Replication engine: Marmot + NATS                  ║\n'
printf '  ╚══════════════════════════════════════════════════════╝\n'
printf '\n'

# ── 1. NATS server (embedded transport for Marmot) ───────────────────────────
printf '[nats] starting server on port 4222...\n'
./nats-server -p 4222 &
NATS_PID=$!

# wait until NATS is accepting connections
for i in $(seq 1 20); do
  if nc -z 127.0.0.1 4222 2>/dev/null; then
    printf '[nats] ready\n'
    break
  fi
  sleep 0.3
done

# ── 2. Marmot sidecar ────────────────────────────────────────────────────────
printf '[marmot] starting central sidecar (node_id=1)...\n'
./marmot -config "$MARMOT_CFG" &

sleep 1
printf '[marmot] sidecar running · replicating: medications · patients · transactions · inventory_ledger\n'

# ── 3. Application server ────────────────────────────────────────────────────
printf '[app] starting central node on :9090...\n'
printf '\n'

go run ./cmd/app -config "$APP_CFG"

wait
