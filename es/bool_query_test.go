package es_test

import (
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

////   Bool   ////

func Test_Bool_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Bool)
}

func Test_Bool_method_should_create_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.boolType", b)
}

func Test_Bool_should_have_MinimumShouldMatch_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.MinimumShouldMatch)
}

func Test_Bool_MinimumShouldMatch_should_create_json_with_minimum_should_match_field_inside_bool(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			MinimumShouldMatch(7),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"minimum_should_match\":7}}}", bodyJSON)
}

func Test_Bool_should_have_Boost_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Boost)
}

func Test_Bool_Boost_should_create_json_with_minimum_should_match_field_inside_bool(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Boost(3.1415),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"boost\":3.1415}}}", bodyJSON)
}

func Test_Bool_should_have_Filter_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Filter)
}

func Test_Bool_should_have_Must_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Must)
}

func Test_Bool_should_have_MustNot_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.MustNot)
}

func Test_Bool_should_have_Should_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Should)
}

////   Bool.Filter   ////

func Test_Filter_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	filter := b.Filter()

	// Then
	assert.NotNil(t, filter)
	assert.IsTypeString(t, "es.boolType", filter)
}

func Test_Filter_method_should_add_filter_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["filter"]
	filter := b.Filter()
	_, afterExists := b["filter"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_Filter_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		Filter(
			es.Term("id", 12345),
		)

	// When
	filter, exists := b["filter"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.filterType", filter)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"filter\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////   Bool.Must   ////

func Test_Must_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	must := b.Must()

	// Then
	assert.NotNil(t, must)
	assert.IsTypeString(t, "es.boolType", must)
}

func Test_Must_method_should_add_must_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["must"]
	filter := b.Must()
	_, afterExists := b["must"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_Must_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		Must(
			es.Term("id", 12345),
		)

	// When
	must, exists := b["must"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.mustType", must)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"must\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////   Bool.MustNot   ////

func Test_MustNot_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	mustNot := b.MustNot()

	// Then
	assert.NotNil(t, mustNot)
	assert.IsTypeString(t, "es.boolType", mustNot)
}

func Test_MustNot_method_should_add_must_not_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["must_not"]
	filter := b.MustNot()
	_, afterExists := b["must_not"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_MustNot_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		MustNot(
			es.Term("id", 12345),
		)

	// When
	mustNot, exists := b["must_not"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.mustNotType", mustNot)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"must_not\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////   Bool.Should   ////

func Test_Should_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	should := b.Should()

	// Then
	assert.NotNil(t, should)
	assert.IsTypeString(t, "es.boolType", should)
}

func Test_Should_method_should_add_should_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["should"]
	filter := b.Should()
	_, afterExists := b["should"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_Should_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		Should(
			es.Term("id", 12345),
		)

	// When
	should, exists := b["should"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.shouldType", should)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"should\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}
