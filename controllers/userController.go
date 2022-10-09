package controllers

import (
	"dataViewer/database"
	"dataViewer/entities"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func checkIfUserExists(userId string) bool {
	var user entities.User
	database.Instance.First(&user, userId)
	if user.UserId == 0 {
		return false
	}
	return true
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if checkIfUserExists(userId) == false {
		json.NewEncoder(w).Encode("User not found!")
		return
	}
	var user entities.User
	database.Instance.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []entities.User
	database.Instance.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if checkIfUserExists(userId) == false {
		json.NewEncoder(w).Encode("User not found!")
		return
	}

	var user entities.User
	database.Instance.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["userId"]
	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found!")
		return
	}

	var user entities.User
	database.Instance.Delete(&user, userId)
	json.NewEncoder(w).Encode("User deleted successfully!")
}
