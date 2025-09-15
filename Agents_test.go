package aspace

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

var (
	agentPerson *Agent
	client      *ASClient
	agentID     int
)

func TestAgents(t *testing.T) {
	flag.Parse()
	var err error
	client, err = NewClient(goaspacetest.Config, goaspacetest.Environment)
	if err != nil {
		t.Error(err)
	}
	t.Run("test unmarshaling agents", func(t *testing.T) {

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

		apiResponse, err := client.CreateAgent("people", *agentPerson)
		if err != nil {
			t.Errorf("Error creating agent: %v", err)
		}

		if apiResponse.Status != "Created" {
			t.Errorf("Expected status 'Created', got '%s'", apiResponse.Status)
		}

		agentID = apiResponse.ID
	})

	t.Run("test get a person agent", func(t *testing.T) {
		agentPerson, err = client.GetAgent("people", agentID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully requested and serialized person agent %s\n", agentPerson.URI)

	})

	t.Run("test update an agent", func(t *testing.T) {
		updatedSortName := "Updated Sort Name"
		agentPerson.Names[0].SortName = updatedSortName

		apiResponse, err := client.UpdateAgent("people", agentID, agentPerson)
		if err != nil {
			t.Error(err)
		}

		if apiResponse.Status != "Updated" {
			t.Logf("Expected status 'Updated', got '%s'", apiResponse.Status)
		}

		t.Logf("Successfully updated agent %s\n", agentPerson.URI)
	})

	t.Run("test delete an agent", func(t *testing.T) {
		apiResponse, err := client.DeleteAgent("people", agentID)
		if err != nil {
			t.Error(err)
		}

		if apiResponse.Status != "Deleted" {
			t.Errorf("Expected status 'Deleted', got '%s'", apiResponse.Status)
		}

		t.Logf("Successfully deleted agent with ID %d\n", agentID)
	})

	t.Run("Test get a corporate entitiy agent", func(t *testing.T) {
		agentId, err := client.GetRandomAgentID("corporate_entities")
		if err != nil {
			t.Error(err)
		}

		agent, err := client.GetAgent("corporate_entities", agentId)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Successfully requested and serialized corporate entity agent %s: %s\n", agent.URI, agent.Names[0].SortName)
		}

	})

	/*





		t.Run("Test a family agent", func(t *testing.T) {
			agentId, err := client.GetRandomAgentID("families")
			if err != nil {
				t.Error(err)
			}

			agent, err := client.GetAgent("families", agentId)
			if err != nil {
				t.Error(err)
			} else {
				t.Logf("Successfully requested and serialized family agent %s: %s\n", agent.URI, agent.Names[0].SortName)
			}

		})


	*/
}
