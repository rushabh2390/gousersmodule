package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rushabh2390/gousersmodule/pkg/models"
	"github.com/rushabh2390/gousersmodule/pkg/utils"

)

// @Summary Get Users
// @Description Get list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary Get User by ID
// @Description Get details of a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id} [get]
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		res, _ := json.Marshal(&models.ErrorResponse{
			Message: err.Error()})
		w.Header().Set("Content-Type", "application.json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		fmt.Println("error while parsing")
	}
	userDetails, _, err := models.GetUserById(Id)
	if err != nil {
		res, _ := json.Marshal(&models.ErrorResponse{
			Message: err.Error()})
		w.Header().Set("Content-Type", "application.json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		// log.Fatal("get error while getting User by id")
	}
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary Create User
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	createuser := &models.User{}
	utils.ParseBody(r, createuser)
	u := createuser.CreateUser()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary Delete User
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userid"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		res, _ := json.Marshal(&models.ErrorResponse{
			Message: err.Error()})
		w.Header().Set("Content-Type", "application.json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		fmt.Println("error while parsing")
	}
	userDetails, err := models.DeleteUser(Id)
	if err != nil {
		res, _ := json.Marshal(&models.ErrorResponse{
			Message: err.Error()})
		w.Header().Set("Content-Type", "application.json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		// log.Fatal("get error while getting User by id")
	}
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary Update User
// @Description Update an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateuser := &models.User{}
	utils.ParseBody(r, updateuser)
	u, err := updateuser.UpdateUser()
	if err != nil {
		res, _ := json.Marshal(&models.ErrorResponse{
			Message: "get error while getting User by id"})
		w.Header().Set("Content-Type", "application.json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		// log.Fatal("")
	}
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary Login User
// @Description Authenticates a user and returns a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body models.Credentials true "Login Credentials"
// @Success 200 {object} models.JWTResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /login [post]
func LoginUser(w http.ResponseWriter, r *http.Request) {
	logincredentials := &models.Credentials{}
	utils.ParseBody(r, logincredentials)
	claim, err := logincredentials.LoginUser()
	if err != nil {
		res, _ := json.Marshal(&models.ErrorResponse{
			Message: "get error while getting User by id"})
		w.Header().Set("Content-Type", "application.json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
		// log.Fatal("")
	}
	res, _ := json.Marshal(claim)
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
