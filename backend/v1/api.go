package v1

import (
	"fmt"

	"github.com/gorilla/mux"
)

func V1(router *mux.Router, prefix string) {
	// articles
	router.Handle(fmt.Sprintf("%s/articles/{page:[0-9]+}/{num:[0-9]+}", prefix), CorsMiddleware(ApiArticles)).Methods("GET", "OPTIONS")
	router.Handle(fmt.Sprintf("%s/articles/{tag:[a-zA-Z]+}/{page:[0-9]+}/{num:[0-9]+}", prefix), CorsMiddleware(ApiArticlesByTag)).Methods("GET", "OPTIONS")
	router.Handle(fmt.Sprintf("%s/article/{id:[0-9]+}", prefix), CorsMiddleware(ApiArticle)).Methods("GET", "OPTIONS")
	// events
	router.Handle(fmt.Sprintf("%s/events", prefix), CorsMiddleware(ApiEvents)).Methods("GET", "OPTIONS")
}
