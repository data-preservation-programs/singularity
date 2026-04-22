// SPDX-Lidense-Identifier: Apache-2.0 OR MIT
pragma solidity ^0.8.30;

import {BURN_ACTOR_ID, BURN_ADDRESS, DATACAP_TOKEN_ACTOR_ID} from "../FVMActors.sol";
import {CBOR_CODEC, EMPTY_CODEC} from "../FVMCodec.sol";
import {EXIT_SUCCESS, INSUFFICIENT_FUNDS} from "../FVMErrors.sol";
import {NO_FLAGS} from "../FVMFlags.sol";
import {DATACAP_TRANSFER, SEND} from "../FVMMethod.sol";

contract FVMCallActorById {
    /// @dev CBOR `TransferReturn` for a single-allocation DataCap -> VerifReg
    /// transfer: [from_balance(empty), to_balance(empty), recipient_data] where
    /// recipient_data is a CBOR `VerifregResponse`:
    /// [allocationResults([1, []]), extensionResults([0, []]), [allocId=66]]
    /// multi-piece flows need a different fixture
    bytes constant DEFAULT_DATACAP_TRANSFER_RETURN = hex"8340404a83820180820080811842";

    struct Message {
        uint64 method;
        uint256 value;
        uint64 flags;
        uint64 codec;
        bytes params;
        uint64 actorId;
    }

    fallback() external payable {
        Message memory m;
        (m.method, m.value, m.flags, m.codec, m.params, m.actorId) =
            abi.decode(msg.data, (uint64, uint256, uint64, uint64, bytes, uint64));

        bytes memory response;
        if (m.actorId == BURN_ACTOR_ID) {
            response = _handleBurn(m);
        } else if (m.actorId == DATACAP_TOKEN_ACTOR_ID) {
            response = _handleDataCap(m);
        } else {
            revert("FVMCallActorById: unsupported actor");
        }

        assembly ("memory-safe") {
            return(add(response, 0x20), mload(response))
        }
    }

    function _handleBurn(Message memory m) private returns (bytes memory) {
        require(m.method == SEND, "FVMCallActorById: burn requires method 0 (send)");
        require(m.flags == NO_FLAGS, "FVMCallActorById: Only non-readonly calls supported");
        require(m.codec == EMPTY_CODEC, "FVMCallActorById: Only no-codec calls supported");
        require(m.params.length == 0, "FVMCallActorById: No params expected");

        (bool success,) = BURN_ADDRESS.call{value: m.value}("");
        return success
            ? abi.encode(EXIT_SUCCESS, EMPTY_CODEC, bytes(""))
            : abi.encode(INSUFFICIENT_FUNDS, EMPTY_CODEC, bytes(""));
    }

    function _handleDataCap(Message memory m) private pure returns (bytes memory) {
        require(m.method == DATACAP_TRANSFER, "FVMCallActorById: DataCap only supports Transfer");
        return abi.encode(EXIT_SUCCESS, CBOR_CODEC, DEFAULT_DATACAP_TRANSFER_RETURN);
    }
}
