package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
