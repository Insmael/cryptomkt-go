package conn

type Pagination struct {
	Previous interface{}
	Next     interface{}
	Limit    int
	Page     int
}

type MarketListResponse struct {
	Status  string
	Message string
	Data    []string
}

type TickerResponse struct {
	Status  string
	Message string
	Data    []Ticker
}

type Ticker struct {
	High      string
	Volume    string
	Low       string
	Ask       string
	Timestamp string
	Bid       string
	LastPrice string `json:"last_price"`
	Market    string
}
type TickerData struct {
	Data []Ticker
}

type BalancesResponse struct {
	Status  string
	Message string
	Data    []Balance
}

type Balance struct {
	Wallet    string
	Available string
	Balance   string
}

type AccountResponse struct {
	Status  string
	Message string
	Data    Account
}

type Account struct {
	Name         string
	Email        string
	Rate         Rate
	BankAccounts []BankAccount `json:"bank_accounts"`
}

type Rate struct {
	MarketMaker string `json:"market_maker"`
	MarketTaker string `json:"market_taker"`
}

type BankAccount struct {
	Id          int
	Bank        string
	Description string
	Country     string
	Number      string
}

type InstantResponse struct {
	Status  string
	Message string
	Data    Quantity
}

type Quantity struct {
	Obtained string
	Required string
}

type TransactionsResponse struct {
	client     *Client
	Status     string
	Message    string
	Pagination Pagination
	Data       []Transaction
}

type Transaction struct {
	Id         string
	Type       int
	Amount     string
	FeePercent string `json:"fee_percent"`
	FeeAmount  string `json:"fee_amount"`
	Balance    string
	Date       string
	Hash       string
	Address    string
	Memo       string
}
