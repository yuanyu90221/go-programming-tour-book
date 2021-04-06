package cmd

import (
	"log"

	"github.com/spf13/cobra"
	timer "github.com/yuanyu90221/go-programming-tour-book/tour/internal/timer"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "時間格式處理",
	Long:  "時間格式處理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "取得目前時間",
	Long:  "取得目前時間",
	Run: func(cmd *cobra.Command, args []string) {

		nowTime := timer.GetNowTime()
		log.Printf("輸出結果: %s, %d", nowTime.Format("2006-01-0215:04:05"), nowTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
}
