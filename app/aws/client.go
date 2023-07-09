package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsDns "update_dns_record/app/aws/dns"
	"update_dns_record/app/domain/dns"
)

type Client struct {
	config aws.Config
}

func NewClient(config aws.Config) *Client {
	return &Client{config: config}
}

func (c Client) Dns(hostedZone string) dns.Client {
	return awsDns.NewClient(hostedZone, c.config)
}
