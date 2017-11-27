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

    function Arrays() {
        uint8 i;
        owners[0] = 0xe92a2a4e3f4c378495145619f2975ce8c60819c2;
        owners[1] = 0x14dd8d9c759a6827aacbf726085ef13a357989ec;
        owners[2] = 0xa1f0a100522350ee2a044fe69831cf469c0f7123;
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