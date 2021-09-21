package promise

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type promiseRouter struct {
	categoryRouter *categoryRouter
	stepRouter     *stepRouter
	taskRouter     *taskRouter
}

func NewPromiseRouter(database *sql.DB, routerGroup *gin.RouterGroup) *promiseRouter {
	return &promiseRouter{
		categoryRouter: newCategoryRouter(database, routerGroup),
		stepRouter:     newStepRouter(database, routerGroup),
		taskRouter:     newTaskRouter(database, routerGroup)}
}
