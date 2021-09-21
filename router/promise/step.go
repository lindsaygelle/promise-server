package promise

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/service/promise"
)

type stepRouter struct {
	routerGroup *gin.RouterGroup
	stepService promise.StepService
}

func newStepRouter(database *sql.DB, routerGroup *gin.RouterGroup) *stepRouter {
	return &stepRouter{
		routerGroup: routerGroup,
		stepService: promise.NewStepService(database)}
}
