package pasiencontroller

import (
	"html/template"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)

}
func Add(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/pasien/add.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}
func Edit(response http.ResponseWriter, request *http.Request) {

}
func Delete(response http.ResponseWriter, request *http.Request) {

}
