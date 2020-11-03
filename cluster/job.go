package cluster

// TODO: Once we have generics, refactor all Describe* functions :)

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JobDescription struct {
	Duration    int64  `json:"duration" header:"duration"`
	EndTime     int64  `json:"end-time" header:"end-time"`
	IsStoppable bool   `json:"isStoppable" header:"isStoppable"`
	Jid         string `json:"jid" header:"jid"`
	Name        string `json:"name" header:"name"`
	Now         int64  `json:"now" header:"now"`
	Plan        struct {
		Jid   string `json:"jid"`
		Name  string `json:"name"`
		Nodes []struct {
			Description string `json:"description"`
			ID          string `json:"id"`
			Inputs      []struct {
				Exchange     string `json:"exchange"`
				ID           string `json:"id"`
				Num          int64  `json:"num"`
				ShipStrategy string `json:"ship_strategy"`
			} `json:"inputs"`
			Operator            string   `json:"operator"`
			OperatorStrategy    string   `json:"operator_strategy"`
			OptimizerProperties struct{} `json:"optimizer_properties"`
			Parallelism         int64    `json:"parallelism"`
		} `json:"nodes"`
	} `json:"plan"`
	StartTime    int64  `json:"start-time" header:"start-time"`
	State        string `json:"state" header:"state"`
	StatusCounts struct {
		Canceled    int64 `json:"CANCELED"`
		Canceling   int64 `json:"CANCELING"`
		Created     int64 `json:"CREATED"`
		Deploying   int64 `json:"DEPLOYING"`
		Failed      int64 `json:"FAILED"`
		Finished    int64 `json:"FINISHED"`
		Reconciling int64 `json:"RECONCILING"`
		Running     int64 `json:"RUNNING"`
		Scheduled   int64 `json:"SCHEDULED"`
	} `json:"status-counts"`
	Timestamps struct {
		Canceled    int64 `json:"CANCELED"`
		Cancelling  int64 `json:"CANCELLING"`
		Created     int64 `json:"CREATED"`
		Failed      int64 `json:"FAILED"`
		Failing     int64 `json:"FAILING"`
		Finished    int64 `json:"FINISHED"`
		Reconciling int64 `json:"RECONCILING"`
		Restarting  int64 `json:"RESTARTING"`
		Running     int64 `json:"RUNNING"`
		Suspended   int64 `json:"SUSPENDED"`
	} `json:"timestamps"`
	Vertices []struct {
		Duration int64  `json:"duration"`
		EndTime  int64  `json:"end-time"`
		ID       string `json:"id"`
		Metrics  struct {
			ReadBytes            int64 `json:"read-bytes"`
			ReadBytesComplete    bool  `json:"read-bytes-complete"`
			ReadRecords          int64 `json:"read-records"`
			ReadRecordsComplete  bool  `json:"read-records-complete"`
			WriteBytes           int64 `json:"write-bytes"`
			WriteBytesComplete   bool  `json:"write-bytes-complete"`
			WriteRecords         int64 `json:"write-records"`
			WriteRecordsComplete bool  `json:"write-records-complete"`
		} `json:"metrics"`
		Name        string `json:"name"`
		Parallelism int64  `json:"parallelism"`
		StartTime   int64  `json:"start-time"`
		Status      string `json:"status"`
		Tasks       struct {
			Canceled    int64 `json:"CANCELED"`
			Canceling   int64 `json:"CANCELING"`
			Created     int64 `json:"CREATED"`
			Deploying   int64 `json:"DEPLOYING"`
			Failed      int64 `json:"FAILED"`
			Finished    int64 `json:"FINISHED"`
			Reconciling int64 `json:"RECONCILING"`
			Running     int64 `json:"RUNNING"`
			Scheduled   int64 `json:"SCHEDULED"`
		} `json:"tasks"`
	} `json:"vertices"`
}

func (c *Cluster) DescribeJob(jid string) (JobDescription, error) {
	//TODO: Respect headers
	if len(jid) != 32 {
		return JobDescription{}, fmt.Errorf("invalid jid: %v", jid)
	}
	re, err := http.Get(c.Jobs.URL.String() + "/" + jid)
	if err != nil {
		return JobDescription{}, err
	}
	defer re.Body.Close()

	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return JobDescription{}, err
	}

	s := JobDescription{}
	err = json.Unmarshal(body, &s)
	if err != nil {
		return JobDescription{}, err
	}
	return s, nil

}
