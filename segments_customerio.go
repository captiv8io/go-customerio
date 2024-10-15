package customerio

import (
	"context"
	"fmt"
)

// IDType is the type of ids you want to use.
// All the values in the ids array must be of this type.
// If you don't provide this parameter, we assume that the ids array contains id values.
// Enum values:
//   - "id"
//   - "email"
//   - "cio_id"
type IDType string

const (
	IDTypeID      IDType = "id"
	IDTypeEmail   IDType = "email"
	IDTypeCioID   IDType = "cio_id"
	DefaultIDType        = IDTypeID
)

// AddPeopleToSegment adds people to a segment.
func (c *CustomerIO) AddPeopleToSegment(ctx context.Context, segmentID int, ids []string) error {
	if segmentID == 0 {
		return ParamError{Param: "segmentID"}
	}
	if len(ids) == 0 {
		return ParamError{Param: "ids"}
	}
	return c.request(ctx, "POST",
		fmt.Sprintf("%s/api/v1/segments/%d/add_customers%s", c.URL, segmentID, c.getQueryParams()),
		map[string]interface{}{
			"ids": ids,
		})
}

// RemovePeopleFromSegment removes people from a segment
func (c *CustomerIO) RemovePeopleFromSegment(ctx context.Context, segmentID int, ids []string) error {
	if segmentID == 0 {
		return ParamError{Param: "segmentID"}
	}
	if len(ids) == 0 {
		return ParamError{Param: "ids"}
	}
	return c.request(ctx, "POST",
		fmt.Sprintf("%s/api/v1/segments/%d/remove_customers%s", c.URL, segmentID, c.getQueryParams()),
		map[string]interface{}{
			"ids": ids,
		})
}

// getQueryParams returns the query parameter for the request.
func (c *CustomerIO) getQueryParams() string {
	if c.IDType == "" {
		return ""
	}

	// Check if the IDType is valid and construct query parameter accordingly
	switch IDType(c.IDType) {
	case IDTypeEmail, IDTypeCioID:
		return fmt.Sprintf("?id_type=%s", c.IDType)
	default:
		return fmt.Sprintf("?id_type=%s", DefaultIDType)
	}
}
