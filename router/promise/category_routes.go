package promise

import "github.com/gin-gonic/gin"

func (c *categoryRouter) Get(context *gin.Context) {
	c.categoryService.Get(context.Param("id"))
}

func (c *categoryRouter) GetAllByProfile(context *gin.Context) {
	c.categoryService.GetAllByProfile(context.Param("id"))
}
