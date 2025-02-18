package namespaces

import (
	"github.com/chnsz/golangsdk"
)

type Namespace struct {
	// Name of the Namespace
	Name string `json:"name"`
	// Creator Name of the Namespace
	CreatorName string `json:"creator_name"`
	// Auth permission of the Namespace
	Auth int `json:"auth"`
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a namespace.
func (r commonResult) Extract() (*Namespace, error) {
	var s Namespace
	err := r.ExtractInto(&s)
	return &s, err
}

// CreateResult represents the result of a create operation.
type CreateResult struct {
	golangsdk.ErrResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Network.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
