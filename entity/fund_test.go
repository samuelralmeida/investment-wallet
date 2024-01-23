package entity

/*
import (
	"reflect"
	"testing"
)

func TestFunds_Boxes(t *testing.T) {
	tests := []struct {
		name string
		f    *Funds
		want *Boxes
	}{
		{
			name: "group funds by box name",
			f: &Funds{
				Fund{
					Name:         "fund1",
					Cnpj:         "cnpj1",
					Bank:         "bank1",
					BoxName:      "box-name-a",
					CategoryName: "category-name1",
					Amount:       10.0,
					Value:        15.0,
					Share:        1,
					Checkpoints:  1,
				},
				Fund{
					Name:         "fund2",
					Cnpj:         "cnpj2",
					Bank:         "bank2",
					BoxName:      "box-name-a",
					CategoryName: "category-name2",
					Amount:       20.0,
					Value:        15.0,
					Share:        2,
					Checkpoints:  1,
				},
				Fund{
					Name:         "fund3",
					Cnpj:         "cnpj3",
					Bank:         "bank3",
					BoxName:      "box-name-b",
					CategoryName: "category-name3",
					Amount:       20.0,
					Value:        50.0,
					Share:        3,
					Checkpoints:  1,
				},
			},
			want: &Boxes{
				Box{
					Name:                "box-name-a",
					Amount:              30.0,
					Value:               30.0,
					QuantityFunds:       2,
					QuantityCheckpoints: 2,
					QuantityShares:      3,
				},
				Box{
					Name:                "box-name-b",
					Amount:              20,
					Value:               50,
					QuantityFunds:       1,
					QuantityCheckpoints: 1,
					QuantityShares:      3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Boxes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Funds.Boxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
