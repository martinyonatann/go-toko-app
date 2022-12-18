package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	logger "github.com/rs/zerolog/log"

	"github.com/gorilla/mux"
	"github.com/martinyonatann/go-invoice/api/presenter"
	"github.com/martinyonatann/go-invoice/usecase/user"
	"github.com/urfave/negroni"
)

func createUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createUserPayload := user.CreateUserRequest{}

		err := json.NewDecoder(r.Body).Decode(&createUserPayload)
		if err != nil {
			logger.Err(err).Msg("createUser_decode")

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed Create User"))
			return
		}

		userData, err := service.CreateUser(r.Context(), createUserPayload)
		if err != nil {
			logger.Err(err).Msg("[handlerUser][CreateUser]")

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(presenter.GenerateFailedResponse(http.StatusInternalServerError, err.Error()))

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(presenter.GenerateSuccessResponse(http.StatusOK, userData))
	})
}

func getUserById(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			logger.Err(errors.New("user_id is required")).Msg("getUserById_varsid")

			panic(err)
		}

		getUserPayload := user.GetUserRequest{UserID: id}

		userData, err := service.GetUser(r.Context(), getUserPayload)
		if err != nil {
			logger.Err(err).Msg("[handlerUser][getUserById]")

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(presenter.GenerateFailedResponse(http.StatusInternalServerError, err.Error()))

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(presenter.GenerateSuccessResponse(http.StatusOK, userData))
	})
}

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/v1/user", n.With(
		negroni.Wrap(createUser(service)),
	)).Methods(http.MethodPost, http.MethodOptions).Name("createUser")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(getUserById(service)),
	)).Methods(http.MethodGet, http.MethodOptions).Name("getUser")
}
