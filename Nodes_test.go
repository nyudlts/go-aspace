package aspace

import "testing"

func TestNodes(t *testing.T) {
	t.Run("test get root node", func(t *testing.T) {
		rootNode, err := testClient.GetRootNode(testRepoID, testResourceID)
		if err != nil {
			t.Fatalf("Failed to get root node: %v", err)
		}

		t.Logf("Root Node: %+v", rootNode)
	})
}
