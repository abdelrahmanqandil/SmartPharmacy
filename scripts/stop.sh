#!/usr/bin/env bash
echo "[stop] killing app, marmot, nats-server..."
pkill -f "go run ./cmd/app" 2>/dev/null || true
pkill -f "cmd/app"           2>/dev/null || true
pkill -9 -f "marmot"         2>/dev/null || true
pkill -f "nats-server"       2>/dev/null || true
echo "[stop] done"
