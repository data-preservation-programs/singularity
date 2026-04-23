// SPDX-License-Identifier: Apache-2.0 OR MIT
pragma solidity ^0.8.30;

library FVMAddress {
    /// @notice Creates an f0 (ID) address in bytes using unsigned LEB128 encoding
    function f0(uint64 actorId) internal pure returns (bytes memory buffer) {
        // Max size: 1 protocol byte + 10 bytes for uint64 LEB128 encoding
        buffer = new bytes(11);
        buffer[0] = 0x00; // Protocol byte for f0

        uint256 i = 1;

        do {
            uint8 byteVal = uint8(actorId & 0x7F); // Take 7 bits
            actorId >>= 7; // Shift right by 7 bits
            if (actorId != 0) {
                byteVal |= 0x80; // Set MSB if more bytes follow
            }
            buffer[i++] = bytes1(byteVal);
        } while (actorId != 0);

        assembly ("memory-safe") {
            mstore(buffer, i) // Set the correct length of the bytes array
        }
    }

    /// @notice Creates an f4 (delegated) address in bytes
    /// @dev NOTE: Only supports namespaces < 128 (single byte varint).
    function f4(uint8 namespace, bytes20 subaddress) internal pure returns (bytes memory) {
        return abi.encodePacked(uint8(0x04), namespace, subaddress);
    }

    /// @notice Creates an f410 address for a Solidity address
    function f410(address addr) internal pure returns (bytes memory) {
        return f4(0x0a, bytes20(addr));
    }
}
