package aspace

import (
	"encoding/csv"
	"fmt"
	"io"
)

// HEADER_ROW is the list of expected work order field names
var HEADER_ROW = []string{"Resource ID", "Ref ID", "URI", "Container Indicator 1", "Container Indicator 2", "Container Indicator 3", "Title", "Component ID"}

// Get a Header Row as a string
func GetWorkOrderHeader() string {
	var header string
	for i, h := range HEADER_ROW {
		if i < len(HEADER_ROW)-1 {
			header = fmt.Sprintf(header+"%s\t", h)
		} else {
			header = header + h
		}
	}
	return header
}

// A WorkOrder stores the header and constituent rows for a parsed work order
type WorkOrder struct {
	Header []string
	Rows   []WorkOrderRow
}

// A WorkOrderRow holds the values for a single row in the parsed WorkOrder
type WorkOrderRow struct {
	fields []string
}

// assertHeaderFields ensures that the fields in the the work order being processed match expectations
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

// NewWorkOrder returns a pointer to a newly constructed WorkOrder
func NewWorkOrder() *WorkOrder {
	// this function is here in case we need to perform further initialization
	return new(WorkOrder)
}

// Load reads and parses data from an io.Reader storing them in the WorkOrder.
// This method returns an error if the header fields do not match expectations.
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

// GetResourceID returns the "Resource ID" value for the WorkOrderRow
func (wor *WorkOrderRow) GetResourceID() string {
	return wor.fields[0]
}

// GetRefID returns the "Ref ID" value for the WorkOrderRow
func (wor *WorkOrderRow) GetRefID() string {
	return wor.fields[1]
}

// GetURI returns the "URI" value for the WorkOrderRow
func (wor *WorkOrderRow) GetURI() string {
	return wor.fields[2]
}

// GetContainerIndicator1 returns the "Container Indicator 1" value for the WorkOrderRow
func (wor *WorkOrderRow) GetContainerIndicator1() string {
	return wor.fields[3]
}

// GetContainerIndicator2 returns the "Container Indicator 2" value for the WorkOrderRow
func (wor *WorkOrderRow) GetContainerIndicator2() string {
	return wor.fields[4]
}

// GetContainerIndicator3 returns the "Container Indicator 3" value for the WorkOrderRow
func (wor *WorkOrderRow) GetContainerIndicator3() string {
	return wor.fields[5]
}

// GetTitle returns the "Title" value for the WorkOrderRow
func (wor *WorkOrderRow) GetTitle() string {
	return wor.fields[6]
}

// GetComponentID returns the "Component ID" value for the WorkOrderRow
func (wor *WorkOrderRow) GetComponentID() string {
	return wor.fields[7]
}
