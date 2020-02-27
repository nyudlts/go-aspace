package go_aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var client *ASClient

type AspaceInfo struct {
	DatabaseProductName    string `json:"databaseProductName"`
	DatabaseProductVersion string `json:"databaseProductVersion"`
	RubyVersion            string `json:"ruby_version"`
	HostOS                 string `json:"host_os"`
	HostCPU                string `json:"host_cpu"`
	Build                  string `json:"build"`
	ArchivesSpaceVersion   string `json:"archivesSpaceVersion"`
}

func GetAspaceInfo() (AspaceInfo, error) {
	var aspaceInfo AspaceInfo

	err := checkClient()
	if err != nil {
		return aspaceInfo, err
	}

	response, err := client.Get("", false)
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
	endpoint := fmt.Sprintf("/repositories/%d/resources?all_ids=true", repositoryId)
	response, err := client.Get(endpoint, true)
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

func GetResourceByID(repositoryId int, resourceId int) (map[string]interface{}, error) {
	var resource map[string]interface{}
	err := checkClient()
	if err != nil {
		return resource, err
	}
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId)
	response, err := client.Get(endpoint, true)

	if err != nil {
		return resource, err
	}

	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return resource, err
	}
	return resource, nil
}

//private functions

func checkClient() error {
	if client == nil {
		ASC, err := NewClient(10)
		if err != nil {
			return err
		}
		client = ASC
	}

	return nil
}

func (client *ASClient) Get(endpoint string, authenticated bool) (*http.Response, error) {

	var response *http.Response

	url := client.rootURL + endpoint

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return response, err
	}

	if authenticated {
		request.Header.Set("X-ArchivesSpace-Session", client.sessionKey)
	}

	response, err = client.nclient.Do(request)

	if err != nil {
		return response, err
	}

	if response.StatusCode != 200 {
		return response, fmt.Errorf("Did not return a 200 while authenticating, recieved a %d", response.StatusCode)
	}

	return response, nil

}
