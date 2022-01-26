package nowcoder

import "testing"

func TestGoodStock(t *testing.T) {
    type args struct {
        prices []int
    }
    tests := []struct {
        name       string
        args       args
        wantBuy    int
        wantSell   int
        wantProfit int
    }{
        {
            name: "",
            args: args{
                prices: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            },
            wantBuy:    0,
            wantSell:   9,
            wantProfit: 9,
        },
        {
            name: "",
            args: args{
                prices: []int{8, 9, 2, 5, 4, 7, 1},
            },
            wantBuy:    2,
            wantSell:   5,
            wantProfit: 5,
        },
        {
            name: "",
            args: args{
                prices: []int{2, 4, 1},
            },
            wantBuy:    0,
            wantSell:   1,
            wantProfit: 2,
        },
        {
            name: "",
            args: args{
                prices: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
            },
            wantBuy:    0,
            wantSell:   0,
            wantProfit: 0,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotBuy, gotSell, gotProfit := GoodStock(tt.args.prices)
            if gotBuy != tt.wantBuy {
                t.Errorf("GoodStock() gotBuy = %v, want %v", gotBuy, tt.wantBuy)
            }
            if gotSell != tt.wantSell {
                t.Errorf("GoodStock() gotSell = %v, want %v", gotSell, tt.wantSell)
            }
            if gotProfit != tt.wantProfit {
                t.Errorf("GoodStock() gotProfit = %v, want %v", gotProfit, tt.wantProfit)
            }
        })
    }
}
