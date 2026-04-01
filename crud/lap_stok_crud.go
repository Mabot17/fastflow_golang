package crud

import (
	"bytes"
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"inventaris-app/config"
	"inventaris-app/schema"
	"inventaris-app/model"
)

func GetLapStock(req schema.GetLapStockRequest) ([]model.LapStock, error) {

	query := `
	SELECT 
		p.prd_nama,
		CASE 
			WHEN ks.ks_ref_type = 'stok_masuk' THEN sm.sm_supplier
			WHEN ks.ks_ref_type = 'stok_keluar' THEN sk.sk_pelanggan
			ELSE '-'
		END as sumber,

		CASE WHEN ks.ks_type = 'IN' THEN ks.ks_qty ELSE 0 END as masuk,
		CASE WHEN ks.ks_type = 'OUT' THEN ks.ks_qty ELSE 0 END as keluar,

		ks.ks_created_at

	FROM inventaris.kartu_stok ks
	JOIN inventaris.produk p ON p.prd_id = ks.ks_prd_id
	LEFT JOIN inventaris.stok_masuk sm ON sm.sm_id = ks.ks_ref_id AND ks.ks_ref_type = 'stok_masuk'
	LEFT JOIN inventaris.stok_keluar sk ON sk.sk_id = ks.ks_ref_id AND ks.ks_ref_type = 'stok_keluar'
	WHERE 1=1
	`

	args := []interface{}{}
	idx := 1

	if req.ProductName != "" {
		query += fmt.Sprintf(" AND p.prd_nama ILIKE $%d", idx)
		args = append(args, "%"+req.ProductName+"%")
		idx++
	}

	if req.StartDate != "" {
		query += fmt.Sprintf(" AND ks.ks_created_at >= $%d", idx)
		args = append(args, req.StartDate)
		idx++
	}

	if req.EndDate != "" {
		query += fmt.Sprintf(" AND ks.ks_created_at <= $%d", idx)
		args = append(args, req.EndDate)
		idx++
	}

	query += " ORDER BY ks.ks_created_at ASC"

	rows, err := config.DB.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.LapStock
	var saldo float64 = 0

	for rows.Next() {
		var r model.LapStock

		err := rows.Scan(
			&r.ProductName,
			&r.Sumber,
			&r.Masuk,
			&r.Keluar,
			&r.Tanggal,
		)

		if err != nil {
			return nil, err
		}

		// hitung saldo berjalan
		saldo += r.Masuk - r.Keluar
		r.Saldo = saldo

		result = append(result, r)
	}

	return result, nil
}

func ExportLapStockExcel(req schema.GetLapStockRequest) (*bytes.Buffer, error) {

	data, err := GetLapStock(req)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Sheet1"

	// =======================
	// HEADER
	// =======================

	headers := []string{
		"No",
		"Produk",
		"Sumber",
		"Masuk",
		"Keluar",
		"Saldo",
		"Tanggal",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// =======================
	// DATA
	// =======================

	for i, row := range data {

		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), i+1)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), row.ProductName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), row.Sumber)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), row.Masuk)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", i+2), row.Keluar)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", i+2), row.Saldo)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", i+2), row.Tanggal)
	}

	// =======================
	// BUFFER
	// =======================

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buf, nil
}