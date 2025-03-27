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
	Host:     "dpg-cvir2ja4d50c73c51ph0-a",
	Port:     "5432",
	User:     "birthyuser",
	Password: "RbFFCqZB8FyLWkkzHf1ZyRH63EeQXNW4",
	DBName:   "birthydb",
	SSLMode:  "disable",
}
