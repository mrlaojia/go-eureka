package v1

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

/**
* @Author: jack.walker
* @File: get_instance.go
* @CreateDate: 2025/11/21 16:53
* @Version: 1.0.0
* @Description:
 */

func (e *EurekaClientV1) GetInstance(i *Instance) (*Instance, error) {
	url := fmt.Sprintf("%s/eureka/apps/%s/%s", e.eurekaHost, strings.ToUpper(i.App), i.InstanceID)
	if e.Debug {
		fmt.Println("GetInstance url:", url)
	}

	resp, err := e.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Eureka returned status %s", resp.Status)
	}

	ins := &Instance{}
	if err := xml.NewDecoder(resp.Body).Decode(ins); err != nil {
		return nil, err
	}

	if e.Debug {
		log.Println("EurekaStatus Response:", resp.Status)
		log.Printf("EurekaStatus: %v", ins)
	}
	return ins, nil
}
