package nowcoder

import (
    "reflect"
    "testing"
)

func TestReadVLRTree(t *testing.T) {
    type args struct {
        tree *TreeNode
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "",
            args: args{
                tree: &TreeNode{
                    Val: 0,
                    Left: &TreeNode{
                        Val:   1,
                        Left:  nil,
                        Right: nil,
                    },
                    Right: &TreeNode{
                        Val:   2,
                        Left:  nil,
                        Right: nil,
                    },
                },
            },
            want: []int{0, 1, 2},
        },
        {
            name: "",
            args: args{
                tree: &TreeNode{
                    Val: 0,
                    Left: &TreeNode{
                        Val: 1,
                        Left: &TreeNode{
                            Val:   3,
                            Left:  nil,
                            Right: nil,
                        },
                        Right: &TreeNode{
                            Val:   4,
                            Left:  nil,
                            Right: nil,
                        },
                    },
                    Right: &TreeNode{
                        Val: 2,
                        Left: &TreeNode{
                            Val:   5,
                            Left:  nil,
                            Right: nil,
                        },
                        Right: &TreeNode{
                            Val:   6,
                            Left:  nil,
                            Right: nil,
                        },
                    },
                },
            },
            want: []int{0, 1, 3, 4, 2, 5, 6},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ReadVLRTree(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ReadVLRTree() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestReadLDRTree(t *testing.T) {
    type args struct {
        tree *TreeNode
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "",
            args: args{
                tree: &TreeNode{
                    Val: 0,
                    Left: &TreeNode{
                        Val:   1,
                        Left:  nil,
                        Right: nil,
                    },
                    Right: &TreeNode{
                        Val:   2,
                        Left:  nil,
                        Right: nil,
                    },
                },
            },
            want: []int{1, 0, 2},
        },
        {
            name: "",
            args: args{
                tree: &TreeNode{
                    Val: 0,
                    Left: &TreeNode{
                        Val: 1,
                        Left: &TreeNode{
                            Val:   3,
                            Left:  nil,
                            Right: nil,
                        },
                        Right: &TreeNode{
                            Val:   4,
                            Left:  nil,
                            Right: nil,
                        },
                    },
                    Right: &TreeNode{
                        Val: 2,
                        Left: &TreeNode{
                            Val:   5,
                            Left:  nil,
                            Right: nil,
                        },
                        Right: &TreeNode{
                            Val:   6,
                            Left:  nil,
                            Right: nil,
                        },
                    },
                },
            },
            want: []int{3, 1, 4, 0, 5, 2, 6},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ReadLDRTree(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ReadLDRTree() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestReadLRDTree(t *testing.T) {
    type args struct {
        tree *TreeNode
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "",
            args: args{
                tree: &TreeNode{
                    Val: 0,
                    Left: &TreeNode{
                        Val:   1,
                        Left:  nil,
                        Right: nil,
                    },
                    Right: &TreeNode{
                        Val:   2,
                        Left:  nil,
                        Right: nil,
                    },
                },
            },
            want: []int{1, 2, 0},
        },
        {
            name: "",
            args: args{
                tree: &TreeNode{
                    Val: 0,
                    Left: &TreeNode{
                        Val: 1,
                        Left: &TreeNode{
                            Val:   3,
                            Left:  nil,
                            Right: nil,
                        },
                        Right: &TreeNode{
                            Val:   4,
                            Left:  nil,
                            Right: nil,
                        },
                    },
                    Right: &TreeNode{
                        Val: 2,
                        Left: &TreeNode{
                            Val:   5,
                            Left:  nil,
                            Right: nil,
                        },
                        Right: &TreeNode{
                            Val:   6,
                            Left:  nil,
                            Right: nil,
                        },
                    },
                },
            },
            want: []int{3, 4, 1, 5, 6, 2, 0},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ReadLRDTree(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ReadLRDTree() = %v, want %v", got, tt.want)
            }
        })
    }
}
