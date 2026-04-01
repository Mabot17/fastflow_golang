package model

import "time"

type LapStock struct {
	ProductName string
	Sumber      string
	Masuk       float64
	Keluar      float64
	Saldo       float64
	Tanggal     time.Time
}