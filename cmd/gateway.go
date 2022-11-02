/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/bisegni/sciarc/services"
	"github.com/spf13/cobra"
)

type NodeFlag struct {
	join_ip         []string
	publishing_port int16
}

var NodeFlags = &NodeFlag{}
var node services.Node

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Start a sciarc gateway node",
	Long: `A gateway node is a part of a frontend sciarc cluster:

Gateway is a management and user request central node`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node called")
		node = services.Node{}
		node.Start()
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
