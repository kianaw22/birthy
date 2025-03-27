package config

var AppConfig = struct {
	TargetAmount       int
	ZarinpalMerchantID string
	CallbackURL        string
}{
	TargetAmount:       50000000,
	ZarinpalMerchantID: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
	CallbackURL:        "http://localhost:8080/verify",
}

var DBConfig = struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}{
	Host:     "localhost",
	Port:     "5432",
	User:     "postgres",
	Password: "1",
	DBName:   "postgres",
	SSLMode:  "disable",
}
