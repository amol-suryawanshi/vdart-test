package entity

//Currencies array
type Currencies []Currency

//Symbols array
type Symbols []Symbol

//Tickers array
type Tickers []Ticker

//AllSymbolResponse API Response
type AllSymbolResponse struct {
	Currencies []SymbolResponse `json:"curremcies"`
}
