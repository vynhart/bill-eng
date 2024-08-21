package bill

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

type LoanMock struct {
	interest float32
	interestType string
}

func (lm *LoanMock) GetId() string {
	return uuid.New().String()
}

func (lm *LoanMock) GetRepaymentTerm() string {
	return "weekly"
}

func (lm *LoanMock) GetInstallment() int {
	return 50
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

type LoanServiceMock struct {
	loan Loan
}

func (lsm *LoanServiceMock) GetLoan(id string) (Loan) {
	return lsm.loan
}

func TestGenerateBillForLoanSuccess(t *testing.T) {
	loan := &LoanMock{
		interest: 10.0,
		interestType: "flat_annual",
	}

	srv := NewService()
	srv.LoanSrv = &LoanServiceMock{
		loan: loan,
	}

	err := srv.GenerateBillForLoan(loan)
	if err != nil {
		t.Fatal("GenerateBilLForLoan return unexpected error: ", err)
	}
}