package main

import (
	"RandomNumberGenerators/csd/console"
	"RandomNumberGenerators/csd/err"
	"RandomNumberGenerators/jsondata"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func writeToFile(info *jsondata.Info, i int, wg *sync.WaitGroup) {
	minVal := info.Min
	maxVal := info.Max
	count := info.Count
	path := fmt.Sprintf("%s-%d", info.BasePath, i+1)
	f, e := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)

	if e != nil {
		err.ExitFailureError("OpenFile", e)
	}

	defer func() { _ = f.Close() }()

	for i := 0; i < count; i++ {
		var val int32 = int32(rand.Intn(maxVal-minVal+1) + minVal)

		e = binary.Write(f, binary.LittleEndian, &val)
		if e != nil {
			err.ExitFailureError("Write", e)
		}
	}

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
