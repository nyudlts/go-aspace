package aspace

import (
	"flag"
	"testing"
)

func TestSubjects(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
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
		t.Logf("Succsefully requested and serialized %s: %s\n", subject.URI, subject.Title)
	}
}
