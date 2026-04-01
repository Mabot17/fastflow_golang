package crud

import (
	"context"
	"errors"
	"inventaris-app/config"
	"inventaris-app/schema"
)

func CreateStockOut(req schema.CreateStockOutRequest) error {
	tx, err := config.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var skID int

	// =======================
	// INSERT HEADER
	// =======================

	err = tx.QueryRow(context.Background(),
		`INSERT INTO inventaris.stok_keluar (sk_status, sk_pelanggan)
		 VALUES ('DRAFT', $1)
		 RETURNING sk_id`,
		req.Pelanggan).
		Scan(&skID)

	if err != nil {
		return err
	}

	// =======================
	// PROCESS ITEMS
	// =======================

	for _, item := range req.Items {

		var physical, reserved float64

		// =======================
		// GET / CREATE INVENTORY
		// =======================

		err := tx.QueryRow(context.Background(),
			`SELECT inv_physical_stock, inv_reserved_stock
			 FROM inventaris.inventaris_stok
			 WHERE inv_prd_id=$1
			 FOR UPDATE`,
			item.ProductID).
			Scan(&physical, &reserved)

		if err != nil {

			// 🔥 jika belum ada → buat row dulu
			if err.Error() == "no rows in result set" {

				_, err = tx.Exec(context.Background(),
					`INSERT INTO inventaris.inventaris_stok
					 (inv_prd_id, inv_physical_stock, inv_reserved_stock, inv_available_stock)
					 VALUES ($1, 0, 0, 0)`,
					item.ProductID)

				if err != nil {
					return err
				}

				physical = 0
				reserved = 0

			} else {
				return err
			}
		}

		available := physical - reserved

		// =======================
		// VALIDATION
		// =======================

		if available < item.Qty {
			return errors.New("stock not enough")
		}

		// =======================
		// RESERVE STOCK
		// =======================

		_, err = tx.Exec(context.Background(),
			`UPDATE inventaris.inventaris_stok
			 SET 
			   inv_reserved_stock = inv_reserved_stock + $1,
			   inv_available_stock = inv_available_stock - $1,
			   inv_updated_at = NOW()
			 WHERE inv_prd_id=$2`,
			item.Qty, item.ProductID)

		if err != nil {
			return err
		}

		// =======================
		// INSERT DETAIL
		// =======================

		_, err = tx.Exec(context.Background(),
			`INSERT INTO inventaris.stok_keluar_produk
			 (skp_sk_id, skp_prd_id, skp_qty)
			 VALUES ($1,$2,$3)`,
			skID, item.ProductID, item.Qty)

		if err != nil {
			return err
		}
	}

	// =======================
	// COMMIT
	// =======================

	return tx.Commit(context.Background())
}

func CancelStockOut(id string) error {
	tx, err := config.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var status string
	err = tx.QueryRow(context.Background(),
		`SELECT sk_status 
		 FROM inventaris.stok_keluar 
		 WHERE sk_id=$1
		 FOR UPDATE`, id).
		Scan(&status)

	if err != nil {
		return err
	}

	if status == "DONE" {
		return errors.New("cannot cancel DONE")
	}

	if status == "CANCELLED" {
		return errors.New("already cancelled")
	}

	rows, err := tx.Query(context.Background(),
		`SELECT skp_prd_id, skp_qty
		 FROM inventaris.stok_keluar_produk
		 WHERE skp_sk_id=$1`, id)

	if err != nil {
		return err
	}

	var items []struct {
		ProductID int
		Qty       float64
	}

	for rows.Next() {
		var item struct {
			ProductID int
			Qty       float64
		}
		rows.Scan(&item.ProductID, &item.Qty)
		items = append(items, item)
	}
	rows.Close()

	for _, item := range items {
		_, err := tx.Exec(context.Background(),
			`UPDATE inventaris.inventaris_stok
			 SET 
			   inv_reserved_stock = inv_reserved_stock - $1,
			   inv_available_stock = inv_available_stock + $1,
			   inv_updated_at = NOW()
			 WHERE inv_prd_id=$2`,
			item.Qty, item.ProductID)

		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(context.Background(),
		`UPDATE inventaris.stok_keluar
		 SET sk_status='CANCELLED',
		     sk_updated_at=NOW()
		 WHERE sk_id=$1`, id)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func FinishStockOut(id string) error {
	tx, err := config.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var status string
	err = tx.QueryRow(context.Background(),
		`SELECT sk_status 
		 FROM inventaris.stok_keluar 
		 WHERE sk_id=$1
		 FOR UPDATE`, id).
		Scan(&status)

	if err != nil {
		return err
	}

	if status == "DONE" {
		return errors.New("already DONE")
	}

	rows, err := tx.Query(context.Background(),
		`SELECT skp_prd_id, skp_qty
		 FROM inventaris.stok_keluar_produk
		 WHERE skp_sk_id=$1`, id)

	if err != nil {
		return err
	}

	var items []struct {
		ProductID int
		Qty       float64
	}

	for rows.Next() {
		var item struct {
			ProductID int
			Qty       float64
		}
		rows.Scan(&item.ProductID, &item.Qty)
		items = append(items, item)
	}
	rows.Close()

	for _, item := range items {

		// final commit
		_, err := tx.Exec(context.Background(),
			`UPDATE inventaris.inventaris_stok
			 SET 
			   inv_physical_stock = inv_physical_stock - $1,
			   inv_reserved_stock = inv_reserved_stock - $1,
			   inv_updated_at = NOW()
			 WHERE inv_prd_id=$2`,
			item.Qty, item.ProductID)

		if err != nil {
			return err
		}

		// kartu stok (ONLY DONE)
		_, err = tx.Exec(context.Background(),
			`INSERT INTO inventaris.kartu_stok
			 (ks_prd_id, ks_type, ks_qty, ks_ref_id, ks_ref_type)
			 VALUES ($1,'OUT',$2,$3,'stok_keluar')`,
			item.ProductID, item.Qty, id)

		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(context.Background(),
		`UPDATE inventaris.stok_keluar
		 SET sk_status='DONE',
		     sk_updated_at=NOW()
		 WHERE sk_id=$1`, id)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}