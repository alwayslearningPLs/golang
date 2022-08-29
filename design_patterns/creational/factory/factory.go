package factory

import "errors"

// ErrPaymentDoNotExists is returned when the user try to create an object which is not
// inside the possible Payment options.
var ErrPaymentDoNotExists = errors.New("payment do not exists")

// PaymentMethod to return when calling the factory
type PaymentMethod interface {
	// Pay will be used to pay with the corresponding payment method
	Pay(float32) string
}

// Payment type to identify the allowed ones
type Payment int

// This are the payments availables to the user
const (
	Cash Payment = iota
	Debit
)

type CashPM struct{}
type DebitPM struct{}

func (c CashPM) Pay(cash float32) string {
	return "cash"
}

func (c DebitPM) Pay(cash float32) string {
	return "debit"
}

// GetPaymentMethod will return an instance of a PaymentMethod available
func GetPaymentMethod(payment Payment) (PaymentMethod, error) {
	var (
		pm  PaymentMethod
		err error
	)

	switch payment {
	case Cash:
		pm = CashPM{}
	case Debit:
		pm = DebitPM{}
	default:
		err = ErrPaymentDoNotExists
	}

	return pm, err
}
