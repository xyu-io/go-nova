package config

import (
	"github.com/aceld/zinx/zconf"
	"github.com/go-nova/pkg/common/config"
	"github.com/go-nova/pkg/common/config/file"
	"github.com/go-nova/pkg/common/dao"
	"github.com/go-nova/pkg/common/datas/base_struct"
	log2 "github.com/go-nova/pkg/core/log"
	"github.com/go-nova/pkg/utils/fileops"
	"github.com/tomwright/dasel"
	"github.com/tomwright/dasel/storage"
)

const defaultConfigName = "./cmd/polaris/config.yml"

var AppCfg *Config

type Config struct {
	App      AppConfig          `yaml:"app" json:"app" xml:"app"`
	Server   ServerConfig       `yaml:"server" json:"server" xml:"server"`
	Zinx     *zconf.Config      `yaml:"zinx" json:"zinx" xml:"zinx"`
	Database dao.Option         `yaml:"database" json:"database" xml:"database"`
	Apis     []base_struct.APIs `yaml:"apis" json:"apis" xml:"apis"`
}

type (
	AppConfig struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Debug   bool   `json:"debug"`
	}
	ServerConfig struct {
		Host          string `json:"host"`
		Port          int    `json:"port"`
		ProfileSwitch bool   `json:"profile_switch"`
		ProfilePort   int    `json:"profile_port"`
	}
	DatabaseConfig struct {
		Username       string   `json:"username,omitempty"`
		Password       string   `json:"password,omitempty"`
		Driver         string   `json:"driver"`
		Addr           string   `json:"addr"`
		DBNames        []string `json:"db_names"` // 多数据库
		DatabasePrefix string   `json:"database_prefix"`
	}
)

func InitConfig(configName string) {
	filePath := defaultConfigName
	if configName != "" {
		filePath = configName
	}
	//filePrefix := strings.Split(configName, ".")[0]
	// 读取配置文件
	c := config.New(
		// 创建文件源
		config.WithSource(file.NewSource(filePath)),
		//config.WithSource(file.NewSource(filePrefix+".json")),
	)

	if err := c.Load(); err != nil {
		log2.Errorw("error", err.Error())
		return
	}
	// 解析配置文件
	var cfg Config
	if err := c.Scan(&cfg); err != nil {
		log2.Errorw("error", err.Error())
		return
	}
	if &cfg.Server != nil {
		fileops.WriteJSON(cfg, "config.json")
	}
	AppCfg = &cfg
	//UpdateWithDasel(filePath)
}

type Info struct {
	Name string `json:"name"`
	Rgb  string `json:"rgb"`
	Tag  int    `json:"tag"`
}

func UpdateWithDasel(fileName string) {
	filePath := defaultConfigName
	if fileName != "" {
		filePath = fileName
	}
	parser, err := storage.NewReadParserFromFilename(filePath)
	if err != nil {
		log2.Errorw("error", "could not get parser, "+err.Error())
		return
	}

	data, err := storage.LoadFromFile(filePath, parser)
	if err != nil {
		log2.Errorw("error", "could not load value from file, "+err.Error())
		return
	}

	rootNode := dasel.New(data)
	if err := rootNode.Put(".server.name", "Nova"); err != nil {
		panic(err)
	}

	err = rootNode.WriteToFile("new_"+filePath, "json", nil)
	if err != nil {
		log2.Errorw("error", err.Error())
	}

	// k, err := dasel.New(value).Query(".colourCodes")

	//err = fileops.WriteJSON(dasel.New(value).OriginalValue, "XXX.json")
	//var cfg []*Info
	//bytes, err := json.Marshal(k.OriginalValue)
	//if err != nil {
	//	return
	//}
	//err = fileops.WriteJSON(k.OriginalValue, "12.json")
	//err = json.Unmarshal(bytes, &cfg)
	//if err != nil {
	//	fmt.Println("zzz", err.Error())
	//}
	//fmt.Println("----", cfg)

	//for _, v := range k.InterfaceValue().([]interface{}) {
	//	if _, ok := v.(map[interface{}]interface{}); !ok {
	//		continue
	//	}
	//	for key, value := range v.(map[interface{}]interface{}) {
	//		fmt.Println(key, ":", value)
	//	}
	//}
}
