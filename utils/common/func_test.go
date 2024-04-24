package common

import (
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	res, err := GetOutboundIP()
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}