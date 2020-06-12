package example

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/utils"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func RegisterService(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, _ := client.RegisterInstance(param)
	fmt.Println(success)
}

func DeregisterService(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	success, _ := client.DeregisterInstance(param)
	fmt.Println(success)
}

func GetService(client naming_client.INamingClient) {
	service, _ := client.GetService(vo.GetServiceParam{
		ServiceName: "demo.go",
		Clusters:    []string{"a"},
	})
	fmt.Println(utils.ToJsonString(service))
}

func Subscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Subscribe(param)
}

func UnSubscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Unsubscribe(param)
}

