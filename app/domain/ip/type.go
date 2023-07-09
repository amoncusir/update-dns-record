package ip

type IpType uint8

const (
	IPv4 IpType = 4
	IPv6 IpType = 6
)

type Ip struct {
	IpType IpType
	Value  string
}
