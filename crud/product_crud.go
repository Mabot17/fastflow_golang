package crud

import (
	"context"
	"inventaris-app/config"
	"inventaris-app/model"
	"inventaris-app/schema"
)

func GetAllProducts() ([]model.Product, error) {
	rows, err := config.DB.Query(context.Background(),
		`SELECT prd_id, prd_nama, prd_sku FROM inventaris.produk`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var p model.Product
		rows.Scan(&p.ID, &p.Name, &p.SKU)
		products = append(products, p)
	}

	return products, nil
}

func GetProductByID(id string) (model.Product, error) {
	var p model.Product

	err := config.DB.QueryRow(context.Background(),
		`SELECT prd_id, prd_nama, prd_sku 
		 FROM inventaris.produk WHERE prd_id=$1`, id).
		Scan(&p.ID, &p.Name, &p.SKU)

	return p, err
}

func CreateProduct(req schema.CreateProductRequest) (model.Product, error) {
	var p model.Product

	err := config.DB.QueryRow(context.Background(),
		`INSERT INTO inventaris.produk (prd_nama, prd_sku)
		 VALUES ($1,$2)
		 RETURNING prd_id, prd_nama, prd_sku`,
		req.Name, req.SKU).
		Scan(&p.ID, &p.Name, &p.SKU)

	return p, err
}

func UpdateProduct(id string, req schema.UpdateProductRequest) (model.Product, error) {
	var p model.Product

	err := config.DB.QueryRow(context.Background(),
		`UPDATE inventaris.produk 
		 SET prd_nama=$1, prd_sku=$2
		 WHERE prd_id=$3
		 RETURNING prd_id, prd_nama, prd_sku`,
		req.Name, req.SKU, id).
		Scan(&p.ID, &p.Name, &p.SKU)

	return p, err
}

func PatchProduct(id string, req schema.UpdateProductRequest) (model.Product, error) {
	var p model.Product

	err := config.DB.QueryRow(context.Background(),
		`UPDATE inventaris.produk
		 SET 
		   prd_nama = COALESCE($1, prd_nama),
		   prd_sku  = COALESCE($2, prd_sku)
		 WHERE prd_id=$3
		 RETURNING prd_id, prd_nama, prd_sku`,
		req.Name, req.SKU, id).
		Scan(&p.ID, &p.Name, &p.SKU)

	return p, err
}

func DeleteProduct(id string) error {
	_, err := config.DB.Exec(context.Background(),
		`DELETE FROM inventaris.produk WHERE prd_id=$1`, id)
	return err
}