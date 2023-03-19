package pasiencontroller

import (
	"html/template"
	"net/http"

	"github.com/zoelabbb/go-crud-mysql/entities"
	"github.com/zoelabbb/go-crud-mysql/libraries"
	"github.com/zoelabbb/go-crud-mysql/models"
)

var validation = libraries.NewValidation() // Memanggil library validation
var pasienModel = models.NewPasienModel()  // Memanggil function pasien models

func Index(response http.ResponseWriter, request *http.Request) {

	// Memanggil FindAll() pada pasienmodel.go
	pasien, _ := pasienModel.FindAll()

	data := map[string]interface{}{
		"pasien": pasien,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
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

		var data = make(map[string]interface{})

		// Melakukan validasi terlebih dahulu
		// Sebelum menambahkan data
		vErrors := validation.Struct(pasien) // Mengambil struct pasien & memvalidasi inputan data yang ada di dalamnya
		if vErrors != nil {
			data["pasien"] = pasien // Mengembalikan data agar tidak hilang
			data["validations"] = vErrors
		} else {
			data["pesan"] = "Data pasien berhasil disimpan !!"
			// Melakukan insert data ke DB
			pasienModel.Create(pasien)
		}

		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}

}
func Edit(response http.ResponseWriter, request *http.Request) {

}
func Delete(response http.ResponseWriter, request *http.Request) {

}
