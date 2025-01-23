package model

type PairEvent struct {
	ContractEvent
	PairAddress      string  `json:"pairAddress"`
	BaseTokenAddress string  `json:"baseTokenAddress"`
	MemeTokenAddress string  `json:"memeTokenAddress"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Decimals         int64   `json:"decimals"`
	TotalSupply      float64 `json:"totalSupply"`
	IpfsLink         string  `json:"ipfsLink"`
	CreatorWallet    string  `json:"creatorWallet"`
	AgentWallet      string  `json:"agentWallet"`
	InitialLiquidity float64 `json:"initialLiquidity"`
	MaxLiquidity     float64 `json:"maxLiquidity"`
}
