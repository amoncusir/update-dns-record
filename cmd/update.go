package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
	"log"
	"update_dns_record/app/aws"
	"update_dns_record/app/domain/dns"
	"update_dns_record/app/ip"
	"update_dns_record/util/arrays"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update current DNS A register to new IP",
	Run: func(cmd *cobra.Command, args []string) {

		ipFlag, err := cmd.Flags().GetString("ip")

		if err != nil {
			log.Fatalf("Error to answer for public Ip. %v\n", err)
		}

		if ipFlag == "" {
			answer := ip.NewAnswer(cmd.Flag("server").Value.String())

			publicIp, err := answer.PublicIp()

			if err != nil {
				log.Fatalf("Error to answer for public Ip. %v\n", err)
			}

			ipFlag = publicIp.Value
		}

		cfg, err := config.LoadDefaultConfig(context.TODO())

		if err != nil {
			log.Fatalf("failed to load configuration, %v", err)
		}

		hostedZone, err := cmd.Flags().GetString("hostedZone")

		if err != nil {
			log.Fatalf("failed to load configuration, %v", err)
		}

		dnsClient := aws.NewClient(cfg).Dns(hostedZone)

		name, err := cmd.Flags().GetString("name")

		if err != nil {
			log.Fatalf("failed to load configuration, %v", err)
		}

		err = dnsClient.CreateOrUpdateRecord(&dns.Record{
			Name:   name,
			Type:   "A",
			TTL:    3600,
			Values: arrays.SliceOf(ipFlag),
		})

		if err != nil {
			log.Fatalf("failed to load configuration, %v", err)
		}

		fmt.Printf("Updated dns values:\n\tName: %s\n\tValue: %s", name, ipFlag)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().String("ip", "", "Set Ip to add to DNS register")

	updateCmd.Flags().StringP("hostedZone", "z", "", "HostedZone to put the changes")
	updateCmd.Flags().StringP("name", "n", "", "DNS name")

	updateCmd.MarkFlagRequired("hostedZone")
	updateCmd.MarkFlagRequired("name")
}
