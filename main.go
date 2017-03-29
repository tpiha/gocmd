package main

var conf *Config

var clip *CliParser

func main() {
	conf = initConfig()

	clip = initCli()

	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd")
	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd1")
	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd2")
	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd3")
	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd4")
	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd5")
	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd6")
	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd7")
	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd8")
	RunSSHCommand("picostats", "46.101.254.59", "touch gocmd9")
}
