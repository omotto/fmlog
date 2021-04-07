[![GoDoc](http://godoc.org/github.com/omotto/fmlog?status.png)](http://godoc.org/github.com/omotto/fmlog)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/omotto/fmlog)](https://pkg.go.dev/github.com/omotto/fmlog)
[![Build Status](https://travis-ci.com/omotto/fmlog.svg?branch=master)](https://travis-ci.com/omotto/fmlog)
[![Coverage Status](https://coveralls.io/repos/github/omotto/fmlog/badge.svg)](https://coveralls.io/github/omotto/fmlog)
[![Go Report Card](https://goreportcard.com/badge/github.com/omotto/fmlog)](https://goreportcard.com/report/github.com/omotto/fmlog)

# fmlog

Package fmlog implements basic JSON output format logger

### Installation

To download the specific tagged release, run:

```
go get github.com/omotto/fmlog
```

Import it in your program as:

```
import "github.com/omotto/fmlog"
```

### Usage

```

type myError struct {
	Code 	int			`json:"code"`
	Msg		string		`json:"msg"`
}

func (e *myError) Error() string {
	msg, _ := json.Marshal(e)
	return string(msg)
}

func NewError(code int, msg string) *myError {
	return &myError{
		Code: code,
		Msg: msg,
	}
}

....

    log := fmlog.NewLogger(os.Stdout)
    log.Log(fmlog.Info, NewError(123, "info message"))


```