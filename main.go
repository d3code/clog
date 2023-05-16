package clog

import (
    "fmt"
    "github.com/d3code/clog/color"
    "strings"
)

var ShowDebugLogs = false

func Debug(input ...string) {
    if !ShowDebugLogs {
        return
    }

    message := strings.Join(input, " ")
    if len(message) == 0 {
        return
    }

    message = color.RemoveColor(message)
    message = color.ColorString(message, "grey")

    Info(message)
}

func Warn(input ...string) {
    message := strings.Join(input, " ")
    if len(message) == 0 {
        message = "Unknown warning"
    }

    message = color.RemoveColor(message)
    message = color.ColorString("WARNING: "+message, "yellow")

    Info(message)
}

func Error(input ...string) {
    message := strings.Join(input, " ")
    if len(message) == 0 {
        message = "Unknown error"
    }

    message = color.RemoveColor(message)
    message = color.ColorString("ERROR: "+message, "red")

    Info(message)
}

func Info(input ...string) {
    message := strings.Join(input, " ")
    if len(message) == 0 {
        return
    }

    message = color.ColorMatchTemplate(message)
    fmt.Println(message)
}

// InfoF uses fmt.Sprintf to log a templated message.
func InfoF(format string, input ...any) {
    message := fmt.Sprintf(format, input...)
    if len(message) == 0 {
        return
    }

    message = color.ColorMatchTemplate(message)
    fmt.Println(message)
}

func InfoL(inputLines ...string) {
    message := strings.Join(inputLines, "\n")
    if len(message) == 0 {
        return
    }

    message = color.ColorMatchTemplate(message)
    fmt.Println(message)
}

func Underline(title string, message string) {
    if message != "" {
        title = fmt.Sprintf("{{ %s | bold }} {{ %s | blue }}", title, message)
    }
    title = color.ColorMatchTemplate(title)

    underline := color.RemoveColor(title)
    underline = strings.Repeat("-", len(underline))
    underline = color.ColorString(underline, "grey")

    fmt.Println()
    InfoL(title, underline)
}

func UnderlineF(format string, input ...any) {
    title := fmt.Sprintf(format, input...)
    Underline(title, "")
}
