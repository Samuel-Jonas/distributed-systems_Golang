package controllers

import (
	"dataViewer/database"
	"dataViewer/entities"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Return struct {
	Sucess bool
	Id     string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)

	result := db.Create(&user)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(errors.New("User not created"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	jsonString := Return{Sucess: true, Id: user.UserId.String()}
	message, _ := json.Marshal(jsonString)
	json.NewEncoder(w).Encode(string(message))

}

func checkIfUserExists(userId string) bool {
	db := database.OpenConnection()
	var user entities.User
	result := db.First(&user, "user_id = ?", userId)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	params := mux.Vars(r)
	userId := params["id"]
	fmt.Println(params)

	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found!")
		return
	}
	var user entities.User
	db.First(&user, "user_id = ?", userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(user)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	var users []entities.User
	db.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	params := mux.Vars(r)
	userId := params["id"]
	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found!")
		return
	}

	var user entities.User
	result := db.Model(&user).Where("user_id = ?", userId).Updates(user)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User data not update")
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["id"]
	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found!")
		return
	}

	var user entities.User
	result := db.Where("user_id = ?", userId).Delete(&user)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User not deleted")
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("User deleted successfully!")
}
