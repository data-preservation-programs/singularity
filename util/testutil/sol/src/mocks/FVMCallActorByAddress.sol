// SPDX-Lidense-Identifier: Apache-2.0 OR MIT
pragma solidity ^0.8.30;

import {EMPTY_CODEC} from "../FVMCodec.sol";
import {EXIT_SUCCESS, INSUFFICIENT_FUNDS} from "../FVMErrors.sol";
import {NO_FLAGS} from "../FVMFlags.sol";
import {SEND} from "../FVMMethod.sol";

contract FVMCallActorByAddress {
    fallback() external payable {
        (uint64 method, uint256 value, uint64 flags, uint64 codec, bytes memory params, bytes memory filAddress) =
            abi.decode(msg.data, (uint64, uint256, uint64, uint64, bytes, bytes));

        // Verify this is a burn operation (actor ID 99, method 0)
        require(filAddress.length > 2, "FVMCallActorByAddress: Invalid short address");
        require(filAddress[0] == 0x04, "FVMCallActorByAddress: Only f4 addresses supported");
        require(filAddress[1] == 0x0a, "FVMCallActorByAddress: Only f410 addresses supported");
        require(filAddress.length == 22, "FVMCallActorByAddress: Invalid f410 address length");

        require(method == SEND, "FVMCallActorByAddress: Only method 0 (send) supported");
        require(flags == NO_FLAGS, "FVMCallActorByAddress: Only non-readonly calls supported");
        require(codec == EMPTY_CODEC, "FVMCallActorByAddress: Only no-codec calls supported");
        require(params.length == 0, "FVMCallActorByAddress: No params expected");

        address payable recipient;
        assembly ("memory-safe") {
            recipient := mload(add(22, filAddress))
        }

        // Perform the transfer
        (bool success,) = recipient.call{value: value}("");

        // Prepare the response in FVM format: exit_code(i256) | codec(u64) | return_value(bytes)
        bytes memory response;
        if (success) {
            // Success: exit code 0
            response = abi.encode(EXIT_SUCCESS, EMPTY_CODEC, bytes(""));
        } else {
            // Failure: exit code -5 (InsufficientFunds)
            response = abi.encode(INSUFFICIENT_FUNDS, EMPTY_CODEC, bytes(""));
        }

        // Return the response using assembly to properly handle delegatecall return
        assembly ("memory-safe") {
            return(add(response, 0x20), mload(response))
        }
    }
}
