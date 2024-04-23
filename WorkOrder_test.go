package aspace

import (
	"testing"

	"os"
	"path/filepath"
)

const fixtureRoot = "./goaspace_testing/testdata"

func assertStringsEqual(want, got string, t *testing.T) {
	if want != got {
		t.Errorf("want: %s , got: %s", want, got)
	}
}

func createAndLoadWorkOrder(path string, t *testing.T) *WorkOrder {
	var wo WorkOrder

	r, err := os.Open(path)
	if err != nil {
		t.Errorf("problem opening %s", path)
	}
	defer r.Close()

	// note: intentionally ignoring any error here
	wo.Load(r)

	return &wo
}

//------------------------------------------------------------------------------
// begin tests
//------------------------------------------------------------------------------

func TestHeader(t *testing.T) {

	wo := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	err := wo.assertHeaderFields()
	if err != nil {
		t.Error(err)
	}

	wo = createAndLoadWorkOrder(filepath.Join(fixtureRoot, "bad_header_wo.tsv"), t)
	err = wo.assertHeaderFields()
	if err == nil {
		t.Errorf("expected header assertion error, but assertion passed")
	}

	want := "[header mismatch: expected: 'Resource ID' got: 'XResource ID' header mismatch: expected: 'Ref ID' got: 'YRef ID' header mismatch: expected: 'URI' got: 'ZURI' header mismatch: expected: 'Container Indicator 1' got: 'AContainer Indicator 1' header mismatch: expected: 'Container Indicator 2' got: 'BContainer Indicator 2' header mismatch: expected: 'Container Indicator 3' got: 'CContainer Indicator 3' header mismatch: expected: 'Title' got: 'DTitle' header mismatch: expected: 'Component ID' got: 'EComponent ID']"
	if err.Error() != want {
		t.Errorf("unexpected header assertion error message.\nwanted:\n'%s'\n but got:\n'%s'", want, err.Error())
	}
}

func TestWorkOrderRowAccessors(t *testing.T) {
	const (
		wantIdx = 0
		gotIdx  = 1
		msgIdx  = 2
	)

	sut := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	scenarios := [][]string{
		{"DLTS.2022", sut.Rows[1].GetResourceID(), "Incorrect Resource ID"},
		{"4a6f56d2b69962a05792478cae78e888", sut.Rows[2].GetRefID(), "Incorrect Ref ID"},
		{"/repositories/3/archival_objects/979520", sut.Rows[0].GetURI(), "Incorrect URI"},
		{"V01", sut.Rows[1].GetContainerIndicator1(), "Incorrect Container Indicator 1"},
		{"I02", sut.Rows[2].GetContainerIndicator2(), "Incorrect Container Indicator 2"},
		{"A03", sut.Rows[0].GetContainerIndicator3(), "Incorrect Container Indicator 3"},
		{"Video Test", sut.Rows[1].GetTitle(), "Incorrect Title"},
		{"cuid39671", sut.Rows[2].GetComponentID(), "Incorrect Component ID"},
	}

	for _, scenario := range scenarios {
		if scenario[wantIdx] != scenario[gotIdx] {
			t.Errorf("unexpected result: %s: want: '%s', got: '%s'", scenario[msgIdx], scenario[wantIdx], scenario[gotIdx])
		}
	}
}
