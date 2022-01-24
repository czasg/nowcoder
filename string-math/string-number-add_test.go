package stringmath

import "testing"

func TestStringNumberAdd(t *testing.T) {
    type args struct {
        x string
        y string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "",
            args: args{
                x: "1",
                y: "1",
            },
            want: "2",
        },
        {
            name: "",
            args: args{
                x: "001",
                y: "01",
            },
            want: "002",
        },
        {
            name: "",
            args: args{
                x: "123456789",
                y: "123456789123456789123456789123456789123456789123456789123456789123456789",
            },
            want: "123456789123456789123456789123456789123456789123456789123456789246913578",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := StringNumberAdd(tt.args.x, tt.args.y); got != tt.want {
                t.Errorf("StringNumberAdd() = %v, want %v", got, tt.want)
            }
        })
    }
}
