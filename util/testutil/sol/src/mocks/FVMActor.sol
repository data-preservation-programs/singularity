// SPDX-License-Identifier: Apache-2.0 OR MIT
pragma solidity ^0.8.30;

import {
    SYSTEM_ACTOR_ID,
    INIT_ACTOR_ID,
    REWARD_ACTOR_ID,
    CRON_ACTOR_ID,
    STORAGE_POWER_ACTOR_ID,
    STORAGE_MARKET_ACTOR_ID,
    VERIFIED_REGISTRY_ACTOR_ID,
    DATACAP_TOKEN_ACTOR_ID,
    EAM_ACTOR_ID,
    BURN_ACTOR_ID
} from "../FVMActors.sol";
import {FVMAddress} from "../FVMAddress.sol";

contract FVMActor {
    using FVMAddress for uint64;

    /// @dev keccak256 of f0(0) = keccak256(hex"0000") — protocol byte 0x00 + ULEB128(0) = 0x00
    bytes32 constant SYSTEM_ACTOR_HASH = keccak256(hex"0000");

    // Mapping from Filecoin address bytes to actor ID
    mapping(bytes32 => uint64) public addressMocks;

    constructor() {
        // Protocol 0 (f0) mappings for system singleton actors
        _mockf0(SYSTEM_ACTOR_ID);
        _mockf0(INIT_ACTOR_ID);
        _mockf0(REWARD_ACTOR_ID);
        _mockf0(CRON_ACTOR_ID);
        _mockf0(STORAGE_POWER_ACTOR_ID);
        _mockf0(STORAGE_MARKET_ACTOR_ID);
        _mockf0(VERIFIED_REGISTRY_ACTOR_ID);
        _mockf0(DATACAP_TOKEN_ACTOR_ID);
        _mockf0(EAM_ACTOR_ID);
        _mockf0(BURN_ACTOR_ID);
    }

    /**
     * @notice Mocks an ID-based address (f0)
     * @dev NOTE: This simple encoding only works for actorId <= 127.
     * Filecoin uses LEB128 Varint encoding for ID addresses. For values <= 127,
     * the Varint is a single byte. For values > 127, the Varint expands to
     * multiple bytes, and this 2-byte packing [0x00, id] would be invalid.
     */
    function _mockf0(uint64 actorId) internal {
        // Protocol 0 address = protocol byte (0x00) + actor ID
        bytes memory filAddress = actorId.f0();
        addressMocks[keccak256(filAddress)] = actorId;
    }

    /// @notice Mock a Filecoin address resolution
    /// @param filAddress The Filecoin address bytes
    /// @param actorId The actor ID to return (0 means doesn't exist)
    function mockResolveAddress(bytes memory filAddress, uint64 actorId) external {
        addressMocks[keccak256(filAddress)] = actorId;
    }

    /// @notice Mock a Solidity address resolution
    /// @param addr The Solidity address
    /// @param actorId The actor ID to return (0 means doesn't exist)
    function mockResolveAddress(address addr, uint64 actorId) external {
        bytes memory filAddress = abi.encodePacked(uint8(0x04), uint8(0x0a), addr);
        addressMocks[keccak256(filAddress)] = actorId;
    }

    fallback() external {
        bytes memory filAddress = msg.data;

        // Basic validation:  address must be non-empty
        require(filAddress.length > 0, "Invalid address:  empty");

        // Check first byte for valid protocol
        // f0 = 0x00, f1 = 0x01, f2 = 0x02, f3 = 0x03, f4 = 0x04
        uint8 protocol = uint8(filAddress[0]);
        require(protocol <= 0x04, "Invalid address: unknown protocol");

        bytes32 key;
        // Equivalent to `keccak256(filAddress)`.
        // Hash the data section of `bytes memory` directly:
        // length at offset 0x00, data at 0x20.
        assembly {
            key := keccak256(add(filAddress, 0x20), mload(filAddress))
        }

        // Look up mocked actor ID
        uint64 actorId = addressMocks[key];

        // If actor exists (actorId > 0), return it as ABI-encoded uint64
        // Special case: SYSTEM_ACTOR_ID is 0, but we want to allow it to be mocked and returned
        if (actorId > 0 || key == SYSTEM_ACTOR_HASH) {
            bytes memory response = abi.encode(uint256(actorId));
            assembly ("memory-safe") {
                return(add(response, 0x20), 32)
            }
        }

        // If actor doesn't exist (actorId == 0), return empty
        assembly ("memory-safe") {
            stop()
        }
    }
}
