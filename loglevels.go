package main

import "fmt"

const (
	Warn int = iota
	Debug
	Info
	Error
	Fatal
)

func main() {
	setLevel(Warn)
	setLevel(Debug)
	setLevel(Info)
	setLevel(Error)
	setLevel(Fatal)
	setLevel(7)
}

func setLevel(level int) {
	switch level {
	case Warn:
		{
			fmt.Printf("Warn Level %d\n", level)
		}
	case Debug:
		{
			fmt.Printf("Debug level %d\n", level)
		}
	case Info:
		{
			fmt.Printf("Info level %d\n", level)
		}
	case Error:
		{
			fmt.Printf("Error level %d\n", level)
		}
	case Fatal:
		{
			fmt.Printf("Fatal level %d\n", level)
		}
	default:
		{
			fmt.Printf("Level %v\n", level)
		}

	}
}
