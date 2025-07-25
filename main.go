package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/keshav78-78/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	feed, err := urlToFeed("https://www.zoho.com/blog/feed")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "3000" // Default if not found anywhere
		log.Println("PORT not found, using default:", portString)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Println("DB_URL is not found in the enviroment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to database:", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	fmt.Println("PORT IS:", portString)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://* "},
		AllowedMethods:   []string{"GET", "POS", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.handlerGetPostsForUser))
	v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Delete("/feed_follows/{feedfollowID}", apiCfg.MiddlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("server starting on the port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}
