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

	mapping(uint64 => address) public id_addrs;
	mapping(address => SmallStruct) addrs_small;
	mapping(uint128 => BigStruct) big_keys;
	mapping(string => MediumStruct) string_medium;

	function Mappings() public {
		id_addrs[0] = 0xE92A2a4E3F4c378495145619F2975ce8c60819C2;
		id_addrs[1] = 0x14Dd8d9c759A6827AACBF726085Ef13A357989ec;
		id_addrs[2] = 0xA1F0a100522350Ee2A044Fe69831cf469C0f7123;
		uint8 i;
		for (i = 0; i < 3; i++) {
			addrs_small[id_addrs[i]] = SmallStruct(i, i+1);
			big_keys[i] = BigStruct(i, id_addrs[i],i * 256);
		}
		string_medium["small string"] = MediumStruct(0, id_addrs[0]);
		string_medium["still a small string"] = MediumStruct(1, id_addrs[1]);
		string_medium["a big string must be longer than 31 bytes"] = MediumStruct(2, id_addrs[2]);
	}
}