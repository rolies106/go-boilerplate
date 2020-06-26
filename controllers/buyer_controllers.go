package controllers

import (
	"github.com/gorilla/schema"

	"mortred/models"
	u "mortred/utils"
	"net/http"
)

// Register Buyer
var BuyerCreate = func(w http.ResponseWriter, r *http.Request) {

	// Parse form and prepare struct
	r.ParseForm()
	account := &models.UserBuyers{}

	// err := json.NewDecoder(buyerData.Values).Decode(account) //decode the request body into struct and failed if any error occur
	err := schema.NewDecoder().Decode(account, r.PostForm)

	if err != nil {
		u.Log("error", err)
	}

	if account.IsEmailExists() {
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Email already exists"))
		return
	}

	if err := models.GetDB().Create(&account).Error; err != nil {
		if account.BaseModel.IsHasStructError() {
			u.Respond(w, http.StatusBadRequest, u.ErrorMessage(false, "Invalid Request", account.BaseModel.ErrorsMap))
		} else {
			u.Respond(w, http.StatusBadRequest, u.Message(false, err.Error()))
		}
		return
	}

	u.Respond(w, http.StatusCreated, u.MessageWithData(true, "Registered Successfully", account))
	return
}

// Buyer Authenticate
var BuyerAuthenticate = func(w http.ResponseWriter, r *http.Request) {

	// Parse form and prepare struct
	r.ParseForm()
	account := &models.UserBuyers{}

	err := schema.NewDecoder().Decode(account, r.PostForm) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Invalid request"))
		return
	}

	// Login user
	resp, success := models.Login(account.Email, account.Password)

	if success == false {
		u.Respond(w, http.StatusBadRequest, resp)
		return
	}

	u.Respond(w, http.StatusOK, u.MessageWithData(true, "Login Successfully", resp))
}
