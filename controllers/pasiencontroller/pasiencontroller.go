package pasiencontroller

import (
	"html/template"
	"net/http"

	"github.com/zoelabbb/go-crud-mysql/entities"
	"github.com/zoelabbb/go-crud-mysql/models"
)

var pasienModel = models.NewPasienModel() // Memanggil function pasien models

func Index(response http.ResponseWriter, request *http.Request) {

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)

}
func Add(response http.ResponseWriter, request *http.Request) {

	// Pengecekan request method pada form
	if request.Method == http.MethodGet {

		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		// Jika methodnya POST
		// Maka proses penyimpanan ke database
		request.ParseForm() //Melakukan parsing form

		var pasien entities.Pasien                              //Melakukan import struct entities di pasien.go
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")   //Mengambil inputan nama pada form
		pasien.NIK = request.Form.Get("nik")                    //Mengambil inputan nik pada form
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin") //Mengambil inputan jenis kelamin pada form
		pasien.TempatLahir = request.Form.Get("tempat_lahir")   //Mengambil inputan tempat lahir pada form
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir") //Mengambil inputan tanggal lahir pada form
		pasien.Alamat = request.Form.Get("alamat")              //Mengambil inputan alamat pada form
		pasien.NoHp = request.Form.Get("no_Hp")                 //Mengambil inputan no hp pada form

		// Memanggil fungsi pasien model
		pasienModel.Create(pasien)
		data := map[string]interface{}{
			"pesan": "Data Pasien Berhasil di Simpan !!",
		}

		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}

}
func Edit(response http.ResponseWriter, request *http.Request) {

}
func Delete(response http.ResponseWriter, request *http.Request) {

}
