package service

import (
	"bootcamp/data"
	"net/http"
)

type IWalletService interface {
	GetWalletsByID(string) (int, int)
	GetWallets() map[string]int
	WalletCreate(string) (map[string]int, int)
	WalletTransaction(string, int) int
}

type WalletService struct {
	data                 data.IWalletData
	InitialBalanceAmount int
	MinimumBalanceAmount int
}

func (w *WalletService) GetWalletsByID(username string) (int, int) {
	result, ok := w.data.GetWalletsById(username)

	if !ok {
		return 0, http.StatusInternalServerError
	}

	return result, http.StatusOK

}

func (w *WalletService) GetWallets() map[string]int {

	return w.data.GetWallets()
}

func (w *WalletService) WalletCreate(username string) (map[string]int, int) {

	if _, ok := w.data.GetWalletsById(username); !ok && username != "" {
		newWallet := w.data.CreateWallet(username, w.InitialBalanceAmount)
		return newWallet, http.StatusOK
	}

	return nil, http.StatusOK
}

func (w *WalletService) WalletTransaction(username string, money int) int {
	wallets := w.data.GetWallets()
	oldBalance, ok := wallets[username]
	if ok != true {
		return http.StatusInternalServerError
	}

	newBalance := oldBalance + money
	if newBalance >= w.MinimumBalanceAmount {
		w.data.CreateWallet(username, newBalance)
		return http.StatusOK
	}

	return http.StatusInternalServerError
}

func NewWalletService(walletData data.IWalletData, MinimumBalance, InitialBalance int) IWalletService {
	return &WalletService{data: walletData, MinimumBalanceAmount: MinimumBalance, InitialBalanceAmount: InitialBalance}
}
