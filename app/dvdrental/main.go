package main

import (
	"database/sql"
	"dvdrental/graph"
	"dvdrental/graph/generated"
	"dvdrental/repository"
	"dvdrental/service"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	initConfig()
	database := initDatabase()

	svc := service.NewService(database)

	r := gin.Default()
	initGraphQLHandler(r, svc)

	addr := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	if err := r.Run(addr); err != nil {
		panic(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initDatabase() repository.Querier {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(2 * time.Second)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfully created connection to database")

	return repository.New(db)
}

func initGraphQLHandler(r *gin.Engine, service service.Service) {
	r.POST("/graphql", graphqlHandler(service))
	r.GET("/playground", playgroundHandler())
}

func graphqlHandler(service service.Service) gin.HandlerFunc {
	resolver := graph.NewResolver(service)
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/graphql")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
