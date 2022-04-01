package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connection()
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	userInDatabase, err := repository.FindByEmail(user.Email)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(userInDatabase.Password, user.Password); err != nil {
		responses.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userInDatabase.ID)
	if err != nil {
		responses.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
