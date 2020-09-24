package api

import (
	"encoding/json"
	"log"
	"net/http"
	http2 "rt/internal/rt/delivery/http"
	"rt/models"

	"github.com/gorilla/mux"
)

func productGet(service http2.RT) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading product"
		res, err := service.GetProduct()

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			log.Println("Product not recieved:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if res == nil {
			log.Println("Product data is empty:", err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println("Marshalling error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func productFind(service http2.RT) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error former offer"

		var conditions []models.Condition

		err := json.NewDecoder(r.Body).Decode(&conditions)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		product, err := service.GetProduct()

		offer, err := service.GetOffer(product, conditions)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			log.Println("Product not recieved:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if offer == nil {
			log.Println("Product data is empty:", err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(offer); err != nil {
			log.Println("Marshalling error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

//MakeProductHandlers make url handlers
func MakeProductHandlers(r *mux.Router, ch chan bool, service http2.RT) {

	r.Handle("/v1/product/get", productGet(service)).Methods("GET").Name("productGet")

	r.Handle("/v1/product/find", productFind(service)).Methods("POST").Name("productFind")

	r.HandleFunc("/v1/api/shutdown", shutdown(ch))
}

func shutdown(shutdownReq chan bool) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Shutdown server"))
		go func() {
			shutdownReq <- true
		}()
	}

}
