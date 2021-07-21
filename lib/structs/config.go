package structs

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Postgres Database `yaml:"postgres"`
}

type Database struct {
	ConnectionString string `yaml:"conn_string"`
}
