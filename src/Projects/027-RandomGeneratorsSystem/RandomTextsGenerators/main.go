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
	"io"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

func infoClient(server string) (int, string) {
	req, e := http.NewRequest("GET", server+"/info/random", nil)
	if e != nil {
		fmt.Printf("Error in request:%s", e.Error())
		return http.StatusInternalServerError, ""
	}

	client := http.Client{Timeout: 20 * time.Second}
	res, e := client.Do(req)

	if e != nil {
		fmt.Printf("Client error:%s", e.Error())
		return http.StatusInternalServerError, ""
	}

	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return res.StatusCode, ""
	}

	data, e := io.ReadAll(res.Body)

	if e != nil {
		fmt.Printf("Data error:%s", e.Error())
		return http.StatusInternalServerError, ""
	}

	return res.StatusCode, string(data)
}

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

	_ = writer.Flush()
	wg.Done()
}

func generateTexts(info *jsondata.Info) {
	c := info.Count
	var wg sync.WaitGroup

	for i := 0; i < c; i++ {
		wg.Add(1)
		go writeToFile(info, i, &wg)
	}

	wg.Wait()
}

func checkArguments() {
	if len(os.Args) != 2 {
		fmt.Printf("Wrong number of arguments!...")
		os.Exit(1)
	}
}

func main() {
	checkArguments()
	server := os.Args[1]
	fmt.Printf("Server:%s\n", server)
	info := jsondata.Info{}

	for {
		_ = console.ReadString("Push enter to generate:")
		status, jsonStr := infoClient(server)
		if status == http.StatusOK {
			e := json.Unmarshal([]byte(jsonStr), &info)
			if e == nil {
				generateTexts(&info)
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "internal problem occurred:%s\n", e.Error())
			}
		} else {
			fmt.Printf("Status Code:%d\n", status)
		}
	}
}
