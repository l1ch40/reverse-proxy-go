/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"reverse-proxy/server"
	"github.com/spf13/cobra"
)


var serverURL string
var clientURL string
var port int

var rootCmd = &cobra.Command{
	Use:   "reverse-proxy",
	Short: "A Simple And Stupid Reverse Proxy With Golang",
	Long: `
	    This is a simple Http reverse proxy service.

	    When you specify the client address, use the prefix /fe to access it.
	    When you specify the server side address, use the prefix /be to access it.
	`,
	Run: func(cmd *cobra.Command, args []string) {
	    urls := make(map[string]string)
	    if serverURL == "" && clientURL == "" {
		fmt.Printf("Please fill in the client url or server url\n")
		return
	    }
	    if serverURL != "" {
		urls["/be"] = serverURL
	    }
	    if clientURL != "" {
		urls["/fe"] = clientURL
	    }
	    server.Start(urls, port)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVar(&port, "port", 8080, "Proxy Port")
	rootCmd.Flags().StringVar(&serverURL, "serverURL", "", "Server URL")
	rootCmd.Flags().StringVar(&clientURL, "clientURL", "", "Client URL")
}


