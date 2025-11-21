package v1

import "testing"

/**
* @Author: jack.walker
* @File: get_instance_test.go
* @CreateDate: 2025/11/21 16:56
* @Version: 1.0.0
* @Description:
 */

func TestEurekaClientV1_GetInstance(t *testing.T) {
	client := CreateEurekaClientV1("http://10.0.199.177:8080")
	client.Debug = true
	instance := client.BuildInstance("LAOJIA2", "1.1.1.1", 80)
	
	ins, err := client.GetInstance(instance)
	if err != nil {
		t.Error(err)
	}
	
	t.Log(ins.App, ins.InstanceID, ins.Status)
}
