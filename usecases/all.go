package usecases

import (
	"context"
	"fmt"
	"vdart-test/config"
	"vdart-test/entity"
)

//CreateAllSymbolUC creates instance of AllSymbol
func CreateAllSymbolUC(reader AllReader) *AllSymbolUC {
	return &AllSymbolUC{
		reader: reader,
	}
}

//Reader to fetch symbol related values
type AllReader interface {
	GetCurrencies(ctx context.Context) (entity.Currencies, error)
	GetSymbols(ctx context.Context) (entity.Symbols, error)
	GetTickers(ctx context.Context) (entity.Tickers, error)
}

//AuthAccessToken implements GithubAction interface in handler package
type AllSymbolUC struct {
	reader AllReader
}

//Redirect creates a redirection url
func (ar *AllSymbolUC) FetchAllSymbolValue(ctx context.Context) (entity.AllSymbolResponse, error) {
	currencies, err := ar.reader.GetCurrencies(ctx)
	if err != nil {
		fmt.Println("error while fetching currencies : ", err)
	}
	symbols, err := ar.reader.GetSymbols(ctx)
	if err != nil {
		fmt.Println("error while fetching symbols : ", err)
	}
	tickers, err := ar.reader.GetTickers(ctx)
	if err != nil {
		fmt.Println("error while fetching tickers : ", err)
	}
	resp, err := buildResponse(currencies, symbols, tickers)
	return resp, err
}

func buildResponse(currencies entity.Currencies, symbols entity.Symbols, tickers entity.Tickers) (entity.AllSymbolResponse, error) {
	response := make([]entity.SymbolResponse, 0)

	for _, cur := range config.ServerConfig.Currencies {
		curr := cur.Currency
		symb := cur.Symbol
		symbresp := entity.SymbolResponse{}

		for i := range currencies {
			if curr == currencies[i].ID {
				symbresp.ID = symb
				symbresp.FullName = currencies[i].FullName
			}
		}

		for i := range symbols {
			if symb == symbols[i].ID {
				symbresp.FreeCurrency = symbols[i].FreeCurrency
			}
		}

		for i := range tickers {
			if symb == tickers[i].Symbol {
				symbresp.Ask = tickers[i].Ask
				symbresp.Bid = tickers[i].Bid
				symbresp.High = tickers[i].High
				symbresp.Last = tickers[i].Last
				symbresp.Low = tickers[i].Low
				symbresp.Open = tickers[i].Open
			}
		}

		response = append(response, symbresp)
	}

	return entity.AllSymbolResponse{Currencies: response}, nil
}
