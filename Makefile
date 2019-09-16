GO = go
V  = 0
Q  = $(if $(filter 1,$V),,@)
M  = $(shell printf "\033[34;1m▶\033[0m")
OS = $(shell uname -s | tr A-Z a-z)

.PHONY: all
all:

all: build

build:
	$Q $(GO) build ./...

generate:
	$Q $(GO) generate ./...

# Tests
TESTFLAGS = -race -v
TESTSUITE = ./...
.PHONY: test
test:
	$Q $(GO) test $(TESTFLAGS) `go list $(TESTSUITE)`

# Misc
.PHONY: clean
clean: ; $(info $(M) cleaning...) @ ## Clean up everything
