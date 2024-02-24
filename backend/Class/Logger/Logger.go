package Logger

import "time"

var introducer = " > "
var separator = " | "

func GetCurrentDatetime() string {
	currentDatetime := time.Now()
	return currentDatetime.Format("15:04:05 02-01-2006")
}

func Write(message string) {
	println(message)
}

func Debug(message string) {
	println(GetCurrentDatetime() + introducer + "DEBUG" + separator + message)
}

func Info(message string) {
	println(GetCurrentDatetime() + introducer + "INFO" + separator + message)
}

func Success(message string) {
	println(GetCurrentDatetime() + introducer + "SUCCESS" + separator + message)
}

func Warning(message string) {
	println(GetCurrentDatetime() + introducer + "WARNING" + separator + message)
}

func Error(message string) {
	println(GetCurrentDatetime() + introducer + "ERROR" + separator + message)
}

func Raise(err string) {
	println(err)
}

func Separator() {
	println("----------------------------------")
}
