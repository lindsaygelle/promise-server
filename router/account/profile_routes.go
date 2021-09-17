package account

import "github.com/gin-gonic/gin"

func (p *profileRouter) GetProfile(context *gin.Context) {
	p.ProfileService.Get(context.Param("id"))
}

func (p *profileRouter) GetProfiles(context *gin.Context) {
	p.ProfileService.GetAll()
}
