package main

import (
	"example/config"
	"example/src/controllers"
	"example/src/middlewares"
	"example/src/models"
	"example/src/repositories"
	"example/src/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	dbCtx, err := InitDatabase(config.DatabaseConfig())
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	engine := gin.Default()
	setup(engine, dbCtx)

	log.Fatal(engine.Run())
}

func setup(engine *gin.Engine, dbCtx *gorm.DB) {
	timeoutCtx := time.Duration(viper.GetInt("TIMEOUT")) * time.Second

	/* Middlewares */
	engine.Use(middlewares.FilterTenantMiddleware())

	/* Article */
	articleRepo := repositories.NewArticleRepository(dbCtx)
	articleService := services.NewArticleService(articleRepo, timeoutCtx)
	controllers.NewArticleController(engine, articleService)
}

func InitDatabase(db *models.Database) (*gorm.DB, error) {
	dbURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Database,
	)

	fmt.Println(dbURL)

	dbCtx, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	/* auto migrate model */
	dbCtx.AutoMigrate(&models.Article{})

	return dbCtx, nil
}