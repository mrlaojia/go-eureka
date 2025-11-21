package v1

import (
	"testing"
	"time"
)

/**
* @Author: jack.walker
* @File: heartbeat_test.go
* @CreateDate: 2025/11/21 16:49
* @Version: 1.0.0
* @Description:
 */

func TestEurekaClientV1_SendHeartbeat(t *testing.T) {
	client := CreateEurekaClientV1("http://10.0.199.177:8080")
	client.Debug = true
	instance := client.BuildInstance("laojia", "1.1.1.1", 80)
	
	//发送心跳 3 次
	for i := 0; i < 3; i++ {
		time.Sleep(5 * time.Second)
		err := client.SendHeartbeat(instance)
		if err != nil {
			t.Error(err)
		}
	}
	
	t.Log("SendHeartbeat success")
}
