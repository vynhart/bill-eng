package loan

import "time"

// this loan type is expected to implement interface
// in bill.Loan.
type Loan struct {}

func (lm *Loan) GetId() string {
	// todo: implement this function
	return ""
}

func (lm *Loan) GetRepaymentTerm() string {
	// todo: implement this function
	return ""
}

func (lm *Loan) GetInstallment() int {
	// todo: implement this function
	return 0
}

func (lm *Loan) GetInterestType() string {
	// todo: implement this function
	return ""
}

func (lm *Loan) GetInterest() float32 {
	// todo: implement this function
	return 0.0
}

func (lm *Loan) GetAmount() uint {
	// todo: implement this function
	return 0
}

func (lm *Loan) GetCustomerId() string {
	// todo: implement this function
	return ""
}

func (lm *Loan) GetDisbursedAt() time.Time {
	// todo: implement this function
	return time.Time{}
}