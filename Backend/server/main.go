package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Query struct {
	Query string `json:"input"`
}

func main() {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	r.Options("/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	})
	r.Post("/search", search)

	http.ListenAndServe(":3000", r)

}

//handlers

func search(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	var body Query
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}
	bodyJson, err := json.Marshal(body.Query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyJson))
	// body := map[string]interface{}{}
	// json.NewDecoder(r.Body).Decode(&body)
	// fmt.Println(body["input"])

	query := `{
		"query":{
			"bool":{
				"must":[
					{
						"query_string":{
							"query": ` + string(bodyJson) + `
						}
					}
				]
			}
		},
		"sort":[
			"-@timestamp"
		],
		"from":0,
		"size":100,
		"aggs":{
			"histogram":{
				"auto_date_histogram":{
					"field":"@timestamp",
					"buckets":100
				}
			}
		}
	}`
	fmt.Println(strings.NewReader(query))

	req, err := http.NewRequest("POST", "http://localhost:4080/es/emails/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	results, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(results)

}
