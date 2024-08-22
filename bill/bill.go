package bill

import (
	"time"
)

const STATUS_WAITING_FOR_PAYMENT = "waiting_for_payment"
const STATUS_PAID = "paid"
const STATUS_OVERDUE = "overdue"

type Bill struct {
	Id        string    `json:"id"`
	LoanId    string    `json:"loan_id"`
	StartDate time.Time `json:"start_date"`
	DueDate   time.Time `json:"due_date"`
	PaidAt    time.Time `json:"paid_at"`
	Amount    uint      `json:"amount"`
	Status    string    `json:"status"`
	CustomerId string `json:"customer_id"`
	PaymentId string    `json:"payment_id"`
}

func (b *Bill) IsPaid() bool {
	// assuming that system was built after 2020
	return b.PaidAt.After(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC))
}
