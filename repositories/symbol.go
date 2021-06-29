package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"vdart-test/config"
	"vdart-test/entity"
	httclient "vdart-test/httpclient"
)

const (
	currencyURL = "%s/api/2/public/currency/%s"
	symbolURL   = "%s/api/2/public/symbol/%s"
	tickerURL   = "%s/api/2/public/ticker/%s"
)

//CreateSymbolReader creates OAuthReader
func CreateSymbolReader() *SymbolReader {
	return &SymbolReader{}
}

//SymbolReader implementation
type SymbolReader struct{}

//GetCurrency get currency values
func (sr *SymbolReader) GetCurrency(ctx context.Context) (entity.Currency, error) {
	var currency entity.Currency
	curr, ok := ctx.Value("currency").(string)
	if !ok || curr == "" {
		return entity.Currency{}, errors.New("currency not found")
	}
	url := fmt.Sprintf(currencyURL, config.ServerConfig.BaseURL, curr)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	resp, err := httclient.CallRestAPI(http.MethodGet, url, headers, nil)
	if err != nil {
		return entity.Currency{}, err
	}

	apiOutput, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.Currency{}, err
	}
	err = json.Unmarshal(apiOutput, &currency)
	if err != nil {
		return entity.Currency{}, err
	}
	return currency, nil
}

//GetSymbol - get the symbol values
func (sr *SymbolReader) GetSymbol(ctx context.Context) (entity.Symbol, error) {
	var symbol entity.Symbol
	symb, ok := ctx.Value("symbol").(string)
	if !ok || symb == "" {
		return entity.Symbol{}, errors.New("symbol not found")
	}
	url := fmt.Sprintf(symbolURL, config.ServerConfig.BaseURL, symb)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	resp, err := httclient.CallRestAPI(http.MethodGet, url, headers, nil)
	if err != nil {
		return entity.Symbol{}, err
	}

	apiOutput, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.Symbol{}, err
	}
	err = json.Unmarshal(apiOutput, &symbol)
	if err != nil {
		return entity.Symbol{}, err
	}
	return symbol, nil
}

//GetTicker - ticker values
func (sr *SymbolReader) GetTicker(ctx context.Context) (entity.Ticker, error) {
	var ticker entity.Ticker
	symb, ok := ctx.Value("symbol").(string)
	if !ok || symb == "" {
		return entity.Ticker{}, errors.New("symbol not found")
	}
	url := fmt.Sprintf(tickerURL, config.ServerConfig.BaseURL, symb)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	resp, err := httclient.CallRestAPI(http.MethodGet, url, headers, nil)
	if err != nil {
		return entity.Ticker{}, err
	}

	apiOutput, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.Ticker{}, err
	}
	err = json.Unmarshal(apiOutput, &ticker)
	if err != nil {
		return entity.Ticker{}, err
	}
	return ticker, nil
}
