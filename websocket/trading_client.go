package websocket

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cryptomarket/cryptomarket-go/args"
	"github.com/cryptomarket/cryptomarket-go/internal"
	"github.com/cryptomarket/cryptomarket-go/models"
)

// SpotTradingClient connects via websocket to cryptomarket to enable the user to manage orders. uses SHA256 as auth method and authenticates automatically.
type SpotTradingClient struct {
	clientBase
}

// NewSpotTradingClient returns a new spot trading client if the connection with the
// cryptomarket server is successful and if the authentication is successfull.
// return error otherwise.
// Arguments:
//  apiKey // The API key
//  apiSecret // The API secret
//  window // Maximum difference between the creation of the request and the moment of request processing in milliseconds. Max is 60_000. Defaul is 10_000 (use 0 as argument for default)
func NewSpotTradingClient(apiKey, apiSecret string, window int) (*SpotTradingClient, error) {

	client := &SpotTradingClient{
		clientBase: clientBase{
			wsManager: newWSManager("/api/3/ws/trading"),
			chanCache: newChanCache(),
			window:    window,
		},
	}

	// connect to streaming
	if err := client.wsManager.connect(); err != nil {
		return nil, fmt.Errorf("Error in websocket client connection: %s", err)
	}
	// handle incomming data
	go client.handle(client.wsManager.rcv)

	if err := client.authenticate(apiKey, apiSecret); err != nil {
		return nil, err
	}
	return client, nil
}

