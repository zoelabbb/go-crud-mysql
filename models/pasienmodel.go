package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/zoelabbb/go-crud-mysql/config"
	"github.com/zoelabbb/go-crud-mysql/entities"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PasienModel{
		conn: conn,
	}
}

func (p *PasienModel) Create(pasien entities.Pasien) bool {
	// Syntax SQL DB
	result, err := p.conn.Exec("insert into pasiens (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values(?,?,?,?,?,?,?)",
		pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp)

	if err != nil {
		fmt.Println(err)
		return false
	}
	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0
}

// Mengambil semua data pasien
func (p *PasienModel) FindAll() ([]entities.Pasien, error) {
	rows, err := p.conn.Query("select * from pasiens") // Syntax SQL (table pasiens)
	if err != nil {
		return []entities.Pasien{}, err
	}

	defer rows.Close()

	var dataPasien []entities.Pasien

	// Tabel DB
	for rows.Next() {
		var pasiens entities.Pasien
		rows.Scan(&pasiens.Id, &pasiens.NamaLengkap, &pasiens.NIK, &pasiens.JenisKelamin, &pasiens.TempatLahir, &pasiens.TanggalLahir, &pasiens.Alamat, &pasiens.NoHp)

		if pasiens.JenisKelamin == "1" {
			pasiens.JenisKelamin = "Laki-laki"
		} else {
			pasiens.JenisKelamin = "Perempuan"
		}

		tgl_lahir, _ := time.Parse("2006-01-02", pasiens.TanggalLahir)
		pasiens.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataPasien = append(dataPasien, pasiens)
	}

	return dataPasien, nil
}
