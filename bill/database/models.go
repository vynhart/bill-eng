// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type BillsStatus string

const (
	BillsStatusWaitingForPayment BillsStatus = "waiting_for_payment"
	BillsStatusPaid              BillsStatus = "paid"
)

func (e *BillsStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BillsStatus(s)
	case string:
		*e = BillsStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for BillsStatus: %T", src)
	}
	return nil
}

type NullBillsStatus struct {
	BillsStatus BillsStatus
	Valid       bool // Valid is true if BillsStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBillsStatus) Scan(value interface{}) error {
	if value == nil {
		ns.BillsStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BillsStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBillsStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BillsStatus), nil
}

type Bill struct {
	ID         uint32
	LoanID     string
	CustomerID string
	StartDate  time.Time
	DueDate    time.Time
	PaidAt     sql.NullTime
	Amount     uint32
	Status     NullBillsStatus
	PaymentID  sql.NullString
}
