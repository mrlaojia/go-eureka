package heartbeat

import (
	v1 "github.com/mrlaojia/go-eureka/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"time"
)

/**
* @Author: jack.walker
* @File: cmd.go
* @CreateDate: 2025/11/22 08:26
* @Version: 1.0.0
* @Description:
 */

// HeartbeatCmd 创建一个命令
var HeartbeatCmd = &cobra.Command{
	Use:   "heartbeat",
	Short: "向 eureka 发送应用信息的心跳",
	Long: `向 eureka 发送应用信息的心跳，instanceID=ip:port

示例：
go-eureka.exe heartbeat --app=jack2 --ip=1.1.1.1 --port=2025 -e=10.0.199.177:8080 -i 5 -c 2
`,

	// 位置参数，这里接受
	// 位置参数，这里接受
	RunE: run,
}

var (
	app, ip               string
	port, count, interval int
)

// 参数设置
func init() {
	HeartbeatCmd.Flags().StringVar(&app, "app", "", "必须，app名字")
	HeartbeatCmd.Flags().StringVar(&ip, "ip", "", "必须，IP地址")
	HeartbeatCmd.Flags().IntVar(&port, "port", 8080, "必须，服务端口")
	HeartbeatCmd.Flags().IntVarP(&count, "count", "c", 3, "发送几次心跳，0: 一直发送")
	HeartbeatCmd.Flags().IntVarP(&interval, "interval", "i", 5, "发送心跳的间隔，单位: 秒")

	HeartbeatCmd.PersistentFlags().SortFlags = false
	HeartbeatCmd.Flags().SortFlags = false

	HeartbeatCmd.MarkFlagRequired("app")
	HeartbeatCmd.MarkFlagRequired("ip")
	HeartbeatCmd.MarkFlagRequired("port")

}

// 执行逻辑
func run(cmd *cobra.Command, args []string) error {

	client := v1.CreateEurekaClientV1(viper.GetString("eureka"))
	client.Debug = viper.GetBool("verbose")

	ins := client.BuildInstance(app, ip, port)

	//不存在也会发送心跳成功
	var times = new(int64)
	*times = 1

	if count <= 0 {
		for {
			err := send(client, ins, times)
			if err != nil {
				return err
			}
		}
	} else {
		for i := 0; i < count; i++ {
			err := send(client, ins, times)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func send(client *v1.EurekaClientV1, ins *v1.Instance, times *int64) error {
	err := client.SendHeartbeat(ins)
	if err != nil {
		return err
	}
	log.Printf("[%v] SendHeartbeat success. ", *times)
	time.Sleep(time.Duration(interval) * time.Second)
	*times++
	return nil
}
