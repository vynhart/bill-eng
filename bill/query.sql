-- name: FindBill :one
SELECT * FROM bills WHERE id = ? LIMIT 1;

-- name: ListBillByCustId :many
SELECT * FROM bills WHERE customer_id = ? LIMIT 1;

-- name: CreateBill :exec
INSERT INTO bills (loan_id, customer_id, start_date, due_date, amount) VALUES (?, ?, ?, ?, ?);

-- name: MarkBillAsPaid :exec
UPDATE bills SET paid_at = ? WHERE id = ?;

-- name: CountCustomerOverdue :one
SELECT COUNT(*) FROM bills WHERE customer_id = ? AND due_date < NOW() AND paid_at IS NULL;

-- name: FindBillByLoanId :one
SELECT * FROM bills WHERE loan_id = ? LIMIT 1;


