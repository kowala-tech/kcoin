pragma solidity ^0.4.0;

contract SmallInts {
	int8 i1 = -127;
	int16 i2;
	int24 i3;
	int32 i4;
	int40 i5;
	int48 i6;
	int56 i7;
	int64 i8;
	uint8 u1 = 129;
	uint16 u2;
	uint24 u3;
	uint32 u4;
	uint40 u5;
	uint48 u6;
	uint56 u7;
	uint64 u8;
	bool bool_byte = true;

	function SmallInts() public {
		uint256 n = 1;
		i2 = int16(i1) * 256 - int16(n); n++;
		i3 = int24(i2) * 256 - int24(n); n++;
		i4 =  int32(i3) * 256 - int32(n); n++;
		i5 = int40(i4) * 256 - int40(n); n++;
		i6 = int48(i5) * 256 - int48(n); n++;
		i7 = int56(i6) * 256 - int56(n); n++;
		i8 = int64(i7) * 256 - int64(n); n++;
		u2 = uint16(u1) * 256 - uint16(n); n++;
		u3 = uint24(u2) * 256 - uint24(n); n++;
		u4 = uint32(u3) * 256 - uint32(n); n++;
		u5 = uint40(u4) * 256 - uint40(n); n++;
		u6 = uint48(u5) * 256 - uint48(n); n++;
		u7 = uint56(u6) * 256 - uint56(n); n++;
		u8 = uint64(u7) * 256 - uint64(n); n++;

	}
}