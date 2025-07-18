# Generated by Claude Code Bootstrap
# This Makefile uses the include pattern to prevent merge conflicts
# Human-maintained targets should be added here
# Agent-generated targets are in .maestro/makefiles/

# Include core makefile
-include .maestro/makefiles/core.mk

# Include platform-specific makefiles
-include .maestro/makefiles/go.mk
-include .maestro/makefiles/node.mk
-include .maestro/makefiles/python.mk
-include .maestro/makefiles/react.mk
-include .maestro/makefiles/docker.mk

# Default targets if no makefiles exist
.PHONY: help

# Only define fallback targets if core.mk doesn't exist
ifeq ($(wildcard .maestro/makefiles/core.mk),)
build:
	@echo "No build configured. Run bootstrap to set up build system."

test:
	@echo "No tests configured. Run bootstrap to set up test system."

lint:
	@echo "No linting configured. Run bootstrap to set up linting system."

run:
	@echo "No run target configured. Run bootstrap to set up run system."
endif

help:
	@echo "Available targets:"
	@echo "  build      - Build the project"
	@echo "  test       - Run tests"
	@echo "  lint       - Run linting"
	@echo "  run        - Run the application"
	@echo "  clean      - Clean build artifacts"
	@echo "  help       - Show this help message"
	@echo ""
	@echo "Platform-specific help:"
	@echo "  go-help    - Show Go targets (if available)"
	@echo "  node-help  - Show Node.js targets (if available)"
	@echo "  python-help - Show Python targets (if available)"
	@echo "  react-help - Show React targets (if available)"
	@echo "  docker-help - Show Docker targets (if available)"
	@echo ""
	@echo "This project uses Claude Code Bootstrap for build management."
	@echo "Agent-generated targets are in .maestro/makefiles/"
