package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
)

////    Simple    ////

func createSimpleQuery() string {
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.Term("id", 123456),
			),
	)

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createSimpleQueryPureGo() string {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"id": 123456,
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

func Benchmark_Simple_Builder(b *testing.B) {
	createSimpleQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createSimpleQuery()
	}
}

func Benchmark_Simple_PureGo(b *testing.B) {
	createSimpleQueryPureGo()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createSimpleQueryPureGo()
	}
}
