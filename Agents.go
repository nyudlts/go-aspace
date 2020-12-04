package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (a *ASClient) GetAgentIds(agentType string) ([]int, error) {
	var agentIDs = []int{}
	endpoint := fmt.Sprintf("/agents/%s?all_ids=true", agentType)
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

func (a *ASClient) GetAgent(agentType string, agentID int) (Agent, error) {
	agent := Agent{}
	endpoint := fmt.Sprintf("/agents/%s/%d", agentType, agentID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return agent, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return agent, err
	}

	err = json.Unmarshal(body, &agent)
	return agent, nil
}

func (a *ASClient) CreateAgent(agentType string, agent Agent) (string, error) {

	endpoint := fmt.Sprintf("/agents/%s/", agentType)
	body, err := json.Marshal(agent)
	if err != nil {
		return "", err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return "", err
	}

	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(r), err
}

func (a *ASClient) UpdateAgent(agentType string, agentId int, agent Agent) (string, error) {

	endpoint := fmt.Sprintf("/agents/%s/%d", agentType, agentId)
	body, err := json.Marshal(agent)
	if err != nil {
		return "", err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return "", err
	}

	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(r), err
}

func (a *ASClient) DeleteAgent(agentType string, agentId int) (string, error) {

	endpoint := fmt.Sprintf("/agents/%s/%d", agentType, agentId)
	response, err := a.delete(endpoint)
	if err != nil {
		return "", err
	}

	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(r), err
}

func (a *ASClient) GetRandomAgentID(agentType string) (int, error) {
	agent := 0
	agentIDs, err := a.GetAgentIds(agentType)
	if err != nil {
		return agent, err
	}
	return agentIDs[rGen.Intn(len(agentIDs))], nil
}
