package server

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/hightemp/ip_info_service/internal/config"
	"github.com/hightemp/ip_info_service/internal/logger"
	"github.com/hightemp/ip_info_service/internal/models/ip_range"
	"github.com/hightemp/ip_info_service/internal/utils"
)

func lookupHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")

	if ip == "" {
		http.Error(w, "IP is required", http.StatusBadRequest)
		return
	}

	if net.ParseIP(ip) == nil {
		http.Error(w, "Invalid IP address", http.StatusBadRequest)
		return
	}

	country, organization := ip_range.SearchIpInfo(ip)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"country": country, "organization": organization})
}

type AddCountryRequest struct {
	IpStart string `json:"ip_start"`
	IpEnd   string `json:"ip_end"`
	Name    string `json:"name"`
}

func addCountryHandler(w http.ResponseWriter, r *http.Request) {
	var data AddCountryRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ip_range.AddCountry(data.Name, utils.IpStringToInt(data.IpStart), utils.IpStringToInt(data.IpEnd))
}

func addOrganizationHandler(w http.ResponseWriter, r *http.Request) {

}

func Start() {
	http.HandleFunc("/lookup", lookupHandler)
	http.HandleFunc("/ranges/add/country", addCountryHandler)
	http.HandleFunc("/ranges/add/organization", addOrganizationHandler)

	logger.LogInfo("Starting server on port %s", config.Get().Port)
	if err := http.ListenAndServe(":"+config.Get().Port, nil); err != nil {
		logger.Panic("Can't start server: %v", err)
	}
}
