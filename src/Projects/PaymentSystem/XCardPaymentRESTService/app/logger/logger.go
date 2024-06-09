package logger

import (
	"io"
	"log"
)

var (
	TraceLog   *log.Logger
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func InitLoggers(traceWriter io.Writer, infoWriter io.Writer, warningWriter io.Writer, errorWriter io.Writer) {
	flag := log.Ldate | log.Ltime | log.Lshortfile

	TraceLog = log.New(traceWriter, "TRACE:", flag)
	InfoLog = log.New(infoWriter, "INFO:", flag)
	WarningLog = log.New(warningWriter, "WARN:", flag)
	ErrorLog = log.New(errorWriter, "ERR:", flag)
}
