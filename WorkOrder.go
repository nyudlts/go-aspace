package aspace

import (
	"encoding/csv"
	"fmt"
	"io"
)

// This is the expected set of columns in a work order
var HEADER_ROW = []string{"Resource ID", "Ref ID", "URI", "Container Indicator 1", "Container Indicator 2", "Container Indicator 3", "Title", "Component ID"}

type WorkOrder struct {
	header []string
	rows   []WorkOrderRow
}

type WorkOrderRow struct {
	fields []string
}

func (wo *WorkOrder) assertHeaderFields() error {
	if len(wo.header) != len(HEADER_ROW) {
		return fmt.Errorf("number of columns in work order header in work order does match expectations")
	}

	var errors []string
	for i, v := range HEADER_ROW {
		if wo.header[i] != v {
			errors = append(errors, fmt.Sprintf("header mismatch: expected: '%s' got: '%s'", v, wo.header[i]))
		}
	}

	if len(errors) != 0 {
		return fmt.Errorf("%s", errors)
	}
	return nil
}

func newWorkOrderRow(record []string) WorkOrderRow {
	var wo WorkOrderRow
	wo.fields = record
	return wo
}

func NewWorkOrder() *WorkOrder {
	// this function is here in case we need to perform further initialization
	return new(WorkOrder)
}

func (wo *WorkOrder) Load(r io.Reader) error {
	wor := csv.NewReader(r)
	wor.Comma = rune('\t')

	headerRow := true
	for {
		record, err := wor.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if headerRow {
			wo.header = record
			err = wo.assertHeaderFields()
			if err != nil {
				return err
			}

			headerRow = false
			wo.header = record
			continue
		}

		// add the row to the work order
		wo.rows = append(wo.rows, newWorkOrderRow(record))
	}
	return nil
}

// accessors

func (wor *WorkOrderRow) GetResourceID() string {
	return wor.fields[0]
}

func (wor *WorkOrderRow) GetRefID() string {
	return wor.fields[1]
}

func (wor *WorkOrderRow) GetURI() string {
	return wor.fields[2]
}

func (wor *WorkOrderRow) GetContainerIndicator1() string {
	return wor.fields[3]
}

func (wor *WorkOrderRow) GetContainerIndicator2() string {
	return wor.fields[4]
}

func (wor *WorkOrderRow) GetContainerIndicator3() string {
	return wor.fields[5]
}

func (wor *WorkOrderRow) GetTitle() string {
	return wor.fields[6]
}

func (wor *WorkOrderRow) GetComponentID() string {
	return wor.fields[7]
}
