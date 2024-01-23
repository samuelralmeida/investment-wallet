package entity

/*
import (
	"reflect"
	"testing"
	"time"
)

func TestShares_Funds(t *testing.T) {
	tests := []struct {
		name string
		ii   *Shares
		want *Funds
	}{
		{
			name: "generate correct fund amount and value",
			ii: &Shares{
				{
					Investiment: Investiment{
						ID:       "uuid1",
						Date:     time.Now(),
						Box:      "box-name",
						Category: "category-name",
						Name:     "fund-name",
						Cnpj:     "cnpj",
						Bank:     "bank",
						Amount:   10.0,
						Wallet:   "wallet",
					},
					CheckpointDate:  &[]time.Time{time.Now()}[0],
					CheckpointValue: &[]float64{50.0}[0],
				},
				{
					Investiment: Investiment{
						ID:       "uuid2",
						Date:     time.Now(),
						Box:      "box-name",
						Category: "category-name",
						Name:     "fund-name",
						Cnpj:     "cnpj",
						Bank:     "bank",
						Amount:   5.0,
						Wallet:   "wallet",
					},
					CheckpointDate:  &[]time.Time{time.Now()}[0],
					CheckpointValue: &[]float64{5.0}[0],
				},
				{
					Investiment: Investiment{
						ID:       "uuid3",
						Date:     time.Now(),
						Box:      "box-name",
						Category: "category-name",
						Name:     "fund-name",
						Cnpj:     "cnpj",
						Bank:     "bank",
						Amount:   15.0,
						Wallet:   "wallet",
					},
					CheckpointDate:  &[]time.Time{time.Now()}[0],
					CheckpointValue: &[]float64{25.0}[0],
				},
			},
			want: &Funds{
				Fund{
					Name:         "fund-name",
					Cnpj:         "cnpj",
					Bank:         "bank",
					BoxName:      "box-name",
					CategoryName: "category-name",
					Amount:       30.0,
					Value:        80.0,
					Share:        3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ii.Funds(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shares.Funds() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
