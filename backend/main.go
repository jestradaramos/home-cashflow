package main

import (
	"home-cashflow/backend/home"
)

func main() {
	h := home.NewHome("", 2, 1.5, 700000, 2000, 500)
	if h != nil {
		return
	}
}
