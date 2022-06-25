package cards

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	Amount         int
	Currency       string
	BankReturnCode string
	LastFoorDigit  string
}

func  (c *Card) Charge(amount int, currency string) (*stripe.PaymentIntent, string, error) {
	return c.createPaymentIntent(amount, currency)
}

func (c *Card) createPaymentIntent(amount int, currency string) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""

		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = stripeErrMessage(stripeErr.Code)
		}
		return nil, msg, err
	}

	return pi, "", nil
}

func stripeErrMessage(code stripe.ErrorCode) string {
	var msg string

	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Payment was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "Payment intent has expired"
	case stripe.ErrorCodeAmountTooLarge:
		msg = "Payment amount is too large"
	case stripe.ErrorCodeAmountTooSmall:
		msg = "Payment amount is too small"
	case stripe.ErrorCodeInvalidCVC:
		msg = "Invalid CVC"
	case stripe.ErrorCodeInvalidNumber:
		msg = "Invalid card number"
	case stripe.ErrorCodeInvalidExpiryMonth:
		msg = "Invalid expiry month"
	case stripe.ErrorCodeInvalidExpiryYear:
		msg = "Invalid expiry year"
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "Invalid postal code"
	case stripe.ErrorCodeIncorrectAddress:
		msg = "Incorrect address"
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Balance insufficient"
	default:
		msg = "Payment was declined"

	}
	return msg
}
