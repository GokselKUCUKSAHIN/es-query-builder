package intermediate

import (
	"encoding/json"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
)

////    Nested Example    ////

func createNestedQuery() string {
	query := es.NewQuery(
		es.Nested("driver",
			es.Nested("driver.vehicle",
				es.Bool().
					Must(
						es.Term("driver.vehicle.make", "Powell Motors"),
						es.Term("driver.vehicle.model", "Canyonero"),
					),
			),
		),
	)

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createNestedQueryPureGo() string {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"nested": map[string]interface{}{
				"path": "driver",
				"query": map[string]interface{}{
					"nested": map[string]interface{}{
						"path": "driver.vehicle",
						"query": map[string]interface{}{
							"bool": map[string]interface{}{
								"must": []map[string]interface{}{
									{
										"term": map[string]interface{}{
											"driver.vehicle.make": "Powell Motors",
										},
									},
									{
										"term": map[string]interface{}{
											"driver.vehicle.model": "Canyonero",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Benchmark_Nested_Example_Builder(b *testing.B) {
	createNestedQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createNestedQuery()
	}
}

func Benchmark_Nested_Example_PureGo(b *testing.B) {
	createNestedQueryPureGo()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createNestedQueryPureGo()
	}
}
