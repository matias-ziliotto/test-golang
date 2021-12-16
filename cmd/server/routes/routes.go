package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/matias-ziliotto/test-golang/cmd/server/handler"
	"github.com/matias-ziliotto/test-golang/cmd/server/middleware"
	"github.com/matias-ziliotto/test-golang/internal/user"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.rg.Use(middleware.TokenAuth())

	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repo := user.NewRepository(r.db)
	service := user.NewService(repo)
	handler := handler.NewUser(service)

	r.rg.GET("/users", handler.GetAll)
}
