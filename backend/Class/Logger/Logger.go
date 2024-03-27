package Logger

import (
	"fmt"
	"time"
)

var introducer = " > "
var separator = " | "

func getCurrentDatetime() string {
	currentDatetime := time.Now()
	return currentDatetime.Format("15:04:05 02-01-2006")
}

func Write(args ...any) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func Debug(message string) {
	println(getCurrentDatetime() + introducer + "DEBUG" + separator + message)
}

func Info(message string) {
	println(getCurrentDatetime() + introducer + "INFO" + separator + message)
}

func Success(message string) {
	println(getCurrentDatetime() + introducer + "SUCCESS" + separator + message)
}

func Warning(message string) {
	println(getCurrentDatetime() + introducer + "WARNING" + separator + message)
}

func Error(message string) {
	println(getCurrentDatetime() + introducer + "ERROR" + separator + message)
}

func Raise(err string) {
	println(err)
}

func Separator() {
	println("----------------------------------")
}

// Welcome to Display a welcome message at the start
func Welcome() {
	println("------{ Welcome on Helm Registry }------")
	println("Ahoy!")
	println("A simple Helm chart registry")
	println()
	println("Version :", "1.0.0")
	println()
	println("App is going to start ...")
}
