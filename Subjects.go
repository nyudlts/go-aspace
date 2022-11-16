package aspace

import (
	"encoding/json"
	"fmt"
	"io"
)

func (a *ASClient) GetSubject(subjectID int) (Subject, error) {
	subject := Subject{}

	endpoint := fmt.Sprintf("/subjects/%d", subjectID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return subject, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return subject, err
	}

	err = json.Unmarshal(body, &subject)
	if err != nil {
		return subject, err
	}

	return subject, nil
}

func (a *ASClient) DeleteSubject(subjectID int) (string, error) {
	msg := ""
	endpoint := fmt.Sprintf("/subjects/%d", subjectID)
	response, err := a.delete(endpoint)
	if err != nil {
		return msg, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return msg, err
	}

	return string(body), err
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
