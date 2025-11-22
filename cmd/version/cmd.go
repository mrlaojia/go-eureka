package version

import (
	"github.com/spf13/cobra"
	"log"
)

/**
* @Author: jack.walker
* @File: cmd.go
* @CreateDate: 2025/11/22 08:26
* @Version: 1.0.0
* @Description:
 */

// VersionCmd 创建一个命令
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number for eureka-client`,
	
	// 位置参数，这里接受
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("v1.0")
	},
}
