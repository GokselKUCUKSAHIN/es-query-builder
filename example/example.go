//nolint:all
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
)

func mockGetDocumentsEs(query string) (string, error) {
	return fmt.Sprintf("query result for '%v'", query), nil
}

func main() {
	id := 42
	queryString, err := json.Marshal(buildQuery(id))
	if err != nil {
		log.Fatal(err.Error())
	}

	documentsEs, err := mockGetDocumentsEs(string(queryString))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("query result: %s\n", documentsEs)

	// Output:
	// query result for '{"_source":{"includes":["id","type","indexedAt","chapters"]},"query":{"bool":{"filter":[{"terms":{"type":["DOC","FILE"]}}],"must":[{"bool":{"should":[{"term":{"doc.id":42}},{"term":{"file.fileId":42}}]}}]}},"size":45,"sort":[{"name":{"order":"asc"}}]}'

	// Formatted query string:
	// {
	//   "_source": {
	//     "includes": [
	//       "id",
	//       "type",
	//       "indexedAt",
	//       "chapters"
	//     ]
	//   },
	//   "query": {
	//     "bool": {
	//       "filter": [
	//         {
	//           "terms": {
	//             "type": [
	//               "DOC",
	//               "FILE"
	//             ]
	//           }
	//         }
	//       ],
	//       "must": [
	//         {
	//           "bool": {
	//             "should": [
	//               {
	//                 "term": {
	//                   "doc.id": 42
	//                 }
	//               },
	//               {
	//                 "term": {
	//                   "file.fileId": 42
	//                 }
	//               }
	//             ]
	//           }
	//         }
	//       ]
	//     }
	//   },
	//   "size": 45,
	//   "sort": [
	//     {
	//       "name": {
	//         "order": "asc"
	//       }
	//     }
	//   ]
	// }
}

func buildQuery(id int) es.Object {
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Bool().
					Should(
						es.Term("doc.id", id),
						es.Term("file.fileId", id),
					),
			).
			Filter(
				es.Terms("type", "DOC", "FILE"),
			),
	)
	query.Size(45)
	query.Sort(es.Sort("name", order.Asc))
	query.Source().
		Includes("id", "type", "indexedAt", "chapters")
	return query
}
