package set

import (
	"errors"
	v1 "github.com/mrlaojia/go-eureka/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/**
* @Author: jack.walker
* @File: cmd.go
* @CreateDate: 2025/11/22 08:26
* @Version: 1.0.0
* @Description:
 */

// SetCmd 创建一个命令
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "设置应用的状态",
	Long: `设置应用的状态，instanceID=ip:port

示例：
go-eureka.exe set --app=jack2 --ip=1.1.1.1 --port=2025 -e=10.0.199.177:8080 -s 3

应用状态:
* 0=UNKNOWN
* 1=UP
* 2=DOWN
* 3=STARTING
* 4=OUT_OF_SERVICE

`,

	// 位置参数，这里接受
	// 位置参数，这里接受
	RunE: run,
}

var (
	app, ip      string
	port, status int
)

// Eureka 状态常量
const (
	EurekaUnknown = iota
	EurekaUp
	EurekaDown
	EurekaStarting
	EurekaOutOfService
)

// 状态名称映射
var eurekaStatusNames = map[int]string{
	EurekaUnknown:      "UNKNOWN",
	EurekaUp:           "UP",
	EurekaDown:         "DOWN",
	EurekaStarting:     "STARTING",
	EurekaOutOfService: "OUT_OF_SERVICE",
}

// 获取 Eureka 状态
func getEurekaStatus(value int) (string, error) {
	name, ok := eurekaStatusNames[value]
	if !ok {
		return "", errors.New("invalid Eureka status value")
	}
	// 这里可以存到全局变量或者结构体
	return name, nil
}

// 参数设置
func init() {
	SetCmd.Flags().StringVar(&app, "app", "", "必须，app名字")
	SetCmd.Flags().StringVar(&ip, "ip", "", "必须，IP地址")
	SetCmd.Flags().IntVar(&port, "port", 8080, "必须，服务端口")
	SetCmd.Flags().IntVarP(&status, "status", "s", 1, "应用状态。0=UNKNOWN,1=UP,2=DOWN,3=STARTING,4=OUT_OF_SERVICE")

	SetCmd.PersistentFlags().SortFlags = false
	SetCmd.Flags().SortFlags = false

	SetCmd.MarkFlagRequired("app")
	SetCmd.MarkFlagRequired("ip")
	SetCmd.MarkFlagRequired("port")

}

// 执行逻辑
func run(cmd *cobra.Command, args []string) error {

	client := v1.CreateEurekaClientV1(viper.GetString("eureka"))
	client.Debug = viper.GetBool("verbose")

	ins := client.BuildInstance(app, ip, port)

	s, err := getEurekaStatus(status)
	if err != nil {
		return err
	}

	return client.ChangeStatus(ins, s)
}
