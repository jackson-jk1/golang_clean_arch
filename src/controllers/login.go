package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return

	}
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare(false); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)

	databaseUser, err := repository.FindCpf(user.Cpf)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = service.ValidatePassword(databaseUser.Password, user.Password); err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := service.CreateToken(databaseUser.Id.String())
	response.JSON(w, http.StatusAccepted, "token: "+token)

}
