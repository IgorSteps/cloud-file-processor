package entities

type Currency int

const (
	USD Currency = iota
	EUR
	GBR
)

var currencyName = map[Currency]string{
	USD: "USD",
	EUR: "EUR",
	GBR: "GBR",
}

var currencyValue = map[string]Currency{
	"USD": USD,
	"EUR": EUR,
	"GBR": GBR,
}

func (c Currency) String() (string, error) {
	if name, ok := currencyName[c]; ok {
		return name, nil
	}
	return "", NewInvalidError("unknown currency")
}

// ParseCurrency maps a string to the Currency enum.
func ParseCurrency(s string) (Currency, error) {
	if val, ok := currencyValue[s]; ok {
		return val, nil
	}
	return 0, NewInvalidError("unknown currency")
}
