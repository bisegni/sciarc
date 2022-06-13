/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type NodeFlag struct {
	join_ip         []string
	publishing_port int16
}

var NodeFlags = &NodeFlag{}

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Start a dcomp node",
	Long: `A node is a part of a dcomp cluster, it can be a gateway or a worker node.:

Gateway is a management node and worker node is the node act to run code.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node called")
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	// Here you will define your flags and configuration settings.

	// Specifies the ip of other nodes in the cluster
	nodeCmd.PersistentFlags().StringArrayVarP(&NodeFlags.join_ip, "jip", "j", nil, "The ip list of the other member nodes, at least one")
	nodeCmd.PersistentFlags().Int16VarP(&NodeFlags.publishing_port, "port", "p", 4000, "The port where node is")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
