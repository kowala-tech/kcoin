pragma solidity ^0.4.0;

contract Arrays {
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

    address[3] owners;
    address[] votes;

    SmallStruct[3] small_fixed;
    SmallStruct[] small_dynamic;

    MediumStruct[3] medium_fixed;
    MediumStruct[] medium_dynamic;

    BigStruct[3] big_fixed;
    BigStruct[] big_dynamic;

    function Arrays() public {
        uint8 i;
        owners[0] = 0xE92A2a4E3F4c378495145619F2975ce8c60819C2;
        owners[1] = 0x14Dd8d9c759A6827AACBF726085Ef13A357989ec;
        owners[2] = 0xA1F0a100522350Ee2A044Fe69831cf469C0f7123;
        for (i = 0; i < 3; i++) { votes.push(owners[i]); }
        for (i = 0; i < 3; i++) {
            small_fixed[i] = SmallStruct(i, i + 1);
            small_dynamic.push(SmallStruct(i, i + 1));
            medium_fixed[i] = MediumStruct(i, owners[i]);
            medium_dynamic.push(MediumStruct(i, owners[i]));
            big_fixed[i] = BigStruct(i, owners[i], i * 256);
            big_dynamic.push(BigStruct(i, owners[i], i * 256));
        }
    }
}