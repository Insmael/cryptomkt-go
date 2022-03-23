package args

import (
	"fmt"
	"strings"
)

// Argument are functions that serves as arguments for the diferent
// requests to the server, either rest request, or websocket requests.
// Argument works by extending a given map with a particular key.
// For example Symbol("EOSETH") returns an Argument that extends
// a map assigning "EOSETH" to the key "symbol".
// usually, a function returning an Argument is what's used as a
// request parameter. just like the Symbol(val string) function
// in the example above.
type Argument func(map[string]interface{})

func fromSnakeCaseToCamelCase(s string) string {
	snakeParts := strings.Split(s, "_")
	camelParts := make([]string, 0)
	for _, snakePart := range snakeParts {
		camelParts = append(camelParts, strings.Title(snakePart))
	}
	return strings.Join(camelParts, "")
}

// BuildParams makes a map with the Arguments functions,
// and check for the presence of "requireds" keys in the map,
// raising an error if some required keys are not present.
func BuildParams(
	arguments []Argument,
	requireds ...string,
) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	for _, argFunc := range arguments {
		argFunc(params)
	}
	missing := []string{}
	for _, required := range requireds {
		if _, ok := params[required]; !ok {
			missing = append(missing, required)
		}
	}
	if len(missing) > 0 {
		missingAsCamelCase := make([]string, 0)
		for _, miss := range missing {
			missingAsCamelCase = append(
				missingAsCamelCase,
				fromSnakeCaseToCamelCase(miss),
			)
		}
		return nil, fmt.Errorf(
			"CryptomarketSDKError: missing arguments: %v", missingAsCamelCase,
		)
	}
	return params, nil
}

func Currencies(val []string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameCurrencies] = val
	}
}

func Currency(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameCurrency] = val
	}
}

func Symbols(val []string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameSymbols] = val
	}
}

func Symbol(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameSymbol] = val
	}
}

func Sort(val SortType) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameSort] = val
	}
}

func SortBy(val SortByType) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameSortBy] = val
	}
}

func From(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameFrom] = val
	}
}

func To(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameTo] = val
	}
}

func Till(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameTill] = val
	}
}

func Limit(val int) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameLimit] = val
	}
}

func Offset(val int) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameOffset] = val
	}
}

func Volume(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameVolume] = val
	}
}

func Period(val PeriodType) Argument {
	return func(params map[string]interface{}) {
		params[ArgNamePeriod] = val
	}
}

func ClientOrderID(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameClientOrderID] = val
	}
}

func Side(val SideType) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameSide] = val
	}
}

func Quantity(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameQuantity] = val
	}
}

func Price(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNamePrice] = val
	}
}

func StopPrice(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameStopPrice] = val
	}
}

func TimeInForce(val TimeInForceType) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameTimeInForce] = val
	}
}

func ExpireTime(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameExpireTime] = val
	}
}

func StrictValidate(val bool) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameStrictValidate] = val
	}
}

func PostOnly(val bool) Argument {
	return func(params map[string]interface{}) {
		params[ArgNamePostOnly] = val
	}
}

func OrderID(val int64) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameOrderID] = val
	}
}

func Amount(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameAmount] = val
	}
}

func Address(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameAddress] = val
	}
}

func PaymentID(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNamePaymentID] = val
	}
}

func IncludeFee(val bool) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameIncludeFee] = val
	}
}

func AutoCommit(val bool) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameAutoCommit] = val
	}
}

func PublicComment(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNamePublicComment] = val
	}
}

func FromCurrency(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameFromCurrency] = val
	}
}

func ToCurrency(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameToCurrency] = val
	}
}

func TransferType(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameTransferType] = val
	}
}

func Identifier(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameIdentifier] = val
	}
}

func IdentifyBy(val IdentifyByType) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameIdentifyBy] = val
	}
}

func ShowSenders(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameShowSenders] = val
	}
}

func ID(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameID] = val
	}
}

func Source(val AccountType) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameSource] = val
	}
}

func Destination(val AccountType) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameDestination] = val
	}
}

func Since(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameSince] = val
	}
}

func Untill(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameUntil] = val
	}
}

func Depth(val string) Argument {
	return func(params map[string]interface{}) {
		params[argNameDepth] = val
	}
}

func TakeRate(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameTakeRate] = val
	}
}

func MakeRate(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameMakeRate] = val
	}
}

func NewClientOrderID(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameNewClientOrderID] = val
	}
}

func UseOffchain(val string) Argument {
	return func(params map[string]interface{}) {
		params[ArgNameNewClientOrderID] = val
	}
}

// TODO: fix
func RequestClientID(val string) Argument {
	return func(params map[string]interface{}) {
		params["request_client_order_id"] = val
	}
}
