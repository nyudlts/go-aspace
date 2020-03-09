package go_aspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var RepositoryIDs = [3]int{2, 3, 6}

type AspaceInfo struct {
	DatabaseProductName    string `json:"databaseProductName"`
	DatabaseProductVersion string `json:"databaseProductVersion"`
	RubyVersion            string `json:"ruby_version"`
	HostOS                 string `json:"host_os"`
	HostCPU                string `json:"host_cpu"`
	Build                  string `json:"build"`
	ArchivesSpaceVersion   string `json:"archivesSpaceVersion"`
}

func (a AspaceInfo) String() string {
	msg := fmt.Sprintf("== ArchivesSpace Version: %s\n", a.ArchivesSpaceVersion)
	msg = msg + fmt.Sprintf("== Database Type: %s\n", a.DatabaseProductName)
	msg = msg + fmt.Sprintf("== Database Version: %s\n", a.DatabaseProductVersion)
	msg = msg + fmt.Sprintf("== Ruby Version: %s\n", a.RubyVersion)
	msg = msg + fmt.Sprintf("== Host OS: %s\n", a.HostOS)
	msg = msg + fmt.Sprintf("== Host CPU: %s\n", a.HostCPU)
	msg = msg + fmt.Sprintf("== Java Version: %s\n", a.Build)
	return msg
}

func (a *ASClient) GetAspaceInfo() (AspaceInfo, error) {
	var aspaceInfo AspaceInfo
	response, err := a.get("", false)
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

func (a *ASClient) GetResourceIDsByRepository(repositoryId int) ([]int, error) {
	var repositoryIds []int
	endpoint := fmt.Sprintf("/repositories/%d/resources?all_ids=true", repositoryId)
	response, err := a.get(endpoint, true)
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

func (a *ASClient) GetResourceByID(repositoryId int, resourceId int) (Resource, error) {
	var resource map[string]interface{}
	r := Resource{}

	endpoint := fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId)
	response, err := a.get(endpoint, true)

	if err != nil {
		return r, err
	}

	body, _ := ioutil.ReadAll(response.Body)

	//check for error response
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return r, err
	}

	if resource["error"] != nil {
		return r, fmt.Errorf("%s", resource["error"])
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (a *ASClient) PostResource(repositoryId int, resourceId int, body string) (*http.Response, error) {
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId)
	response, err := a.post(endpoint, true, body)
	if err != nil {
		return response, err
	} else {
		return response, nil
	}
}

//private functions

func (a *ASClient) do(request *http.Request, authenticated bool) (*http.Response, error) {
	var response *http.Response

	if authenticated {
		request.Header.Add("X-ArchivesSpace-Session", a.sessionKey)
	}

	response, err := a.nclient.Do(request)
	if response.StatusCode != 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return response, fmt.Errorf("ArchivesSpace responded with a non-200:\nstatus-code: %d\n%s", response.StatusCode, string(body))
	}
	if err != nil {
		return response, err
	}

	return response, nil
}
func (a *ASClient) get(endpoint string, authenticated bool) (*http.Response, error) {
	var response *http.Response
	url := a.rootURL + endpoint

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, err
	}

	response, err = a.do(request, authenticated)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (a *ASClient) post(endpoint string, authenticated bool, body string) (*http.Response, error) {
	var response *http.Response
	url := a.rootURL + endpoint
	request, err := http.NewRequest("Post", url, bytes.NewBufferString(body))
	if err != nil {
		return response, err
	}

	response, err = a.do(request, authenticated)
	if err != nil {
		return response, err
	}

	return response, nil
}
