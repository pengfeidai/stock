package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Server `yaml:"server"`
	Mysql  `yaml:"mysql"`
	Log    `yaml:"log"`
	Url    `yaml:"url"`
}

type Server struct {
	Port      int    `yaml:"port"`
	Mode      string `yaml:"mode"`
	LimitNum  int    `yaml:"limitNum"`
	UserMongo bool   `yaml:"useMongo"`
	UserRedis bool   `yaml:"useRedis"`
}

type Mysql struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Path         string `yaml:"path"`
	Database     string `yaml:"database"`
	Config       string `yaml:"config"`
	Driver       string `yaml:"driver"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	Log          bool   `yaml:"log"`
}

type Log struct {
	Debug    bool          `yaml:"debug"`
	MaxAge   time.Duration `yaml:"maxAge"`
	FileName string        `yaml:"fileName"`
	DirName  string        `yaml:"dirName"`
}

type Url struct {
	Prefix string `yaml:"prefix"`
}

var Conf *Yaml

const defaultConfigFile = "config.yaml"

func init() {
	c := &Yaml{}
	configFile := flag.String("c", defaultConfigFile, "help config path")
	flag.Parse()
	yamlFile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(fmt.Errorf("get yamlFile error: %s", err))
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Printf("config yamlFile load Init success.")
	Conf = c
}
