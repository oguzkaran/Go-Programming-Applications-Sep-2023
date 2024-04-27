package bitwise

import (
	"fmt"
)

func BitStrUInt8(a uint8) string {
	return fmt.Sprintf("%08b", a)
}

func BitStrUInt16(a uint16) string {
	return fmt.Sprintf("%016b", a)
}

func BitStrUInt32(a uint32) string {
	return fmt.Sprintf("%032b", a)
}

func BitStrUInt64(a uint64) string {
	return fmt.Sprintf("%064b", a)
}

func ClearBitUInt8(a, k uint8) uint8 {
	return a &^ (uint8(1) << k)
}

func ClearBitUInt16(a, k uint16) uint16 {
	return a &^ (uint16(1) << k)
}
func ClearBitUInt32(a, k uint32) uint32 {
	return a &^ (uint32(1) << k)
}

func ClearBitUInt64(a, k uint64) uint64 {
	return a &^ (uint64(1) << k)
}

func IsClearUInt32(a, k uint32) bool {
	return !IsSetUInt32(a, k)
}

func IsSetUInt32(a, k uint32) bool {
	return a&uint32(1)<<k != 0
}

func SetBitUInt8(a, k uint8) uint8 {
	return a | uint8(1)<<k
}

func SetBitUInt16(a, k uint16) uint16 {
	return a | uint16(1)<<k
}

func SetBitUInt32(a, k uint32) uint32 {
	return a | uint32(1)<<k
}

func SetBitUInt64(a, k uint64) uint64 {
	return a | uint64(1)<<k
}

func ToggleBitUInt32(a, k uint32) uint32 {
	return a ^ uint32(1)<<k
}

//...
