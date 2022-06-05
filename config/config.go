package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	PORT       string `yaml:"port"`
	EnableLogs bool   `yaml:"enableLogs"`
}

type DB struct {
	Name string
	Path string
}

type Config struct {
	Environment string
	Server      Server
	Logging     bool
	Feeds       Feeds
	DB          DB
}

type Feeds struct {
	URL []string
}

func (c *Config) defaults() *Config {

	var urls []string
	urls = append(urls, "https://www.theverge.com/rss/index.xml", "https://www.wired.com/feed/category/business/latest/rss", "https://www.wired.com/feed/tag/ai/latest/rss")

	config := &Config{
		Environment: "Debug",
		Server: Server{
			PORT:       "9999",
			EnableLogs: false,
		},
		Logging: false,
		Feeds: Feeds{
			URL: urls,
		},
		DB: DB{
			Name: "feeds.db",
			Path: "",
		},
	}

	return config
}

func (c *Config) load() (*Config, error) {
	//Config file exist. Load the config.
	log.Println("Config file exists....Loading")
	f, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("ERROR: Cannot read config file - %v", err)
		return nil, err
	}

	err = yaml.Unmarshal([]byte(f), &c)
	if err != nil {
		log.Fatalf("ERROR: Configuration Error: %v", err)
		return nil, err
	}

	return c, nil
}

func InitConfig() (*Config, error) {
	var c *Config

	if _, err := os.Stat("config.yml"); err != nil {

		//If config file does not exists.
		config := c.defaults()
		con, err := yaml.Marshal(&config)
		if err != nil {
			log.Fatalf("Error while parsing the config file: %v", err)
			return nil, err
		}

		f, err := os.Create("config.yml")
		if err != nil {
			log.Fatalf("Error while creating the config file: %v", err)
			return nil, err
		}

		defer f.Close()

		d := []byte(string(con))
		w, err := f.Write(d)
		if err != nil {
			log.Fatalf("Error while writing to config file: %v", err)
			return nil, err
		}

		log.Printf("Config saved successfully - %d bytes written", w)

		c, err = c.load()
		if err != nil {
			log.Fatalf("ERROR: Cannot load configuration: %v", err)
		}
	} else if os.IsNotExist(err) {
		log.Fatalf("ERROR: File not found - %v", err)
		return nil, err
	} else {
		c, err = c.load()
		if err != nil {
			log.Fatalf("ERROR: Cannot load configurations: %v", err)
		}
	}

	return c, nil

}
