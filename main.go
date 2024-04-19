package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Veri yapısı
type Data struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Timestamp int64         `json:"timestamp"`
	Price     float64       `json:"price"`
}

var cache = make(map[string]Data)

func main() {
	// MongoDB bağlantısı
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// MongoDB veritabanı ve koleksiyon seçimi
	db := session.DB("crypto_data")
	collection := db.C("binance_prices")

	// Yönlendirici oluşturma
	router := mux.NewRouter()

	// API endpointi
	router.HandleFunc("/price", func(w http.ResponseWriter, r *http.Request) {
		// Eğer cache'te veri varsa ve veri 59 saniyeden daha kısa süre önce alınmışsa, önbellekten veriyi döndür
		cachedData, ok := cache["price"]
		if ok && time.Now().Unix()-cachedData.Timestamp < 59 {
			jsonResponse(w, http.StatusOK, cachedData)
			return
		}

		// MongoDB'den veriyi al
		var result Data
		err := collection.Find(nil).Sort("-timestamp").One(&result)
		if err != nil {
			log.Println("MongoDB'den veri alınamadı:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Önbelleğe al
		cache["price"] = result

		// JSON yanıtını gönder
		jsonResponse(w, http.StatusOK, result)
	}).Methods("GET")

	// HTTP sunucusu başlatma
	log.Fatal(http.ListenAndServe(":8080", router))
}

// JSON yanıtı gönderme
func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
