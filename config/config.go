package config

type DatabaseConfig struct {
	User 	   string `json:"user"`
	Password string `json:"password"`
	Host 	   string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
}

var C = &DatabaseConfig{}