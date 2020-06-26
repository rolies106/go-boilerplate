package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	u "mortred/utils"

	"errors"
	"time"
)

//a struct to rep user account
type UserBuyers struct {
	BaseModel `json:"-" gorm:"-" schema:"-" structs:"-"`

	ID            uint64     `json:"id" gorm:"primary_key"`
	Email         string     `json:"email" schema:"email" validate:"email,required"`
	Password      string     `json:"-" schema:"password" validate:"required"`
	Token         string     `json:"token" sql:"-" schema:"-"`
	CompanyId     uint64     `json:"company_id" schema:"company_id" validate:"required"`
	Username      string     `json:"username" schema:"username" validate:"required"`
	VerifiedEmail bool       `json:"verified_email" schema:"verified_email"`
	Status        bool       `json:"status" schema:"status"`
	Name          string     `json:"name" schema:"name" validate:"required,min=3,max=255"`
	Gender        string     `json:"gender" schema:"gender"`
	Phone         string     `json:"phone" schema:"phone" validate:"required"`
	Mobile        string     `json:"mobile" schema:"mobile"`
	JobTitle      string     `json:"job_title" schema:"job_title"`
	Avatar        uint64     `json:"avatar" schema:"avatar"`
	LastLoginAt   *time.Time `json:"last_login_at" schema:"last_login_at"`
	CreatedAt     *time.Time `json:"created_at" schema:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" schema:"-"`
	DeletedAt     *time.Time `json:"deleted_at" schema:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (this *UserBuyers) TableName() string {
	return "user_buyers"
}

// Before create callback
func (this *UserBuyers) BeforeCreate() (err error) {

	// Struct validation
	checkStruct := validate.Struct(this)
	this.validateStruct(checkStruct)

	// Create password for new user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(this.Password), bcrypt.DefaultCost)
	this.Password = string(hashedPassword)

	if this.BaseModel.IsHasStructError() {
		err = errors.New("Invalid data request")
	}

	return
}

// After create callback
func (this *UserBuyers) AfterCreate(scope *gorm.Scope) (err error) {

	token := &Token{}
	this.Token = token.CreateJWTToken(this.TableName(), this)
	return
}

// Check email must be unique
func (this *UserBuyers) IsEmailExists() bool {
	//Email must be unique
	temp := &UserBuyers{}

	//check for errors and duplicate emails
	errDB := GetDB().Table(this.TableName()).Where("email = ?", this.Email).First(temp).Error
	if errDB != nil && errDB != gorm.ErrRecordNotFound {
		panic("Connection error. Please retry")
	}

	if temp.Email != "" {
		return true
	}

	return false
}

// User login
func Login(email, password string) (interface{}, bool) {

	account := &UserBuyers{}
	err := GetDB().Table(account.TableName()).Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found"), false
		}
		return u.Message(false, "Connection error. Please retry"), false
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again"), false
	}

	//Create JWT token
	token := &Token{}
	account.Token = token.CreateJWTToken(account.TableName(), account) //Store the token in the response

	return account, true
}
