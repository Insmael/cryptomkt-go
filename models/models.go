package models

// SideType is the side of the order or trade
type SideType string

type SortByType string

type SortType string

// OrderStatus is the status of an order in the exchange
type OrderStatus string

// OrderType is the type of an order
type OrderType string

// TimeInForceType is the TimeInForce of an order
type TimeInForceType string

// TransactionStatus is the status of transaction
type TransactionStatus string

// TransactionType is the type of a transaction
type TransactionType string

// TransactionSubType is the sub type of a transaction
type TransactionSubType string

type AccountType string

type TransferBy string

type UseOffchainType string

// ReportType shows the type of report
type ReportType string

const (
	SideSell SideType = "sell"
	SideBuy  SideType = "buy"

	SortByDate SortByType = "created_at"
	SortByID   SortByType = "id"

	SortASC  SortType = "ASC"
	SortDESC SortType = "DESC"

	OrderStatusNew             OrderStatus = "new"
	OrderStatusSuspended       OrderStatus = "suspended"
	OrderStatusPartiallyFilled OrderStatus = "partiallyFilled"
	OrderStatusFilled          OrderStatus = "filled"
	OrderStatusCanceled        OrderStatus = "canceled"
	OrderStatusExpired         OrderStatus = "expired"

	OrderLimit            OrderType = "limit"
	OrderMarket           OrderType = "market"
	OrderStopLimit        OrderType = "stopLimit"
	OrderStopMarket       OrderType = "stopMarket"
	OrdertakeProfitLimit  OrderType = "takeProfitMarket"
	OrdertakeProfitMarket OrderType = "takeProfitMarket"

	TimeInForceGTC TimeInForceType = "GTC"
	TimeInForceIOC TimeInForceType = "IOC"
	TimeInForceFOK TimeInForceType = "FOK"
	TimeInForceDAY TimeInForceType = "DAY"
	TimeInForceGTD TimeInForceType = "GTD"

	TransactionStatusCreated    TransactionStatus = "CREATED"
	TransactionStatusPending    TransactionStatus = "PENDING"
	TransactionStatusFailed     TransactionStatus = "FAILED"
	TransactionStatusSuccess    TransactionStatus = "SUCCESS"
	TransactionStatusRolledBack TransactionStatus = "ROLLED_BACK"

	TransactionTypeDeposit  TransactionType = "DEPOSIT"
	TransactionTypeWithdraw TransactionType = "WITHDRAW"
	TransactionTypeTransfer TransactionType = "TRANSFER"
	TransactionTypeSwap     TransactionType = "SWAP"

	TransactionSubtypeUnclassified          TransactionSubType = "UNCLASSIFIED"
	TransactionSubtypeBlockchain            TransactionSubType = "BLOCKCHAIN"
	TransactionSubtypeAirdrop               TransactionSubType = "AIRDROP"
	TransactionSubtypeAffiliate             TransactionSubType = "AFFILIATE"
	TransactionSubtypeStaking               TransactionSubType = "STAKING"
	TransactionSubtypeBuy_crypto            TransactionSubType = "BUY_CRYPTO"
	TransactionSubtypeOffchain              TransactionSubType = "OFFCHAIN"
	TransactionSubtypeFiat                  TransactionSubType = "FIAT"
	TransactionSubtypeSub_account           TransactionSubType = "SUB_ACCOUNT"
	TransactionSubtypeWallet_to_spot        TransactionSubType = "WALLET_TO_SPOT"
	TransactionSubtypeSpot_to_wallet        TransactionSubType = "SPOT_TO_WALLET"
	TransactionSubtypeWallet_to_derivatives TransactionSubType = "WALLET_TO_DERIVATIVES"
	TransactionSubtypeDerivatives_to_wallet TransactionSubType = "DERIVATIVES_TO_WALLET"
	TransactionSubtypeChain_switch_from     TransactionSubType = "CHAIN_SWITCH_FROM"
	TransactionSubtypeChain_switch_to       TransactionSubType = "CHAIN_SWITCH_TO"
	TransactionSubtypeInstant_exchange      TransactionSubType = "INSTANT_EXCHANGE"

	AccountWallet AccountType = "wallet"
	AccountSpot   AccountType = "spot"

	TransferByEmail    TransferBy = "email"
	TransferByUsername TransferBy = "username"

	UseOffchainNever      UseOffchainType = "never"
	UseOffchainOptionally UseOffchainType = "optionally"
	UseOffchainRequired   UseOffchainType = "required"

	ReportStatus    ReportType = "status"
	ReportNew       ReportType = "new"
	ReportCanceled  ReportType = "canceled"
	ReportExpired   ReportType = "expired"
	ReportSuspended ReportType = "suspended"
	ReportTrade     ReportType = "trade"
	ReportReplaced  ReportType = "replaced"
)

