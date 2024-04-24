package consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)

type DiscoveryConfig struct {
	ID      string
	Name    string
	Tags    []string
	Port    int
	Address string
	Ip      string
}

func (dc *DiscoveryConfig) GetId() string {
	if dc.ID == "" {
		return dc.BuildId()
	}

	return dc.ID
}

func (dc *DiscoveryConfig) BuildId() string {
	return fmt.Sprintf("%s-%s-%d", dc.Name, dc.Ip, dc.Port) // 服务唯一ID
}

func (dc *DiscoveryConfig) GetCheckTcpAddr() string {
	return fmt.Sprintf("%s:%d", dc.Address, dc.Port)
}

type Client struct {
	client *consulapi.Client
	cfg    DiscoveryConfig
}

func NewConsul(config DiscoveryConfig) (*Client, error) {
	cfg := consulapi.DefaultConfig()
	cfg.Address = config.Address
	c, err := consulapi.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{client: c, cfg: config}, nil
}
