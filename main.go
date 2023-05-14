package clog

import (
    "fmt"
    "github.com/d3code/clog/color"
    "strings"
)

func Debug(title string, input ...string) {
    message := strings.Join(input, " ")
    message = color.RemoveColor(message)
    message = fmt.Sprintf("{{ [ %s ] %s | grey }}", title, message)

    Info(message)
}

func Warn(input ...string) {
    message := strings.Join(input, " ")
    message = color.RemoveColor(message)
    message = fmt.Sprintf("{{ [ WARNING ] | yellow }} %s", message)

    Info(message)
}

func Error(input ...string) {
    message := strings.Join(input, " ")
    message = color.RemoveColor(message)
    message = fmt.Sprintf("{{ [ ERROR ] | red }} %s", message)

    Info(message)
}

func Info(input ...string) {
    message := strings.Join(input, " ")
    message = color.ColorMatchTemplate(message)

    fmt.Println(message)
}

func InfoF(format string, input ...any) {
    message := fmt.Sprintf(format, input...)
    message = color.ColorMatchTemplate(message)

    fmt.Println(message)
}

func InfoL(inputLines ...string) {
    message := strings.Join(inputLines, "\n")
    message = color.ColorMatchTemplate(message)

    fmt.Println(message)
}

func Underline(message ...string) {
    title := strings.Join(message, " ")
    title = color.ColorMatchTemplate(title)

    underline := color.RemoveColor(title)
    underline = strings.Repeat("-", len(underline))

    fmt.Println()
    InfoL(title, underline)
}

func UnderlineF(format string, input ...any) {
    title := fmt.Sprintf(format, input...)
    title = color.ColorMatchTemplate(title)

    underline := color.RemoveColor(title)
    underline = strings.Repeat("-", len(underline))

    fmt.Println()
    InfoL(title, underline)
}
