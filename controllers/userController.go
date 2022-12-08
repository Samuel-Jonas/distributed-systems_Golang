package controllers

import (
	"dataViewer/database"
	"dataViewer/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Return struct {
	Sucess bool   `json:"success"`
	Id     string `json:"id"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Unable to read the body: %v\n", err)
	}

	var user entities.User
	json.Unmarshal(reqBody, &user)

	if e := db.Create(&user).Error; e != nil {
		log.Println("Unable to create new user")
		jsonString := Return{Sucess: false, Id: " "}
		message, _ := json.Marshal(jsonString)
		json.NewEncoder(w).Encode(json.RawMessage(message))
		return
	}
	fmt.Println("Endpoint hit! Create new user")

	w.WriteHeader(http.StatusCreated)
	jsonString := Return{Sucess: true, Id: user.UserId.String()}
	message, _ := json.Marshal(jsonString)
	json.NewEncoder(w).Encode(json.RawMessage(message))

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
	var user entities.User
	db.First(&user, params["user_id"])

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading the body %v\n", err)
		return
	}

	json.Unmarshal(reqBody, &user)

	if result := db.Model(&user).Updates(user).Error; result != nil {
		log.Println("Unable to update user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Unable to update user")
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
	var user entities.User

	if err := db.First(&user, params["user_id"]).Error; err != nil {
		log.Printf("Unable to find id %v\n", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found!")
		return
	}

	if result := db.Delete(&user).Error; result != nil {
		log.Println("Unable to delete user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User not deleted")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("User deleted successfully!")
}
