package v1

import (
	"fmt"
	"log"
	"strings"
)

/**
* @Author: jack.walker
* @File: change_status.go
* @CreateDate: 2025/11/21 16:38
* @Version: 1.0.0
* @Description:
 */

// ChangeStatus  更改 application 状态
// UP:             实例运行正常，可被客户端调用
// DOWN:           实例不可用，客户端不应调用
// STARTING:       实例正在启动，尚未可用
// OUT_OF_SERVICE: 实例手动下线或维护，客户端不应调用
// UNKNOWN:        状态未知，通常 Eureka 还未同步
func (e *EurekaClientV1) ChangeStatus(instance *Instance, status string) error {
	url := fmt.Sprintf("%s/eureka/apps/%s/%s/status?value=%s", e.eurekaHost, instance.App, instance.InstanceID, strings.ToUpper(status))
	if e.Debug {
		log.Printf("ChangeStatus request Url: %v", url)
	}

	resp, err := e.put(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("eureka server returned status %s", resp.Status)
	}

	if e.Debug {
		log.Println("Change Status Response:", resp.Status)
	}

	return nil
}
