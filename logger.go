/*
 * Copyright (c) 2024. Altafino Ltd
 * Content:
 * Comment:
 */

package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
)

type Level int

const (
	InfoLevel     Level = 0
	HttpLevel     Level = 1 // Special Info Level for eg HTTP
	ErrorLevel    Level = 2
	CriticalLevel Level = 3
	DebugLevel    Level = 4
	OnlyLevel     Level = 5
	Disabled      Level = 6
)

func (level Level) String() string {
	// declare an array of string
	names := [...]string{
		"INFO    ",
		"HTTP    ",
		"ERROR   ",
		"CRITICAL",
		"DEBUG   ",
		"ONLY",
		"DISABLED",
	}

	if level < InfoLevel || level > Disabled {
		return "UndefinedLevel"
	}
	return names[level]
}

type Output int

const (
	Terminal Output = 0
	Json     Output = 1
)

type TerminalStyle int

const (
	FlatStyle TerminalStyle = 0
	JsonStyle TerminalStyle = 1
)

type Settings struct {
	Level         Level
	Output        Output
	TerminalStyle TerminalStyle
}

var LoggerSettings Settings

func InitLogger(settings Settings) {

	loggerSettings := Settings{
		Level:         settings.Level,
		Output:        settings.Output,
		TerminalStyle: settings.TerminalStyle,
	}
	LoggerSettings = loggerSettings

}

func Only(interf ...interface{}) {
	if checkLevel(OnlyLevel) {
		printlog(OnlyLevel, interf)
	}
}

func Info(interf ...interface{}) {
	if checkLevel(InfoLevel) {
		printlog(InfoLevel, interf)
	}
}

func Http(interf ...interface{}) {
	if checkLevel(HttpLevel) {
		printlog(HttpLevel, interf)
	}
}

func Debug(interf ...interface{}) {
	if checkLevel(DebugLevel) {
		printlog(DebugLevel, interf)
	}
}

func Error(interf ...interface{}) {

		printlog(ErrorLevel, interf)

}

func Critical(interf ...interface{}) {

		printlog(CriticalLevel, interf)

}

func printlog(level Level, interf ...interface{}) {

	if LoggerSettings.Level == Disabled {
		return
	}

	fmt.Print(getTime())

	switch level {
	case InfoLevel:
		color.Set(color.FgGreen)
	case HttpLevel:
		color.Set(color.FgHiMagenta)
	case DebugLevel:
		color.Set(color.FgBlue)
	case ErrorLevel:
		color.Set(color.FgRed)
	case OnlyLevel:
		color.Set(color.BgHiGreen)
		color.Set(color.FgBlack)
	case CriticalLevel:
		color.Set(color.BgHiYellow)
		color.Set(color.FgRed)
	default:
		color.Set(color.BgBlue)
	}

	fmt.Print(" [", level.String(), "]")
	color.Unset()
	for _, v := range interf {
		if LoggerSettings.TerminalStyle == FlatStyle {
			fmt.Print(" ", v)
		} else {
			s, err := prettyjson.Marshal(v)
			if err != nil {
				fmt.Print(" ", v)
			} else {
				fmt.Println(string(s))
			}
		}
	}
	fmt.Println()
}

func getTime() string {
	return time.Now().Format(time.RFC822)
}

func checkLevel(level Level) bool {
	if LoggerSettings.Level >= level {
		return true
	}
	if level == CriticalLevel || level == ErrorLevel || level==OnlyLevel {
		return true
	}
	return false
}
