package intermediate

import (
	"encoding/json"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
)

////    TY Example    ////

func createTyExampleQuery(brandIds []int64, storefrontIds []string) string {
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.Term("type", "LegalRule"),
				es.TermsArray("brandId", brandIds),
				es.TermsArray("allowedStorefronts.storefrontId", storefrontIds),
			),
	)
	query.Size(1)
	query.SourceFalse()

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createTyExampleQueryPureGo(brandIds []int64, storefrontIds []string) string {
	query := map[string]interface{}{
		"size": 1,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []interface{}{
					map[string]interface{}{
						"term": map[string]interface{}{
							"type": "LegalRule",
						},
					},
					map[string]interface{}{
						"terms": map[string]interface{}{
							"brandId": brandIds,
						},
					},
					map[string]interface{}{
						"terms": map[string]interface{}{
							"allowedStorefronts.storefrontId": storefrontIds,
						},
					},
				},
			},
		},
		"_source": false,
	}

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Benchmark_Ty_Example_Builder(b *testing.B) {
	brandIds := []int64{11, 22, 33, 44}
	storefrontIds := []string{"35", "36", "43", "48", "49", "50"}
	createTyExampleQuery(brandIds, storefrontIds)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createTyExampleQuery(brandIds, storefrontIds)
	}
}

func Benchmark_Ty_Example_PureGo(b *testing.B) {
	brandIds := []int64{11, 22, 33, 44}
	storefrontIds := []string{"35", "36", "43", "48", "49", "50"}
	createTyExampleQueryPureGo(brandIds, storefrontIds)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createTyExampleQueryPureGo(brandIds, storefrontIds)
	}
}
