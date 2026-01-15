package domains

type RedisConfig struct {
	Prefix   string
	Addr     string
	Password string
	DB       int
}
