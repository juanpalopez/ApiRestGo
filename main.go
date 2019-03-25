package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// MarketProduct is struct for json
type MarketProduct struct {
	IDMarketProduct int     `json:"idMarketProduct,omitempty"`
	IDMarket        int     `json:"idMarket,omitempty"`
	Title           string  `json:"title,omitempty"`
	Price           float32 `json:"price,omitempty"`
	ListPrice       float32 `json:"list_price,omitempty"`
	ImgURL          string  `json:"img_url,omitempty"`
	URL             string  `json:"url,omitempty"`
	IsDiscounted    string  `json:"is_discounted,omitempty"`
	Ean             int64   `json:"ean,omitempty"`
	OfficialTitle   string  `json:"official_title,omitempty"`
	ImageCorrect    float64 `json:"image_correct,omitempty"`
}

var marketProducts []MarketProduct

func main() {
	marketProducts = append(marketProducts, MarketProduct{IDMarketProduct: 9,
		IDMarket:      1,
		Title:         "Coca Cola Light 600 Ml - Pack 6",
		Price:         169.9,
		ImgURL:        "https://www.disco.com.ar/DiscoComprasArchivos/Archivos/662463.jpg",
		URL:           "https://www.disco.com.ar/prod/231947/coca-cola-light-600-ml---pack-6",
		IsDiscounted:  "0",
		Ean:           7790895063541,
		OfficialTitle: "Gaseosa Coca-Cola Light - Sabor Liviano 600 ml Multipack x 6",
		ImageCorrect:  0.301430985172064})

	marketProducts = append(marketProducts, MarketProduct{IDMarketProduct: 10,
		IDMarket:      1,
		Title:         "Coca-Cola Sin Azúcar 354 Ml",
		Price:         35.9,
		ImgURL:        "https://www.disco.com.ar/DiscoComprasArchivos/Archivos/662464.jpg",
		URL:           "https://www.disco.com.ar/prod/264711/coca-cola-sin-azúcar-354-ml",
		IsDiscounted:  "0",
		Ean:           7790895067587,
		OfficialTitle: "Gaseosa Coca-Cola Sin Azúcar Lata 354 ml",
		ImageCorrect:  0.980977972726114})

	router := mux.NewRouter()
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))

}

// GetProducts get all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(marketProducts)
}

// GetProduct by id
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err == nil {
		errors.New("Couldnt cast string to integer")
	}

	for _, item := range marketProducts {

		if item.IDMarketProduct == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&MarketProduct{})
}

// CreateProduct a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {}

// DeleteProduct by id
func DeleteProduct(w http.ResponseWriter, r *http.Request) {}

// UpdateProduct by id
func UpdateProduct(w http.ResponseWriter, r *http.Request) {}
