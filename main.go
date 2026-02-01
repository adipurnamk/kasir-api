package main

import (
	"kasir-api/databases"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

// ubah Config
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := databases.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup routes on default mux (existing product handlers remain intact)
	http.HandleFunc("/healthz", handlers.HealthzHandler)
	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)

	// Create Chi router for categories and mount the default mux so product routes keep working
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Get("/api/categories", categoryHandler.GetAll)
	r.Post("/api/categories", categoryHandler.Create)
	r.Get("/api/categories/{id}", categoryHandler.GetByID)
	r.Put("/api/categories/{id}", categoryHandler.Update)
	r.Delete("/api/categories/{id}", categoryHandler.Delete)

	// Mount existing DefaultServeMux so old handlers continue to work
	r.Mount("/", http.DefaultServeMux)

	log.Printf("Server running on port %s\n", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, r))

}
