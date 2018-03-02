package storage

type VarInt struct {
	value []uint8
}

func NewVarInt(v int64) *VarInt {
	result := VarInt{
		value: make([]uint8, 9),
	}

	result.value[8] = uint8(v & 0x7F)
	result.value[7] = uint8(v >> 6 & 0x7F)

	return &result
}

func (i *VarInt) ToInt64() int64 {
	result := int64(0)
	result = int64(i.value[0])
	result = result<<7 | int64(i.value[1])
	result = result<<7 | int64(i.value[2])
	result = result<<7 | int64(i.value[3])
	result = result<<7 | int64(i.value[4])
	result = result<<7 | int64(i.value[5])
	result = result<<7 | int64(i.value[6])
	result = result<<7 | int64(i.value[7])
	result = result<<7 | int64(i.value[8])
	return result
}