// Currency is the abstraction for a digital currency
type Currency struct {
	FullName          string    `json:"full_name"`
	PayinEnabled      bool      `json:"payin_enabled"`
	PayoutEnabled     bool      `json:"payout_enabled"`
	TransferEnabled   bool      `json:"transfer_enabled"`
	PrecisionTransfer string    `json:"precision_transfer"`
	Networks          []Network `json:"networks"`
}

type Network struct {
	Network            string `json:"network"`
	Protocol           string `json:"protocol"`
	Default            bool   `json:"default"`
	PayinEnabled       bool   `json:"payin_enabled"`
	PayoutEnabled      bool   `json:"payout_enabled"`
	PrecisionPayout    string `json:"presicion_payout"`
	PayoutFee          string `str:"payout_fee"`
	PayoutIsPaymentID  bool   `json:"payout_is_payment_id"`
	PayinPaymentID     bool   `json:"payin_payment_id"`
	PayinConfirmations int    `json:"payin_confirmation"`
	AddressRegex       string `json:"address_confirmation"`
	PaymentIDRegex     string `json:"payment_id_regex"`
	LowProcessingTime  string `json:"low_processing_time"`
	HighProcessingTime string `json:"high_processing_time"`
	AvgProcessingTime  string `json:"avg_processing_time"`
}

// Balance is the amount of currency a user have
type Balance struct {
	Currency       string `json:"currency"`
	Available      string `json:"available"`
	Reserved       string `json:"reserved"`
	ReservedMargin string `json:"reserved_margin"`
}

// Ticker is a snapshot of a symbol
type Ticker struct {
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Low         string `json:"low"`
	High        string `json:"high"`
	Open        string `json:"open"`
	Volume      string `json:"volume"`
	VolumeQuote string `json:"volume_quote"`
	Timestamp   string `json:"timestamp"`
}

type TickerPrice struct {
	Price     string `json:"price"`
	Timestamp string `json:"timestamp"`
}

type QuotationPrice struct {
	Currency  string `json:"currency"`
	Price     string `json:"price"`
	Timestamp string `json:"timestamp"`
}

type QuotationPriceHistory struct {
	Currency string         `json:"currency"`
	History  []HistoryPoint `json:"history"`
}

