pragma solidity ^0.4.0;

contract SmallInts {
	int8 i1 = -127;
	int16 i2 = int16(i1 - 1) * 256 + 1;
	int24 i3 = int24(i2 - 1) * 256 + 1;
	int32 i4 =  int32(i3 - 1) * 256 + 1;
	int40 i5 = int40(i4 - 1) * 256 + 1;
	int48 i6 = int48(i5 - 1) * 256 + 1;
	int56 i7 = int56(i6 - 1) * 256 + 1;
	int64 i8 = int64(i7 - 1) * 256 + 1;

	uint8 u1 = 129;
	uint16 u2 = uint16(u1 - 1) * 256 + 1;
	uint24 u3 = uint24(u2 - 1) * 256 + 1;
	uint32 u4 = uint32(u3 - 1) * 256 + 1;
	uint40 u5 = uint40(u4 - 1) * 256 + 1;
	uint48 u6 = uint48(u5 - 1) * 256 + 1;
	uint56 u7 = uint56(u6 - 1) * 256 + 1;
	uint64 u8 = uint64(u7 - 1) * 256 + 1;

	bool bool_byte = true;
}