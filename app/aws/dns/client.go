package dns

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"update_dns_record/app/domain/dns"
	"update_dns_record/util/arrays"
)

func NewClient(hostedZone string, config aws.Config) dns.Client {
	client := route53.NewFromConfig(config)
	return &AwsSdkHostedZoneClient{
		config:     config,
		client:     client,
		hostedZone: hostedZone,
	}
}

type AwsSdkHostedZoneClient struct {
	config     aws.Config
	client     *route53.Client
	hostedZone string
}

func (c AwsSdkHostedZoneClient) ListRecords() (*[]*dns.Record, error) {

	ctx := context.TODO()

	paginator := route53.NewListResourceRecordSetsPaginator(c.client, &route53.ListResourceRecordSetsInput{
		HostedZoneId: &c.hostedZone,
	})

	records := make([]*dns.Record, 0)

	for paginator.HasMorePages() {
		outputs, err := paginator.NextPage(ctx)

		if err != nil {
			return nil, err
		}

		for _, r := range *toRecord(outputs) {
			records = append(records, r)
		}
	}

	return &records, nil
}

func (c AwsSdkHostedZoneClient) CreateOrUpdateRecord(record *dns.Record) error {

	ctx := context.TODO()

	changes := make([]types.Change, 1)

	changes[0] = *toChange(record)

	input := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: &c.hostedZone,
		ChangeBatch: &types.ChangeBatch{
			Changes: changes,
			Comment: nil,
		},
	}

	_, err := c.client.ChangeResourceRecordSets(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

func toRecord(output *route53.ListResourceRecordSetsOutput) *[]*dns.Record {
	records := make([]*dns.Record, len(output.ResourceRecordSets))
	for i, record := range output.ResourceRecordSets {
		r := &dns.Record{
			Name: *record.Name,
			Type: string(record.Type),
			TTL:  *record.TTL,
			Values: arrays.OnMap(record.ResourceRecords, func(t types.ResourceRecord) string {
				return *t.Value
			}),
		}
		records[i] = r
	}

	return &records
}

func toChange(record *dns.Record) *types.Change {
	return &types.Change{
		Action: types.ChangeActionUpsert,
		ResourceRecordSet: &types.ResourceRecordSet{
			Name: &record.Name,
			Type: types.RRType(record.Type),
			ResourceRecords: arrays.OnMap(record.Values, func(t string) types.ResourceRecord {
				return types.ResourceRecord{Value: &t}
			}),
			TTL: &record.TTL,
		},
	}
}
