package aspace

import (
	"encoding/json"
	"flag"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

var (
	testRepoID       int
	testResourceID   = 1
	testAccessionIDs = []AccessionEntry{}
	testAccessionID  int
	testAccession    *Accession
)

var defaultAccession = `
{
  "jsonmodel_type": "accession",
  "is_slug_auto": true,
  "accession_date": "2007-03-15",
  "extents": [
    {
      "jsonmodel_type": "extent",
      "portion": "part",
      "number": "60",
      "extent_type": "cassettes",
      "dimensions": "S406KHU",
      "physical_details": "632RL388775"
    }
  ],
  "restrictions_apply": false,
  "access_restrictions": false,
  "use_restrictions": false,
  "id_0": "575NV197270",
  "id_1": "SD524676B",
  "id_2": "DO962433M",
  "id_3": "HGW604442",
  "title": "Accession Title: 30",
  "content_description": "Description: 11",
  "condition_description": "Description: 12"
}
`

func TestAccessions(t *testing.T) {
	//get a client
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment)
	if err != nil {
		t.Error(err)
	}

	t.Run("test unmarshal default accessions", func(t *testing.T) {
		testAccession = &Accession{}
		err := json.Unmarshal([]byte(defaultAccession), &testAccession)
		if err != nil {
			t.Error(err)
		}

		testAccession.ID0 = RandStringRunes(4)
		testAccession.ID1 = RandStringRunes(4)
		testAccession.ID2 = RandStringRunes(4)
		testAccession.ID3 = RandStringRunes(4)

		t.Logf("%s", testAccession.Title)
	})

	//get a list of accessions
	t.Run("Test get Accession List for Resource", func(t *testing.T) {
		testRepoID, testResourceID, err = client.GetRandomResourceID()
		if err != nil {
			t.Error(err)
		}

		t.Logf("Testing GetAccessionList for Repository ID %d, ResourceID %d", testRepoID, testResourceID)
		testAccessionIDs, err = client.GetAccessionList(testRepoID, testResourceID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Found %d accessions", len(testAccessionIDs))

	})

	t.Run("Test get an accession", func(t *testing.T) {

		t.Logf("Testing GetAccession for Repository ID %d, AccessionID %d", testRepoID, testAccessionIDs[0].AccessionID)
		testAccession2, err := client.GetAccession(testRepoID, testAccessionIDs[0].AccessionID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully unmarshalled %s %s", testAccession2.URI, testAccession2.Title)

	})

	t.Run("Test create an accession", func(t *testing.T) {

		str, err := client.CreateAccession(testRepoID, *testAccession)
		if err != nil {
			t.Error(err)
		}

		apiResponse := &APIResponse{}
		err = json.Unmarshal([]byte(str), apiResponse)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%v", apiResponse)
		}
		testAccessionID = apiResponse.ID
	})

	t.Run("Test update an accession", func(t *testing.T) {
		testAccession.ID0 = RandStringRunes(4)
		testAccession.ID1 = RandStringRunes(4)
		testAccession.ID2 = RandStringRunes(4)
		testAccession.ID3 = RandStringRunes(4)

		str, err := client.UpdateAccession(testRepoID, testAccessionID, *testAccession)
		if err != nil {
			t.Error(err)
		}
		apiResponse := &APIResponse{}
		err = json.Unmarshal([]byte(str), apiResponse)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%v", apiResponse)
		}
	})

	t.Run("Test delete an accession", func(t *testing.T) {
		t.Logf("testing delete accession for RepoID %d, AccessionID %d", testRepoID, testAccessionID)
		str, err := client.DeleteAccession(testRepoID, testAccessionID)
		if err != nil {
			t.Error(err)
		}
		APIResponse := &APIResponse{}
		err = json.Unmarshal([]byte(str), APIResponse)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%v", APIResponse)
		}
	})
}
