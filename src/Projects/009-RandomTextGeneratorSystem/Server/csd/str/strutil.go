package str

import (
	"math/rand"
	"strings"
)

func GenerateRandomText(count int, source string) string {
	builder := strings.Builder{}

	for i := 0; i < count; i++ {
		builder.WriteRune(rune(source[rand.Intn(len(source))]))
	}

	return builder.String()
}

func GenerateRandomTextEN(count int) string {
	return GenerateRandomText(count, "abcdefghijklmnopqrstuVwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func GenerateRandomTextTR(count int) string {
	return GenerateRandomText(count, "abcdefghijklmnopqrstuVwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
