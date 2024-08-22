package worker

type PaymentObserver struct {}

// NOT IMPLEMENTED YET
func (pl *PaymentObserver) ObservePayment() {
	// listen to payment event on queue service such as kafka,
	// evaluate the message, check billing id and amount,
	// update billing id to paid by calling Service.MarkBillingAsPaid
	// from package bill
}
