package consul

import "testing"

func TestClient_RegisterService(t *testing.T) {
	reg := DiscoveryConfig{
		Name:    "demo1",
		Tags:    []string{"demo"},
		Port:    0,
		Address: "",
		Ip:      "",
	}
}
