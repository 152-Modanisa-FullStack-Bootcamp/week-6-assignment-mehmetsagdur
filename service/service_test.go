package service

import (
	"bootcamp/config"
	"bootcamp/data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestWalletService_GetWallets(t *testing.T) {
	t.Run("username empty get all wallets", func(t *testing.T) {
		dataReturn := map[string]int{"test1": config.C.InitialBalanceAmount}

		data := data.NewWallet()
		srv := NewWalletService(data, config.C.MinimumBalanceAmount, config.C.InitialBalanceAmount)
		srv.WalletCreate("test1")
		wallets := srv.GetWallets()

		assert.Equal(t, dataReturn, wallets)

	})

}

func TestWalletService_GetWalletsByID(t *testing.T) {
	data := data.NewWallet()
	srv := NewWalletService(data, config.C.MinimumBalanceAmount, config.C.InitialBalanceAmount)
	srv.WalletCreate("test1")
	srv.WalletCreate("test2")
	srv.WalletCreate("test3")

	t.Run("get wallets by username", func(t *testing.T) {

		wallets, status := srv.GetWalletsByID("test1")

		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, config.C.InitialBalanceAmount, wallets)

	})

	t.Run("calling wallet with non-username", func(t *testing.T) {

		wallets, status := srv.GetWalletsByID("test5")

		assert.Equal(t, 0, wallets)
		assert.Equal(t, http.StatusInternalServerError, status)

	})

}

func TestWalletService_WalletCreate(t *testing.T) {
	t.Run("create wallet with username", func(t *testing.T) {
		data := data.NewWallet()
		srv := NewWalletService(data, config.C.MinimumBalanceAmount, config.C.InitialBalanceAmount)
		newWallet, status := srv.WalletCreate("test1")

		assert.Equal(t, config.C.InitialBalanceAmount, newWallet["test1"])
		assert.Equal(t, http.StatusOK, status)

	})
	t.Run("create wallet with non-username", func(t *testing.T) {
		data := data.NewWallet()
		srv := NewWalletService(data, config.C.MinimumBalanceAmount, config.C.InitialBalanceAmount)
		newWallet, status := srv.WalletCreate("")

		assert.Nil(t, nil, newWallet)
		assert.Equal(t, http.StatusOK, status)

	})
	t.Run("create wallet with already have username", func(t *testing.T) {
		data := data.NewWallet()
		srv := NewWalletService(data, config.C.MinimumBalanceAmount, config.C.InitialBalanceAmount)
		_, _ = srv.WalletCreate("test1")

		newWallet, status := srv.WalletCreate("test1")

		assert.Nil(t, nil, newWallet)
		assert.Equal(t, http.StatusOK, status)

	})
}

func TestWalletService_WalletTransaction(t *testing.T) {
	t.Run("balance update non-username", func(t *testing.T) {
		data := data.NewWallet()
		srv := NewWalletService(data, config.C.MinimumBalanceAmount, config.C.InitialBalanceAmount)
		status := srv.WalletTransaction("test1", 50)

		assert.Equal(t, http.StatusInternalServerError, status)

	})
	t.Run("balance update with have username", func(t *testing.T) {
		data := data.NewWallet()
		srv := NewWalletService(data, config.C.MinimumBalanceAmount, config.C.InitialBalanceAmount)
		_, _ = srv.WalletCreate("test1")

		status := srv.WalletTransaction("test1", 50)

		assert.Equal(t, http.StatusOK, status)

	})
	t.Run("balance update over balance withdraw limit", func(t *testing.T) {
		data := data.NewWallet()
		srv := NewWalletService(data, config.C.MinimumBalanceAmount, config.C.InitialBalanceAmount)
		_, _ = srv.WalletCreate("test1")

		status := srv.WalletTransaction("test1", -200)

		assert.Equal(t, http.StatusInternalServerError, status)

	})

}
