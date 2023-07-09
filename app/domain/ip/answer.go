package ip

type Answer interface {
	PublicIp() (*Ip, error)
}
