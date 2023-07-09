package dns

type Client interface {
	ListRecords() (*[]*Record, error)
	CreateOrUpdateRecord(record *Record) error
}
