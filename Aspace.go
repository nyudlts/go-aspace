package go_aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var aspaceClient *ASClient

type AspaceInfo struct {
	DatabaseProductName    string `json:"databaseProductName"`
	DatabaseProductVersion string `json:"databaseProductVersion"`
	RubyVersion            string `json:"ruby_version"`
	HostOS                 string `json:"host_os"`
	HostCPU                string `json:"host_cpu"`
	Build                  string `json:"build"`
	ArchivesSpaceVersion   string `json:"archivesSpaceVersion"`
}

func checkClient() error {
	if aspaceClient == nil {
		ASC, err := NewClient(10)
		if err != nil {
			return err
		}
		aspaceClient = ASC
	}

	return nil
}

func GetAspaceInfo() (AspaceInfo, error) {
	var aspaceInfo AspaceInfo

	err := checkClient()
	if err != nil {
		return aspaceInfo, err
	}

	response, err := aspaceClient.asGet("", false)
	if err != nil {
		return aspaceInfo, err
	}
	body, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &aspaceInfo)
	if err != nil {
		return aspaceInfo, err
	}
	return aspaceInfo, nil
}

func GetResourceIDsByRepository(repositoryId int) ([]int, error) {
	var repositoryIds []int
	err := checkClient()
	if err != nil {
		return repositoryIds, err
	}
	response, err := aspaceClient.asGet(fmt.Sprintf("/repositories/%d/resources?all_ids=true", repositoryId), true)
	if err != nil {
		return repositoryIds, err
	}
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &repositoryIds)
	if err != nil {
		return repositoryIds, err
	}
	return repositoryIds, nil
}

//private functions
func (asClient *ASClient) asGet(endpoint string, authenticated bool) (*http.Response, error) {

	var response *http.Response

	url := aspaceClient.rootURL + endpoint

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return response, err
	}

	if authenticated {
		request.Header.Set("X-ArchivesSpace-Session", aspaceClient.sessionKey)
	}

	response, err = aspaceClient.nclient.Do(request)

	if err != nil {
		return response, err
	}

	if response.StatusCode != 200 {
		return response, fmt.Errorf("Did not return a 200 while authenticating, recieved a %d", response.StatusCode)
	}

	return response, nil

}
