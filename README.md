# 🚀 fastflow_golang

Backend service untuk sistem inventaris berbasis **Golang + FastFlow**.
Dirancang dengan pendekatan modular, ringan, dan siap berkembang menjadi sistem yang lebih besar.

---

## 🧠 Konsep Dasar

fastflow_golang dibangun dengan prinsip:

* **Separation of Concern** → router, crud, model, schema terpisah
* **Flow-based system** → stok masuk, keluar, dan laporan sebagai aliran data
* **Lightweight architecture** → tanpa over-engineering, tapi scalable

---

## 📁 Struktur Project

```bash
.
├── config/        # konfigurasi database & environment
├── crud/          # logic database (query, transaksi)
├── docs/          # swagger documentation
├── model/         # representasi struct database
├── router/        # handler endpoint
├── schema/        # request/response schema
├── db.sql         # struktur database
├── main.go        # entry point aplikasi
├── router.go      # register semua route
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Menjalankan Project

### 1. Masuk ke folder project

```bash
cd project_be
```

### 2. Jalankan aplikasi

```bash
go run .
```

---

## 📚 Swagger Documentation

### Generate / refresh swagger

```bash
swag init -g main.go -o docs
```

### Jika terjadi error (paksa parsing)

```bash
swag init -g main.go -d . --parseDependency --parseInternal
```

---

## 🔌 Endpoint Utama

### 📦 Produk

* `GET /products`
* `POST /products`
* `PUT /products/{id}`
* `DELETE /products/{id}`

---

### 📥 Stok Masuk

* `GET /stock-in`
* `GET /stock-in/{id}`
* `POST /stock-in`
* `POST /stock-in/{id}/finish`
* `POST /stock-in/{id}/cancel`

---

### 📤 Stok Keluar

* `GET /stock-out`
* `GET /stock-out/{id}`
* `POST /stock-out`
* `POST /stock-out/{id}/finish`
* `POST /stock-out/{id}/cancel`

---

### 📊 Laporan Stok

* `GET /lap-stok`
* `GET /lap-stok/export`

---

## 🔄 Flow Sistem

```text
Produk → Stok Masuk → Stok Keluar → Laporan
```

* **Stok Masuk** → menambah jumlah barang
* **Stok Keluar** → mengurangi jumlah barang
* **Laporan** → merekam jejak pergerakan

---

## 🧠 Catatan Arsitektur

* Menggunakan **transaction (BEGIN / COMMIT / ROLLBACK)** untuk operasi penting
* Query menggunakan **pgx (raw SQL)** untuk kontrol penuh
* Struktur modular memudahkan scaling ke microservice

---

## 🔥 Next Improvement (Opsional)

* Validasi stok (tidak boleh minus saat keluar)
* Pagination & filtering
* Authentication & multi-user
* Audit log transaksi
* Dashboard monitoring

---

## 🌿 Filosofi

Project ini tidak hanya tentang CRUD,
tetapi tentang **aliran data yang mencerminkan realitas sistem**:

* Masuk → pertambahan
* Keluar → pengurangan
* Laporan → kesadaran sistem

---

## 🤝 Kontribusi

Silakan eksplorasi, modifikasi, dan kembangkan.
Sistem ini dibuat untuk tumbuh, bukan hanya berjalan.

---

**fastflow_golang — bukan sekadar backend, tapi fondasi aliran sistem.**
