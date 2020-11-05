package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)



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
