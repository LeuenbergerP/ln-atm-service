package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	BaseUrl  string
	ApiKey   string
	DeviceId string
}

func main() {
	a := App{}
	a.Initialize()
	a.Run(":8010")
}

func (a *App) Initialize() {
	apiKey := os.Getenv("OPENNODE_API_KEY")
	baseUrl := os.Getenv("OPENNODE_BASE_URL")
	devideId := os.Getenv("OPENNODE_DEVICE_ID")
	if baseUrl == "" || apiKey == "" {
		log.Panic("OPENNODE_BASE_URL and OPENNODE_API_KEY must be set")
	}
	if devideId == "" {
		uuid, err := uuid.NewUUID()
		if err != nil {
			log.Panic(err)
		}
		log.Println("App UUID: " + uuid.String())
		devideId = uuid.String()
	}
	log.Println("OPENNODE_BASE_URL: " + baseUrl)
	a.BaseUrl = baseUrl
	a.ApiKey = apiKey
	a.DeviceId = devideId
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) chargeInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Charge ID")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]int{"Id": id})
}

func (a *App) chargeHandler(w http.ResponseWriter, r *http.Request) {
	var cr ChargeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cr); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	c, err := cr.CreateCharge(a.chargeUrl(), a.ApiKey)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating charge: %s", err.Error()))
		return
	}
	fmt.Printf("PayRequest: %s", c.Data.LightningInvoice.PayReq)
	respondWithJSON(w, http.StatusOK, c)
}

// func (a *App) rateHandler(w http.ResponseWriter, r *http.Request) {
// 	res, err := http.Get(a.rateUrl())
// 	if err != nil {
// 		log.Fatal(err)
// 		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting rates: %s", err.Error()))
// 		return
// 	}
// 	defer res.Body.Close()
// 	var rates Rates
// 	if err := json.NewDecoder(res.Body).Decode(&rates); err != nil {
// 		log.Fatal(err)
// 		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding rates: %s", err.Error()))
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, rates)
// }

func (cr *ConcurencyRate) RateForCurrency(url string) (Rate, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var rates Rates
	if err := json.NewDecoder(res.Body).Decode(&rates); err != nil {
		return Rate{}, err
	}
	return rates.Data[cr.Symbol], nil
}

func (c *ChargeRequest) CreateCharge(url string, apiKey string) (ChargeResponse, error) {
	jsonBody, err := json.Marshal(c)
	if err != nil {
		return ChargeResponse{}, err
	}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", apiKey)
	if err != nil {
		return ChargeResponse{}, err
	}
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
		return ChargeResponse{}, err
	}
	defer res.Body.Close()

	// b, err := io.ReadAll(res.Body)
	// // b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(string(b))
	var charge ChargeResponse
	if err := json.NewDecoder(res.Body).Decode(&charge); err != nil {
		return ChargeResponse{}, err
	}
	return charge, nil
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/charge/{id}", a.chargeInfoHandler).Methods(http.MethodGet)
	a.Router.HandleFunc("/charge/", a.chargeHandler).Methods(http.MethodOptions, "POST")
}

func (a *App) rateUrl() string {
	return fmt.Sprintf("%s/rates", a.BaseUrl)
}

func (a *App) chargeUrl() string {
	return fmt.Sprintf("%s/charges", a.BaseUrl)
}
