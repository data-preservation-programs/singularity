UPSTREAM_REPO := filecoin-project/fvm-solidity
UPSTREAM_PIN  := b0404b968b8e78b213803b6835af85e5dcbaea7e

VENDORED_FILES := \
  src/FVMActors.sol \
  src/FVMAddress.sol \
  src/FVMCodec.sol \
  src/FVMErrors.sol \
  src/FVMFlags.sol \
  src/FVMMethod.sol \
  src/FVMPrecompiles.sol \
  src/mocks/FVMActor.sol \
  src/mocks/FVMCallActorByAddress.sol \
  src/mocks/FVMCallActorById.sol

OUTPUT_DIR := ..

.PHONY: build bytecode drift clean

build:
	forge build

# emit deployedBytecode for each precompile mock into util/testutil/<name>_bytecode.txt
bytecode: build
	forge inspect src/mocks/FVMActor.sol:FVMActor                       deployedBytecode > $(OUTPUT_DIR)/fvm_resolve_address_bytecode.txt
	forge inspect src/mocks/FVMCallActorById.sol:FVMCallActorById       deployedBytecode > $(OUTPUT_DIR)/fvm_call_actor_by_id_bytecode.txt
	forge inspect src/mocks/FVMCallActorByAddress.sol:FVMCallActorByAddress deployedBytecode > $(OUTPUT_DIR)/fvm_call_actor_by_address_bytecode.txt

drift:
	@tmp=$$(mktemp -d); trap "rm -rf $$tmp" EXIT; \
	for f in $(VENDORED_FILES); do \
	    curl -fsSL "https://raw.githubusercontent.com/$(UPSTREAM_REPO)/$(UPSTREAM_PIN)/$$f" -o "$$tmp/$$(basename $$f)" || { echo "fetch failed: $$f"; exit 1; }; \
	    if ! diff -u "$$f" "$$tmp/$$(basename $$f)" >/dev/null; then \
	        echo "DRIFT: $$f"; \
	        diff -u "$$f" "$$tmp/$$(basename $$f)"; \
	    fi; \
	done; \
	echo "drift check complete (vs $(UPSTREAM_PIN))"

clean:
	rm -rf out cache
