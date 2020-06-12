package config

import (
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


type config struct {
	Nacos Nacos `yaml:"nacos"`
}

type Nacos struct {
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

func createConfigClient(config *config) config_client.ConfigClient {
	nc := nacos_client.NacosClient{}
	nc.SetServerConfig([]constant.ServerConfig{
		{
			IpAddr:      config.Nacos.IpAddr,
			Port:        config.Nacos.Port,
			ContextPath: config.Nacos.ContextPath,
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


func Do() (content string){
	var config config
	file, err := os.Open("./config/nacos.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	fileBody, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = yaml.Unmarshal(fileBody, &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := createConfigClient(&config)
	content, err = client.GetConfig(vo.ConfigParam{
		DataId: config.Nacos.DataId,
		Group:  config.Nacos.Group,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.ListenConfig(vo.ConfigParam{
		DataId: config.Nacos.DataId,
		Group:  config.Nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", data:" + data)
			content = data
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return

}