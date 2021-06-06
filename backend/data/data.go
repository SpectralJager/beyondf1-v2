package data

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// db info
var (
	host     = os.Getenv("DATABASE_HOST")
	port     = os.Getenv("DATABASE_PORT")
	username = os.Getenv("DATABASE_USERNAME")
	password = os.Getenv("DATABASE_PASSWORD")
	dbname   = os.Getenv("DATABASE_NAME")
)

// additional

// sql statements
const (
	getArticles = `
	select a.id, a.title, a.text, a.created_at, a.image_url, a.source, t.tag
	from articles a 
	left join articles_tags__tags_articles temp on a.id = temp.article_id
	left join tags t on temp.tag_id = t.id
	order by a.created_at desc limit $1 offset $2;
	`
	getArticlesByTag = `
	select a.id, a.title, a.text, a.created_at, a.image_url, a.source, t.tag
	from articles a 
	left join articles_tags__tags_articles temp on a.id = temp.article_id
	left join tags t on temp.tag_id = t.id	
	where t.tag = $1 order by a.created_at desc limit $2 offset $3;
	`
	getArticle = `
	select a.id, a.title, a.text, a.created_at, a.image_url, a.source, t.tag
	from articles a 
	left join articles_tags__tags_articles temp on a.id = temp.article_id
	left join tags t on temp.tag_id = t.id
	where a.id = $1;
	`
	countArticles      = `select count(id) from articles;`
	countArticlesByTag = `
	select count(a.id)
	from articles a 
	left join articles_tags__tags_articles temp on a.id = temp.article_id
	left join tags t on temp.tag_id = t.id
	where t.id = $1;
	`
	getEvents = `select id, name, description, link, date_start, date_end, image_url from events order by date_start desc limit 2;`
)

// structures
type Events struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Desc       string    `json:"description" db:"description"`
	Link       string    `json:"link" db:"link"`
	Image_url  string    `json:"image_url" db:"image_url"`
	Date_start time.Time `json:"date_start" db:"date_start"`
	Date_end   time.Time `json:"date_end" db:"date_end"`
}

type Article struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"createdat" db:"created_at"`
	Tag       string    `json:"tag" db:"tag"`
	ImageUrl  string    `json:"image_url" db:"image_url"`
	Text      string    `json:"text" db:"text"`
	Source    string    `json:"source" db:"source"`
}

// functions
func connectDB() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname,
	)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Cant connect to db")
		log.Println(err)
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
		log.Println(err)
		return &articles, "unable", 0
	}
	defer rows.Close()
	for rows.Next() {
		var article Article
		if err := rows.StructScan(&article); err != nil {
			log.Println("Cant scan row")
			log.Println(err)
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
		log.Println(err)
		return &articles, "unable", 0
	}
	defer rows.Close()
	for rows.Next() {
		var article Article
		if err := rows.StructScan(&article); err != nil {
			log.Println("Cant scan row")
			log.Println(err)
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

func GetEvents() (*[]Events, string) {
	db, err := connectDB()
	if err != nil {
		return new([]Events), "unable"
	}
	defer db.Close()

	var events []Events
	rows, err := db.Queryx(getEvents)
	if err != nil {
		log.Println("Cant get rows!")
		log.Println(err)
		return &events, "unable"
	}
	defer rows.Close()
	for rows.Next() {
		var event Events
		if err := rows.StructScan(&event); err != nil {
			log.Println("Cant scan row")
			log.Println(err)
			continue
		}
		events = append(events, event)
	}
	if len(events) == 0 {
		log.Println("Haven't articles!")
		return &events, "unable"
	}
	return &events, "success"
}
