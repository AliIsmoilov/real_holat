package v1

import (
	"real-holat/pkg/libs"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) MainPageStats(ctx *gin.Context) {
	stats, err := h.service.Report().MainPageStats(ctx.Request.Context())
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.WriteJSONWithSuccess(ctx.Writer, stats)
}
