package es

type regexpType Object

// Regexp creates a new regexpType object with the specified key-value pair.
//
// This function initializes a regexpType object with a single term query, where the
// key is the field name and the value is the term to search for. This is typically
// used to construct a regexp query in search queries.
//
// Example usage:
//
//	t := Regexp("endpoint", "/books/.*")
//	// t now contains a regexpType object with a regexp query for the "endpoint" field.
//
// Parameters:
//   - key: A string representing the field name for the regexp query.
//   - value: The value to be searched for in the specified field. The type is regexp.
//
// Returns:
//
//	A termType object containing the specified term query.
func Regexp[T any](key string, value T) regexpType {
	return regexpType{
		"regexp": Object{
			key: value,
		},
	}
}
