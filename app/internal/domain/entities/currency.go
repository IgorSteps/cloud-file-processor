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

func (c Currency) String() (string, error) {
	if name, ok := currencyName[c]; ok {
		return name, nil
	}
	return "", NewInvalidError("unknown currency")
}
