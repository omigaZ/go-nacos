package config

import (
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/nacos_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/common/http_agent"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var (
	pPath 	= flag.String("path","./config/","config file path")

	Path 	string

	Redis	*redis
	Mysql	*mysql
)

type NCfg struct {
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
}

type mysqlOption struct {
	Protocol string `yaml:"protocol"`
	Charset  string `yaml:"charset"`
}

type mysql struct {
	Flavor      string `yaml:"flavor"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Schema      string `yaml:"schema"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MysqlOption mysqlOption `yaml:"option"`
}

func (i *mysql) Addr() string {
	return fmt.Sprintf("%s:%d", i.Host, i.Port)
}

func (i *mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		i.Username, i.Password, i.Host, i.Port, i.Schema)
}

type redisOption struct {
	MaxIdle     int    `yaml:"max_idle"`
	MaxActive   int    `yaml:"max_active"`
	IdleTimeout int    `yaml:"idle_timeout"`
	Protocol    string `yaml:"protocol"`
}

type redis struct {
	Host 		string	`yaml:"host"`
	Port 		int 	`yaml:"port"`
	Username 	string 	`yaml:"username"`
	Password 	string 	`yaml:"password"`
	Index 		int 	`yaml:"index"`
	RedisOption redisOption `yaml:"option"`
}

func (i *redis) Addr() string {
	return fmt.Sprintf("%s:%d", i.Host, i.Port)
}


type nacos struct {
	Nacos _nacos `yaml:"nacos"`
}

type _nacos struct {
	DataId string `yaml:"dataId"`
	Group string `yaml:"group"`
	IpAddr string `yaml:"ipAddr"`
	ContextPath string `yaml:"contextPath"`
	Port uint64 `yaml:"port"`
}

var clientConfig = constant.ClientConfig{
	TimeoutMs:           10 * 1000,
	BeatInterval:        5 * 1000,
	ListenInterval:      300 * 1000,
	NotLoadCacheAtStart: true,
	//Username:            "nacos",
	//Password:            "nacos",
}

var serverConfig = constant.ServerConfig{
	IpAddr:      "",
	Port:        80,
	ContextPath: "",
}

func createConfigClient(n *nacos) config_client.ConfigClient {
	nc := nacos_client.NacosClient{}
	nc.SetServerConfig([]constant.ServerConfig{
		{
			IpAddr:      n.Nacos.IpAddr,
			Port:        n.Nacos.Port,
			ContextPath: n.Nacos.ContextPath,
		},
	})
	nc.SetClientConfig(constant.ClientConfig{
		TimeoutMs:           10 * 1000,
		BeatInterval:        5 * 1000,
		ListenInterval:      300 * 1000,
		NotLoadCacheAtStart: true,
		//Username:            "nacos",
		//Password:            "nacos",
	})
	nc.SetHttpAgent(&http_agent.HttpAgent{})
	client, _ := config_client.NewConfigClient(&nc)
	return client
}


func init() {
	flag.Parse()

	Path = *pPath

	var (
		n nacos
		ncfg NCfg
	)
	file, err := os.Open(Path + "nacos.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	fileBody, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = yaml.Unmarshal(fileBody, &n)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := createConfigClient(&n)
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: n.Nacos.DataId,
		Group:  n.Nacos.Group,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	var contentByte = []byte(content)
	err = yaml.Unmarshal(contentByte,&ncfg)

	if err != nil {
		log.Fatal(err.Error())
	}

	Redis = &ncfg.Redis
	Mysql = &ncfg.Mysql



	/*err = client.ListenConfig(vo.ConfigParam{
		DataId: n.Nacos.DataId,
		Group:  n.Nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			//TODO if nacos config changed
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", data:" + data)
			content = data
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}*/



}