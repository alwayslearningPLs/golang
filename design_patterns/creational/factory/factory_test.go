package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPaymentMethod(t *testing.T) {
	for _, each := range []struct {
		description string
		input       Payment
		wantPm      PaymentMethod
		wantErr     error
	}{
		{
			description: "get payment method from cash payment",
			input:       Cash,
			wantPm:      CashPM{},
			wantErr:     nil,
		},
		{
			description: "get payment method from debit payment",
			input:       Debit,
			wantPm:      DebitPM{},
			wantErr:     nil,
		},
		{
			description: "error bro",
			input:       3,
			wantPm:      nil,
			wantErr:     ErrPaymentDoNotExists,
		},
	} {
		t.Run(each.description, func(t *testing.T) {
			gotPm, gotErr := GetPaymentMethod(each.input)

			assert.Equal(t, each.wantPm, gotPm)
			assert.Equal(t, each.wantErr, gotErr)
		})
	}
}

func TestPay(t *testing.T) {
	for _, each := range []struct {
		description string
		input       Payment
		want        string
	}{
		{
			description: "pay with cash payment method",
			input:       Cash,
			want:        "cash",
		}, {
			description: "pay with debit payment method",
			input:       Debit,
			want:        "debit",
		},
	} {
		t.Run(each.description, func(t *testing.T) {
			pm, err := GetPaymentMethod(each.input)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, each.want, pm.Pay(23.3))
		})
	}
}
