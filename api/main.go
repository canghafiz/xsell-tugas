package main

import (
	"be/apps"
	"be/dependencies"
	"be/helpers"
	"be/models/domains"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	helpers.FatalError(err)

	if os.Getenv("APP_STATUS") == "Debug" {
		gin.SetMode(gin.DebugMode)
	}

	if os.Getenv("APP_STATUS") == "Production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Env
	port := os.Getenv("APP_PORT")
	jwtKey := os.Getenv("JWT_KEY")
	appName := os.Getenv("APP_NAME")
	mapApiEndpoint := os.Getenv("MAP_API_ENDPOINT")
	mapApiKey := os.Getenv("MAP_API_KEY")

	// Database Config
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db := apps.OpenConnection(dbUser, dbPass, dbHost, dbPort, dbName)

	// Redis Config
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisConfig := domains.RedisConfig{
		Prefix:   os.Getenv("REDIS_PREFIX"),
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		DB:       redisDb,
		Password: os.Getenv("REDIS_PASS"),
	}

	// Smtp Config
	smtp := domains.Smtp{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		From:     os.Getenv("SMTP_FROM"),
	}

	// Other
	validate := validator.New()

	// Dependency
	memberDependency := dependencies.NewMemberDependency(db, validate, redisConfig, jwtKey)
	adminDependency := dependencies.NewAdminDependency(db, validate)
	dependency := dependencies.NewDependency(db, validate, smtp, appName, jwtKey, mapApiEndpoint, mapApiKey)

	// Setup Router
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	routerParent := apps.Router{
		MemberDependency: memberDependency,
		AdminDependency:  adminDependency,
		Dependency:       dependency,

		JwtKey: jwtKey,
		Engine: engine,
	}
	router := apps.NewRouter(routerParent)

	// Run Server
	if port == "" {
		port = ":3001"
	}
	err = router.Engine.Run(port)
	helpers.FatalError(err)
}
