package go_aspace

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type AspaceInfo struct {
	DatabaseProductName string	 `json:"databaseProductName"`
	DatabaseProductVersion string	 `json:"databaseProductVersion"`
	RubyVersion	string	 `json:"ruby_version"`
	HostOS	string	 `json:"host_os"`
	HostCPU	string	 `json:"host_cpu"`
	Build	string	 `json:"build"`
	ArchivesSpaceVersion	string	 `json:"archivesSpaceVersion"`
}

func GetAspaceInfo(client *Client) (AspaceInfo, error) {

	response, err := ASGet(client, "", false)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	var aspaceInfo AspaceInfo
	err = json.Unmarshal(body, &aspaceInfo)
	if err != nil {
		return aspaceInfo, err
	}
	return aspaceInfo, nil
}
