package currency

type CurrencyRate struct {
	Currencies []struct {
		Name  string  `yaml:"name"`
		Value float64 `yaml:"value"`
	}
}
