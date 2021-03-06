package models

import (
// "regexp"
)

type Currency struct {
	CommonModel
	Key         string `json:"key"`
	Code        string `json:"code"`
	Symbol      string `json:"-"`
	Coin        bool   `json:"coin"`
	Precision   int    `json:"precision"`
	Erc20       bool   `json:"erc20"`
	Erc23       bool   `json:"erc23"`
	Visible     bool   `json:"visible"`
	Tradable    bool   `json:"tradable"`
	Depositable bool   `json:"depositable"`
}

func (currency *Currency) IsEthereum() (result bool) {
	if currency.Code == "eth" || currency.Erc20 || currency.Erc23 {
		result = true
	}
	return
}
