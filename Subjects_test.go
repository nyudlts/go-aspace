package aspace

import (
	"flag"
	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
	"testing"
)

func TestSubjects(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(err)
	}

	subjectId, err := client.GetRandomSubjectID()
	if err != nil {
		t.Error(err)
	}

	subject, err := client.GetSubject(subjectId)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Successfully requested and serialized %s: %s\n", subject.URI, subject.Title)
	}
}
