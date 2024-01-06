package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var g_kb = bufio.NewReader(os.Stdin)

func ReadInt(prompt, errorPrompt string) int {
	for {
		fmt.Print(prompt)
		str, err := g_kb.ReadString('\n')

		if err != nil {
			fmt.Println(errorPrompt)
			continue
		}

		str = strings.TrimRight(str, "\r\n")

		result, e := strconv.Atoi(str)

		if e == nil {
			return result
		}
		fmt.Print(errorPrompt)
	}
}

func ReadFloat64(prompt, errorPrompt string) float64 {

	for {
		fmt.Print(prompt)
		str, err := g_kb.ReadString('\n')

		if err != nil {
			fmt.Println(errorPrompt)
			continue
		}

		str = strings.TrimRight(str, "\r\n")
		result, err := strconv.ParseFloat(str, 64)

		if err == nil {
			return result
		}
		fmt.Print(errorPrompt)
	}
}

func ReadString(prompt string) string {
	fmt.Print(prompt)
	str, _ := g_kb.ReadString('\n')

	return strings.TrimRight(str, "\r\n")
}

func Write(fmtStr string, a ...any) error {
	_, err := fmt.Printf(fmtStr, a)

	return err
}

func WriteLine(fmtStr string, a ...any) error {
	return Write(fmtStr+"\n", a)
}

func WriteErr(fmtStr string, a ...any) error {
	_, err := fmt.Fprintf(os.Stderr, fmtStr, a)

	return err
}

func WriteErrLine(fmtStr string, a ...any) error {
	return Write(fmtStr+"\n", a)
}

//...
