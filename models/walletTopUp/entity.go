package walletTopUp

type WalletTopUp struct {
	ID         int    `form:"id" json:"id"`
	KodeYukPay string `form:"kodeYukPay" json:"kodeYukPay"`
	Metode     string `form:"metode" json:"metode"`
	Nominal    int    `form:"nominal" json:"nominal"`
}

type WalletTopUpResponse struct {
	Message string        `form:"message" json:"message"`
	Data    []WalletTopUp `form:"data" json:"data"`
}
