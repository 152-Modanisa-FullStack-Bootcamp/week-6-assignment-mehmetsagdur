package data

import (
	"bootcamp/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalletData_CreateWallet(t *testing.T) {
	d := NewWallet()
	result := d.CreateWallet("test1", config.C.InitialBalanceAmount)

	assert.Equal(t, config.C.InitialBalanceAmount, result["test1"])

}
func TestWalletData_GetWallets(t *testing.T) {
	d := NewWallet()
	result := d.CreateWallet("test1", config.C.InitialBalanceAmount)
	wallets := d.GetWallets()
	assert.Equal(t, result, wallets)
}
func TestWalletData_GetWalletsById(t *testing.T) {
	d := NewWallet()
	_ = d.CreateWallet("test1", config.C.InitialBalanceAmount)
	wallets, ok := d.GetWalletsById("test1")
	assert.Equal(t, config.C.InitialBalanceAmount, wallets)
	assert.Equal(t, true, ok)

}
