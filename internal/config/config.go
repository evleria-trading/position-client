package config

type Сonfig struct {
	PositionServiceUrl string `env:"POSITION_SERVICE_URL" envDefault:"localhost:6001"`
	PriceServiceUrl    string `env:"PRICE_SERVICE_URL" envDefault:"localhost:6101"`
	UserServiceUrl     string `env:"USER_SERVICE_URL" envDefault:"localhost:6201"`
}
