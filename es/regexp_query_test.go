package es_test

import (
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

////   Regexp   ////

func Test_Regexp_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Regexp[any])
}

func Test_Regexp_should_create_json_with_regexp_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Regexp("endpoint", "/books/.*"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"regexp\":{\"endpoint\":\"/books/.*\"}}}", bodyJSON)
}

func Test_Regexp_method_should_create_regexpType(t *testing.T) {
	// Given
	b := es.Regexp("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.regexpType", b)
}
