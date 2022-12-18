package presenter

import (
	"bytes"
	"encoding/json"
	"net/http"

	logger "github.com/rs/zerolog/log"
)

type Response struct {
	ResponseCode int         `json:"status_code"`
	Message      string      `json:"message"`
	Error        string      `json:"error,omitempty"`
	Data         interface{} `json:"data"`
}

func GenerateSuccessResponse(
	statusCode int,
	data interface{},
) []byte {
	var inputBuffer bytes.Buffer

	dataResp := &Response{
		ResponseCode: statusCode,
		Message:      "success",
		Data:         data,
	}

	if err := json.NewEncoder(&inputBuffer).Encode(dataResp); err != nil {
		logger.Err(err).Msg("GenerateSuccessResponse_encoder")
	}

	return inputBuffer.Bytes()
}

func GenerateFailedResponse(
	statusCode int,
	errorText string,
) []byte {
	var inputBuffer bytes.Buffer

	dataResp := &Response{
		ResponseCode: statusCode,
		Message:      http.StatusText(statusCode),
		Error:        errorText,
	}

	if err := json.NewEncoder(&inputBuffer).Encode(dataResp); err != nil {
		logger.Err(err).Msg("GenerateFailedResponse_encoder")
	}

	return inputBuffer.Bytes()
}
