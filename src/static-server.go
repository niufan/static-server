package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type ServerInfo struct {
	Host string
	Port int16
}

type StaticInfo struct {
	Root string
}

type Config struct {
	Server ServerInfo
	Static StaticInfo
}

func (config *Config) getAddress() string {
	return string(config.Server.Host) + ":" + strconv.FormatInt(int64(config.Server.Port), 10)
}

func (config *Config) getStaticRoot() string {
	return config.Static.Root
}

func main() {
	var configPath string
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	flag.StringVar(&configPath, "config", path+"/../config/config.yml", "请设置配置文件路径")
	flag.Parse()
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	config := Config{}
	yaml.Unmarshal(configData, &config)
	http.ListenAndServe(config.getAddress(), http.FileServer(http.Dir(config.getStaticRoot())))
}
