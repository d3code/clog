package clog

import (
    "bytes"
    "github.com/d3code/clog/color"
    "io"
    "os"
    "sync"
    "testing"
)

const format = "\n" +
    color.COLOR_REDBG + "got" + color.COLOR_END + "\n" +
    "[%v]\n" +
    color.COLOR_RED + "want" + color.COLOR_END + "\n" +
    "[%v]\n\n"

func TestInfo(t *testing.T) {
    type args struct {
        input []string
    }
    tests := []struct {
        name       string
        args       args
        wantStdout string
        wantStderr string
    }{
        {name: "Standard", args: args{input: []string{"Hello World!"}}, wantStdout: "Hello World!\n"},
        {name: "Join", args: args{input: []string{"Hello", "World!"}}, wantStdout: "Hello World!\n"},
        {name: "Color", args: args{input: []string{"Hello {{ World | blue}}!"}}, wantStdout: "Hello \x1b[34mWorld\x1b[0m!\n"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            out, err := captureOutput(func() {
                Info(tt.args.input...)
            })

            if out != tt.wantStdout {
                t.Errorf(format, out, tt.wantStdout)
            }
            if err != tt.wantStderr {
                t.Errorf(format, err, tt.wantStderr)
            }
        })
    }
}

func captureOutput(function func()) (string, string) {
    readerStdOut, writerStdOut, errOut := os.Pipe()
    readerStdErr, writerStdErr, errErr := os.Pipe()
    if errOut != nil || errErr != nil {
        panic("Could not set up pipe")
    }

    stdout := os.Stdout
    stderr := os.Stderr
    defer func() {
        os.Stdout = stdout
        os.Stderr = stderr
    }()

    os.Stdout = writerStdOut
    os.Stderr = writerStdErr

    out := make(chan string)
    outErr := make(chan string)

    wg := new(sync.WaitGroup)
    wg.Add(1)
    go func() {
        var buf bytes.Buffer
        var bufErr bytes.Buffer
        wg.Done()

        io.Copy(&buf, readerStdOut)
        io.Copy(&bufErr, readerStdErr)

        out <- buf.String()
        outErr <- bufErr.String()
    }()
    wg.Wait()

    function()
    writerStdOut.Close()
    writerStdErr.Close()

    return <-out, <-outErr
}
