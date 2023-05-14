package clog

import (
    "fmt"
    "github.com/d3code/clog/pkg"
    "strings"
)

func Info(input ...string) {
    message := strings.Join(input, " ")
    message = pkg.ColorMatchTemplate(message)

    fmt.Println(message)
}

func InfoF(format string, input ...any) {
    message := fmt.Sprintf(format, input...)
    message = pkg.ColorMatchTemplate(message)

    fmt.Println(message)
}

func InfoL(inputLines ...string) {
    message := strings.Join(inputLines, "\n")
    message = pkg.ColorMatchTemplate(message)

    fmt.Println(message)
}

func Underline(message ...string) {
    title := strings.Join(message, " ")
    title = pkg.ColorMatchTemplate(title)

    underline := pkg.RemoveColor(title)
    underline = strings.Repeat("-", len(underline))

    fmt.Println()
    InfoL(title, underline)
}

func UnderlineF(format string, input ...any) {
    title := fmt.Sprintf(format, input...)
    title = pkg.ColorMatchTemplate(title)

    underline := pkg.RemoveColor(title)
    underline = strings.Repeat("-", len(underline))

    fmt.Println()
    InfoL(title, underline)
}
