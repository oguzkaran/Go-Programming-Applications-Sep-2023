// {"n":10000, "min":10, "max":20, "count":30, "basePath":"texts/text.txt"}
package main

import (
	"RandomTextsGenerators/csd/console"
	"RandomTextsGenerators/csd/err"
	"RandomTextsGenerators/csd/util/str"
	"RandomTextsGenerators/jsondata"
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func writeToFile(info *jsondata.Info, i int, wg *sync.WaitGroup) {
	minVal := info.Min
	maxVal := info.Max
	n := info.N
	path := fmt.Sprintf("%s-%d", info.BasePath, i+1)
	f, e := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)

	if e != nil {
		err.ExitFailureError("OpenFile", e)
	}

	defer func() { _ = f.Close() }()

	writer := bufio.NewWriter(f)

	for i := 0; i < n; i++ {
		text := str.GenerateRandomTextEN(rand.Intn(maxVal-minVal+1) + minVal)
		_, e = writer.WriteString(fmt.Sprintf("%s\r\n", text))
		if e != nil {
			err.ExitFailureError("Write", e)
		}
	}

	writer.Flush()
	wg.Done()
}

func generateNumbers(info *jsondata.Info) {
	c := info.Count
	var wg sync.WaitGroup

	for i := 0; i < c; i++ {
		wg.Add(1)
		go writeToFile(info, i, &wg)
	}

	wg.Wait()
}

func main() {
	info := jsondata.Info{}

	for {
		jsonStr := console.ReadString("Input config as JSON:")

		if jsonStr == "" {
			break
		}

		e := json.Unmarshal([]byte(jsonStr), &info)
		if e == nil {
			generateNumbers(&info)
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "internal problem occurred:%s\n", e.Error())
		}
	}
}
