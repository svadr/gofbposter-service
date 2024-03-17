package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	supa "github.com/nedpals/supabase-go"
	"github.com/svadr/gofbposter-service/internal/models"
)

type SupabaseClient struct {
	Url string
	Key string
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

type Application struct {
	logger *slog.Logger
	posts  *models.PostModel
	router *gin.Engine
	client *supa.Client
}

func main() {
	// Load environment variables from .env file
	loadEnv()

	// Parse command line flags
	addr := flag.String("addr", ":5005", "HTTP network address")
	flag.Parse()

	// Initialize application dependencies
	app := initializeApp()

	// Start the server
	startServer(app, addr)
}

func initializeApp() *Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	sbConfig := SupabaseClient{
		Url: os.Getenv("SUPABASE_URL"),
		Key: os.Getenv("SUPABASE_KEY"),
	}
	supabaseClient := configDBClient(sbConfig)

	dbConfig := DatabaseConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db := openDB(dbConfig)

	return &Application{
		logger: logger,
		posts:  &models.PostModel{DB: db},
		router: nil,
		client: supabaseClient,
	}
}

func startServer(app *Application, addr *string) {
	app.router = app.routes()

	app.logger.Info("starting gofbposter-log-service", "addr", *addr)

	err := app.router.Run(*addr)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
