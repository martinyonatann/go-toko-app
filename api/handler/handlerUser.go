package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/martinyonatann/go-invoice/usecase/user"
	"github.com/urfave/negroni"
)

type Response struct {
	ResponseCode int         `json:"StatusCode"`
	Message      string      `json:"Message"`
	Data         interface{} `json:"Data"`
}

func createUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createUserPayload := user.CreateUserRequest{}

		err := json.NewDecoder(r.Body).Decode(&createUserPayload)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed Create User"))
			return
		}
		userData, err := service.CreateUser(r.Context(), createUserPayload)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed Create User"))
			return
		}

		resp, _ := generateResponse(http.StatusOK, Response{
			ResponseCode: http.StatusOK,
			Message:      http.StatusText(http.StatusOK),
			Data:         userData,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(resp.Bytes())
	})
}

func generateResponse(statusCode int, dataResponse interface{}) (inputBuffer bytes.Buffer, err error) {
	if dataResponse == nil {
		return inputBuffer, errors.New("failed generate response")
	}

	if err := json.NewEncoder(&inputBuffer).Encode(dataResponse); err != nil {
		return inputBuffer, errors.New("failed encode data response")
	}

	return inputBuffer, err
}

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/v1/user", n.With(
		negroni.Wrap(createUser(service)),
	)).Methods("POST", "OPTIONS").Name("createUser")

	// r.Handle("/v1/user/{id}", n.With(
	// 	negroni.Wrap(getUser(service)),
	// )).Methods("GET", "OPTIONS").Name("getUser")
}
