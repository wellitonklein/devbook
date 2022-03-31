package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusCreated, user)
}

func FindAllUser(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.URL.Query().Get("search"))

	db, err := database.Connection()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	users, err := repository.Find(search)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusOK, users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	user, err := repository.FindById(userId)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIdInToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdInToken {
		responses.ErrorJSON(w, http.StatusForbidden, errors.New("nâo é possível atualizar um usuário que não seja o seu"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("edit"); err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	if err = repository.Update(userId, user); err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIdInToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdInToken {
		responses.ErrorJSON(w, http.StatusForbidden, errors.New("nâo é possível deletar um usuário que não seja o seu"))
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	if err = repository.Delete(userId); err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusNoContent, nil)
}
