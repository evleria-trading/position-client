package config

type Сonfig struct {
	RedisPass string `env:"REDIS_PASS" envDefault:""`
	RedisHost string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort int    `env:"REDIS_PORT" envDefault:"6379"`

	PositionServiceUrl string `env:"POSITION_SERVICE_URL" envDefault:"localhost:6001"`
}
