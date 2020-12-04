package aspace

import (
	"flag"
	"testing"
)

func TestPersonalAgent(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Error(err)
	}

	agentId, err := client.GetRandomAgentID("people")
	if err != nil {
		t.Error(err)
	}

	agent, err := client.GetAgent("people", agentId)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Successfully requested and serialized agent %s: %s\n", agent.URI, agent.Names[0].SortName)
	}

}
