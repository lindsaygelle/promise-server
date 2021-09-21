package account

import "github.com/gin-gonic/gin"

func (p *profileRouter) Get(context *gin.Context) {
	p.profileService.Get(context.Param("id"))
}

func (p *profileRouter) GetAll(context *gin.Context) {
	p.profileService.GetAll()
}
