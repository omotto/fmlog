package fmlog

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
)

type MyError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *MyError) Error() string {
	msg, _ := json.Marshal(e)
	return string(msg)
}

func NewError(code int, msg string) *MyError {
	return &MyError{
		Code: code,
		Msg:  msg,
	}
}

func TestFMLog(t *testing.T) {
	log := NewLogger(os.Stdout)
	log.Log(Warning, errors.New("warning message"))
	log.Log(Error, fmt.Errorf("error message"))
	log.Log(Info, NewError(123, "info message"))
}
