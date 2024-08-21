package bill

import "time"


type LoanService interface {
	GetLoan(loanId string) Loan
}

type Loan interface {
	GetId() string
	GetRepaymentTerm() string // weekly, monthly
	GetInstallment() int // how many times payment should be made
	GetInterestType() string // for now only flat_annual
	GetInterest() float32
	GetAmount() uint
	GetCustomerId() string
	GetDisbursedAt() time.Time
}

