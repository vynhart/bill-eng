package bill

import "time"

type Repository struct {}

func NewRepository() *Repository {
	// fixme
	return &Repository{}
}

func (r *Repository) Insert(bill *Bill) error {
	// fixme
	id := "some-uuid" // insert to db, get the id from transaction
	bill.Id = id
	return nil
}

func (r *Repository) BulkInsert(bills []Bill) error {
	return nil // fixme
}

func (r *Repository) Update(bill *Bill) error {
	// update on db
	return nil
}

func (r *Repository) Read(id string) (*Bill) {
	// fixme
	return &Bill{}
}

func (r *Repository) MarkPaid(bill *Bill, paidAt time.Time) {

}

func (r *Repository) CountOverdue(customerId string) int {
	// fixme
	return 0
}

func (r *Repository) IsBillingExistsForLoanId(loanId string) bool {
	// fixme, implement
	return false
}

func (r *Repository) LockForUpdate(id string) (Lock, error) {
	return Lock{}, nil // fixme
}

func (r *Repository) Start() error {
	return nil // fixme
}

func (r *Repository) Commit() error {
	return nil // fixme
}

func (r *Repository) Abort() error {
	return nil // fixme
}

type Lock struct{}