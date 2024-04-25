package consul

import (
	consulapi "github.com/hashicorp/consul/api"
)

const (
	AgentServiceCheckOpt_Timeout                        string = "5s"
	AgentServiceCheckOpt_Interval                       string = "5s"
	AgentServiceCheckOpt_DeregisterCriticalServiceAfter string = "60s"
)

// 注册服务
func (c *Client) RegisterService() error {
	srv := &consulapi.AgentServiceRegistration{
		ID:      c.cfg.GetId(),
		Name:    c.cfg.Name,
		Tags:    c.cfg.Tags,
		Port:    c.cfg.Port,
		Address: c.cfg.Address,
	}

	// 健康检查
	check := &consulapi.AgentServiceCheck{
		TCP:      c.cfg.TcpAddr,
		Timeout:  AgentServiceCheckOpt_Timeout,
		Interval: AgentServiceCheckOpt_Interval,
		DeregisterCriticalServiceAfter: AgentServiceCheckOpt_DeregisterCriticalServiceAfter,
	}

	srv.Check = check

	return c.client.Agent().ServiceRegister(srv)
}
