package config

type GlobalConfig struct {
	Server server
}
type server struct {
	Port string
}