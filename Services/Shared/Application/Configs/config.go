package configs

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Database struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
	} `yaml:"database"`

	JWT struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
	} `yaml:"database"`
}

/*


server:
  host: "localhost"
  port: 8000


database:
  driver: "mysql"
  host: "locahost"
  user: "user"
  pass: "user"
  port: 3306


*/
