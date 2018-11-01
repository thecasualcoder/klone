package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thecasualcoder/klone/pkg/kubeconfig"
	"github.com/thecasualcoder/klone/pkg/resource/deployment"
)

var cfgFile string
var kubeCfgFile string
var fromCluster string
var toCluster string
var namespace string

var rootCmd = &cobra.Command{
	Use:   "klone",
	Short: "Clone Kubernetes resources",
	Long:  "Clone Kubernetes resources within a cluster or across different clusters",
	Run: func(cmd *cobra.Command, args []string) {
		fromClusterConfig, err := kubeconfig.ConfigFor(fromCluster, kubeCfgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		toClusterConfig, err := kubeconfig.ConfigFor(toCluster, kubeCfgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		d, err := deployment.Get("utils", namespace, fromClusterConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = deployment.Create(d, namespace, toClusterConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func home() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return home
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.klone.yaml)")
	rootCmd.PersistentFlags().StringVar(&kubeCfgFile, "kube-config", filepath.Join(home(), ".kube", "config"), "kube config file (default is $HOME/.kube/config)")
	rootCmd.PersistentFlags().StringVar(&fromCluster, "from-cluster", "", "context of cluster from which resource will be copied")
	rootCmd.PersistentFlags().StringVar(&toCluster, "to-cluster", "", "context of cluster to which resource will be copied")
	rootCmd.PersistentFlags().StringVar(&namespace, "namespace", "", "namespace to work on")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(home())
		viper.SetConfigName(".klone")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
