package promise

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/service/promise"
)

type taskRouter struct {
	routerGroup *gin.RouterGroup
	taskService promise.TaskService
}

func newTaskRouter(database *sql.DB, routerGroup *gin.RouterGroup) *taskRouter {
	return &taskRouter{
		routerGroup: routerGroup,
		taskService: promise.NewTaskService(database)}
}
