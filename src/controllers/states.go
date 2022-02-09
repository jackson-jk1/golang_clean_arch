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

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return

	}

	var state models.State
	if err = json.Unmarshal(requestBody, &state); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = state.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repository := repositories.NewStatesRepository(db)
	res, err := repository.Create(state)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}

func Show(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {

		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewStatesRepository(db)
	states, err := repository.Show()
	if err != nil {

		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, states)

}

func View(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uf := params["stateUF"]
	db, err := database.Connect()
	if err != nil {

		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewStatesRepository(db)
	state, err := repository.Find(uf)
	if err != nil {

		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, state)

}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uf := params["stateUF"]

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return

	}
	var state models.State
	if err = json.Unmarshal(requestBody, &state); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = state.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {

		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewStatesRepository(db)

	res, err := repository.Update(uf, state)
	if err != nil {

		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uf := params["stateUF"]
	db, err := database.Connect()
	if err != nil {

		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewStatesRepository(db)

	res := repository.Delete(uf)
	if res != nil {

		response.Erro(w, http.StatusInternalServerError, res)
		return
	}
	response.JSON(w, http.StatusOK, "successfully deleted")
}
