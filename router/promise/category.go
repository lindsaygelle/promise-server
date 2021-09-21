package promise

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/service/promise"
)

type categoryRouter struct {
	categoryService promise.CategoryService
	routerGroup     *gin.RouterGroup
}

func newCategoryRouter(database *sql.DB, routerGroup *gin.RouterGroup) *categoryRouter {
	return &categoryRouter{
		categoryService: promise.NewCategoryService(database),
		routerGroup:     routerGroup}
}
