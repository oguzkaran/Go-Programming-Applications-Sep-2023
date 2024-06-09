/*
------------------------------------------------------------------------------------------------------------------------
    Log'lama işlemi (logging) için pek çok standart olmayan paket olsa da standart log paketi de kullanılabilir. Log'lama
	işlemi uygulama düzeyinde pek çok durumda kullanılabilmektedir. Aşağıdaki demor örneği inceleyiniz
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"errors"
	"io"
	"log"
	"math/rand/v2"
	"os"
)

var (
	TraceLog   *log.Logger
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func initLoggers(traceWriter io.Writer, infoWriter io.Writer, warningWriter io.Writer, errorWriter io.Writer) {
	flag := log.Ldate | log.Ltime | log.Lshortfile

	TraceLog = log.New(traceWriter, "TRACE:", flag)
	InfoLog = log.New(infoWriter, "INFO:", flag)
	WarningLog = log.New(warningWriter, "WARN:", flag)
	ErrorLog = log.New(errorWriter, "ERR:", flag)
}

func main() {
	initLoggers(io.Discard, os.Stdout, os.Stdout, os.Stderr)
	TraceLog.Println("Application started")
	generateNumbers(10)
	e := errors.New("any error")
	ErrorLog.Println(e.Error())
	TraceLog.Println("Application ended")
}

func generateNumbers(n int) {
	InfoLog.Print("Number generation started:")
	sum := 0

	for i := 0; i < n; i++ {
		val := rand.IntN(100)
		sum += val
	}

	WarningLog.Println("Sum calculated bu not used yet")
	InfoLog.Print("Number generation ended:")
}
