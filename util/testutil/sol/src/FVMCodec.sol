// SPDX-License-Identifier: Apache-2.0 OR MIT
pragma solidity ^0.8.30;

// @dev empty returned data
uint64 constant EMPTY_CODEC = 0;

// @dev actor returned unencoded raw data
uint64 constant RAW_CODEC = 0x55;

// @dev CBOR encoded data (IPLD dag-cbor)
uint64 constant CBOR_CODEC = 0x51;
