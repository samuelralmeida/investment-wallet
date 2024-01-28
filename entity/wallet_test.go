package entity_test

import (
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/samuelralmeida/investment-wallet/entity"
)

var fakeWalletName = "test-wallet"

var fakeFund1 = entity.Fund{
	ID:       "183df4fc-29ca-4371-a993-f4620991497f",
	Name:     "HIX CAPITAL INST FIA",
	Cnpj:     "22662135000155",
	Box:      "VALORIZAÇÃO",
	Flavor:   "FORA DO RADAR",
	Bank:     "BTG",
	MinValue: 100.0,
}

var fakeFund1Investment1 = entity.Investment{
	ID:     "9c7cd261-6063-4d21-996a-f5e2947482c1",
	FundID: fakeFund1.ID,
	Date:   time.Now().AddDate(0, -2, 0),
	Value:  1000.0,
	Wallet: fakeWalletName,
}

var fakeFund1Investment2 = entity.Investment{
	ID:     "cb306582-072b-4635-bf28-88e2f0c4647e",
	FundID: fakeFund1.ID,
	Date:   time.Now().AddDate(0, -1, 0),
	Value:  150.0,
	Wallet: fakeWalletName,
}

var fakeFund1Checkpoint = entity.Checkpoint{
	ID:     "f7b55121-073f-46fe-8b8a-ce2b4ef94416",
	FundID: fakeFund1.ID,
	Date:   time.Now(),
	Value:  1300.0,
}

var fakeFund2 = entity.Fund{
	ID:       "86bec889-ee9d-4b13-bcdc-96df4cb8892a",
	Name:     "BRASIL CAPITAL 30 FICFIA",
	Cnpj:     "14866273000128",
	Box:      "VALORIZAÇÃO",
	Flavor:   "QUALIDADE",
	Bank:     "BTG",
	MinValue: 500.0,
}

var fakeFund2Investment1 = entity.Investment{
	ID:     "1ef78156-4934-4067-a59d-447af914f0fd",
	FundID: fakeFund2.ID,
	Date:   time.Now().AddDate(0, -2, 0),
	Value:  600.50,
	Wallet: fakeWalletName,
}

var fakeFund2Investment2 = entity.Investment{
	ID:     "49e5e632-7cc5-41b6-b4a8-df52553dd224",
	FundID: fakeFund2.ID,
	Date:   time.Now().AddDate(0, -1, 0),
	Value:  303.33,
	Wallet: fakeWalletName,
}

var fakeFund2Checkpoint = entity.Checkpoint{
	ID:     "6bc51af1-715e-41c3-a93f-988b494866b3",
	FundID: fakeFund2.ID,
	Date:   time.Now(),
	Value:  300.10,
}

var fakeFund3 = entity.Fund{
	ID:       "577f73e1-43e5-4436-bc30-4a4b3bfe8a37",
	Name:     "SPX SEAHAWK FICFIRF CRPR LP ACCESS",
	Cnpj:     "35343590000130",
	Box:      "ESTABILIDADE",
	Flavor:   "BAUNILHA",
	Bank:     "BTG",
	MinValue: 500.0,
}

var fakeFund3Investment1 = entity.Investment{
	ID:     "e4ac23a3-b569-4198-b8f2-aca405b378aa",
	FundID: fakeFund3.ID,
	Date:   time.Now().AddDate(0, -2, 0),
	Value:  3034.01,
	Wallet: fakeWalletName,
}

var fakeFund3Checkpoint = entity.Checkpoint{
	ID:     "6bc51af1-715e-41c3-a93f-988b494866b3",
	FundID: fakeFund3.ID,
	Date:   time.Now(),
	Value:  4000.53,
}

var baseWallet = entity.Wallet{
	Name: "test-wallet",
	FundsDetail: []entity.FundDetail{
		{
			Fund:        fakeFund1,
			Investments: entity.Investments{fakeFund1Investment1, fakeFund1Investment2},
			Checkpoint:  fakeFund1Checkpoint,
		},
		{
			Fund:        fakeFund2,
			Investments: entity.Investments{fakeFund2Investment1, fakeFund2Investment2},
			Checkpoint:  fakeFund2Checkpoint,
		},
		{
			Fund:        fakeFund3,
			Investments: entity.Investments{fakeFund3Investment1},
			Checkpoint:  fakeFund3Checkpoint,
		},
	},
}

func TestWallet(t *testing.T) {
	t.Run("total invested value", func(t *testing.T) {
		var got float64 = baseWallet.TotalInvestedValue()
		var want float64 = fakeFund1Investment1.Value + fakeFund1Investment2.Value + fakeFund2Investment1.Value + fakeFund2Investment2.Value + fakeFund3Investment1.Value

		if got != want {
			t.Errorf("Wallet.TotalInvestedValue() = %v, want %v", got, want)
		}
	})

	t.Run("total current value", func(t *testing.T) {
		var got float64 = baseWallet.TotalCurrentValue()
		var want float64 = fakeFund1Checkpoint.Value + fakeFund2Checkpoint.Value + fakeFund3Checkpoint.Value

		if got != want {
			t.Errorf("Wallet.TotalCurrentValue() = %v, want %v", got, want)
		}
	})

	t.Run("boxes", func(t *testing.T) {
		var got = baseWallet.Boxes()
		var want entity.Boxes = entity.Boxes{
			entity.Box{
				Name:          "VALORIZAÇÃO",
				InvestedValue: fakeFund1Investment1.Value + fakeFund1Investment2.Value + fakeFund2Investment1.Value + fakeFund2Investment2.Value,
				CurrentValue:  fakeFund1Checkpoint.Value + fakeFund2Checkpoint.Value,
				FundsDetail:   []entity.FundDetail{baseWallet.FundsDetail[0], baseWallet.FundsDetail[1]},
			},
			entity.Box{
				Name:          "ESTABILIDADE",
				InvestedValue: fakeFund3Investment1.Value,
				CurrentValue:  fakeFund3Checkpoint.Value,
				FundsDetail:   []entity.FundDetail{baseWallet.FundsDetail[2]},
			},
		}

		sort.Slice(got, func(i, j int) bool { return i > j })
		sort.Slice(want, func(i, j int) bool { return i > j })

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wallet.Boxes() = %v, want %v", got, want)
		}
	})

}
