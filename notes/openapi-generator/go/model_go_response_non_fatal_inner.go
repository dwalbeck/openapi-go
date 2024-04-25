/*
 * Order Uploader
 *
 * Takes cleaned and validated order data records as input and enters the record into the database tables
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GoResponseNonFatalInner struct {
	DbTablePath string `json:"db_table_path,omitempty"`

	Error string `json:"error,omitempty"`
}

// AssertGoResponseNonFatalInnerRequired checks if the required fields are not zero-ed
func AssertGoResponseNonFatalInnerRequired(obj GoResponseNonFatalInner) error {
	return nil
}

// AssertGoResponseNonFatalInnerConstraints checks if the values respects the defined constraints
func AssertGoResponseNonFatalInnerConstraints(obj GoResponseNonFatalInner) error {
	return nil
}