// GetSpotTradingBalances gets the user's spot trading balance for all currencies with balance
//
// https://api.exchange.cryptomkt.com/#get-spot-trading-balances
func (client *SpotTradingClient) GetSpotTradingBalances(
	ctx context.Context,
) ([]models.Balance, error) {
	var resp struct {
		Result []models.Balance
	}
	err := client.doRequest(ctx, methodSpotBalances, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

// GetSpotTradingBalanceOfCurrency gets the user spot trading balance of a currency
//
// https://api.exchange.cryptomkt.com/#get-spot-trading-balance-2
//
// Arguments:
//  Currency(string)  // The currency code to query the balance
func (client *SpotTradingClient) GetSpotTradingBalanceOfCurrency(
	ctx context.Context,
	arguments ...args.Argument,
) (*models.Balance, error) {
	var resp struct {
		Result models.Balance
	}
	err := client.doRequest(
		ctx,
		methodSpotBalance,
		arguments,
		[]string{internal.ArgNameCurrency},
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

// GetActiveSpotOrders gets the user's active spot orders
//
// https://api.exchange.cryptomkt.com/#get-all-active-spot-orders
func (client *SpotTradingClient) GetActiveSpotOrders(
	ctx context.Context,
) ([]models.Report, error) {
	var resp struct {
		Result []models.Report
	}
	err := client.doRequest(ctx, methodGetSpotOrders, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

// CreateSpotOrder creates a new spot order
//
// For fee, for price accuracy and quantity, and for order status information see the api docs at https://api.exchange.cryptomkt.com/#create-new-spot-order
//
// https://api.exchange.cryptomkt.com/#place-new-spot-order
//
// Arguments:
//  Symbol(string)  // Trading symbol
//  Side(SideType)  // Either SideBuy or SideSell
//  Quantity(string)  // Order quantity
//  ClientOrderID(string)  // Optional. If given must be unique within the trading day, including all active orders. If not given, is generated by the server
//  Type(OrderType)  // Optional. OrderLimit, OrderMarket, OrderStopLimit, OrderStopMarket, OrderTakeProfitLimit or OrderTakeProfitMarket. Default is OrderLimit
//  TimeInForce(TimeInForceType)  // Optional. TimeInForceGTC, TimeInForceIOC, TimeInForceFOK, TimeInForceDay, TimeInForceGTD. Default to TimeInForceGTC
//  Price(string)  // Optional. Required for OrderLimit and OrderStopLimit. limit price of the order
//  StopPrice(string)  // Optional. Required for OrderStopLimit and OrderStopMarket orders. stop price of the order
//  ExpireTime(string)  // Optional. Required for orders with timeInForceGDT
//  StrictValidate(bool)  // Optional. If False, the server rounds half down for tickerSize and quantityIncrement. Example of ETHBTC: tickSize = '0.000001', then price '0.046016' is valid, '0.0460165' is invalid
//  PostOnly(bool)  // Optional. If True, your postOnly order causes a match with a pre-existing order as a taker, then the order will be cancelled
//  TakeRate(string)  // Optional. Liquidity taker fee, a fraction of order volume, such as 0.001 (for 0.1% fee). Can only increase the fee. Used for fee markup.
//  MakeRate(string)  // Optional. Liquidity provider fee, a fraction of order volume, such as 0.001 (for 0.1% fee). Can only increase the fee. Used for fee markup.
func (client *SpotTradingClient) CreateSpotOrder(
	ctx context.Context,
	arguments ...args.Argument,
) (*models.Report, error) {
	var resp struct {
		Result models.Report
	}
	err := client.doRequest(
		ctx,
		methodCreateSpotOrder,
		arguments,
		[]string{
			internal.ArgNameSymbol,
			internal.ArgNameSide,
			internal.ArgNameQuantity,
		},
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

// CreateSpotOrderList creates a list of spot orders and returns a list of reports of the created orders, or a possible error
//
// Types or contingency:
//
//  - ContingencyTypeAllOrNone (ContingencyTypeAON) (AON)
//  - ContingencyTypeOneCancelOther (ContingencyTypeOCO) (OCO)
//  - ContingencyOneTriggerOneCancelOther (ContingencyTypeOTOCO) (OTOCO)
//
// Restriction in the number of orders:
//
//  - An AON list must have 2 or 3 orders
//  - An OCO list must have 2 or 3 orders
//  - An OTOCO must have 3 or 4 orders
//
// Symbol restrictions:
//
//  - For an AON order list, the symbol code of orders must be unique for each order in the list.
//  - For an OCO order list, there are no symbol code restrictions.
//  - For an OTOCO order list, the symbol code of orders must be the same for all orders in the list (placing orders in different order books is not supported).
//
// ORDER_TYPE restrictions:
//  - For an AON order list, orders must be OrderLimit or OrderMarket
//  - For an OCO order list, orders must be OrderLimit, OrderStopLimit, OrderStopMarket, OrderTakeProfitLimit or OrderTakeProfitMarket.
//  - An OCO order list cannot include more than one limit order (the same applies to secondary orders in an OTOCO order list).
//  - For an OTOCO order list, the first order must be OrderLimit, OrderMarket, OrderStopLimit, OrderStopMarket, OrderTakeProfitLimit or OrderTakeProfitMarket.
//  - For an OTOCO order list, the secondary orders have the same restrictions as an OCO order
//  - Default is OrderTypeLimit
//
// https://api.exchange.cryptomkt.com/#create-new-spot-order-list-2
//
// Arguments:
//  OrderListID(string)  // order list identifier. If ommited, it will be generated by the system. Must be equal to the client order id of the first order in the request
//  Contingency(ContingencyType)  // order list type. ContingencyAON, ContingencyOCO or ContingencyOTOCO
//  Orders([]Order)  // the list of orders, orders from the args package
func (client *SpotTradingClient) CreateSpotOrderList(
	ctx context.Context,
	arguments ...args.Argument,
) ([]models.Order, error) {
	var resp struct {
		Result []models.Order
	}
	err := client.doRequest(
		ctx,
		methodCreateSpotOrderList,
		arguments,
		[]string{internal.ArgNameOrderListID, internal.ArgNameContingencyType, internal.ArgNameOrders},
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

// CancelSpotOrders cancels a spot order
//
// https://api.exchange.cryptomkt.com/#cancel-spot-order-2
//
// Arguments:
//  ClientOrderID(string)  // the client order id of the order to cancel
func (client *SpotTradingClient) CancelSpotOrder(
	ctx context.Context,
	arguments ...args.Argument,
) (*models.Report, error) {
	var resp struct {
		Result models.Report
	}
	err := client.doRequest(
		ctx,
		methodCancelSpotOrder,
		arguments,
		[]string{internal.ArgNameClientOrderID},
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

// CancelAllSpotOrders cancel all active spot orders and returns the ones that could not be canceled
//
// https://api.exchange.cryptomkt.com/#cancel-spot-orders
func (client *SpotTradingClient) CancelAllSpotOrders(
	ctx context.Context,
) ([]models.Order, error) {
	var resp struct {
		Result []models.Order
	}
	err := client.doRequest(
		ctx,
		methodCancelSpotOrders,
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

// ReplaceSpotOrder changes the parameters of an existing order, quantity or price
//
// https://api.exchange.cryptomkt.com/#cancel-replace-spot-order
//
// Arguments:
//  ClientOrderID(string)  // the client order id of the order to change
//  NewClientOrderID(string)  // the new client order id for the modified order. must be unique within the trading day
//  Quantity(string)  // new order quantity
//  Price(string)  // new order price
//  StrictValidate(bool)  //  price and quantity will be checked for the incrementation with tick size and quantity step. See symbol's tick_size and quantity_increment
func (client *SpotTradingClient) ReplaceSpotOrder(
	ctx context.Context,
	arguments ...args.Argument,
) (*models.Report, error) {
	var resp struct {
		Result models.Report
	}
	err := client.doRequest(
		ctx,
		methodReplaceSpotOrder,
		arguments,
		[]string{
			internal.ArgNameClientOrderID,
			internal.ArgNameNewClientOrderID,
			internal.ArgNamePrice,
			internal.ArgNameQuantity,
		},
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

// GetTradingCommissions gets the personal trading commission rates for all symbols
//
// https://api.exchange.cryptomkt.com/#get-spot-fees
func (client *SpotTradingClient) GetTradingCommissions(
	ctx context.Context,
) ([]models.TradingCommission, error) {
	var resp struct {
		Result []models.TradingCommission
	}
	err := client.doRequest(ctx, methodSpotFees, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

// GetTradingCommissionOfSymbol gets the personal trading commission rate of a symbol
//
// https://api.exchange.cryptomkt.com/#get-spot-fee
//
// Arguments:
//  Symbol(string)  // The symbol of the commission rate
func (client *SpotTradingClient) GetSpotFee(
	ctx context.Context,
	arguments ...args.Argument,
) (*models.TradingCommission, error) {
	var resp struct {
		Result models.TradingCommission
	}
	err := client.doRequest(ctx, methodSpotFee, arguments, []string{internal.ArgNameSymbol}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

///////////////////
// Subscriptions //
///////////////////

// SubscribeToReports subscribe to a feed of execution reports of the user's orders
//
// recieves a snapshot and updates notifications
//
// https://api.exchange.cryptomkt.com/#socket-spot-trading
func (client *SpotTradingClient) SubscribeToReports() (notificationCh chan models.Notification[[]models.Report], err error) {
	dataCh, err := client.doSubscription(methodSubscribeSpotReports, nil, nil)
	if err != nil {
		return nil, err
	}
	notificationCh = make(chan models.Notification[[]models.Report])
	go func() {
		defer close(notificationCh)
		sendMap := map[string]func([]byte){
			methodSpotOrders: func(data []byte) {
				var snapshot struct {
					Params []models.Report
				}
				json.Unmarshal(data, &snapshot)
				notificationCh <- models.Notification[[]models.Report]{
					Data:             snapshot.Params,
					NotificationType: args.NotificationSnapshot,
				}
			},
			methodSpotOrder: func(data []byte) {
				var update struct {
					Params models.Report
				}
				json.Unmarshal(data, &update)
				notificationCh <- models.Notification[[]models.Report]{
					Data:             []models.Report{update.Params},
					NotificationType: args.NotificationUpdate,
				}
			},
		}
		// the first time it recieves a list of reports
		var method struct {
			Method string
		}
		for data := range dataCh {
			json.Unmarshal(data, &method)
			sendMap[method.Method](data)
		}
	}()
	return notificationCh, nil
}

// UnsubscribeToReports stop recieveing the report feed subscription
//
// https://api.exchange.cryptomkt.com/#socket-spot-trading
func (client *SpotTradingClient) UnsubscribeToReports() (err error) {
	return client.doUnsubscription(methodSpotUnsubscribe, nil, nil)
}
