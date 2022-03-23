package args

// SideType is an order side or a trade side of an order
type SideType string

const (
	// SideTypeSell is the sell side for an order or a trade
	SideTypeSell SideType = "sell"
	// SideTypeBuy is the buy side for an order or a trade
	SideTypeBuy SideType = "buy"
)

// OrderType is a type of order
type OrderType string

// OrderTypes
const (
	OrderLimit            OrderType = "limit"
	OrderMarket           OrderType = "market"
	OrderStopLimit        OrderType = "stopLimit"
	OrderStopMarket       OrderType = "stopMarket"
	OrderTakeProfitLimit  OrderType = "takeProfitLimit"
	OrderTakeProfitMarket OrderType = "takeProfitMarket"
)

// TimeInForceType is the time in force of an order
type TimeInForceType string

// types of time in force
const (
	TimeInForceGTC TimeInForceType = "GTC" // Good Till Cancel
	TimeInForceIOC TimeInForceType = "IOC" // Immediate or Cancel
	TimeInForceFOK TimeInForceType = "FOK" // Fill or Kill
	TimeInForceDAY TimeInForceType = "Day" // valid during Day
	TimeInForceGTD TimeInForceType = "GTD" // Good Till Date
)

// SortType is the sorting direction of a query
type SortType string

const (
	// SortTypeASC is the ascending sorting direction of a query
	SortASC SortType = "ASC"
	// SortTypeDESC is the descending sorting direction of a query
	SortDESC SortType = "DESC"
)

// SortByType is the parameter for sorting
type SortByType string

const (
	// SortByTypeTimestamp is the sorting field for pagination, sorting by timestamp
	SortByTimestamp SortByType = "timestamp"
	// SortByTypeID is the sorting field for pagination, sorting by id
	SortByID SortByType = "id"

	SortByCreatedAt SortByType = "created_at"
)

// PeriodType is the period of a candle
type PeriodType string

// candle periods
const (
	Period1Minute   PeriodType = "M1"
	Period3Minutes  PeriodType = "M3"
	Period5Minutes  PeriodType = "M5"
	Period15Minutes PeriodType = "M15"
	Period30Minutes PeriodType = "M30"
	Period1Hour     PeriodType = "H1"
	Period4Hours    PeriodType = "H4"
	Period1Day      PeriodType = "D1"
	Period7Days     PeriodType = "D7"
	Period1Month    PeriodType = "1M"
)

// MarginType is the type of margin of a trade
type MarginType string

// IdentifyByType for transfers
type IdentifyByType string

// identify by types
const (
	IdentifyByEmail    IdentifyByType = "email"
	IdentifyByUsername IdentifyByType = "username"
)

type AccountType string

const (
	AccountWallet AccountType = "wallet"
	AccountSpot   AccountType = "spot"
)

type UseOffchainType string

const (
	UseOffchainNever     UseOffchainType = "never"
	UseOffchainOptionaly UseOffchainType = "optionaly"
	UseOffChainRequired  UseOffchainType = "required"
)

type TransactionType string

const (
	TransactionDeposit  TransactionType = "DEPOSIT"
	TransactionWithdraw TransactionType = "WITHDRAW"
	TransactionTransfer TransactionType = "TRANSFER"
	TransactionSwap     TransactionType = "SWAP"
)

type TransactionSubtypeType string

const (
	TransactionSubtyeUnclassified        TransactionSubtypeType = "UNCLASSIFIED"
	TransactionSubtyeBlockchain          TransactionSubtypeType = "BLOCKCHAIN"
	TransactionSubtyeAirdrop             TransactionSubtypeType = "AIRDROP"
	TransactionSubtyeAffiliate           TransactionSubtypeType = "AFFILIATE"
	TransactionSubtyeStaking             TransactionSubtypeType = "STAKING"
	TransactionSubtyeBuyCrypto           TransactionSubtypeType = "BUY_CRYPTO"
	TransactionSubtyeOffchain            TransactionSubtypeType = "OFFCHAIN"
	TransactionSubtyeFiat                TransactionSubtypeType = "FIAT"
	TransactionSubtyeSubAccount          TransactionSubtypeType = "SUB_ACCOUNT"
	TransactionSubtyeWalletToSpot        TransactionSubtypeType = "WALLET_TO_SPOT"
	TransactionSubtyeSpotToWallet        TransactionSubtypeType = "SPOT_TO_WALLET"
	TransactionSubtyeWalletToDerivatives TransactionSubtypeType = "WALLET_TO_DERIVATIVES"
	TransactionSubtyeDerivativesToWallet TransactionSubtypeType = "DERIVATIVES_TO_WALLET"
	TransactionSubtyeChainSwitchFrom     TransactionSubtypeType = "CHAIN_SWITCH_FROM"
	TransactionSubtyeChainSwitchTo       TransactionSubtypeType = "CHAIN_SWITCH_TO"
	TransactionSubtyeInstantExchange     TransactionSubtypeType = "INSTANT_EXCHANGE"
)

type TransactionStatusType string

const (
	TransactionStatusCreated    TransactionStatusType = "CREATED"
	TransactionStatusPending    TransactionStatusType = "PENDING"
	TransactionStatusFailed     TransactionStatusType = "FAILED"
	TransactionStatusSuccess    TransactionStatusType = "SUCCESS"
	TransactionStatusRolledBack TransactionStatusType = "ROLLED_BACK"
)

type AirdropStatusType string

const (
	AirdropStatusAvailable AirdropStatusType = "available"
	AirdropStatusClaimed   AirdropStatusType = "claimed"
	AirdropStatusPending   AirdropStatusType = "pending"
	AirdropStatusCommited  AirdropStatusType = "commited"
)
