// SPDX-License-Identifier: Apache-2.0 OR MIT
pragma solidity ^0.8.30;

// See filecoin-project/ref-fvm shared/src/error/mod.rs

int256 constant EXIT_SUCCESS = 0;

// @dev A syscall parameters was invalid.
int256 constant ILLEGAL_ARGUMENT = -1;
// @dev The actor is not in the correct state to perform the requested operation.
int256 constant ILLEGAL_OPERATION = -2;
// @dev This syscall would exceed some system limit (memory, lookback, call depth, etc.).
int256 constant LIMIT_EXCEEDED = -3;
// @dev A system-level assertion has failed.
int256 constant ASSERTION_FAILED = -4;
// @dev There were insufficient funds to complete the requested operation.
int256 constant INSUFFICIENT_FUNDS = -5;
// @dev A resource was not found.
int256 constant NOT_FOUND = -6;
// @dev The specified IPLD block handle was invalid.
int256 constant INVALID_HANDLE = -7;
// @dev The requested CID shape (multihash codec, multihash length) isn't supported.
int256 constant ILLEGAL_CID = -8;
// @dev The requested IPLD codec isn't supported.
int256 constant ILLEGAL_CODEC = -9;
// @dev The IPLD block did not match the specified IPLD codec.
int256 constant SERIALIZATION = -10;
// @dev The operation is forbidden.
int256 constant FORBIDDEN = -11;
// @dev The passed buffer is too small.
int256 constant BUFFER_TOO_SMALL = -12;
// @dev The actor is executing in a read-only context.
int256 constant READ_ONLY = -13;
