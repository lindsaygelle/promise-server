package promise

import "github.com/gin-gonic/gin"

func (s *stepRouter) GetAllByTask(context *gin.Context) {
	s.stepService.Get(context.Param("id"))
}
