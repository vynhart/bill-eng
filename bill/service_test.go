package bill

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

type LoanMock struct {
	interest float32
	interestType string
	installment int
}

func (lm *LoanMock) GetId() string {
	return uuid.New().String()
}

func (lm *LoanMock) GetRepaymentTerm() string {
	return "weekly"
}

func (lm *LoanMock) GetInstallment() int {
	return lm.installment
}

func (lm *LoanMock) GetInterestType() string {
	return lm.interestType
}

func (lm *LoanMock) GetInterest() float32 {
	return lm.interest
}

func (lm *LoanMock) GetAmount() uint {
	return 5000000
}

func (lm *LoanMock) GetCustomerId() string {
	return uuid.New().String()
}

func (lm *LoanMock) GetDisbursedAt() time.Time {
	return time.Now().Add(time.Hour * -1)
}

func Test_GenerateBillForLoan(t *testing.T) {
	loan := &LoanMock{
		interest: 10.0,
		interestType: "flat_annual",
		installment: 10,
	}

	srv := NewService()
	err := srv.GenerateBillForLoan(loan)
	if err != nil {
		t.Fatal("GenerateBilLForLoan return unexpected error: ", err)
	}

	// second time the method is called it should raise error
	// because the bill should exists on database
	err = srv.GenerateBillForLoan(loan)
	if err == nil {
		t.Fatal("Second time GenerateBillForLoan is called for the same record, it should be failed")
	}
}

