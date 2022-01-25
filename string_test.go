package nowcoder

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
                x: "01",
                y: "1",
            },
            want: "02",
        },
        {
            name: "",
            args: args{
                x: "9",
                y: "9",
            },
            want: "18",
        },
        {
            name: "",
            args: args{
                x: "09",
                y: "9",
            },
            want: "18",
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
                x: "009",
                y: "11",
            },
            want: "020",
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
            if got := AddStringNumber(tt.args.x, tt.args.y); got != tt.want {
                t.Errorf("StringNumberAdd() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestStringReverse(t *testing.T) {
    type args struct {
        v string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "",
            args: args{
                v: "12",
            },
            want: "21",
        },
        {
            name: "",
            args: args{
                v: "123",
            },
            want: "321",
        },
        {
            name: "",
            args: args{
                v: "!@#$%^&*()_+QWERTYUIOPASDFGHJKLZXCVBNM<>?",
            },
            want: "?><MNBVCXZLKJHGFDSAPOIUYTREWQ+_)(*&^%$#@!",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ReverseString(tt.args.v); got != tt.want {
                t.Errorf("StringReverse() = %v, want %v", got, tt.want)
            }
        })
    }
}