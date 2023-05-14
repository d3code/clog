package color

import "testing"

func TestColorMatchTemplate(t *testing.T) {
    type args struct {
        message string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ColorMatchTemplate(tt.args.message); got != tt.want {
                t.Errorf("ColorMatchTemplate() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestColorString(t *testing.T) {
    type args struct {
        text  string
        color string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ColorString(tt.args.text, tt.args.color); got != tt.want {
                t.Errorf("ColorString() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRemoveColor(t *testing.T) {
    type args struct {
        message string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := RemoveColor(tt.args.message); got != tt.want {
                t.Errorf("RemoveColor() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_matchColor(t *testing.T) {
    type args struct {
        color string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := matchColor(tt.args.color); got != tt.want {
                t.Errorf("matchColor() = %v, want %v", got, tt.want)
            }
        })
    }
}
