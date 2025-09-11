package aspace

import (
	"flag"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestAgents(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test a person agent", func(t *testing.T) {
		agentId, err := client.GetRandomAgentID("people")
		if err != nil {
			t.Error(err)
		}

		agent, err := client.GetAgent("people", agentId)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Successfully requested and serialized person agent %s: %s\n", agent.URI, agent.Names[0].SortName)
		}

	})

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

	t.Run("Test a corporate entitiy agent", func(t *testing.T) {
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
}
