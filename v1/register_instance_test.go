package v1

import (
	"testing"
)

/**
* @Author: jack.walker
* @File: register_instance_test.go
* @CreateDate: 2025/11/21 16:24
* @Version: 1.0.0
* @Description:
 */

func TestEurekaClientV1_RegisterInstance(t *testing.T) {
	client := CreateEurekaClientV1("http://10.0.199.177:8080")
	client.Debug = true
	instance := client.BuildInstance("laojia2", "1.1.1.1", 80)
	
	err := client.RegisterInstance(instance)
	if err != nil {
		t.Error(err)
	}
	t.Log("RegisterInstance success")
}
