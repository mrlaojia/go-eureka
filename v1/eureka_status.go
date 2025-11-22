package v1

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
)

/**
* @Author: jack.walker
* @File: eurekaStatus.go
* @CreateDate: 2025/11/21 16:15
* @Version: 1.0.0
* @Description:
 */

// EurekaStatus Eureka Server 健康状态
// UP:      Eureka Server 正常，服务注册和查询可用
// DOWN:    Eureka Server 不可用
// UNKNOWN: 未知状态，通常出现在错误或未完全启动
func (e *EurekaClientV1) EurekaStatus() (*EurekaStatus, error) {
	url := fmt.Sprintf("%s/eureka/status", e.eurekaHost)
	if e.Debug {
		log.Printf("EurekaStatus request Url: %v", url)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/xml")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("eureka server returned status %s", resp.Status)
	}
	
	status := &EurekaStatus{}
	if err := xml.NewDecoder(resp.Body).Decode(status); err != nil {
		return nil, err
	}
	
	if e.Debug {
		log.Println("EurekaStatus Response:", resp.Status)
	}
	
	return status, nil
}

// EurekaStatus 用于解析 /eureka/ 或 /eureka/status 返回的 XML
type EurekaStatus struct {
	XMLName      xml.Name     `xml:"com.netflix.eureka.util.StatusInfo"`
	GeneralStats GeneralStats `xml:"generalStats"`
	InstanceInfo InstanceInfo `xml:"instanceInfo"`
}

type GeneralStats struct {
	ServerUptime       string `xml:"server-uptime"`
	Environment        string `xml:"environment"`
	Cpus               string `xml:"num-of-cpus"`
	TotalMemory        string `xml:"total-avail-memory"`
	CurrentMemoryUsage string `xml:"current-memory-usage"`
}

type InstanceInfo struct {
	Status   string `xml:"status"`
	HostName string `xml:"hostName"`
	App      string `xml:"app"`
	IpAddr   string `xml:"ipAddr"`
	Port     Port   `xml:"port"`
}
