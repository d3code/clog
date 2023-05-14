package clog

import "testing"

func TestInfo(t *testing.T) {
    type args struct {
        input []string
    }
    tests := []struct {
        name string
        args args
    }{
        // TODO: Add test cases.
        {name: "Standard", args: struct{ input []string }{input: []string{"Hello", "{{ World | blue}}!"}}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            Info(tt.args.input...)
        })
    }
}

func TestInfoF(t *testing.T) {
    type args struct {
        format string
        input  []any
    }
    tests := []struct {
        name string
        args args
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            InfoF(tt.args.format, tt.args.input...)
        })
    }
}

func TestInfoL(t *testing.T) {
    type args struct {
        inputLines []string
    }
    tests := []struct {
        name string
        args args
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            InfoL(tt.args.inputLines...)
        })
    }
}

func TestUnderline(t *testing.T) {
    type args struct {
        message []string
    }
    tests := []struct {
        name string
        args args
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            Underline(tt.args.message...)
        })
    }
}

func TestUnderlineF(t *testing.T) {
    type args struct {
        format string
        input  []any
    }
    tests := []struct {
        name string
        args args
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            UnderlineF(tt.args.format, tt.args.input...)
        })
    }
}
