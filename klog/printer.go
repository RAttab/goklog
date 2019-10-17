// Copyright (c) 2014 Datacratic. All rights reserved.

package klog

import (
	"fmt"
	"log"
)

// Printer represents a stage in the printing pipeline.
type Printer interface {
	Print(*Line)
}

// PrinterFunc implements the Printer interface for functions.
type PrinterFunc func(*Line)

// Print passes the line to the function.
func (fn PrinterFunc) Print(line *Line) { fn(line) }

// NilPrinter is a noop printer.
// nolint `NilPrinter` is a global variable (gochecknoglobals)
var NilPrinter = PrinterFunc(func(line *Line) {})

// DefaultPrinter is the default printer used by the kprint functions in klog.
// nolint `DefaultPrinter` is a global variable (gochecknoglobals)
var DefaultPrinter = PrinterFunc(LogPrinter)

// DefaultFatalPrinter is the default printer used by the kfatal and kpanic functions
// in klog.
// nolint `DefaultFatalPrinter` is a global variable (gochecknoglobals)
var DefaultFatalPrinter = PrinterFunc(LogPrinter)

// LogPrinter is forwards all lines to the golang standard log library.
func LogPrinter(line *Line) { log.Printf("<%s> %s", line.Key, line.Value) }

// Keyf is a utility formatting functions for key and is a light wrapper around
// fmt.Sprintf.
func Keyf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
