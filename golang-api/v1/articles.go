package v1

import (
	"beyondf1-v2/golang-api/data"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ApiArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Hit endpoint: ApiArticles; request from ip: %s", r.RemoteAddr)
	var articles = new([]data.Article)

	vals := mux.Vars(r)
	num, err := strconv.Atoi(vals["num"])
	if err != nil {
		log.Println("Cant decode request!")
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"articles": articles, "msg": "unable"})
		return
	}
	page, err := strconv.Atoi(vals["page"])
	if err != nil {
		log.Println("Cant decode request!")
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"articles": articles, "msg": "unable"})
		return
	}
	articles, msg, count := data.GetArticles(num, page)
	var next int
	var prev int
	if page <= 1 {
		prev = 0
	} else {
		prev = page - 1
	}
	var statement = len(*articles) + (num * (page - 1))
	if statement >= count {
		next = 0
	} else {
		next = page + 1
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"articles": articles, "msg": msg, "next": next, "prev": prev})
}

func ApiArticlesByTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Hit endpoint: ApiArticlesByTag; request from ip: %s", r.RemoteAddr)
	var articles = new([]data.Article)

	vals := mux.Vars(r)
	num, err := strconv.Atoi(vals["num"])
	if err != nil {
		log.Println("Cant decode request!")
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"articles": articles, "msg": "unable"})
		return
	}
	page, err := strconv.Atoi(vals["page"])
	if err != nil {
		log.Println("Cant decode request!")
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"articles": articles, "msg": "unable"})
		return
	}
	tag := vals["tag"]
	articles, msg, count := data.GetArticlesByTag(num, page, tag)
	var next int
	var prev int
	if page <= 1 {
		prev = 0
	} else {
		prev = page - 1
	}
	var statement = len(*articles) + (num * (page - 1))
	if statement >= count {
		next = 0
	} else {
		next = page + 1
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"articles": articles, "msg": msg, "next": next, "prev": prev})
}

func ApiArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Hit endpoint: ApiArticle; request from ip: %s", r.RemoteAddr)
	var article = new(data.Article)

	vals := mux.Vars(r)
	id, err := strconv.Atoi(vals["id"])
	if err != nil {
		log.Println("Cant decode request!")
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"articles": article, "msg": "unable"})
		return
	}
	article, msg := data.GetArticle(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"articles": article, "msg": msg})
}
