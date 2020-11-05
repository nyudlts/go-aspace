package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

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

func (a *ASClient) GetArchivalObjectById(repositoryId int, aoId int) (ArchivalObject, error) {

	ao := ArchivalObject{}
	endpoint := fmt.Sprintf("/repositories/%d/archival_objects/%d", repositoryId, aoId)

	reponse, err := a.get(endpoint, true)
	if err != nil {
		return ao, err
	}

	body, err := ioutil.ReadAll(reponse.Body)
	if err != nil {
		return ao, err
	}

	err = json.Unmarshal(body, &ao)
	if err != nil {
		return ao, err
	}

	return ao, nil
}

func (a *ASClient) PrintResponse(endpoint string) error {
	response, err := a.get(endpoint, true)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func (a *ASClient) GetResourceByID(repositoryId int, resourceId int) (Resource, error) {

	r := Resource{}

	endpoint := fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId)
	response, err := a.get(endpoint, true)

	if err != nil {
		return r, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (a *ASClient) GetRepositoryList() ([]int, error) {
	repIds := []int{}
	endpoint := "/repositories"
	response, err := a.get(endpoint, false)
	if err != nil {
		return repIds, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return repIds, err
	}

	reps := make([]map[string]interface{}, 1, 1)
	err = json.Unmarshal(body, &reps)

	for i := range reps {
		rep := fmt.Sprintf("%v", reps[i]["uri"])
		repId, err := strconv.Atoi(rep[len(rep)-1:])
		if err != nil {
			return repIds, err
		}
		repIds = append(repIds, repId)
	}

	return repIds, nil
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

func (a *ASClient) GetEADAsByteArray(repositoryId int, resourceId int) ([]byte, error) {
	eadBytes := []byte{}
	endpoint := fmt.Sprintf("/repositories/%d/resource_descriptions/%d.xml?include_unpublished=%t&include_daos=%t&numbered_cs=%t&ead3=%t&print_pdf=%t", repositoryId, resourceId, false, true, false, false, false)
	response, err := a.get(endpoint, true)
	if err != nil {
		return eadBytes, err
	}

	eadBytes, err = ioutil.ReadAll(response.Body)
	return eadBytes, err
}

func (a *ASClient) GetDigitalObjectsByRepositoryId(repositoryId int) ([]int, error) {
	daoIds := []int{}
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects?all_ids=true", repositoryId)
	response, err := a.get(endpoint, true)
	if err != nil {
		return daoIds, err
	}
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &daoIds)
	if err != nil {
		return daoIds, err
	}

	return daoIds, nil
}

func (a *ASClient) GetDigitalObject(repositoryId int, daoId int) (DigitalObject, error) {
	do := DigitalObject{}
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects/%d", repositoryId, daoId)
	response, err := a.get(endpoint, true)
	if err != nil {
		return do, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return do, err
	}
	err = json.Unmarshal(body, &do)
	if err != nil {
		return do, err
	}

	return do, nil
}

func (a *ASClient) UpdateDigitalObject(repositoryId int, daoId int, dao DigitalObject) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects/%d", repositoryId, daoId)
	body, err := json.Marshal(dao)
	if err != nil {
		return "", err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return "", err
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (a *ASClient) CreateDigitalObject(repositoryId int, dao DigitalObject) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects", repositoryId)
	body, err := json.Marshal(dao)
	if err != nil {
		return "", err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return "", err
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (a *ASClient) DeleteDigitalObject(repositoryId int, doId int) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects/%d", repositoryId, doId)
	response, err := a.delete(endpoint)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return string(body), err
	}
	return string(body), nil
}

func (a *ASClient) GetResourceTree(repositoryId int, resourceId int) (ResourceTree, error) {
	tree := ResourceTree{}
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d/tree", repositoryId, resourceId)
	response, err := a.get(endpoint, true)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return tree, err
	}

	err = json.Unmarshal(body, &tree)
	if err != nil {
		return tree, err
	}
	return tree, nil

}

func (a *ASClient) AdvancedSearch(page int, repositoryId int, queryType string, adQuery string) (SearchResult, error) {
	endpoint := fmt.Sprintf(`/repositories/%d/search?type[]=%s&page=%d&aq=%s`, repositoryId, queryType, page, adQuery)
	//fmt.Println(endpoint)
	response, err := a.get(endpoint, true)
	if err != nil {
		return SearchResult{}, err
	}
	body := response.Body
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return SearchResult{}, err
	}

	results := SearchResult{}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return SearchResult{}, err
	}

	return results, nil

}

func (a *ASClient) Search(repositoryId int, searchType string, query string, page int) (SearchResult, error) {

	endpoint := fmt.Sprintf(`/repositories/%d/search?type[]=%s&q=%s&page=%d`, repositoryId, searchType, query, page)
	response, err := a.get(endpoint, true)
	if err != nil {
		return SearchResult{}, err
	}
	body := response.Body
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return SearchResult{}, err
	}

	results := SearchResult{}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return SearchResult{}, err
	}

	return results, nil
}

func (a *ASClient) SerializeEAD(repositoryId int, resourceId int, daos bool, unpub bool, num_cs bool, ead3 bool, pdf bool) ([]byte, error) {
	var ead []byte = []byte{}

	endpoint := fmt.Sprintf("/repositories/%d/resource_descriptions/%d.xml?include_unpublished=%t&include_daos=%t&numbered_cs=%t&ead3=%t&print_pdf=%t", repositoryId, resourceId, unpub, daos, num_cs, ead3, pdf)
	response, err := a.get(endpoint, true)
	if err != nil {
		return ead, err
	}

	ead, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return ead, err
	}

	return ead, nil

}

func (a *ASClient) GetEndpoint(e string) ([]byte, error) {
	fmt.Println(e)
	response, err := a.get(e, true)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func (a *ASClient) GetPeopleAgentIds() ([]int, error) {
	var agentIDs = []int{}
	endpoint := "/agents/people?all_ids=true"
	response, err := a.get(endpoint, true)
	if err != nil {
		return agentIDs, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return agentIDs, err
	}

	err = json.Unmarshal(body, &agentIDs)
	if err != nil {
		return agentIDs, err
	}

	return agentIDs, nil
}

func (a *ASClient) GetPeopleAgent(agentID int) ([]byte, error) {
	agent := []byte{}
	endpoint := fmt.Sprintf("/agents/people/%d", agentID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return agent, err
	}
	agent, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return agent, err
	}

	return agent, nil
}
