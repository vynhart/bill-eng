// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const countCustomerOverdue = `-- name: CountCustomerOverdue :one
SELECT COUNT(*) FROM bills WHERE customer_id = ? AND due_date < NOW() AND paid_at IS NULL
`

func (q *Queries) CountCustomerOverdue(ctx context.Context, customerID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countCustomerOverdue, customerID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createBill = `-- name: CreateBill :exec
INSERT INTO bills (loan_id, customer_id, start_date, due_date, amount) VALUES (?, ?, ?, ?, ?)
`

type CreateBillParams struct {
	LoanID     string
	CustomerID string
	StartDate  time.Time
	DueDate    time.Time
	Amount     sql.NullInt32
}

func (q *Queries) CreateBill(ctx context.Context, arg CreateBillParams) error {
	_, err := q.db.ExecContext(ctx, createBill,
		arg.LoanID,
		arg.CustomerID,
		arg.StartDate,
		arg.DueDate,
		arg.Amount,
	)
	return err
}

const findBill = `-- name: FindBill :one
SELECT id, loan_id, customer_id, start_date, due_date, paid_at, amount, status, payment_id FROM bills WHERE id = ? LIMIT 1
`

func (q *Queries) FindBill(ctx context.Context, id uint32) (Bill, error) {
	row := q.db.QueryRowContext(ctx, findBill, id)
	var i Bill
	err := row.Scan(
		&i.ID,
		&i.LoanID,
		&i.CustomerID,
		&i.StartDate,
		&i.DueDate,
		&i.PaidAt,
		&i.Amount,
		&i.Status,
		&i.PaymentID,
	)
	return i, err
}

const findBillByLoanId = `-- name: FindBillByLoanId :one
SELECT id, loan_id, customer_id, start_date, due_date, paid_at, amount, status, payment_id FROM bills WHERE loan_id = ? LIMIT 1
`

func (q *Queries) FindBillByLoanId(ctx context.Context, loanID string) (Bill, error) {
	row := q.db.QueryRowContext(ctx, findBillByLoanId, loanID)
	var i Bill
	err := row.Scan(
		&i.ID,
		&i.LoanID,
		&i.CustomerID,
		&i.StartDate,
		&i.DueDate,
		&i.PaidAt,
		&i.Amount,
		&i.Status,
		&i.PaymentID,
	)
	return i, err
}

const listBillByCustId = `-- name: ListBillByCustId :many
SELECT id, loan_id, customer_id, start_date, due_date, paid_at, amount, status, payment_id FROM bills WHERE customer_id = ? LIMIT 1
`

func (q *Queries) ListBillByCustId(ctx context.Context, customerID string) ([]Bill, error) {
	rows, err := q.db.QueryContext(ctx, listBillByCustId, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bill
	for rows.Next() {
		var i Bill
		if err := rows.Scan(
			&i.ID,
			&i.LoanID,
			&i.CustomerID,
			&i.StartDate,
			&i.DueDate,
			&i.PaidAt,
			&i.Amount,
			&i.Status,
			&i.PaymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markBillAsPaid = `-- name: MarkBillAsPaid :exec
UPDATE bills SET paid_at = ? WHERE id = ?
`

type MarkBillAsPaidParams struct {
	PaidAt sql.NullTime
	ID     uint32
}

func (q *Queries) MarkBillAsPaid(ctx context.Context, arg MarkBillAsPaidParams) error {
	_, err := q.db.ExecContext(ctx, markBillAsPaid, arg.PaidAt, arg.ID)
	return err
}
