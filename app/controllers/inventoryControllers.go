package controllers

import (
	"net/http"
	u "main/utils"
	"main/models"
	"encoding/json"
)

var CreateInventory = func(w http.ResponseWriter, r *http.Request) {

	inventory := &models.Inventory{}
	err := json.NewDecoder(r.Body).Decode(inventory) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	response := inventory.Create()
	u.Respond(w, response)

}

var ListInventory = func(w http.ResponseWriter, r *http.Request) {

	response := models.List(10)
	u.Respond(w, response)

}
