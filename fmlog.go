package fmlog

import (
	"encoding/json"
	"io"
	"runtime"
	"sync"
	"time"
)

// LogType defines log output type as Info, Warning or Error
type LogType string

const(
	// Info message level
	Info 		LogType = "INFO"
	// Warning error level
	Warning 			= "WARNING"
	// Error output log level
	Error 				= "ERROR"
)

// Logger class
type Logger struct {
	mu     	sync.Mutex
	out 	io.Writer
}

// NewLogger constructor
func NewLogger(out io.Writer) *Logger {
	return &Logger{out: out}
}

// Log method of Logger class
func (l *Logger) Log(logType LogType, err error) {
	var logMessage struct {
		Message  	interface{} `json:"message"`
		Type     	LogType 	`json:"type"`
		Function 	string  	`json:"function"`
		FileName 	string  	`json:"fileName"`
		Line     	int     	`json:"line"`
		TimeStamp	string  	`json:"timeStamp"`
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	pc, fn, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()

	var js map[string]interface{}
	if json.Unmarshal([]byte(err.Error()), &js) == nil {
		logMessage.Message = js
	} else {
		logMessage.Message = err.Error()
	}
	logMessage.FileName = fn
	logMessage.Type 	= logType
	logMessage.Line		= line
	logMessage.Function = funcName
	logMessage.TimeStamp= time.Now().UTC().String()

	if msg, e := json.Marshal(&logMessage); e == nil {
		_, _ = io.WriteString(l.out, string(msg))
	}
}


