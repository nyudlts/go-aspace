package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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

func (a *ASClient) Search(repositoryId int, searchType string, search string) {

	endpoint := fmt.Sprintf(`/repositories/%d/search?type[]=%s&page=1`, repositoryId, searchType)
	response, err := a.get(endpoint, true)
	if err != nil {
		panic(err)
	}
	body := response.Body
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyBytes))
}

func (a *ASClient) SerializeEAD(repositoryId int, resoureId int, loc string, daos bool, unpub bool, num_cs bool, ead3 bool, pdf bool, filename string) error {
	var ext string
	endpoint := fmt.Sprintf("/repositories/%d/resource_descriptions/%d.xml?include_unpublished=%t&include_daos=%t&numbered_cs=%t&ead3=%t&print_pdf=%t", repositoryId, resoureId, unpub, daos, num_cs, ead3, pdf)
	response, err := a.get(endpoint, true)
	if err != nil {
		return err
	}

	bodybytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if pdf {
		ext = "pdf"
	} else {
		ext = "xml"
	}

	err = writeEADtoFile(bodybytes, fmt.Sprintf("%s.%s", filename, ext), loc)
	return nil

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

//private functions
func writeEADtoFile(ead []byte, title string, loc string) error {
	f, err := os.Create(fmt.Sprintf("%s/%s", loc, title))
	defer f.Close()
	if err != nil {
		return nil
	}
	f.Write(ead)
	return nil
}
