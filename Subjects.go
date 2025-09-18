package aspace

import (
	"encoding/json"
	"fmt"
	"io"
)

func (a *ASClient) CreateSubject(subject *Subject) (*APIResponse, error) {
	endpoint := "/subjects"
	body, err := json.Marshal(subject)
	if err != nil {
		return nil, err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	apiResponse := &APIResponse{}
	if err = json.Unmarshal(responseBody, apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}
	return apiResponse, nil
}

func (a *ASClient) GetSubject(subjectID int) (*Subject, error) {

	endpoint := fmt.Sprintf("/subjects/%d", subjectID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	subject := &Subject{}
	err = json.Unmarshal(body, &subject)
	if err != nil {
		return subject, err
	}

	return subject, nil
}

func (a *ASClient) UpdateSubject(subjectID int, subject *Subject) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/subjects/%d", subjectID)
	body, err := json.Marshal(subject)
	if err != nil {
		return nil, err
	}

	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := &APIResponse{}
	if err = json.Unmarshal(responseBody, apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	return apiResponse, nil
}

func (a *ASClient) DeleteSubject(subjectID int) (*APIResponse, error) {

	endpoint := fmt.Sprintf("/subjects/%d", subjectID)
	response, err := a.delete(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := &APIResponse{}
	if err := json.Unmarshal(body, apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}
	return apiResponse, nil
}

func (a *ASClient) GetSubjectIDs() ([]int, error) {
	subjectIDs := []int{}

	endpoint := "/subjects?all_ids=true"
	response, err := a.get(endpoint, true)
	if err != nil {
		return subjectIDs, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return subjectIDs, err
	}

	err = json.Unmarshal(body, &subjectIDs)
	if err != nil {
		return subjectIDs, err
	}

	return subjectIDs, nil
}

func (a *ASClient) GetRandomSubjectID() (int, error) {
	subjectID := 0
	subjectIDs, err := a.GetSubjectIDs()
	if err != nil {
		return subjectID, err
	}

	return subjectIDs[rGen.Intn(len(subjectIDs))], nil
}
