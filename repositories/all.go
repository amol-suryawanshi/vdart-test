package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"vdart-test/config"
	"vdart-test/entity"
	httclient "vdart-test/httpclient"
)

const (
	currenciesURL = "%s/api/2/public/currency"
	symbolsURL    = "%s/api/2/public/symbol"
	tickersURL    = "%s/api/2/public/ticker"
)

//CreateAllSymbolImpl creates instance of AllSymbolImpl
func CreateAllSymbolImpl() *AllSymbolImpl {
	return &AllSymbolImpl{}
}

//AllSymbolImpl implements
type AllSymbolImpl struct {
}

//GetCurrencies get currency values
func (sr *AllSymbolImpl) GetCurrencies(ctx context.Context) (entity.Currencies, error) {
	var currencies entity.Currencies
	url := fmt.Sprintf(currenciesURL, config.ServerConfig.BaseURL)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	resp, err := httclient.CallRestAPI(http.MethodGet, url, headers, nil)
	if err != nil {
		return entity.Currencies{}, err
	}

	apiOutput, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.Currencies{}, err
	}
	err = json.Unmarshal(apiOutput, &currencies)
	if err != nil {
		return entity.Currencies{}, err
	}
	return currencies, nil
}

//GetSymbols - get the symbol values
func (sr *AllSymbolImpl) GetSymbols(ctx context.Context) (entity.Symbols, error) {
	var symbols entity.Symbols
	url := fmt.Sprintf(symbolsURL, config.ServerConfig.BaseURL)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	resp, err := httclient.CallRestAPI(http.MethodGet, url, headers, nil)
	if err != nil {
		return entity.Symbols{}, err
	}

	apiOutput, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.Symbols{}, err
	}
	err = json.Unmarshal(apiOutput, &symbols)
	if err != nil {
		return entity.Symbols{}, err
	}
	return symbols, nil
}

//GetTickers - ticker values
func (sr *AllSymbolImpl) GetTickers(ctx context.Context) (entity.Tickers, error) {
	var tickers entity.Tickers
	url := fmt.Sprintf(tickersURL, config.ServerConfig.BaseURL)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	resp, err := httclient.CallRestAPI(http.MethodGet, url, headers, nil)
	if err != nil {
		return entity.Tickers{}, err
	}

	apiOutput, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.Tickers{}, err
	}
	err = json.Unmarshal(apiOutput, &tickers)
	if err != nil {
		return entity.Tickers{}, err
	}
	return tickers, nil
}
