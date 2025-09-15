package aspace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestPersonAgents(t *testing.T) {

	var (
		agentPerson *Agent
		agentID     int
	)

	t.Run("test unmarshaling person agents", func(t *testing.T) {

		person, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "agent_person.json"))
		if err != nil {
			t.Error(err)
		}

		agentPerson = &Agent{}
		err = json.Unmarshal(person, agentPerson)
		if err != nil {
			t.Errorf("Error unmarshalling person agent: %v", err)
		}
	})

	t.Run("test create an agent", func(t *testing.T) {
		var err error
		apiResponse, err := testClient.CreateAgent("people", *agentPerson)
		if err != nil {
			t.Errorf("Error creating agent: %v", err)
		}

		if apiResponse.Status != "Created" {
			t.Errorf("Expected status 'Created', got '%s'", apiResponse.Status)
		}

		agentID = apiResponse.ID
	})

	t.Run("test get a person agent", func(t *testing.T) {
		var err error
		agentPerson, err = testClient.GetAgent("people", agentID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully requested and serialized person agent %s\n", agentPerson.URI)

	})

	t.Run("test update an person agent", func(t *testing.T) {
		updatedSortName := "Updated Sort Name"
		agentPerson.Names[0].SortName = updatedSortName

		apiResponse, err := testClient.UpdateAgent("people", agentID, agentPerson)
		if err != nil {
			t.Error(err)
		}

		if apiResponse.Status != "Updated" {
			t.Logf("Expected status 'Updated', got '%s'", apiResponse.Status)
		}

		t.Logf("Successfully updated agent %s\n", agentPerson.URI)
	})

	t.Run("test delete an person agent", func(t *testing.T) {
		var err error
		apiResponse, err := testClient.DeleteAgent("people", agentID)
		if err != nil {
			t.Error(err)
		}

		if apiResponse.Status != "Deleted" {
			t.Errorf("Expected status 'Deleted', got '%s'", apiResponse.Status)
		}

		t.Logf("Successfully deleted agent with ID %d\n", agentID)
	})

	t.Run("test get a random person agent", func(t *testing.T) {
		agentId, err := testClient.GetRandomAgentID("people")
		if err != nil {
			t.Error(err)
		}

		agent, err := testClient.GetAgent("people", agentId)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Successfully requested and serialized person agent %s: %s\n", agent.URI, agent.Names[0].SortName)
		}
	})
}

func TestFamilyAgents(t *testing.T) {
	var (
		agentFamily *Agent
	)
	t.Run("test unmarshalling example family agent", func(t *testing.T) {
		family, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "agent_family.json"))
		if err != nil {
			t.Error(err)
		}
		agentFamily = &Agent{}
		if err := json.Unmarshal(family, agentFamily); err != nil {
			t.Errorf("Error unmarshalling family agent: %v", err)
		}

		t.Logf("Successfully unmarshalled family agent: %s", agentFamily.Names[0].SortName)
	})

}

func TestCorporateEntityAgents(t *testing.T) {
	agentCorporate := &Agent{}
	t.Run("test unmarshalling example corporate entity agent", func(t *testing.T) {
		corporateEntity, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "agent_corporate.json"))
		if err != nil {
			t.Error(err)
		}
		agentCorporate = &Agent{}
		if err := json.Unmarshal(corporateEntity, agentCorporate); err != nil {
			t.Errorf("Error unmarshalling corporate entity agent: %v", err)
		}
		t.Logf("Successfully unmarshalled corporate entity agent: %s", agentCorporate.Names[0].SortName)
	})
}

func TestSoftwareAgents(t *testing.T) {
	var agentSoftware = &Agent{}
	t.Run("test unmarshalling example software agent", func(t *testing.T) {
		software, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "agent_software.json"))
		if err != nil {
			t.Error(err)
		}

		if err := json.Unmarshal(software, agentSoftware); err != nil {
			t.Errorf("Error unmarshalling software agent: %v", err)
		}
		t.Logf("Successfully unmarshalled software agent: %s", agentSoftware.Names[0].SortName)
	})

	t.Run("Test get a random software agent", func(t *testing.T) {
		agentId, err := testClient.GetRandomAgentID("software")
		if err != nil {
			t.Error(err)
		}

		agent, err := testClient.GetAgent("software", agentId)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Successfully requested and serialized software agent %s: %s\n", agent.URI, agent.Names[0].SortName)
		}
	})
}
