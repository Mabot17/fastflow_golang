-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               PostgreSQL 18.3 on x86_64-windows, compiled by msvc-19.44.35223, 64-bit
-- Server OS:                    
-- HeidiSQL Version:             12.0.0.6468
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table inventaris.inventaris_stok
DROP TABLE IF EXISTS "inventaris_stok";
CREATE TABLE IF NOT EXISTS "inventaris_stok" (
	"inv_id" INTEGER NOT NULL DEFAULT 'nextval(''inventaris_stok_inv_id_seq''::regclass)',
	"inv_prd_id" INTEGER NOT NULL,
	"inv_physical_stock" NUMERIC(12,2) NULL DEFAULT '0',
	"inv_reserved_stock" NUMERIC(12,2) NULL DEFAULT '0',
	"inv_updated_at" TIMESTAMP NULL DEFAULT 'CURRENT_TIMESTAMP',
	"inv_available_stock" NUMERIC(12,2) NULL DEFAULT '0',
	PRIMARY KEY ("inv_id"),
	UNIQUE INDEX "inventaris_stok_inv_prd_id_key" ("inv_prd_id"),
	CONSTRAINT "inventaris_stok_inv_prd_id_fkey" FOREIGN KEY ("inv_prd_id") REFERENCES "produk" ("prd_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Dumping data for table inventaris.inventaris_stok: -1 rows
/*!40000 ALTER TABLE "inventaris_stok" DISABLE KEYS */;
INSERT INTO "inventaris_stok" ("inv_id", "inv_prd_id", "inv_physical_stock", "inv_reserved_stock", "inv_updated_at", "inv_available_stock") VALUES
	(5, 4, 6.00, 0.00, '2026-04-01 20:54:27.115476', 6.00),
	(7, 3, 26.00, 2.00, '2026-04-01 21:00:37.443902', 24.00),
	(6, 5, 12.00, 6.00, '2026-04-01 21:00:37.443902', 6.00);
/*!40000 ALTER TABLE "inventaris_stok" ENABLE KEYS */;

-- Dumping structure for table inventaris.kartu_stok
DROP TABLE IF EXISTS "kartu_stok";
CREATE TABLE IF NOT EXISTS "kartu_stok" (
	"ks_id" INTEGER NOT NULL DEFAULT 'nextval(''kartu_stok_ks_id_seq''::regclass)',
	"ks_prd_id" INTEGER NOT NULL,
	"ks_type" VARCHAR(20) NOT NULL,
	"ks_qty" NUMERIC(12,2) NOT NULL,
	"ks_ref_id" INTEGER NULL DEFAULT NULL,
	"ks_ref_type" VARCHAR(50) NULL DEFAULT NULL,
	"ks_created_at" TIMESTAMP NULL DEFAULT 'CURRENT_TIMESTAMP',
	PRIMARY KEY ("ks_id"),
	CONSTRAINT "kartu_stok_ks_prd_id_fkey" FOREIGN KEY ("ks_prd_id") REFERENCES "produk" ("prd_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Dumping data for table inventaris.kartu_stok: -1 rows
/*!40000 ALTER TABLE "kartu_stok" DISABLE KEYS */;
INSERT INTO "kartu_stok" ("ks_id", "ks_prd_id", "ks_type", "ks_qty", "ks_ref_id", "ks_ref_type", "ks_created_at") VALUES
	(10, 4, 'IN', 4.00, 14, 'stok_masuk', '2026-04-01 20:40:21.763863'),
	(11, 5, 'IN', 15.00, 14, 'stok_masuk', '2026-04-01 20:40:21.763863'),
	(12, 3, 'IN', 16.00, 14, 'stok_masuk', '2026-04-01 20:40:21.763863'),
	(13, 4, 'IN', 3.00, 13, 'stok_masuk', '2026-04-01 20:40:43.279244'),
	(14, 5, 'IN', 10.00, 13, 'stok_masuk', '2026-04-01 20:40:43.279244'),
	(15, 3, 'IN', 11.00, 13, 'stok_masuk', '2026-04-01 20:40:43.279244'),
	(16, 4, 'OUT', 1.00, 2, 'stok_keluar', '2026-04-01 20:54:27.115476'),
	(17, 5, 'OUT', 12.00, 6, 'stok_keluar', '2026-04-01 20:55:04.717339'),
	(18, 3, 'OUT', 1.00, 11, 'stok_keluar', '2026-04-01 21:00:37.443902'),
	(19, 5, 'OUT', 1.00, 11, 'stok_keluar', '2026-04-01 21:00:37.443902');
/*!40000 ALTER TABLE "kartu_stok" ENABLE KEYS */;

-- Dumping structure for table inventaris.produk
DROP TABLE IF EXISTS "produk";
CREATE TABLE IF NOT EXISTS "produk" (
	"prd_id" INTEGER NOT NULL DEFAULT 'nextval(''produk_prd_id_seq''::regclass)',
	"prd_nama" VARCHAR(100) NOT NULL,
	"prd_sku" VARCHAR(50) NOT NULL,
	"prd_created_at" TIMESTAMP NULL DEFAULT 'CURRENT_TIMESTAMP',
	PRIMARY KEY ("prd_id"),
	UNIQUE INDEX "produk_prd_sku_key" ("prd_sku")
);

-- Dumping data for table inventaris.produk: -1 rows
/*!40000 ALTER TABLE "produk" DISABLE KEYS */;
INSERT INTO "produk" ("prd_id", "prd_nama", "prd_sku", "prd_created_at") VALUES
	(2, 'Ban Mobil B', 'SKU-BAN-B', '2026-04-01 15:38:13.801802'),
	(4, 'Ban Mobil C', 'SKU-BAN-C', '2026-04-01 16:19:22.584021'),
	(5, 'Ban Mobil D', 'SKU-BAN-D', '2026-04-01 16:19:31.60773'),
	(1, 'Ban Partial', '', '2026-04-01 15:38:13.801802'),
	(3, 'Ban Mobil E', 'SKU-BAN-E', '2026-04-01 16:38:51.326311'),
	(7, 'oke', '123', '2026-04-01 22:04:10.925944');
/*!40000 ALTER TABLE "produk" ENABLE KEYS */;

-- Dumping structure for table inventaris.stok_keluar
DROP TABLE IF EXISTS "stok_keluar";
CREATE TABLE IF NOT EXISTS "stok_keluar" (
	"sk_id" INTEGER NOT NULL DEFAULT 'nextval(''stok_keluar_sk_id_seq''::regclass)',
	"sk_status" VARCHAR(20) NOT NULL,
	"sk_created_at" TIMESTAMP NULL DEFAULT 'CURRENT_TIMESTAMP',
	"sk_updated_at" TIMESTAMP NULL DEFAULT 'CURRENT_TIMESTAMP',
	"sk_pelanggan" VARCHAR(250) NULL DEFAULT 'NULL::character varying',
	PRIMARY KEY ("sk_id")
);

-- Dumping data for table inventaris.stok_keluar: -1 rows
/*!40000 ALTER TABLE "stok_keluar" DISABLE KEYS */;
INSERT INTO "stok_keluar" ("sk_id", "sk_status", "sk_created_at", "sk_updated_at", "sk_pelanggan") VALUES
	(2, 'DONE', '2026-04-01 20:51:50.63535', '2026-04-01 20:54:27.115476', 'coba'),
	(6, 'DONE', '2026-04-01 20:53:46.836809', '2026-04-01 20:55:04.717339', 'cobalah'),
	(8, 'DRAFT', '2026-04-01 20:55:46.716521', '2026-04-01 20:55:46.716521', 'cobalah'),
	(12, 'DRAFT', '2026-04-01 21:00:12.761439', '2026-04-01 21:00:12.761439', 'oke'),
	(11, 'DONE', '2026-04-01 20:59:45.463003', '2026-04-01 21:00:37.443902', 'rohman');
/*!40000 ALTER TABLE "stok_keluar" ENABLE KEYS */;

-- Dumping structure for table inventaris.stok_keluar_produk
DROP TABLE IF EXISTS "stok_keluar_produk";
CREATE TABLE IF NOT EXISTS "stok_keluar_produk" (
	"skp_id" INTEGER NOT NULL DEFAULT 'nextval(''stok_keluar_produk_skp_id_seq''::regclass)',
	"skp_sk_id" INTEGER NOT NULL,
	"skp_prd_id" INTEGER NOT NULL,
	"skp_qty" NUMERIC(12,2) NOT NULL,
	PRIMARY KEY ("skp_id"),
	CONSTRAINT "stok_keluar_produk_skp_prd_id_fkey" FOREIGN KEY ("skp_prd_id") REFERENCES "produk" ("prd_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT "stok_keluar_produk_skp_sk_id_fkey" FOREIGN KEY ("skp_sk_id") REFERENCES "stok_keluar" ("sk_id") ON UPDATE NO ACTION ON DELETE CASCADE
);

-- Dumping data for table inventaris.stok_keluar_produk: -1 rows
/*!40000 ALTER TABLE "stok_keluar_produk" DISABLE KEYS */;
INSERT INTO "stok_keluar_produk" ("skp_id", "skp_sk_id", "skp_prd_id", "skp_qty") VALUES
	(2, 2, 4, 1.00),
	(6, 6, 5, 12.00),
	(8, 8, 5, 4.00),
	(9, 11, 3, 1.00),
	(10, 11, 5, 1.00),
	(11, 12, 3, 2.00),
	(12, 12, 5, 2.00);
/*!40000 ALTER TABLE "stok_keluar_produk" ENABLE KEYS */;

-- Dumping structure for table inventaris.stok_masuk
DROP TABLE IF EXISTS "stok_masuk";
CREATE TABLE IF NOT EXISTS "stok_masuk" (
	"sm_id" INTEGER NOT NULL DEFAULT 'nextval(''stok_masuk_sm_id_seq''::regclass)',
	"sm_status" VARCHAR(20) NOT NULL,
	"sm_created_at" TIMESTAMP NULL DEFAULT 'CURRENT_TIMESTAMP',
	"sm_updated_at" TIMESTAMP NULL DEFAULT 'CURRENT_TIMESTAMP',
	"sm_supplier" VARCHAR(250) NULL DEFAULT 'NULL::character varying',
	PRIMARY KEY ("sm_id")
);

-- Dumping data for table inventaris.stok_masuk: -1 rows
/*!40000 ALTER TABLE "stok_masuk" DISABLE KEYS */;
INSERT INTO "stok_masuk" ("sm_id", "sm_status", "sm_created_at", "sm_updated_at", "sm_supplier") VALUES
	(14, 'DONE', '2026-04-01 20:40:08.002842', '2026-04-01 20:40:21.763863', 'djarum'),
	(13, 'DONE', '2026-04-01 20:39:34.433259', '2026-04-01 20:40:43.279244', 'sampurna'),
	(15, 'CREATED', '2026-04-01 20:51:27.667315', '2026-04-01 20:51:27.667315', 'string');
/*!40000 ALTER TABLE "stok_masuk" ENABLE KEYS */;

-- Dumping structure for table inventaris.stok_masuk_produk
DROP TABLE IF EXISTS "stok_masuk_produk";
CREATE TABLE IF NOT EXISTS "stok_masuk_produk" (
	"smp_id" INTEGER NOT NULL DEFAULT 'nextval(''stok_masuk_produk_smp_id_seq''::regclass)',
	"smp_sm_id" INTEGER NOT NULL,
	"smp_prd_id" INTEGER NOT NULL,
	"smp_qty" NUMERIC(12,2) NOT NULL,
	PRIMARY KEY ("smp_id"),
	CONSTRAINT "stok_masuk_produk_smp_prd_id_fkey" FOREIGN KEY ("smp_prd_id") REFERENCES "produk" ("prd_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT "stok_masuk_produk_smp_sm_id_fkey" FOREIGN KEY ("smp_sm_id") REFERENCES "stok_masuk" ("sm_id") ON UPDATE NO ACTION ON DELETE CASCADE
);

-- Dumping data for table inventaris.stok_masuk_produk: -1 rows
/*!40000 ALTER TABLE "stok_masuk_produk" DISABLE KEYS */;
INSERT INTO "stok_masuk_produk" ("smp_id", "smp_sm_id", "smp_prd_id", "smp_qty") VALUES
	(13, 13, 4, 3.00),
	(14, 13, 5, 10.00),
	(15, 13, 3, 11.00),
	(16, 14, 4, 4.00),
	(17, 14, 5, 15.00),
	(18, 14, 3, 16.00),
	(19, 15, 3, 1.00);
/*!40000 ALTER TABLE "stok_masuk_produk" ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
