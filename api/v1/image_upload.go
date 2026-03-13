package v1

import (
	"fmt"
	"path/filepath"
	"real-holat/pkg/libs"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadImage uploads an image to Cloudflare R2 and returns a signed URL valid for 1 year
func (h *handlerV1) UploadImage(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	if err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "file is required")
		return
	}

	// Validate size (max 200 KB)
	if file.Size > 200*1024 {
		libs.HandleBadRequestErr(ctx.Writer, fmt.Errorf("image too large: image must be < 200 KB"))
		return
	}

	// Validate MIME type
	mime := file.Header.Get("Content-Type")
	if !strings.HasPrefix(mime, "image/") {
		libs.HandleBadRequestErr(ctx.Writer, fmt.Errorf("invalid file type: only image uploads are allowed"))
		return
	}

	//  Open file
	f, err := file.Open()
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}
	defer f.Close()

	// Generate unique object key
	ext := filepath.Ext(file.Filename)
	key := fmt.Sprintf("uploads/%d%s", time.Now().UnixNano(), ext)

	//  Upload using service
	publicURL, err := h.service.R2().UploadImage(ctx.Request.Context(), "safar", key, f, mime)
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return

	}

	libs.WriteJSONWithSuccess(ctx.Writer, gin.H{"url": publicURL})
}