type HistoryPoint struct {
	Timestamp string `json:"timestamp"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	Min       string `json:"min"`
	Max       string `json:"max"`
}

// PublicTrade is the available information from public trades
type PublicTrade struct {
	ID        int64  `json:"id"`
	Price     string `json:"price"`
	Quantity  string `json:"qty"`
	Side      string `json:"side"`
	Timestamp string `json:"timestamp"`
}

// BookLevel agregates orders by price in a symbol
type BookLevel struct {
	Price  string `json:"price"`
	Amount string `json:"amount"`
}

// OrderBook is the current state of a symbol
type OrderBookJson struct {
	Ask       [][]string `json:"ask"`
	Bid       [][]string `json:"bid"`
	Timestamp string     `json:"timestamp"`
}

// OrderBook is the current state of a symbol
type OrderBook struct {
	Ask       []BookLevel `json:"ask"`
	Bid       []BookLevel `json:"bid"`
	Timestamp string      `json:"timestamp"`
}

// TradingCommission is the asociated cost to trade in the exchange
type TradingCommission struct {
	Symbol   string `json:"symbol"`
	TakeRate string `json:"take_rate"`
	MakeRate string `json:"make_rate"`
}

// Symbol is a market made of two currencies being exchanged
type Symbol struct {
	Type               string `json:"type"`
	BaseCurrency       string `json:"base_currency"`
	QuoteCurrency      string `json:"quote_currency"`
	Status             string `json:"status"`
	QuantityIncrement  string `json:"quantity_increment"`
	TickSize           string `json:"tick_size"`
	TakeRate           string `json:"take_rate"`
	MakeRate           string `json:"make_rate"`
	FeeCurrency        string `json:"fee_currency"`
	MarginTrading      bool   `json:"margin_trading"`
	MaxInitialLeverage string `json:"max_initial_leverage"`
}

// SpotOrder is the abstraction of an order in a symbol in the exchange
type SpotOrder struct {
	ID                    int64          `json:"id"`
	ClientOrderID         string         `json:"client_order_id"`
	Symbol                string         `json:"symbol"`
	Side                  string         `json:"side"`
	Status                string         `json:"status"`
	Type                  string         `json:"type"`
	TimeInForce           string         `json:"time_in_force"`
	Quantity              string         `json:"quantity"`
	QuantityCumulative    string         `json:"quantity_cumulative"`
	Price                 string         `json:"price"`
	StopPrice             string         `json:"stop_price"`
	ExpireTime            string         `json:"expire_time"`
	PostOnly              bool           `json:"post_only"`
	OriginalClientOrderID string         `json:"original_client_order_id"`
	CreatedAt             string         `json:"created_at"`
	UpdatedAt             string         `json:"updated_at"`
	Trades                []TradeOfOrder `json:"trades"`
}

// TradeOfOrder is the trade information of trades of an order
type TradeOfOrder struct {
	ID        int64  `json:"id"`
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	Fee       string `json:"fee"`
	Taker     bool   `json:"taker"`
	Timestamp string `json:"timestamp"`
}

// Trade is a movement of currency where the user takes part
type Trade struct {
	ID            int64    `json:"id"`
	OrderID       int64    `json:"order_id"`
	ClientOrderID string   `json:"client_order_id"`
	Symbol        string   `json:"symbol"`
	Side          SideType `json:"side"`
	Quantity      string   `json:"quantity"`
	Price         string   `json:"price"`
	Fee           string   `json:"fee"`
	Timestamp     string   `json:"timestamp"`
	Taker         bool     `json:"taker"`
}

// Transaction is a movement of currency,
// not in the market, but related on the exchange
type Transaction struct {
	ID        int64              `json:"id,result"`
	Status    TransactionStatus  `json:"status"`
	Type      TransactionType    `json:"type"`
	SubType   TransactionSubType `json:"subtype"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	Native    NativeTransaction  `json:"native"`
	Meta      MetaTransaction    `json:"meta"`
}

type NativeTransaction struct {
	ID            string   `json:"tx_id"`
	Index         int64    `json:"index"`
	Currency      string   `json:"currency"`
	Amount        string   `json:"amount"`
	Fee           string   `json:"fee"`
	Address       string   `json:"address"`
	PaymentID     string   `json:"payment_id"`
	Hash          string   `json:"hash"`
	OffchainID    string   `json:"offchain_id"`
	Confirmations int64    `json:"confirmations"`
	PublicComment string   `json:"public_comment"`
	ErrorCode     string   `json:"error_code"`
	Senders       []string `json:"senders"`
}

