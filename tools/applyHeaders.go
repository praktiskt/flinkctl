package tools

import (
	"strings"

	"github.com/magnusfurugard/flinkctl/config"
	"github.com/parnurzeal/gorequest"
)

func ApplyBasicAuthToRequest(request *gorequest.SuperAgent) *gorequest.SuperAgent {
	conf := config.GetBasicAuth()
	if len(conf.Username) != 0 && len(conf.Password) != 0 {
		request = request.SetBasicAuth(conf.Username, conf.Password)
	}
	return request
}

func ApplyHeadersToRequest(request *gorequest.SuperAgent) *gorequest.SuperAgent {
	for _, header := range config.GetHeaders() {
		parts := strings.Split(header, ": ")
		if len(parts) != 2 {
			panic("could not split header: " + header)
		}
		request = request.AppendHeader(parts[0], parts[1])
	}
	request = ApplyBasicAuthToRequest(request)
	return request
}
