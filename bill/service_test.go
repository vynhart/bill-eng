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
	id string
	custId string
	interest float32
	interestType string
	installment int
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
	return 5000000
}

func (lm *LoanMock) GetCustomerId() string {
	return lm.custId
}

func (lm *LoanMock) GetDisbursedAt() time.Time {
	return time.Now().Add(time.Hour * -1)
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

func Test_GenerateBillForLoan(t *testing.T) {
	loan := &LoanMock{
		id: uuid.New().String(),
		custId: uuid.New().String(),
		interest: 10.0,
		interestType: "flat_annual",
		installment: 10,
	}

	srv := NewService(getConfigTest())
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

