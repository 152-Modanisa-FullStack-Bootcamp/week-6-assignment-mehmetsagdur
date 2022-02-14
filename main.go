package main

import (
	"bootcamp/config"
	"bootcamp/data"
	"bootcamp/handler"
	"bootcamp/service"
	"fmt"
	"net/http"
)

func main() {

	d := data.NewWallet()
	s := service.NewWalletService(d, config.C.MinimumBalanceAmount, config.C.InitialBalanceAmount)
	h := handler.NewHandler(s)

	http.HandleFunc("/", h.Wallets)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}

}
