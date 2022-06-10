// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

// TODO natspec docs
contract AsciiNormalizer {
	// Each index in idnamap corresponds to the ascii code point.
	// Value at index location is either:
	//  0x00: sendtinel value for DISALLOWED
	//  0x01: sendtinel value for VALID
	//  other: mapped code point
	// Using bytes1[] allows for indexed based public access, yet
	// storage array is still tightly packed.
	bytes1[] public idnamap;

	constructor (bytes memory asciimap) {
		for(uint i = 0; i < asciimap.length; i += 2) {
			bytes1 r = asciimap[i+1];
			for(uint8 j = 0; j < uint8(asciimap[i]); j++) {
				idnamap.push(r);
			}
		}
	}

	// TODO natspec docs
	function namehash(string memory domain) external view returns (string memory, bytes32) {
		uint256 i = bytes(domain).length;

		// Process labels.
		uint256 lastDot = i;
		bytes32 node = bytes32(0);
		for (; i > 0; i--) {
			bytes1 c = bytes(domain)[i-1];

			if (c == '.') {
				node = keccak256(abi.encodePacked(node, labelHash(domain, i, lastDot)));
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
		return (domain, keccak256(abi.encodePacked(node, labelHash(domain, i, lastDot))));
	}

	function labelHash(string memory domain, uint start, uint end) internal pure returns (bytes32 hash) {
		assembly ("memory-safe") {
			hash := keccak256(add(add(domain, 0x20), start), sub(end, start))
		}
	}
}
