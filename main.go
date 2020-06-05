package main

import (
	"database/sql"
	"encoding/json"
	"golangRestApi/database"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

var databaseConnection *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}


func catch(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {

	databaseConnection = database.InitDB()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/products",AllProductos)

	http.ListenAndServe(":3000", r)

	defer databaseConnection.Close() // Con el defer la conexión se cerrará cuando finalice la función contenedora


}

func AllProductos(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description,'')
				FROM products`
	results,err := databaseConnection.Query(sql)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description) // Mapeamos en el slice
		catch(err)
		products = append(products, product)
	}
	respondwithJSON(w,http.StatusOK,products)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
