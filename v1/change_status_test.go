package v1

import "testing"

/**
* @Author: jack.walker
* @File: change_status_test.go
* @CreateDate: 2025/11/21 16:43
* @Version: 1.0.0
* @Description:
 */

func TestEurekaClientV1_ChangeStatus(t *testing.T) {
	client := CreateEurekaClientV1("http://10.0.199.177:8080")
	client.Debug = true
	instance := client.BuildInstance("LAOJIA2", "1.1.1.1", 80)
	
	err := client.ChangeStatus(instance, "OUT_OF_SERVICE")
	if err != nil {
		t.Error(err)
	}
	t.Log("ChangeStatus success")
}
