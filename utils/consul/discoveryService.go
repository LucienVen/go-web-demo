package consul

import consulapi "github.com/hashicorp/consul/api"

// 服务发现

const (
	ServerAddr string = "8.134.52.222:8500"
)

func Discovery(serviceName string) ([]*consulapi.ServiceEntry, error) {
	config := consulapi.DefaultConfig()
	config.Address = ServerAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}

	service, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	return service, nil



}