package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/RamiroCuenca/go-moneyManager/common"
	"github.com/RamiroCuenca/go-moneyManager/user/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db := common.GetConnection()
	// defer db.Close()

	db.Find(&users)

	json, _ := json.Marshal(users)

	common.SendResponse(w, http.StatusOK, json)
}
