package aspace

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
	"github.com/lestrrat-go/libxml2/parser"
	"github.com/lestrrat-go/libxml2/xsd"
	"github.com/nyudlts/go-aspace/box"
	"strconv"
	"strings"
)

var p = parser.New()

var LibraryVersion = "v0.3.0"

func PrintClientVersion() {
	fmt.Println("Go Aspace", LibraryVersion)
}

func Seed() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func RandInt(min int, max int) int {
	return min + math_rand.Intn(max-min)
}

func ValidateEAD(fa []byte) error {

	eadxsd, err := xsd.Parse(box.Box.Get("/ead.xsd"))
	if err != nil {
		return err
	}
	doc, err := p.Parse(fa)
	if err != nil {
		return err
	}
	if err := eadxsd.Validate(doc); err != nil {
		return err
	}
	return nil
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
