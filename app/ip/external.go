package ip

import (
	"io"
	"net/http"
	"update_dns_record/app/domain/ip"
)

func NewAnswer(url string) ip.Answer {
	return &ExternalApiProvider{url: url}
}

type ExternalApiProvider struct {
	url string
}

func (e ExternalApiProvider) PublicIp() (*ip.Ip, error) {

	resp, err := http.Get(e.url)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &ip.Ip{IpType: ip.IPv4, Value: string(body)}, nil
}
