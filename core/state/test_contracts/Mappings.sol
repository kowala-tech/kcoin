pragma solidity ^0.4.0;

contract Mappings {
    struct SmallStruct {
		uint64 id;
		uint32 nonce;
	}

    struct MediumStruct {
		uint64 id;
		address addr;
	}

	struct BigStruct {
		uint128 id;
		address addr;
		uint128 nonce;
	}

	mapping(uint64 => address) id_addrs;
	mapping(address => SmallStruct) addrs_small;
	mapping(uint128 => BigStruct) big_keys;
	mapping(string => MediumStruct) small_medium;

	function Mappings() {
		id_addrs[0] = 0xe92a2a4e3f4c378495145619f2975ce8c60819c2;
		id_addrs[1] = 0x14dd8d9c759a6827aacbf726085ef13a357989ec;
		id_addrs[2] = 0xa1f0a100522350ee2a044fe69831cf469c0f7123;
		uint8 i;
		for (i = 0; i < 3; i++) {
			addrs_small[id_addrs[i]] = SmallStruct(i, i+1);
			big_keys[i] = BigStruct(i, id_addrs[i],i * 256);
		}
		small_medium["small string"] = MediumStruct(0, id_addrs[0]);
		small_medium["still a small string"] = MediumStruct(1, id_addrs[1]);
		small_medium["a big string must be longer than 31 bytes"] = MediumStruct(2, id_addrs[2]);
	}
}