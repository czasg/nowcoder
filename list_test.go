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
                Val: 3,
                Next: &ListNode{
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
            ReorderList(tt.args.head)
            if !reflect.DeepEqual(tt.args.head, tt.want) {
                t.Errorf("ReorderList() = %v, want %v", tt.args.head, tt.want)
            }
        })
    }
}

func TestReadList(t *testing.T) {
    type args struct {
        head *ListNode
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "",
            args: args{
                head: &ListNode{
                    Val:  0,
                    Next: nil,
                },
            },
            want: []int{0},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ReadList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ReadList() = %v, want %v", got, tt.want)
            }
        })
    }
}
