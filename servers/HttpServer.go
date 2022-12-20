package servers

import (
	"context"
	"database/sql"
	"fmt"
	"microtwo/database"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"microtwo/repositories"
)

type HttpServer struct {
	db                 *sql.DB
	port               string
	router             *mux.Router
	CustomerRepository repositories.CustomerRepository
	UserRepository     repositories.UserRepository
}

func NewHttpServer(ctx context.Context) (*HttpServer, error) {
	// Get a database handle.
	db, err := database.GetDbSqlConnection()
	if err != nil {
		return nil, err
	}

	// Set port server will listen on
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}

	// Set port server will listen on
	port := ":" + os.Getenv("PORT")

	// Create repositories
	customerRepository, _ := repositories.NewCustomerRepository(db)
	userRepository, _ := repositories.NewUserRepository(db)

	// Return a new HttpServer
	return &HttpServer{
		db,
		port,
		mux.NewRouter(),
		customerRepository,
		userRepository,
	}, nil
}

func (s *HttpServer) Start(router func(h *HttpServer, r *mux.Router)) {

	// Add routes
	router(s, s.router)

	// Add CORS support
	handlerRoutes := cors.Default().Handler(s.router)

	// Start the server
	fmt.Println("Server listening on port", s.port)
	if err := http.ListenAndServe(s.port, handlerRoutes); err != nil {
		panic(err)
	}
}

func (s *HttpServer) Close() {
	// Close the database connection
	s.db.Close()
}

func (s *HttpServer) GetDb() *sql.DB {
	return s.db
}

func (s *HttpServer) SetDb(db *sql.DB) {
	s.db = db
}
