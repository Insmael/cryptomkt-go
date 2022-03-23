package rest

const (
	// public endpoints
	endpointCurrency     = "public/currency"
	endpointSymbol       = "public/symbol"
	endpointTicker       = "public/ticker"
	endpointPrices       = "public/price/rate"
	endpointPriceHistory = "public/price/history"
	endpointPriceTicker  = "public/price/ticker"
	endpointTrade        = "public/trades"
	endpointOrderbook    = "public/orderbook"
	endpointCandle       = "public/candles"
	// trading endpoints
	endpointTradingBalance    = "spot/balance"
	endpointOrder             = "spot/order"
	endpointTradingCommission = "spot/fee"
	// trading history endpoints
	endpointOrderHistory = "spot/history/order"
	endpointTradeHistory = "spot/history/trades"
	// wallet management
	endpointWalletBalance              = "wallet/balance"
	endpointCryptoAdress               = "wallet/crypto/address"
	endpointCryptoAdressRecentDeposit  = "wallet/crypto/address/recent-deposit"
	endpointCryptoAdressRecentWithdraw = "wallet/crypto/address/recent-withdraw"
	endpointCryptoAdressCheckMine      = "wallet/crypto/address/check-mine"
	endpointCryptoWithdraw             = "wallet/crypto/withdraw"
	endpointConvert                    = "wallet/convert"
	endpointAccountTranser             = "wallet/transfer"
	endpointInternalWithdraw           = "wallet/internal/withdraw"
	endpointTransactions               = "wallet/transactions"
	endpointCryptoCheckOffchain        = "wallet/crypto/check-offchain-available"
	endpointEstimateWithdrawFee        = "wallet/crypto/fee/estimate"
	endpointAirdrops                   = "wallet/airdrops"
	endpointAmountLocks                = "wallet/amount-locks"
)
