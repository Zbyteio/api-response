package response

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorStruct struct {
	ErrorMsg string
	RespCode int
}

type Response struct {
	Status bool          `json:"status"`
	Data   interface{}   `json:"data"`
	Error  ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ApiResponse(c *gin.Context, errorMap map[int]ErrorStruct, result bool, data interface{}, errCode int) {
	var resp Response
	resp.Status = result
	resp.Data = data
	resp.Error.Code = errCode
	errMsg, exists := errorMap[errCode]
	if !exists {
		errMsg = ErrorStruct{
			ErrorMsg: "failed to perform the operation due to some generic error at server side",
			RespCode: http.StatusInternalServerError,
		}
	}
	if !(result) {
		resp.Error.Message = errMsg.ErrorMsg
	} else {
		resp.Error = ErrorResponse{}
	}
	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln("error in marshalling response")
		panic(err)
	}
	c.Data(errMsg.RespCode, "application/json", response)
}
