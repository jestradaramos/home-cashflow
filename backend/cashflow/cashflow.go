package cashflow

import (
	. "home-cashflow/backend/home"
	. "home-cashflow/backend/mortgage"
)

type Cashflow struct {
	Mortgage
	Home
	Net float32
}

// Rename and do calculations
func NewCashflow(m Mortgage, h Home) *Cashflow {
	return &Cashflow{m, h, 0}
}
