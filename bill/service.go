package bill

import (
	"fmt"
	"time"
)

type Service struct {
	repo    *Repository
}

func NewService() *Service {
	return &Service{
		repo: NewRepository(),
	}
}

const INTEREST_TYPE_FLAT_ANNUAL = "flat_annual"
const WEEKLY = "weekly"
const MONTHLY = "monthly"

func (s *Service) GenerateBillForLoan(loan Loan) error {
	if s.repo.IsBillingExistsForLoanId(loan.GetId()) {
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

	bills := []Bill{}
	billAmount := loanAndIntr / uint(loan.GetInstallment())
	for i := 0; i < loan.GetInstallment(); i++ {
		bill := Bill{
			LoanId:    loan.GetId(),
			StartDate: startDate,
			DueDate:   dueDate,
			Amount:    billAmount,
			Status:    STATUS_WAITING_FOR_PAYMENT,
		}
		bills = append(bills, bill)
		startDate := startOfDay(bill.DueDate.Add(time.Hour * 24))
		dueDate, _ = calcDueDate(startDate, loan.GetRepaymentTerm())
	}

	s.repo.BulkInsert(bills)

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
