package str

import (
	"math/rand"
	"strings"
)

func generateRandomText(count int, source []rune) string {
	builder := strings.Builder{}

	for i := 0; i < count; i++ {
		builder.WriteRune(source[rand.Intn(len(source))])
	}

	return builder.String()
}

func GenerateRandomText(count int, source string) string {
	return generateRandomText(count, []rune(source))
}

func GenerateRandomTextEN(count int) string {
	return GenerateRandomText(count, "abcdefghijklmnopqrstuVwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func GenerateRandomTextTR(count int) string {
	return generateRandomText(count, []rune("abcçdefgğhıijklmnoöprsştuüvyzABCÇDEFGĞHIİJKLMNOÖPRSŞTUÜVYZ"))
}
