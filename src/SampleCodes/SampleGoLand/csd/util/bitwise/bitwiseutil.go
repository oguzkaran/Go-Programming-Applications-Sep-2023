package bitwise

import (
	"fmt"
)

func BitStr(a int, size int) string {
	return fmt.Sprintf(fmt.Sprintf("%%0%db", size), a)
}

func BitStrInt32(a int32) string {
	return BitStr(int(a), 32)
}

func BitStrUInt32(a uint32) string {
	return BitStr(int(a), 32)
}

func BitStrInt64(a int64) string {
	return BitStr(int(a), 64)
}

func BitStrUInt64(a uint64) string {
	return BitStr(int(a), 64)
}
