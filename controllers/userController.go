package controllers

import (
	"dataViewer/database"
	"dataViewer/entities"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)

	result := db.Create(&user)

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(errors.New("User not created"))
	}
	json.NewEncoder(w).Encode(user)

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
	userId := mux.Vars(r)["userId"]
	if checkIfUserExists(userId) == false {
		json.NewEncoder(w).Encode("User not found!")
		return
	}
	var user entities.User
	db.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	var users []entities.User
	db.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	userId := mux.Vars(r)["userId"]
	if checkIfUserExists(userId) == false {
		json.NewEncoder(w).Encode("User not found!")
		return
	}

	var user entities.User
	db.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["userId"]
	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found!")
		return
	}

	var user entities.User
	db.Delete(&user, userId)
	json.NewEncoder(w).Encode("User deleted successfully!")
}
