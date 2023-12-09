/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/daifukuninja/petit-misskey-go/infrastructure/setting"
	"github.com/daifukuninja/petit-misskey-go/view"
	"github.com/daifukuninja/petit-misskey-go/view/meta"
	"github.com/spf13/cobra"
)

// metaCmd represents the meta command
var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("meta run")

		key, _ := cmd.Flags().GetString("key")

		setting := setting.NewUserSetting()
		instance := setting.GetInstanceByKey(key)

		model := meta.InitializeModel(instance)

		view.Run(model)
	},
}

func init() {
	rootCmd.AddCommand(metaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// metaCmd.Flags().StringP("key", "k", "", "Instance Key")
}
