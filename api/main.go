package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

type Service struct {
	DB        *sqlx.DB
	secretKey []byte
}

func main() {
	var (
		addr      string // Defaut adress to listen on
		dsn       string // Defaut data Source Name
		secretKey []byte // Defaut key that can be add in a flag
	)

	secretKeyS := string(secretKey)
	flag.StringVar(&addr, "addr", "127.0.0.1:8080", "addr to listen on")
	flag.StringVar(&dsn, "dsn", "root:@tcp(127.0.0.1:3306)/mangi?parseTime=true", "path to the database to use")
	flag.StringVar(&secretKeyS, "secretKey", "default secret key", "secretKey for JWT")
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
		DB:        db,
		secretKey: secretKey,
	}

	var mux = http.NewServeMux()

	mux.HandleFunc("POST /register", s.Register)
	mux.HandleFunc("POST /login", s.Login)
	mux.HandleFunc("DELETE /user", s.UserDelete)
	mux.HandleFunc("GET /user", s.ShowUser)
	mux.HandleFunc("PUT /user", s.UserUpdate)
	mux.HandleFunc("PUT /password", s.UpdatePassword)
	mux.HandleFunc("PUT /email", s.UpdateEmail)
	mux.HandleFunc("PUT /user/ustensil", s.UpdateUserUstensil)
	mux.HandleFunc("PUT /user/preference", s.UpdateUserPreference)
	mux.HandleFunc("GET /favorites", s.FetchRecipesFavorites)
	mux.HandleFunc("POST /favorite", s.CreateFavoritesRecipes)
	mux.HandleFunc("GET /user/rgpd-data", s.FetchUserDataRGDP)

	mux.HandleFunc("POST /user/home/create", s.CreateUserHome)
	mux.HandleFunc("PUT /user/home/delete", s.DeleteUserHome)
	mux.HandleFunc("POST /home/invitation", s.InviteUserHome)
	mux.HandleFunc("PUT /home/invitation", s.AcceptHomeInvitation)
	mux.HandleFunc("POST /home", s.ShowHome)
	mux.HandleFunc("GET /homes", s.FetchHomes)
	mux.HandleFunc("GET /home/shopping-lists", s.FetchHomeShoppingLists)

	mux.HandleFunc("POST /recipe", s.CreateRecipe)
	mux.HandleFunc("DELETE /recipe", s.DeleteRecipe)
	mux.HandleFunc("PUT /recipe", s.UpdateRecipe)
	mux.HandleFunc("GET /recipe", s.ShowRecipe)
	mux.HandleFunc("GET /recipes", s.FetchRecipes)
	mux.HandleFunc("GET /items", s.FetchItems)
	mux.HandleFunc("GET /recipes/list", s.FetchRecipeNameList)
	mux.HandleFunc("GET /ustensils/list", s.FetchUstensilsList)
	mux.HandleFunc("GET /categories/list", s.FetchCategoriesList)

	mux.HandleFunc("POST /meal", s.CreateMeal)
	mux.HandleFunc("GET /meals", s.FetchMeals)
	mux.HandleFunc("DELETE /meal", s.DeleteMeal)
	mux.HandleFunc("PUT /meal", s.UpdateMeal)
	mux.HandleFunc("GET /meal", s.ShowMeal)

	mux.HandleFunc("POST /compute-shopping-list", s.ComputeShoppingList)
	mux.HandleFunc("PUT /shopping-list/add-items", s.ShoppingListAddingItem)
	mux.HandleFunc("PUT /shopping-list/delete-items", s.ShoppingListDeletingItem)
	mux.HandleFunc("PUT /shopping-list/delete", s.DeleteShoppingList)
	mux.HandleFunc("PUT /shopping-list", s.UpdateShoppingList)
	mux.HandleFunc("POST /shopping-list", s.ShowShoppingList)
	mux.HandleFunc("GET /shopping-lists", s.FetchShoppingLists)

	handler := cors.AllowAll().Handler(mux)
	// CORS => handler
	var srv = &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Println("listen on addr", addr)
	// Start the HTTP server.
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println("listening", err)
		return
	}
}
