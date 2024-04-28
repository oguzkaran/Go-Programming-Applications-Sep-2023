package bitwise

import (
	"fmt"
	"strings"
)

func BitStrUInt8(a uint8) string {
	return fmt.Sprintf("%08b", a)
}

func BitStrUInt16(a uint16) string {
	return fmt.Sprintf("%016b", a)
}

func BitStrUInt32(a uint32) string {
	var sb strings.Builder

	for i := 31; i >= 0; i-- {
		if IsSetUInt32(a, uint32(i)) {
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
	}

	return sb.String()
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
func ClearBitCountUInt32(a uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if IsClearUInt32(a, uint32(i)) {
			count++
		}
	}

	return count
}

func HighestClearBitIndexUInt32(a uint32) int {
	for i := 31; i >= 0; i-- {
		if IsClearUInt32(a, uint32(i)) {
			return i
		}
	}

	return -1
}

func HighestSetBitIndexUInt32(a uint32) int {
	for i := 31; i >= 0; i-- {
		if IsSetUInt32(a, uint32(i)) {
			return i
		}
	}

	return -1
}

func IndicesOfClearBitsUInt32(a uint32) []int {
	result := *new([]int)
	for i := 0; i < 32; i++ {
		if IsClearUInt32(a, uint32(i)) {
			result = append(result, i)
		}
	}
	return result
}

func IndicesOfSetBitsUInt32(a uint32) []int {
	result := *new([]int)

	for i := 0; i < 32; i++ {
		if IsSetUInt32(a, uint32(i)) {
			result = append(result, i)
		}
	}
	return result
}

func IsClearUInt32(a, k uint32) bool {
	return !IsSetUInt32(a, k)
}

func IsSetUInt32(a, k uint32) bool {
	return (a & (1 << k)) != 0
}

func IsPowerOfTwo(a uint32) bool {
	if a == 0 {
		return false
	}
	return (a & (a - 1)) == 0
}

func LowestClearBitIndexUInt32(a uint32) int {
	for i := 0; i < 32; i++ {
		if IsClearUInt32(a, uint32(i)) {
			return i
		}
	}

	return -1
}

func LowestSetBitIndexUInt32(a uint32) int {
	for i := 0; i < 32; i++ {
		if IsSetUInt32(a, uint32(i)) {
			return i
		}
	}

	return -1
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

func SetBitCountUInt32(a uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if IsSetUInt32(a, uint32(i)) {
			count++
		}
	}

	return count
}

func ToggleBitUInt32(a, k uint32) uint32 {
	return a ^ uint32(1)<<k
}

//...
