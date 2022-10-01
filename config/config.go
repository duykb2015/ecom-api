package config

import "github.com/caarlos0/env/v6"

type (
	Config struct {
		MySQL
		App
	}

	MySQL struct {
		URI             string `env:"MYSQL_URI" envDefault:"root:@tcp(localhost:3306)/ecommerce?parseTime=true"`
		MaxIdleConns    int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"100"`
		MaxOpenConns    int    `env:"MYSQL_OPEN_CONNS" envDefault:"100"`
		ConnMaxLifeTime int    `env:"MYSQL_CONN_MAX_LIFE_TIME" envDefault:"149"`
		Timeout         int    `env:"MYSQL_TIMEOUT" envDefault:"100"`
		Debug           bool   `env:"MYSQL_DEBUG" envDefault:"true"`
	}

	RedisUC struct {
		Addrs    string `env:"REDIS_UC_ADDRS,notEmpty"`
		PoolSize int    `env:"REDIS_UC_POOL_SIZE" envDefault:"10"`
	}

	App struct {
		JWTSecretKey string `env:"SECRET_KEY"`
		APIKey       string
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
