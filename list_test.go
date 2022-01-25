package nowcoder

import (
    "reflect"
    "testing"
)

func TestMiddleLeftList(t *testing.T) {
    type args struct {
        head *ListNode
    }
    tests := []struct {
        name string
        args args
        want *ListNode
    }{
        {
            name: "",
            args: args{},
            want: nil,
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val:  0,
                    Next: nil,
                },
            },
            want: &ListNode{
                Val:  0,
                Next: nil,
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val:  1,
                        Next: nil,
                    },
                },
            },
            want: &ListNode{
                Val: 0,
                Next: &ListNode{
                    Val:  1,
                    Next: nil,
                },
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val:  2,
                            Next: nil,
                        },
                    },
                },
            },
            want: &ListNode{
                Val: 1,
                Next: &ListNode{
                    Val:  2,
                    Next: nil,
                },
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val: 2,
                            Next: &ListNode{
                                Val:  3,
                                Next: nil,
                            },
                        },
                    },
                },
            },
            want: &ListNode{
                Val: 1,
                Next: &ListNode{
                    Val: 2,
                    Next: &ListNode{
                        Val:  3,
                        Next: nil,
                    },
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := MiddleLeftList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("MiddleLeftList() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestMiddleRightList(t *testing.T) {
    type args struct {
        head *ListNode
    }
    tests := []struct {
        name string
        args args
        want *ListNode
    }{
        {
            name: "",
            args: args{},
            want: nil,
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val:  0,
                    Next: nil,
                },
            },
            want: &ListNode{
                Val:  0,
                Next: nil,
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val:  1,
                        Next: nil,
                    },
                },
            },
            want: &ListNode{
                Val:  1,
                Next: nil,
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val:  2,
                            Next: nil,
                        },
                    },
                },
            },
            want: &ListNode{
                Val: 1,
                Next: &ListNode{
                    Val:  2,
                    Next: nil,
                },
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val: 2,
                            Next: &ListNode{
                                Val:  3,
                                Next: nil,
                            },
                        },
                    },
                },
            },
            want: &ListNode{
                Val: 2,
                Next: &ListNode{
                    Val:  3,
                    Next: nil,
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := MiddleRightList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("MiddleRightList() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestReverseList(t *testing.T) {
    type args struct {
        head *ListNode
    }
    tests := []struct {
        name string
        args args
        want *ListNode
    }{
        {
            name: "",
            args: args{},
            want: nil,
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val:  0,
                    Next: nil,
                },
            },
            want: &ListNode{
                Val:  0,
                Next: nil,
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val:  1,
                        Next: nil,
                    },
                },
            },
            want: &ListNode{
                Val: 1,
                Next: &ListNode{
                    Val:  0,
                    Next: nil,
                },
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val:  2,
                            Next: nil,
                        },
                    },
                },
            },
            want: &ListNode{
                Val: 2,
                Next: &ListNode{
                    Val: 1,
                    Next: &ListNode{
                        Val:  0,
                        Next: nil,
                    },
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ReverseList(tt.args.head); got != nil && !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ReverseList() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestReorderList(t *testing.T) {
    type args struct {
        head *ListNode
    }
    tests := []struct {
        name string
        args args
        want *ListNode
    }{
        {
            name: "",
            args: args{},
            want: nil,
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val:  0,
                    Next: nil,
                },
            },
            want: &ListNode{
                Val:  0,
                Next: nil,
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val:  1,
                        Next: nil,
                    },
                },
            },
            want: &ListNode{
                Val: 0,
                Next: &ListNode{
                    Val:  1,
                    Next: nil,
                },
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val:  2,
                            Next: nil,
                        },
                    },
                },
            },
            want: &ListNode{
                Val: 0,
                Next: &ListNode{
                    Val: 2,
                    Next: &ListNode{
                        Val:  1,
                        Next: nil,
                    },
                },
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val: 2,
                            Next: &ListNode{
                                Val:  3,
                                Next: nil,
                            },
                        },
                    },
                },
            },
            want: &ListNode{
                Val: 0,
                Next: &ListNode{
                    Val: 3,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val:  2,
                            Next: nil,
                        },
                    },
                },
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val: 2,
                            Next: &ListNode{
                                Val: 3,
                                Next: &ListNode{
                                    Val:  4,
                                    Next: nil,
                                },
                            },
                        },
                    },
                },
            },
            want: &ListNode{
                Val: 0,
                Next: &ListNode{
                    Val: 4,
                    Next: &ListNode{
                        Val: 1,
                        Next: &ListNode{
                            Val: 3,
                            Next: &ListNode{
                                Val:  2,
                                Next: nil,
                            },
                        },
                    },
                },
            },
        },
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val: 1,
                    Next: &ListNode{
                        Val: 2,
                        Next: &ListNode{
                            Val: 3,
                            Next: &ListNode{
                                Val:  4,
                                Next: nil,
                            },
                        },
                    },
                },
            },
            want: &ListNode{
                Val: 1,
                Next: &ListNode{
                    Val: 4,
                    Next: &ListNode{
                        Val: 2,
                        Next: &ListNode{
                            Val:  3,
                            Next: nil,
                        },
                    },
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ReorderList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ReorderList() = %v, want %v", got, tt.want)
            }
        })
    }
}
