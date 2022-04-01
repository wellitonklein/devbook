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

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	publication := models.Publication{}
	if err = json.Unmarshal(bodyRequest, &publication); err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = userId

	if err = publication.Prepare(); err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryPublication(db)
	publication.ID, err = repository.Create(publication)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusCreated, publication)
}

func FindAllPublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryPublication(db)
	publications, err := repository.Find(userId)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusOK, publications)
}

func FindPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
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

	repository := repositories.NewRepositoryPublication(db)
	pubication, err := repository.FindById(publicationId)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusOK, pubication)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
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

	repository := repositories.NewRepositoryPublication(db)
	publicationInDatabase, err := repository.FindById(publicationId)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	if publicationInDatabase.AuthorID != userId {
		responses.ErrorJSON(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicação que não seja sua"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	publication := models.Publication{}
	if err = json.Unmarshal(bodyRequest, &publication); err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = publication.Prepare(); err != nil {
		responses.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(publicationId, publication); err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusNoContent, nil)
}

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
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

	repository := repositories.NewRepositoryPublication(db)
	publicationInDatabase, err := repository.FindById(publicationId)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	if publicationInDatabase.AuthorID != userId {
		responses.ErrorJSON(w, http.StatusForbidden, errors.New("não é possível excluir uma publicação que não seja sua"))
		return
	}

	if err = repository.Delete(publicationId); err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusNoContent, nil)
}

func FindByUser(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewRepositoryPublication(db)
	publications, err := repository.FindByUser(userId)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusOK, publications)
}

func Like(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
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

	repository := repositories.NewRepositoryPublication(db)

	if err = repository.Like(publicationId); err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusNoContent, nil)
}

func Unlike(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
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

	repository := repositories.NewRepositoryPublication(db)

	if err = repository.Unlike(publicationId); err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessJSON(w, http.StatusNoContent, nil)
}
