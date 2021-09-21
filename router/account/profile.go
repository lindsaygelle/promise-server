package account

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/service/account"
)

type profileRouter struct {
	profileService account.ProfileService
	routerGroup    *gin.RouterGroup
}

func newProfileRouter(database *sql.DB, routerGroup *gin.RouterGroup) *profileRouter {
	return &profileRouter{
		profileService: account.NewProfileService(database),
		routerGroup:    routerGroup}
}