type MetaTransaction struct {
	FiatToCrypto map[string]interface{} `json:"fiat_to_crypto"`

	ID                string `json:"id"`
	ProviderName      string `json:"provider_name"`
	OrderType         string `json:"order_type"`
	SourceCurrency    string `json:"source_currency"`
	TargetCurrency    string `json:"target_currency"`
	WalletAddress     string `json:"wallet_address"`
	TransactionHash   string `json:"tx_hash"`
	TargetAmount      string `json:"target_amount"`
	SourceAmount      string `json:"source_amount"`
	Status            string `json:"status"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
	PaymentMethodType string `json:"payment_method_type"`
}

// CryptoAddress is an crypto address
type CryptoAddress struct {
	Currency  string `json:"currency"`
	Address   string `json:"address"`
	PaymentID string `json:"payment_id"`
	PublicKey string `json:"publicKey"`
}

// PayoutCryptoAddress is for external crypto addresses
type PayoutCryptoAddress struct {
	Address   string `json:"address"`
	PaymentID string `json:"payment_id"`
}

// Candle is an OHLC representation of the market
// This version uses Max instead of High nad Min instead of Low
type Candle struct {
	Timestamp   string `json:"timestamp"`
	Open        string `json:"open"`
	Close       string `json:"close"`
	Min         string `json:"min"`
	Max         string `json:"max"`
	Volume      string `json:"volume"`
	VolumeQuote string `json:"volume_quote"`
}

// Error is an error from the exchange
type Error struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

// ErrorMetadata is the data asociated with an error
// from the exchange
type ErrorMetadata struct {
	Timestamp string `json:"timestamp"`
	Path      string `json:"path"`
	Error     *Error `json:"error"`
	RequestID string `json:"request_id"`
	Status    int    `json:"status"`
}

// Report is used for websocket trading reports.
type Report struct {
	ID                           int64           `json:"id"`
	ClientOrderID                string          `json:"client_order_id"`
	Symbol                       string          `json:"symbol"`
	Side                         SideType        `json:"side"`
	Status                       OrderStatus     `json:"status"`
	Type                         OrderType       `json:"type"`
	TimeInForce                  TimeInForceType `json:"timeInForce"`
	ExpireTime                   string          `json:"expireTime"`
	Quantity                     string          `json:"quantity"`
	Price                        string          `json:"price"`
	StopPrice                    string          `json:"stopPrice"`
	PostOnly                     bool            `json:"postOnly"`
	CumQuantity                  string          `json:"cumQuantity"`
	CreatedAt                    string          `json:"createdAt"`
	UpdatedAt                    string          `json:"updatedAt"`
	ReportType                   ReportType      `json:"reportType"`
	TradeID                      int64           `json:"trade_id"`
	TradeQuantity                string          `json:"tradeQuantity"`
	TradePrice                   string          `json:"tradePrice"`
	TradeFee                     string          `json:"tradeFee"`
	OriginalRequestClientOrderID string          `json:"originalRequestClientOrder_id"`
}

type Airdrop struct {
	ID            int64  `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Currency      string `json:"currency"`
	BaseCurrency  string `json:"base_currency"`
	Description   string `json:"description"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	Amount        string `json:"amount"`
	Status        string `json:"status"`
	TransactionID string `json:"transaction_id"`
}

type AmountLock struct {
	ID                int64  `json:"id"`
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	DateEnd           string `json:"date_end"`
	Description       string `json:"description"`
	Cancelled         bool   `json:"cancelled"`
	CancelledAt       string `json:"cancelled_at"`
	CancelDescription string `json:"cancel_description"`
	CreatedAt         string `json:"created_at"`
}

type IDResponse struct {
	ID string `json:"id"`
}

type ResultResponse struct {
	ID string `json:"result"`
}

type ResultListResponse struct {
	IDs []string `json:"result"`
}

type BooleanResponse struct {
	Result bool `json:"result"`
}

type FeeResponse struct {
	Fee string `json:"fee"`
}
