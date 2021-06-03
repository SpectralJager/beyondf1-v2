package data

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// db info
const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "238516"
	dbname   = "beyondf1"
)

// additional

// sql statements
const (
	getArticles        = `select * from api_article order by createdat desc limit $1 offset $2;`
	getArticlesByTag   = `select * from api_article where tag = $1 order by createdat desc limit $2 offset $3;;`
	getArticle         = `select * from api_article where id = $1;`
	countArticles      = `select count(id) from api_article`
	countArticlesByTag = `select count(id) from api_article where tag = $1`
)

// structures
type Article struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"createdat" db:"createdat"`
	Tag       string    `json:"tag" db:"tag"`
	ImageUrl  string    `json:"image_url" db:"image_url"`
	Text      string    `json:"text" db:"text"`
	Source    string    `json:"source" db:"source"`
}

// functions
func connectDB() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname,
	)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Cant connect to db")
		log.Fatalln(err)
		return db, err
	}
	err = db.Ping()
	log.Println("Db connected")
	return db, err
}

func GetArticle(id int) (*Article, string) {
	db, err := connectDB()
	if err != nil {
		return new(Article), "unable"
	}
	defer db.Close()

	var article Article
	row := db.QueryRowx(getArticle, id)
	if err := row.StructScan(&article); err != nil {
		log.Printf("Cant scan row! with id %d\n", id)
		return new(Article), "unable"
	}
	return &article, "success"
}

func GetArticles(n, page int) (*[]Article, string, int) {
	db, err := connectDB()
	if err != nil {
		return new([]Article), "unable", 0
	}
	defer db.Close()

	var articles []Article
	rows, err := db.Queryx(getArticles, n, n*(page-1))
	if err != nil {
		log.Println("Cant get rows!")
		log.Fatalln(err)
		return &articles, "unable", 0
	}
	defer rows.Close()
	for rows.Next() {
		var article Article
		if err := rows.StructScan(&article); err != nil {
			log.Println("Cant scan row")
			log.Fatalln(err)
			continue
		}
		articles = append(articles, article)
	}
	if len(articles) == 0 {
		log.Println("Haven't articles!")
		return &articles, "unable", 0
	}
	var count int
	row := db.QueryRowx(countArticles)
	row.Scan(&count)
	return &articles, "success", count
}

func GetArticlesByTag(n, page int, tag string) (*[]Article, string, int) {
	db, err := connectDB()
	if err != nil {
		return new([]Article), "unable", 0
	}
	defer db.Close()

	var articles []Article
	rows, err := db.Queryx(getArticlesByTag, tag, n, n*(page-1))
	if err != nil {
		log.Println("Cant get rows!")
		log.Fatalln(err)
		return &articles, "unable", 0
	}
	defer rows.Close()
	for rows.Next() {
		var article Article
		if err := rows.StructScan(&article); err != nil {
			log.Println("Cant scan row")
			log.Fatalln(err)
			continue
		}
		articles = append(articles, article)
	}
	if len(articles) == 0 {
		log.Println("Haven't articles!")
		return &articles, "unable", 0
	}
	var count int
	row := db.QueryRowx(countArticlesByTag, tag)
	row.Scan(&count)
	return &articles, "success", count
}
