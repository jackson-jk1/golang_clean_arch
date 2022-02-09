package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	if err = user.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)
	res, err := repository.Create(user)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}
func ViewUser(w http.ResponseWriter, r *http.Request) {

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
