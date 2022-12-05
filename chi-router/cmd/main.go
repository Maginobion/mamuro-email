package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })

	r.Handle("/", http.FileServer(http.Dir("../public")))

	http.ListenAndServe(":3000", r)

}

// func EmailCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		articleID := chi.URLParam(r, "articleID")
// 		article, err := dbGetEmails()
// 		if err != nil {
// 			http.Error(w, http.StatusText(404), 404)
// 			return
// 		}
// 		ctx := context.WithValue(r.Context(), "article", article)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func getEmails(w http.ResponseWriter, r *http.Request) {

	url := "http://localhost:4080/api/mails/_search"

	query := `{
        "search_type": "matchphrase",
        "query":
        {
            "term": "john",
        },
        "from": 0,
        "max_results": 20,
        "_source": []
    }`
	req, err := http.NewRequest("POST", url, strings.NewReader(query))

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

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)

	w.Write(body)

}
