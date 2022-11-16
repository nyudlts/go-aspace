package aspace

import (
	"flag"
	"fmt"
	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
	"testing"
)

var topContainerID int
var repositoryID int
var resourceID int

func TestTopContainers(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(err)
	}

	repositoryID, _ = client.GetRandomRepository()
	t.Log("Testing on repository", repositoryID)
	resourceIDs, _ := client.GetResourceIDs(repositoryID)
	t.Log("Testing on Resource", resourceID)

	t.Run("Test Get TopContainer IDS", func(t *testing.T) {
		topContainers, err := client.GetTopContainerIDs(repositoryID)
		if err != nil {
			t.Error(err)
		}

		if len(topContainers) <= 0 {
			t.Error("Array of less than 1 returned")
		} else {
			t.Log("returned", len(topContainers), "Top Containers")
		}

		topContainerID = topContainers[0]

	})

	t.Run("Test Get A Top Container", func(t *testing.T) {
		topContainer, err := client.GetTopContainer(repositoryID, topContainerID)
		if err != nil {
			t.Error(err)
		}

		t.Log("Top Container", topContainer.URI, "serialized")
	})

	t.Run("Test Get Top Container Ids for Resource", func(t *testing.T) {
		topContainers, err := client.GetTopContainerIDsForResource(repositoryID, resourceIDs[0])
		if err != nil {
			t.Error(err)
		}
		fmt.Println(topContainers)
	})

	t.Run("Test Get Top Containers for Resource", func(t *testing.T) {
		topContainers, err := client.GetTopContainersForResource(repositoryID, resourceIDs[0])
		if err != nil {
			t.Error(err)
		}
		for k, v := range topContainers {
			fmt.Println(k, "->", v.Barcode)
		}
	})
}
