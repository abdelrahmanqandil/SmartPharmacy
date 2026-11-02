.PHONY: central edge stop clean seed-central seed-edge deps build

# ── OS / arch detection ───────────────────────────────────────────────────────
OS   := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/')

MARMOT_VERSION    := 0.9.3
NATS_VERSION      := 2.10.7

# ── targets ───────────────────────────────────────────────────────────────────
build:
	go build -o app ./cmd/app

central:
	@bash scripts/start-central.sh

edge:
	@bash scripts/start-edge.sh

stop:
	@bash scripts/stop.sh

seed-central:
	@echo "Seeding central.db with demo data..."
	python3 -c "import sqlite3; c=sqlite3.connect('central.db'); c.executescript(open('scripts/seed-central.sql').read()); c.commit(); c.close()"
	@echo "Done."

seed-edge:
	@echo "Seeding edge.db with demo data..."
	python3 -c "import sqlite3; c=sqlite3.connect('edge.db'); c.executescript(open('scripts/seed-edge.sql').read()); c.commit(); c.close()"
	@echo "Done."

clean:
	@echo "Removing databases and Marmot state..."
	-pkill -9 -f "nats-server|marmot" 2>/dev/null || true
	rm -rf marmot-data marmot-data-edge central.db edge.db app
	@echo "Clean complete."

# ── deps: download marmot + nats-server ──────────────────────────────────────
deps:
	@echo "Downloading Marmot v$(MARMOT_VERSION) for $(OS)/$(ARCH)..."
	@if [ "$(OS)" = "linux" ]; then \
		curl -fsSL "https://github.com/maxpert/marmot/releases/download/v$(MARMOT_VERSION)/marmot-v$(MARMOT_VERSION)-linux-$(ARCH).tar.gz" | tar -xz marmot; \
	elif [ "$(OS)" = "darwin" ]; then \
		curl -fsSL "https://github.com/maxpert/marmot/releases/download/v$(MARMOT_VERSION)/marmot-v$(MARMOT_VERSION)-darwin-$(ARCH).tar.gz" | tar -xz marmot; \
	else \
		echo "Unsupported OS: $(OS). Download marmot manually from https://github.com/maxpert/marmot/releases"; \
	fi
	@chmod +x marmot 2>/dev/null || true
	@echo ""
	@echo "Downloading NATS Server v$(NATS_VERSION) for $(OS)/$(ARCH)..."
	@if [ "$(OS)" = "linux" ]; then \
		curl -fsSL "https://github.com/nats-io/nats-server/releases/download/v$(NATS_VERSION)/nats-server-v$(NATS_VERSION)-linux-$(ARCH).tar.gz" | tar -xz --strip-components=1 "nats-server-v$(NATS_VERSION)-linux-$(ARCH)/nats-server"; \
	elif [ "$(OS)" = "darwin" ]; then \
		curl -fsSL "https://github.com/nats-io/nats-server/releases/download/v$(NATS_VERSION)/nats-server-v$(NATS_VERSION)-darwin-$(ARCH).tar.gz" | tar -xz --strip-components=1 "nats-server-v$(NATS_VERSION)-darwin-$(ARCH)/nats-server"; \
	else \
		echo "Unsupported OS: $(OS). Download nats-server manually from https://github.com/nats-io/nats-server/releases"; \
	fi
	@chmod +x nats-server 2>/dev/null || true
	@echo ""
	@echo "Dependencies ready. Run 'make central' to start the central node."
