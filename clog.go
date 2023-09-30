package clog

import (
    "fmt"
    "strings"

    "github.com/d3code/clog/color"
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

    message = color.RemoveTemplate(message)
    message = color.RemoveColor(message)
    message = color.String(message, "grey")

    Info(message)
}

func Debugf(format string, input ...any) {
    message := fmt.Sprintf(format, input...)
    Debug(message)
}

func Warn(input ...string) {
    message := strings.Join(input, " ")
    if len(message) == 0 {
        message = "Unknown warning"
    }

    message = color.RemoveTemplate(message)
    message = color.RemoveColor(message)
    message = color.String(message, "yellow")

    Info(message)
}

func Warnf(format string, input ...any) {
    message := fmt.Sprintf(format, input...)
    Warn(message)
}

func Error(input ...string) {
    message := strings.Join(input, " ")
    if len(message) == 0 {
        message = "Unknown error"
    }

    message = color.RemoveTemplate(message)
    message = color.RemoveColor(message)
    message = color.String(message, "red")

    Info(message)
}

func Errorf(format string, input ...any) {
    message := fmt.Sprintf(format, input...)
    Error(message)
}

func Info(input ...string) {
    message := strings.Join(input, " ")
    if len(message) == 0 {
        return
    }

    message = color.Template(message)
    fmt.Println(message)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(format string, input ...any) {
    message := fmt.Sprintf(format, input...)
    Info(message)
}

// Infol logs a templated message with a newline.
func Infol(inputLines ...string) {
    message := strings.Join(inputLines, "\n")
    Info(message)
}

// Underline logs a templated message with an underline.
func Underline(title string, message string) {
    if message != "" {
        title = fmt.Sprintf("{{ %s | bold }} {{ %s | blue }}", title, message)
    }
    title = color.Template(title)

    underline := color.RemoveColor(title)
    underline = strings.Repeat("-", len(underline))
    underline = color.String(underline, "grey")

    fmt.Println()
    Infol(title, underline)
}

func UnderlineF(format string, input ...any) {
    title := fmt.Sprintf(format, input...)
    Underline(title, "")
}
