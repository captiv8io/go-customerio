package customerio

import (
	"context"
	"fmt"
	"net/url"
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

	uri, err := url.Parse(fmt.Sprintf("%s/api/v1/segments/%d/add_customers", c.URL, segmentID))
	if err != nil {
		return err
	}
	c.addIDTypeQuery(uri)

	return c.request(ctx, "POST", uri.String(),
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

	uri, err := url.Parse(fmt.Sprintf("%s/api/v1/segments/%d/remove_customers", c.URL, segmentID))
	if err != nil {
		return err
	}
	c.addIDTypeQuery(uri)

	return c.request(ctx, "POST", uri.String(),
		map[string]interface{}{
			"ids": ids,
		})
}

// getQueryParams returns the query parameter for the request.
func (c *CustomerIO) addIDTypeQuery(uri *url.URL) {
	if c.IDType == "" {
		return
	}

	query := uri.Query()
	switch IDType(c.IDType) {
	case IDTypeEmail, IDTypeCioID:
		query.Set("id_type", c.IDType)
	default:
		query.Set("id_type", string(DefaultIDType))
	}
	uri.RawQuery = query.Encode()
}
