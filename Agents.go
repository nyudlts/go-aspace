package aspace

import (
	"encoding/json"
	"fmt"
	"io"
)

func (a *ASClient) GetAgentIds(agentType string) ([]int, error) {
	var agentIDs = []int{}
	endpoint := fmt.Sprintf("/agents/%s?all_ids=true", agentType)
	response, err := a.get(endpoint, true)
	if err != nil {
		return agentIDs, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return agentIDs, err
	}

	err = json.Unmarshal(body, &agentIDs)
	if err != nil {
		return agentIDs, err
	}

	return agentIDs, nil
}

func (a *ASClient) GetAgent(agentType string, agentID int) (*Agent, error) {
	agent := &Agent{}
	endpoint := fmt.Sprintf("/agents/%s/%d", agentType, agentID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &agent)
	if err != nil {
		return nil, err
	}
	return agent, nil
}

func (a *ASClient) CreateAgent(agentType string, agent Agent) (*APIResponse, error) {

	endpoint := fmt.Sprintf("/agents/%s", agentType)
	body, err := json.Marshal(agent)
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

	aPIResponse := &APIResponse{}
	err = json.Unmarshal(responseBody, aPIResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling API response: %v", err)
	}

	return aPIResponse, nil
}

func (a *ASClient) UpdateAgent(agentType string, agentId int, agent *Agent) (*APIResponse, error) {

	endpoint := fmt.Sprintf("/agents/%s/%d", agentType, agentId)
	body, err := json.Marshal(agent)
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

	aPIResponse := &APIResponse{}
	if err = json.Unmarshal(responseBody, aPIResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling API response: %v", err)
	}

	return aPIResponse, nil
}

func (a *ASClient) DeleteAgent(agentType string, agentId int) (*APIResponse, error) {

	endpoint := fmt.Sprintf("/agents/%s/%d", agentType, agentId)
	response, err := a.delete(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	aPIResponse := &APIResponse{}
	if err = json.Unmarshal(responseBody, aPIResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling API response: %v", err)
	}

	return aPIResponse, nil

}

func (a *ASClient) GetRandomAgentID(agentType string) (int, error) {
	agent := 0
	agentIDs, err := a.GetAgentIds(agentType)
	if err != nil {
		return agent, err
	}
	return agentIDs[rGen.Intn(len(agentIDs))], nil
}
