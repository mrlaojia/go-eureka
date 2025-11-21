package v1

import "encoding/xml"

/**
* @Author: jack.walker
* @File: struct.go
* @CreateDate: 2025/11/21 15:34
* @Version: 1.0.0
* @Description: 通过 curl http://10.0.199.177:8080/eureka/apps/xxx 获取
 */

type Instance struct {
	XMLName                 xml.Name   `xml:"instance"`
	InstanceID              string     `xml:"instanceId"`
	HostName                string     `xml:"hostName"`
	App                     string     `xml:"app"`
	IPAddr                  string     `xml:"ipAddr"`
	Status                  string     `xml:"status"`
	Port                    Port       `xml:"port"`
	SecurePort              Port       `xml:"securePort"`
	DataCenterInfo          DataCenter `xml:"dataCenterInfo"`
	LeaseInfo               LeaseInfo  `xml:"leaseInfo,omitempty"`
	Metadata                Metadata   `xml:"metadata,omitempty"`
	HomePageUrl             string     `xml:"homePageUrl"`
	StatusPageUrl           string     `xml:"statusPageUrl"`
	HealthCheckUrl          string     `xml:"healthCheckUrl"`
	VipAddress              string     `xml:"vipAddress"`
	SecureVipAddress        string     `xml:"secureVipAddress"`
	IsCoordinatingDiscovery bool       `xml:"isCoordinatingDiscoveryServer"`
	LastUpdatedTimestamp    int64      `xml:"lastUpdatedTimestamp"`
	LastDirtyTimestamp      int64      `xml:"lastDirtyTimestamp"`
	ActionType              string     `xml:"actionType"`
}

type Port struct {
	Enabled string `xml:"enabled,attr"`
	Value   int    `xml:",chardata"`
}

type DataCenter struct {
	Class string `xml:"class,attr"`
	Name  string `xml:"name"`
}

type LeaseInfo struct {
	RenewalIntervalInSecs int   `xml:"renewalIntervalInSecs"`
	DurationInSecs        int   `xml:"durationInSecs"`
	RegistrationTimestamp int64 `xml:"registrationTimestamp"`
	LastRenewalTimestamp  int64 `xml:"lastRenewalTimestamp"`
	EvictionTimestamp     int64 `xml:"evictionTimestamp"`
	ServiceUpTimestamp    int64 `xml:"serviceUpTimestamp"`
}

type Metadata struct {
	ManagementPort int `xml:"management.port"`
}
