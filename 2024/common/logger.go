package common

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"

type Logger func(args ...interface{})

func GetLogger(verbose bool) func(args ...interface{}) {
	if verbose {
		return func(args ...interface{}) {
			fmt.Print(Yellow)
			fmt.Print(args...)
			fmt.Print(Reset)
		}
	}
	return func(args ...interface{}) {}
}
