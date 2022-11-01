package io

import (
	"github.com/symfony-cli/terminal"
	"io"
)

type Level int8

const (
	VerboseLevel = 2
	DebugLevel   = 4
)

func Write(format string, a ...interface{}) {
	_, _ = terminal.Printfln(format, a...)
}

func Verbose(format string, a ...interface{}) {
	if terminal.GetLogLevel() < VerboseLevel {
		return
	}

	_, _ = terminal.Printfln(format, a...)
}

func Debug(format string, a ...interface{}) {
	if terminal.GetLogLevel() < DebugLevel {
		return
	}

	_, _ = terminal.Printfln(format, a...)
}

func Out() io.Writer {
	return terminal.Stdout
}

func Format(msg string) string {
	return terminal.Format(msg)
}
