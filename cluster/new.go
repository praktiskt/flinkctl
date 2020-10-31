package cluster

import (
	"fmt"
	"net/url"
	"os"

	"github.com/magnusfurugard/flinkctl/tools"
)

type Cluster struct {
	HostURL     url.URL
	ConfigURL   url.URL
	DatasetsURL url.URL
	OverviewURL url.URL
	Jobs        struct {
		URL         url.URL
		MetricsURL  url.URL
		OverviewURL url.URL
	}
	Jars struct {
		URL       url.URL
		UploadURL url.URL
	}
	Jobmanager struct {
		ConfigURL  url.URL
		LogsURL    url.URL
		MetricsURL url.URL
	}
	Taskmanagers struct {
		URL        url.URL
		MetricsURL url.URL
	}
	// TODO: Add support to use custom headers when calling cluster
	Headers struct{}
}

func New(hostURL string) Cluster {
	url, err := url.Parse(hostURL)
	if err != nil {
		fmt.Printf("Invalid cluster url: %v\n", hostURL)
		os.Exit(1)
	}

	h := *url
	cl := Cluster{}

	cl.HostURL = h
	cl.ConfigURL = tools.UrlOrFail(h, "/config")
	cl.ConfigURL = tools.UrlOrFail(h, "/overview")
	cl.DatasetsURL = tools.UrlOrFail(h, "/datasets")

	cl.Jobs.MetricsURL = tools.UrlOrFail(h, "/jobs/metrics")
	cl.Jobs.OverviewURL = tools.UrlOrFail(h, "/jobs/overview")

	cl.Jars.URL = tools.UrlOrFail(h, "/jars")
	cl.Jars.UploadURL = tools.UrlOrFail(h, "/jars/upload")

	cl.Jobmanager.ConfigURL = tools.UrlOrFail(h, "/jobmanager/config")
	cl.Jobmanager.LogsURL = tools.UrlOrFail(h, "/jobmanager/logs")
	cl.Jobmanager.MetricsURL = tools.UrlOrFail(h, "/jobmanager/metrics")

	cl.Taskmanagers.URL = tools.UrlOrFail(h, "/taskmanagers")
	cl.Taskmanagers.MetricsURL = tools.UrlOrFail(h, "/taskmanagers/metrics")

	return cl
}
