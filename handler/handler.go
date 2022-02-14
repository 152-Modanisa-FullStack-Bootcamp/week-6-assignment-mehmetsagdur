package handler

import (
	"bootcamp/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type IHandler interface {
	Wallets(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IWalletService
}

func (h *Handler) Wallets(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	username := strings.ReplaceAll(url, "/", "")

	if r.Method == http.MethodGet {
		if username != "" {
			result, status := h.service.GetWalletsByID(username)

			jsonData, _ := json.Marshal(result)

			//if err != nil {
			//	w.WriteHeader(http.StatusInternalServerError)
			//	w.Write([]byte(err.Error()))
			//	return
			//}

			w.Header().Add("content-type", "application/json")
			w.WriteHeader(status)
			w.Write(jsonData)
		} else {
			result := h.service.GetWallets()

			jsonData, _ := json.Marshal(result)

			//if err != nil {
			//	w.WriteHeader(http.StatusInternalServerError)
			//	w.Write([]byte(err.Error()))
			//}

			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		}

	} else if r.Method == http.MethodPut {

		_, status := h.service.WalletCreate(username)
		w.WriteHeader(status)

	} else if r.Method == http.MethodPost {

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		c := make(map[string]int)
		err = json.Unmarshal(b, &c)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		balance, ok := c["balance"]
		if ok != true {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		status := h.service.WalletTransaction(username, balance)
		w.WriteHeader(status)

	} else {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

}

func NewHandler(service service.IWalletService) IHandler {
	return &Handler{service: service}
}
