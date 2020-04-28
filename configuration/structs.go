package configuration

//Config struct holds the configuration from application.yml
type Config struct {
	ConsulServersSettings struct {
		ConsulServers []string `yaml:"consulServers"`
	} `yaml:"router"`
	RouterAPISettings struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"settings"`
}
