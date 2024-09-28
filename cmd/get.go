/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	controller "github.com/miyamoto2024/cli-pokemon-api/controller"
	"github.com/spf13/cobra"
)

// Used for flags.
var id int

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.MaximumNArgs(0),
	Short: "ポケモンのIDを指定してポケモンの情報を取得します。",
	Long: `ポケモンのIDを指定してポケモンの情報を取得します。現在は一つしか指定できません。
	例：
	------------------------------------------------------------
	go run main.go get -i 1
	------------------------------------------------------------
	go run main.go get -i 20
	------------------------------------------------------------
`,
	Run: func(cmd *cobra.Command, args []string) {
		controller.DispPokemonFetchByIdController(id)
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	// バインド
	getCmd.Flags().IntVarP(&id, "id", "i", 0, "Pokemon ID")
	// 必須
	getCmd.MarkFlagRequired("id")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
