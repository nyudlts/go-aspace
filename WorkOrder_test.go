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

func TestGetResourceID(t *testing.T) {
	var want, got string

	sut := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	want = "DLTS.2022"
	got = sut.Rows[1].GetResourceID()
	assertStringsEqual(want, got, t)
}

func TestGetRefID(t *testing.T) {
	var want, got string

	sut := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	want = "4a6f56d2b69962a05792478cae78e888"
	got = sut.Rows[2].GetRefID()
	assertStringsEqual(want, got, t)
}

func TestGetURI(t *testing.T) {
	var want, got string

	sut := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	want = "/repositories/3/archival_objects/979520"
	got = sut.Rows[0].GetURI()
	assertStringsEqual(want, got, t)
}

func TestGetContainerIndicator1(t *testing.T) {
	var want, got string

	sut := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	want = "V01"
	got = sut.Rows[1].GetContainerIndicator1()
	assertStringsEqual(want, got, t)
}

func TestGetContainerIndicator2(t *testing.T) {
	var want, got string

	sut := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	want = "I02"
	got = sut.Rows[2].GetContainerIndicator2()
	assertStringsEqual(want, got, t)
}

func TestGetContainerIndicator3(t *testing.T) {
	var want, got string

	sut := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	want = "A03"
	got = sut.Rows[0].GetContainerIndicator3()
	assertStringsEqual(want, got, t)
}

func TestGetContainerTitle(t *testing.T) {
	var want, got string

	sut := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	want = "Video Test"
	got = sut.Rows[1].GetTitle()
	assertStringsEqual(want, got, t)
}

func TestGetComponentID(t *testing.T) {
	var want, got string

	sut := createAndLoadWorkOrder(filepath.Join(fixtureRoot, "valid_wo.tsv"), t)

	want = "cuid39671"
	got = sut.Rows[2].GetComponentID()
	assertStringsEqual(want, got, t)
}
