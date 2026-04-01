package crud

import (
	"context"
	"errors"
	"inventaris-app/config"
	"inventaris-app/schema"
)

// =======================
// CREATE STOCK IN
// =======================

func CreateStockIn(req schema.CreateStockInRequest) error {
	tx, err := config.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var stockInID int

	// insert header (dengan pelanggan)
	err = tx.QueryRow(context.Background(),
		`INSERT INTO inventaris.stok_masuk (sm_status, sm_pelanggan)
		 VALUES ('CREATED', $1)
		 RETURNING sm_id`,
		req.Pelanggan).
		Scan(&stockInID)

	if err != nil {
		return err
	}

	// insert multiple produk
	for _, item := range req.Items {
		_, err := tx.Exec(context.Background(),
			`INSERT INTO inventaris.stok_masuk_produk
			 (smp_sm_id, smp_prd_id, smp_qty)
			 VALUES ($1,$2,$3)`,
			stockInID, item.ProductID, item.Qty)

		if err != nil {
			return err
		}
	}

	return tx.Commit(context.Background())
}

// =======================
// FINISH STOCK IN (DONE)
// =======================
func FinishStockIn(id string) error {
	tx, err := config.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// =======================
	// LOCK HEADER
	// =======================

	var status string
	err = tx.QueryRow(context.Background(),
		`SELECT sm_status 
		 FROM inventaris.stok_masuk 
		 WHERE sm_id=$1
		 FOR UPDATE`, id).
		Scan(&status)
		

	if err != nil {
		return err
	}

	if status == "DONE" {
		return errors.New("already DONE")
	}

	// =======================
	// AMBIL DETAIL
	// =======================

	rows, err := tx.Query(context.Background(),
		`SELECT smp_prd_id, smp_qty
		 FROM inventaris.stok_masuk_produk
		 WHERE smp_sm_id=$1`, id)

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

		if err := rows.Scan(&item.ProductID, &item.Qty); err != nil {
			rows.Close()
			return err
		}

		items = append(items, item)
	}
	rows.Close()

	// =======================
	// EXECUTION
	// =======================

	for _, item := range items {

		// 🔥 update semua state
		_, err := tx.Exec(context.Background(),
			`INSERT INTO inventaris.inventaris_stok
			(inv_prd_id, inv_physical_stock, inv_reserved_stock, inv_available_stock)
			VALUES ($1, $2, 0, $2)
			ON CONFLICT (inv_prd_id)
			DO UPDATE SET
			inv_physical_stock = inventaris.inventaris_stok.inv_physical_stock + EXCLUDED.inv_physical_stock,
			inv_available_stock = inventaris.inventaris_stok.inv_available_stock + EXCLUDED.inv_available_stock,
			inv_updated_at = NOW()`,
			item.ProductID, item.Qty)

		if err != nil {
			return err
		}

		// 🔥 kartu stok hanya saat DONE
		_, err = tx.Exec(context.Background(),
			`INSERT INTO inventaris.kartu_stok
			 (ks_prd_id, ks_type, ks_qty, ks_ref_id, ks_ref_type)
			 VALUES ($1,'IN',$2,$3,'stok_masuk')`,
			item.ProductID, item.Qty, id)

		if err != nil {
			return err
		}
	}

	// =======================
	// UPDATE STATUS
	// =======================

	_, err = tx.Exec(context.Background(),
		`UPDATE inventaris.stok_masuk
		 SET sm_status='DONE',
		     sm_updated_at=NOW()
		 WHERE sm_id=$1`, id)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

// =======================
// CANCEL STOCK IN
// =======================

func CancelStockIn(id string) error {
	tx, err := config.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var status string
	err = tx.QueryRow(context.Background(),
		`SELECT sm_status 
		 FROM inventaris.stok_masuk 
		 WHERE sm_id=$1
		 FOR UPDATE`, id).
		Scan(&status)

	if err != nil {
		return err
	}

	if status == "DONE" {
		return errors.New("cannot cancel, already DONE")
	}

	if status == "CANCELLED" {
		return errors.New("already cancelled")
	}

	// 🔥 hanya update status (tidak sentuh stok & kartu_stok)
	_, err = tx.Exec(context.Background(),
		`UPDATE inventaris.stok_masuk
		 SET sm_status='CANCELLED',
		     sm_updated_at=NOW()
		 WHERE sm_id=$1`, id)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}