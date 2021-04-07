/*

Package fmlog implements JSON output format logger

Installation

To download the specific tagged release, run:

	go get github.com/omotto/fmlog

Import it in your program as:

	import "github.com/omotto/fmlog"

Usage

	log := fmlog.NewLogger(os.Stdout)
	// Add new log
	log.Log(fmlog.Warning, errors.New("warning message"))

*/
package fmlog
