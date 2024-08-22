package bill

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type LoanMock struct {
	id           string
	custId       string
	interest     float32
	interestType string
	installment  int
	amount       int
	disbursedAt  time.Time
}

func (lm *LoanMock) GetId() string {
	return lm.id
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
	return uint(lm.amount)
}

func (lm *LoanMock) GetCustomerId() string {
	return lm.custId
}

func (lm *LoanMock) GetDisbursedAt() time.Time {
	return lm.disbursedAt
}

func getConfigTest() map[string]string {
	fileName := "../.env.testing"
	err := godotenv.Load(fileName)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return map[string]string{
		"db_username": os.Getenv("DB_USERNAME"),
		"db_password": os.Getenv("DB_PASSWORD"),
		"db_host":     os.Getenv("DB_HOST"),
		"db_name":     os.Getenv("DB_NAME"),
	}
}

func Test_BillService(t *testing.T) {
	loan := &LoanMock{
		id:           uuid.New().String(),
		custId:       uuid.New().String(),
		interest:     0.1,
		interestType: "flat_annual",
		installment:  10,
		amount:       5000000,
		disbursedAt:  time.Now().Add(time.Hour * -1),
	}

	srv := NewService(getConfigTest())
	err := srv.GenerateBillForLoan(loan)
	if err != nil {
		t.Fatal("GenerateBilLForLoan return unexpected error: ", err)
	}

	err = srv.GenerateBillForLoan(loan)
	if err == nil {
		t.Fatal("Second time GenerateBillForLoan is called for the same record, it should be failed")
	}

	outsAmt := srv.GetOutstandingAmount(loan.GetId())
	if outsAmt != uint32(5500000) {
		t.Fatalf("incorrect outstanding amount %v", outsAmt)
	}

	bills := srv.GetOutstandingBill(loan.GetId())
	if len(bills) == 0 {
		t.Fatalf("outstanding bill is empty")
	}

	outsBill := bills[0]
	t.Log(outsBill.ID)
	err = srv.MarkBillAsPaid(outsBill.ID, time.Now().Add(time.Hour*-1))
	if err != nil {
		t.Fatal("error when marking bill as paid")
	}

	outsAmtAfterPaid := srv.GetOutstandingAmount(loan.GetId())
	if outsAmtAfterPaid != outsAmt-outsBill.Amount {
		t.Fatalf("incorrect outstanding amount after paid %v", outsAmtAfterPaid)
	}

	deliq := srv.IsDelinquent(loan.GetCustomerId())
	if deliq {
		t.Fatalf("Customer should not be deliquent")
	}
}

func Test_BillService_IsDeliquent(t *testing.T) {
	loan := &LoanMock{
		id:           uuid.New().String(),
		custId:       uuid.New().String(),
		interest:     0.1,
		interestType: "flat_annual",
		installment:  10,
		amount:       5000000,
		disbursedAt:  time.Now().Add(time.Hour * -1 * 24 * 16),
	}

	srv := NewService(getConfigTest())
	srv.GenerateBillForLoan(loan)

	deliq := srv.IsDelinquent(loan.GetCustomerId())
	if !deliq {
		t.Fatalf("Customer should be deliquent")
	}
}
