/*
Copyright © 2025 Laojia wukejia555@126.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/mrlaojia/go-eureka/cmd/version"
	"log"
	"os"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eurekaGo",
	Short: "go eureka client",
	Long:  `一个go语言写的 eureka client，对app的注册信息简单控制`,
	
	// 禁用 errors，使用 Rune 的 error 默认会打印
	// 这里禁用，下面我们再使用自己的格式
	SilenceErrors: true,
	// 默认命令错误，会打印 usage，这里禁用
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err.Error()) // 设置 SilenceErrors 后， 这里打印 Rune 的错误。
		os.Exit(1)
	}
}

var eurekaHost string
var verbose bool

func init() {
	
	rootCmd.PersistentFlags().StringVarP(&eurekaHost, "eureka_host", "e", "", "必须，eureka服务器地址，如 10.0.199.177:8080")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "打印更多信息 more message")
	
	// 禁用 Flags 排序
	// 不影响 Global Flags 部分的显式,它还是会自动排序
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.Flags().SortFlags = false
	
	// 绑定到 viper
	viper.BindPFlags(rootCmd.PersistentFlags())
	
	// 添加子命令
	rootCmd.AddCommand(version.VersionCmd)
	
}
