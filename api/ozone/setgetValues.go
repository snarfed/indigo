// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.set.getValues

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// SetGetValues_Output is the output of a tools.ozone.set.getValues call.
type SetGetValues_Output struct {
	Cursor *string          `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Set    *SetDefs_SetView `json:"set" cborgen:"set"`
	Values []string         `json:"values" cborgen:"values"`
}

// SetGetValues calls the XRPC method "tools.ozone.set.getValues".
func SetGetValues(ctx context.Context, c *xrpc.Client, cursor string, limit int64, name string) (*SetGetValues_Output, error) {
	var out SetGetValues_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
		"name":   name,
	}
	if err := c.Do(ctx, xrpc.Query, "", "tools.ozone.set.getValues", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}