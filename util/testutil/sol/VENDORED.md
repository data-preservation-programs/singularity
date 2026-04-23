# Vendored from filecoin-project/fvm-solidity

Source: https://github.com/filecoin-project/fvm-solidity
Pin:    b0404b9 (main, 2026-04-23, includes PR #16 datacap mock)

Vendored verbatim (no local modifications):
- `src/FVMActors.sol`, `FVMAddress.sol`, `FVMCodec.sol`, `FVMErrors.sol`,
  `FVMFlags.sol`, `FVMMethod.sol`, `FVMPrecompiles.sol`
- `src/mocks/FVMActor.sol`, `FVMCallActorByAddress.sol`, `FVMCallActorById.sol`

To audit drift against upstream:

    make -C util/testutil/sol drift
