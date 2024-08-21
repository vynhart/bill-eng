package bill

type PaymentObserver struct {}

func (pl *PaymentObserver) ObservePayment() {
	// listen to queue service, for example kafka
	// evaluate the message, check billing id
	// update billing id to paid.
}