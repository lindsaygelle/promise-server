package account

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type accountRouter struct {
	profileRouter profileRouter
}

func NewAccountRouter(database *sql.DB, routerGroup *gin.RouterGroup) *accountRouter {
	return &accountRouter{
		profileRouter: *newProfileRouter(database, routerGroup)}
}
