package data

type IWalletData interface {
	GetWalletsById(string) (int, bool)
	GetWallets() map[string]int
	CreateWallet(string, int) map[string]int
}

type WalletData struct {
	WalletStorage map[string]int
}

func (k WalletData) GetWalletsById(key string) (int, bool) {
	result, ok := k.WalletStorage[key]
	return result, ok
}

func (k WalletData) GetWallets() map[string]int {
	result := k.WalletStorage
	return result
}

func (k WalletData) CreateWallet(key string, value int) map[string]int {
	k.WalletStorage[key] = value
	return k.WalletStorage
}

func NewWallet() IWalletData {
	return &WalletData{WalletStorage: map[string]int{}}
}
