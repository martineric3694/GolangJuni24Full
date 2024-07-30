package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("File .env tidak terbaca")
	}

	// Konfigurasi koneksi database
	connStr := "user=postgres dbname=trialdb password=postgre host=" + os.Getenv("DB_HOST") + " port=5432 sslmode=disable"

	// Membuka koneksi ke database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Memeriksa koneksi
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Koneksi ke database berhasil!")

	nama := "budi"
	umur := 15

	// Memanggil stored procedure
	// Cara buat procedure insert di postgres
	/*
		CREATE OR REPLACE PROCEDURE public.masukan_data(IN nama_orang character varying, IN umur_orang numeric)
		LANGUAGE plpgsql
		AS $procedure$
		BEGIN
			INSERT INTO data(name,umur) VALUES (nama_orang,umur_orang);
		END;
		$procedure$
	*/
	_, err = db.Exec("CALL masukan_data($1, $2)", nama, umur)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Stored procedure berhasil dipanggil!")

	// Memanggil stored procedure dan menangani hasilnya
	// Cara buat procedure select postgres
	/*
		CREATE OR REPLACE FUNCTION public.kembali_data()
		RETURNS TABLE(id_data integer, nama character varying, umur_data numeric)
		LANGUAGE plpgsql
		AS $function$
		BEGIN
			RETURN QUERY SELECT id, name, umur FROM data;
		END;
		$function$
	*/

	rows, err := db.Query("SELECT id_data,nama,umur_data FROM kembali_data()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id int
		var nama string
		var umur float64

		data := map[string]interface{}{}

		err := rows.Scan(&id, &nama, &umur)
		if err != nil {
			log.Fatal(err)
		}
		data["id"] = id
		data["nama"] = nama
		data["umur"] = umur
		result = append(result, data)

	}

	fmt.Println(result)
}
