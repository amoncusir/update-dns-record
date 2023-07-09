package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"update_dns_record/app/ip"
)

var rawPrint bool

var publicIpCmd = &cobra.Command{
	Use:   "publicIp",
	Short: "Get our network public IP",
	Run: func(cmd *cobra.Command, args []string) {

		server, err := cmd.Flags().GetString("server")

		if err != nil {
			log.Fatalf("Error to answer for public Ip. %v\n", err)
		}

		answer := ip.NewAnswer(server)

		publicIp, err := answer.PublicIp()

		if err != nil {
			log.Fatalf("Error to answer for public Ip. %v\n", err)
		}

		if rawPrint {
			fmt.Print(publicIp.Value)
		} else {
			fmt.Printf("Public Ip: %v\n", publicIp.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(publicIpCmd)

	publicIpCmd.Flags().BoolVarP(&rawPrint, "raw", "r", false, "Return only IP value")
}
