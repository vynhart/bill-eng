package bill

import (
	"amartha_bill_eng/bill/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Service struct {
	db *database.Queries
}

func NewDatabase(config map[string]string) *database.Queries {
	dbUser := config["db_username"]
	dbPassword := config["db_password"]
	dbHost := config["db_host"]
	dbName := config["db_name"]

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		panic(err)
	}
	return database.New(db)
}

func NewService(config map[string]string) *Service {
	return &Service{
		db: NewDatabase(config),
	}
}

const INTEREST_TYPE_FLAT_ANNUAL = "flat_annual"
const WEEKLY = "weekly"
const MONTHLY = "monthly"

func (s *Service) GenerateBillForLoan(loan Loan) error {
	ctx := context.Background()
	bill, err := s.db.FindBillByLoanId(ctx, loan.GetId())
	if err == nil || bill.ID != 0 {
		return fmt.Errorf("billing for loan id %s is already generated", loan.GetId())
	}

	if loan.GetInterestType() == INTEREST_TYPE_FLAT_ANNUAL {
		return s.generateFlatInterestBills(loan)
	}

	return fmt.Errorf("interest type %s is not known", loan.GetInterestType())
}

func (s *Service) OutstandingAmount(loan Loan) int {
	return 0 // fixme
}

func (s *Service) IsDelinquent(custId string) bool {
	return false // fixme
}

func (s *Service) MarkBillAsPaid(billId string) error {
	return nil // fixme
}

func (s *Service) generateFlatInterestBills(loan Loan) error {
	loanAndIntr := loan.GetAmount() + (loan.GetAmount() * uint(loan.GetInterest()))

	startDate := startOfDay(loan.GetDisbursedAt().Add(time.Hour * 24))
	dueDate, err := calcDueDate(startDate, loan.GetRepaymentTerm())
	if err != nil {
		return err
	}

	billAmount := loanAndIntr / uint(loan.GetInstallment())
	for i := 0; i < loan.GetInstallment(); i++ {
		param := database.CreateBillParams{
			LoanID:    loan.GetId(),
			StartDate: startDate,
			DueDate:   dueDate,
			Amount: sql.NullInt32{
				Int32: int32(billAmount),
			},
			CustomerID: loan.GetCustomerId(),
		}
		err := s.db.CreateBill(context.Background(), param)
		if err != nil {
			log.Fatal(err)
		}
		startDate = startOfDay(dueDate.Add(time.Hour * 24))
		dueDate, _ = calcDueDate(startDate, loan.GetRepaymentTerm())
	}

	return nil
}

func startOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func calcDueDate(start time.Time, repaymentTerm string) (time.Time, error) {
	switch repaymentTerm {
	case WEEKLY:
		return start.AddDate(0, 0, 7).Add(time.Second * -1), nil
	case MONTHLY:
		return start.AddDate(0, 1, 0).Add(time.Second * -1), nil
	default:
		return time.Time{}, fmt.Errorf("unkown repayment term %s ", repaymentTerm)
	}
}
