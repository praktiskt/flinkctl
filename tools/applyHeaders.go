package tools

import (
	"strings"

	"github.com/magnusfurugard/flinkctl/config"
	"github.com/parnurzeal/gorequest"
)

func ApplyHeadersToRequest(request *gorequest.SuperAgent) *gorequest.SuperAgent {
	for _, header := range config.GetHeaders() {
		parts := strings.Split(header, ": ")
		if len(parts) != 2 {
			panic("could not split header: " + header)
		}
		request = request.AppendHeader(parts[0], parts[1])
	}
	return request
}
