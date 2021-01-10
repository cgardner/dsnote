package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootDir string
	rootCmd = &cobra.Command{
		Use: "dsnote",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running Root Command")
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "Root Directory for notes (default is $HOME/.config/dsnote)")
}

func initConfig() {
	if rootDir == "" {
		rootDir = os.ExpandEnv("$HOME/.config/dsnote/")
	}

	_, err := ioutil.ReadDir(rootDir)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
