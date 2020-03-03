package go_aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestGetAspaceInfo(t *testing.T) {

	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	info, err := a.GetAspaceInfo()
	if err != nil {
		t.Error(err)
	}

	t.Log(info)
}

func TestGetResourceIDsByRepository(t *testing.T) {
	repositoryId := 2

	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	resourceIds, err := a.GetResourceIDsByRepository(repositoryId)
	if err != nil {
		t.Error(err)
	}
	if len(resourceIds) < 0 {
		t.Errorf("There are no resources in repository %d", repositoryId)
	}
}

func TestGetResourceByID(t *testing.T) {

	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	repositoryId := 2
	resources, err := a.GetResourceIDsByRepository(repositoryId)
	if err != nil {
		t.Error(err)
	}

	randomNum := randInt(0, len(resources))
	resourceId := resources[randomNum]

	resource, err := a.GetResourceByID(repositoryId, resourceId)
	if err != nil {
		t.Error(err)
	}

	title := resource.Title
	if len(title) < 0 {
		t.Errorf("Nil title returned")
	}

	lockVersion := resource.Lock_Version

	if lockVersion < 0 {
		t.Errorf("Malformed lock version")
	}

}

func TestASClient_PostResource(t *testing.T) {
	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	repositoryId, resourceId := 2, 2

	resource, err := a.GetResourceByID(repositoryId, resourceId)
	if err != nil {
		t.Error(err)
	}

	id_0 := resource.ID_0
	id_1 := resource.ID_1
	id_2 := resource.ID_2
	id_3 := resource.ID_3
	target := ""
	if id_0 == "" {
		log.Println(fmt.Errorf("Resource %s does not have id0 defined", resource.Title))
	} else {
		target = id_0
	}
	if id_1 != "" {
		target = target + "_" + id_1
	}
	if id_2 != "" {
		target = target + "_" + id_2
	}
	if id_3 != "" {
		target = target + "_" + id_3
	}
	target = strings.ToLower(target)

	resource.EAD_ID = target

	jsonResource, err := json.Marshal(resource)

	r, err := a.PostResource(repositoryId, resourceId, string(jsonResource))
	if err != nil {
		t.Error(r.Status)
		body, _ := ioutil.ReadAll(r.Body)
		t.Error(string(body))
	}

}
