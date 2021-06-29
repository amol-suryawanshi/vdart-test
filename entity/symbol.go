package entity

//Currency - currency respone
type Currency struct {
	ID                    string `json:"id"`                  //"id": "ETH",
	FullName              string `json:"fullName"`            //"fullName": "Ethereum",
	Crypto                bool   `json:"crypto"`              //"crypto": true,
	PayInEnabled          bool   `json:"payinEnabled"`        //"payinEnabled": true,
	PayinPaymentID        bool   `json:"payinPaymentId"`      //"payinPaymentId": false,
	PayinConfirmations    int    `json:"payinConfirmations"`  //"payinConfirmations": 20,
	PayoutEnabled         bool   `json:"payoutEnabled"`       //"payoutEnabled": true,
	PayoutIsPaymentID     bool   `json:"payoutIsPaymentId"`   //"payoutIsPaymentId": false,
	TransferEnabled       bool   `json:"transferEnabled"`     //"transferEnabled": true,
	Delisted              bool   `json:"delisted"`            //"delisted": false,
	PayoutFee             string `json:"payoutFee"`           //"payoutFee": "0.042800000000",
	PayoutMinimalAmount   string `json:"payoutMinimalAmount"` //"payoutMinimalAmount": "0.00958",
	PrecisionPayout       int    `json:"precisionPayout"`     //"precisionPayout": 10,
	PrecisionTransfer     int    `json:"precisionTransfer"`   //"precisionTransfer": 10,
	LowProcessingTime     string `json:"lowProcessingTime"`   //"lowProcessingTime": "3.518",
	HighProcessingTime    string `json:"highProcessingTime"`  //"highProcessingTime": "2411.963",
	AverageProcessingTime string `json:"avgProcessingTime"`   //"avgProcessingTime": "247.1108947"
}

//Symbol - symbol response
type Symbol struct {
	ID                   string `json:"id"`                   //"id": "ETHBTC",
	BaseCurrency         string `json:"baseCurrency"`         //"baseCurrency": "ETH",
	QuoteCurrency        string `json:"quoteCurrency"`        //"quoteCurrency": "BTC",
	QuantityIncrement    string `json:"quantityIncrement"`    //"quantityIncrement": "0.0001",
	TickSize             string `json:"tickSize"`             //"tickSize": "0.000001",
	TakeLiquidityRate    string `json:"takeLiquidityRate"`    //"takeLiquidityRate": "0.002",
	ProvideLiquidityRate string `json:"provideLiquidityRate"` //"provideLiquidityRate": "0.001",
	FreeCurrency         string `json:"feeCurrency"`          //"feeCurrency": "BTC",
	MarginTrading        bool   `json:"marginTrading"`        //"marginTrading": true,
	MaxInitialLeverage   string `json:"maxInitialLeverage"`   //"maxInitialLeverage": "10.00"
}

//Ticker - ticker response
type Ticker struct {
	Symbol      string `json:"symbol"`      //"symbol": "ETHBTC",
	Ask         string `json:"ask"`         //"ask": "0.020572",
	Bid         string `json:"bid"`         //"bid": "0.020566",
	Last        string `json:"last"`        //"last": "0.020574",
	Open        string `json:"open"`        //"open": "0.020913",
	Low         string `json:"low"`         //"low": "0.020388",
	High        string `json:"high"`        //"high": "0.021084",
	Volume      string `json:"volume"`      //"volume": "138444.3666",
	VolumeQuote string `json:"volumeQuote"` //"volumeQuote": "2853.6874972480",
	Timestamp   string `json:"timestamp"`   //"timestamp": "2020-04-02T17:52:36.731Z"
}

//SymbolResponse - response format
type SymbolResponse struct {
	ID           string `json:"id"`          //"id": "ETH",
	FullName     string `json:"fullName"`    //"fullName": "Ethereum",
	Ask          string `json:"ask"`         //"ask": "0.020572",
	Bid          string `json:"bid"`         //"bid": "0.020566",
	Last         string `json:"last"`        //"last": "0.020574",
	Open         string `json:"open"`        //"open": "0.020913",
	Low          string `json:"low"`         //"low": "0.020388",
	High         string `json:"high"`        //"high": "0.021084",
	FreeCurrency string `json:"feeCurrency"` //"feeCurrency": "BTC",
}
