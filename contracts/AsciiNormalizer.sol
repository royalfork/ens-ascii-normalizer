// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

// https://github.com/Arachnid/solidity-stringutils/blob/master/src/strings.sol
struct slice {
	uint _len;
	uint _ptr;
}

function keccak(slice memory self) pure returns (bytes32 ret) {
	assembly {
		ret := keccak256(mload(add(self, 32)), mload(self))
	}
}

contract AsciiNormalizer {
	// Each index in idnamap corresponds to the ascii number.
	// Value at index location is either:
	//  0x00: sendtinel value for DISALLOWED
	//  0x01: sendtinel value for VALID
	//  other: mapped value
	// TODO inspect actual memory?
	bytes1[] public idnamap;

	constructor (bytes memory asciimap) {
		for(uint i = 0; i < asciimap.length; i += 2) {
			bytes1 r = asciimap[i+1];
			for(uint8 j = 0; j < uint8(asciimap[i]); j++) {
				idnamap.push(r);
			}
		}
	}

	// TODO external vs public?
	function valid(string calldata domain) external view returns (bool) {
		for (uint256 i = 0; i < bytes(domain).length; i++) {
			bytes1 c = bytes(domain)[i];
			if (c == '.') {
				continue;
			}
			if (c > 0x80) {
				return false;
			}
			if (uint8(idnamap[uint8(c)]) != 1) {
				return false;
			}
		}
		return true;
	}

	function normalize(string memory domain) external view returns (string memory) {
		for (uint256 i = 0; i < bytes(domain).length; i++) {
			bytes1 c = bytes(domain)[i];
			if (c == '.') {
				continue;
			}

			require(c < 0x80);
			bytes1 r = idnamap[uint8(c)];
			require(r > 0);
			if (uint8(r) > 1) {
				bytes(domain)[i] = r;
			}
		}
		return domain;
	}

	function namehash(string memory domain) external view returns (string memory, bytes32) {
		uint256 i = bytes(domain).length;

		// Process labels.
		uint256 lastDot = i;
		bytes32 node = bytes32(0);
		bytes32 lh;
		for (; i > 0; i--) {
			bytes1 c = bytes(domain)[i-1];

			if (c == '.') {
				lh = labelHash(domain, i, lastDot);
				node = keccak256(abi.encodePacked(node, lh));
				lastDot = i-1;
				continue;
			}

			require(c < 0x80);
			bytes1 r = idnamap[uint8(c)];
			require(r > 0);
			if (uint8(r) > 1) {
				bytes(domain)[i-1] = r;
			}
		}
		lh = labelHash(domain, i, lastDot);
		return (domain, keccak256(abi.encodePacked(node, lh)));
	}


	// Should this return labelhash?
	// TODO using memory will pass by value
	function labelHash(string memory domain, uint start, uint end) internal pure returns (bytes32) {
		// Convert to slice, and return hash.
		uint ptr;
		uint offset = 32+start;
		assembly {
			ptr := add(domain, offset)
		}
		return keccak(slice(end-start, ptr));
	}
}
