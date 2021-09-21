package promise

import "github.com/gin-gonic/gin"

func (t *taskRouter) Get(context *gin.Context) {
	t.taskService.Get(context.Param("id"))
}

func (t *taskRouter) GetAllByProfile(context *gin.Context) {
	t.taskService.GetAllByProfile(context.Param("id"))
}
