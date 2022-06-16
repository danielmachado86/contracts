// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"fmt"
	"time"
)

type PeriodUnits string

const (
	PeriodUnitsDays   PeriodUnits = "days"
	PeriodUnitsMonths PeriodUnits = "months"
	PeriodUnitsYears  PeriodUnits = "years"
)

func (e *PeriodUnits) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PeriodUnits(s)
	case string:
		*e = PeriodUnits(s)
	default:
		return fmt.Errorf("unsupported scan type for PeriodUnits: %T", src)
	}
	return nil
}

type Templates string

const (
	TemplatesRental    Templates = "rental"
	TemplatesFreelance Templates = "freelance"
	TemplatesServices  Templates = "services"
)

func (e *Templates) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Templates(s)
	case string:
		*e = Templates(s)
	default:
		return fmt.Errorf("unsupported scan type for Templates: %T", src)
	}
	return nil
}

type Contract struct {
	ID       int64     `json:"id"`
	Template Templates `json:"template"`
}

type Party struct {
	Username   string    `json:"username"`
	ContractID int64     `json:"contractID"`
	CreatedAt  time.Time `json:"createdAt"`
}

type PeriodParam struct {
	ID         int64       `json:"id"`
	ContractID int64       `json:"contractID"`
	Name       string      `json:"name"`
	Value      int32       `json:"value"`
	Units      PeriodUnits `json:"units"`
}

type PriceParam struct {
	ID         int64   `json:"id"`
	ContractID int64   `json:"contractID"`
	Name       string  `json:"name"`
	Value      float64 `json:"value"`
	Currency   string  `json:"currency"`
}

type TimeParam struct {
	ID         int64     `json:"id"`
	ContractID int64     `json:"contractID"`
	Name       string    `json:"name"`
	Value      time.Time `json:"value"`
}

type User struct {
	Name              string    `json:"name"`
	LastName          string    `json:"lastName"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	HashedPassword    string    `json:"hashedPassword"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	CreatedAt         time.Time `json:"createdAt"`
}
