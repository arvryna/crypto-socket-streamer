package fetcher

type Quote struct {
	Symbol string
	Ask    float64
	Bid    float64
}

type Trade struct {
	Symbol string
	Price  float64
	Size   float64
}

type Bar struct {
	Open  float64 `json:"o"`
	Close float64 `json:"c"`
	High  float64 `json:"h"`
	Low   float64 `json:"l"`
}

type marketData struct {
	Type   string  `json:"T"`
	Symbol string  `json:"S"`
	Price  float64 `json:"p"`
	Size   float64 `json:"s"`
	Ask    float64 `json:"ap"`
	Bid    float64 `json:"bp"`
	Time   string  `json:"t"`
	Open   float64 `json:"o"`
	Close  float64 `json:"c"`
	High   float64 `json:"h"`
	Low    float64 `json:"l"`
}
