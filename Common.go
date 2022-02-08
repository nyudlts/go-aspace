package aspace

import (
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/libxml2/parser"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var p = parser.New()

var LibraryVersion = "v0.3.12b"

var seed = rand.NewSource(time.Now().UnixNano())
var rGen = rand.New(seed)

func PrintClientVersion() {
	fmt.Println("Go Aspace", LibraryVersion)
}

func URISplit(uri string) (int, int, error) {
	splitURI := strings.Split(uri, "/")
	resourceId, err := strconv.Atoi(splitURI[2])
	if err != nil {
		return 0, 0, err
	}
	objectId, err := strconv.Atoi(splitURI[4])
	if err != nil {
		return 0, 0, err
	}
	return resourceId, objectId, nil
}

type AspaceInfo struct {
	DatabaseProductName    string `json:"databaseProductName"`
	DatabaseProductVersion string `json:"databaseProductVersion"`
	RubyVersion            string `json:"ruby_version"`
	HostOS                 string `json:"host_os"`
	HostCPU                string `json:"host_cpu"`
	Build                  string `json:"build"`
	ArchivesSpaceVersion   string `json:"archivesSpaceVersion"`
}

func (a AspaceInfo) String() string {
	msg := fmt.Sprintf("== ArchivesSpace Version: %s\n", a.ArchivesSpaceVersion)
	msg = msg + fmt.Sprintf("== Database Type: %s\n", a.DatabaseProductName)
	msg = msg + fmt.Sprintf("== Database Version: %s\n", a.DatabaseProductVersion)
	msg = msg + fmt.Sprintf("== Ruby Version: %s\n", a.RubyVersion)
	msg = msg + fmt.Sprintf("== Host OS: %s\n", a.HostOS)
	msg = msg + fmt.Sprintf("== Host CPU: %s\n", a.HostCPU)
	msg = msg + fmt.Sprintf("== Java Version: %s\n", a.Build)
	return msg
}

func (a *ASClient) GetAspaceInfo() (AspaceInfo, error) {
	var aspaceInfo AspaceInfo
	response, err := a.get("", false)
	if err != nil {
		return aspaceInfo, err
	}
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &aspaceInfo)
	if err != nil {
		return aspaceInfo, err
	}
	return aspaceInfo, nil
}

func (a *ASClient) PrintResponse(endpoint string) error {
	response, err := a.get(endpoint, true)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func (a *ASClient) GetEndpoint(e string) ([]byte, error) {

	response, err := a.get(e, true)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

//slice contains methods
func containsInt(list []int, id int) bool {
	for _, i := range list {
		if id == i {
			return true
		}
	}
	return false
}
