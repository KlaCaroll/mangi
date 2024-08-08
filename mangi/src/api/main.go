package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

type any = interface{}

type Service struct {
	DB *sqlx.DB
}

func main() {
	var (
		addr string
		dsn  string // Data Source Name
	)
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "addr to listen on")
	flag.StringVar(&dsn, "dsn", "root:@tcp(0.0.0.0:3306)/mangi?parseTime=true", "path to the database to use")
	flag.Parse()

	// INITIALIZE THE DATABASE CONNEXION

	log.Println("opening connection to", dsn)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println("opening connection", err)
		return
	}
	log.Println("opened connection")
	defer db.Close()

	s := Service{
		DB: db,
	}

	var mux = http.NewServeMux()

	mux.HandleFunc("/user/register", s.UserRegister)
	mux.HandleFunc("/user/login", s.UserLogin)

	mux.HandleFunc("/recipe/create", s.CreateRecipe)

	mux.HandleFunc("/meal/create", s.CreateMeal)
	mux.HandleFunc("/meals", s.FetchMeals)

	// Start the HTTP server.
	handler := cors.Default().Handler(mux)
	// CORS => handler
	var srv = &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Println("listen on addr", addr)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println("listening", err)
		return
	}
}

func write(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	raw, _ := json.Marshal(payload)
	w.Write(raw)
}

type apiError struct {
	Code string `json:"code"`
	Err  string `json:"err"`
}

func writeError(w http.ResponseWriter, code string, err error) {
	write(w, apiError{
		Code: code,
		Err:  err.Error(),
	})
}

func read(r *http.Request, payload any) (err error) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, payload)
	if err != nil {
		return err
	}
	return nil
}
