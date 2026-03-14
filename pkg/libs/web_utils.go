package libs

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type errorInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func RespondError(ctx *gin.Context, statusCode int, message string, err error) {
	ctx.JSON(statusCode, gin.H{
		"error": true,
		"data": gin.H{
			"status":  statusCode,
			"message": message + " " + err.Error(),
		},
	})
}

func RespondSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"error": false,
		"data":  data,
	})
}

func BodyParser(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(body)
}

func HandleBadRequestErrWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}

	w.WriteHeader(http.StatusBadRequest)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusBadRequest,
			Message: message + ": " + err.Error(),
		}})
	return err
}

func HandleNotFoundError(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}

	w.WriteHeader(http.StatusNotFound)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusNotFound,
			Message: message + ": " + err.Error(),
		}})
	return err
}

func HandleBadRequestErr(w http.ResponseWriter, err error) error {
	if err == nil {
		return nil
	}

	w.WriteHeader(http.StatusBadRequest)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}})
	return err
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	bytes, _ := json.MarshalIndent(data, "", "  ")

	w.Header().Set("Content-Type", "Application/json")
	w.Write(bytes)
}

func HandleInternalServerError(w http.ResponseWriter, err error) error {
	if err == nil {
		return nil
	}

	w.WriteHeader(http.StatusInternalServerError)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusInternalServerError,
			Message: "internal server error: " + err.Error(),
		}})
	return err
}

func HandleNotFoundErr(w http.ResponseWriter, err error) error {
	if err == nil {
		return nil
	}

	w.WriteHeader(http.StatusNotFound)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusNotFound,
			Message: "not found: " + err.Error(),
		}})
	return err
}

func HandleUnauthorizedErr(w http.ResponseWriter, err error) error {
	if err == nil {
		return nil
	}

	w.WriteHeader(http.StatusUnauthorized)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized: " + err.Error(),
		}})
	return err
}

func WriteJSONWithSuccess(w http.ResponseWriter, data interface{}) {
	data = response{
		Error: false,
		Data:  data,
	}
	bytes, _ := json.MarshalIndent(data, "", "  ")
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func ParsePositiveIntQueryParam(ctx *gin.Context, query string) (int, error) {
	if query == "" {
		return 0, nil
	}

	num, err := strconv.Atoi(ctx.Query(query))
	if err != nil {
		return 0, err
	}
	if num < 1 {
		return 0, &strconv.NumError{
			Func: "ParsePositiveIntQueryParam",
			Num:  ctx.Query(query),
			Err:  strconv.ErrSyntax,
		}
	}
	return num, nil
}

func ParseSrtQueryParamToTime(ctx *gin.Context, query string) (time.Time, error) {
	dateStr := ctx.Query(query)
	if dateStr == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
