package main

import "fmt"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
)

func printRed(msg string)   { fmt.Println(colorRed + msg + colorReset) }
func printGreen(msg string) { fmt.Println(colorGreen + msg + colorReset) }
func printYellow(msg string){ fmt.Println(colorYellow + msg + colorReset) }
func printBlue(msg string)  { fmt.Println(colorBlue + msg + colorReset) }
func printCyan(msg string)  { fmt.Println(colorCyan + msg + colorReset) }
func printGray(msg string)  { fmt.Println(colorGray + msg + colorReset) }
func printBold(msg string)  { fmt.Println("\033[1m" + msg + colorReset) }
