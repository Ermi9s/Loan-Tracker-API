package domain

import "github.com/gin-gonic/gin"

type User_Controller_interface interface {
	GetOneUser() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	FilterUser() gin.HandlerFunc
	UpdatePassword() gin.HandlerFunc
}

type User_Usecase_interface interface {
	GetOneUser(id string) (ResponseUser , error) 
	GetUsers() ([]ResponseUser , error)
	DeleteUser(id string) (error)
	FilterUser(map[string]string) ([]ResponseUser , error)
	UpdatePassword(id string , updated_user UpdatePassword)(ResponseUser , error) 
}

type User_Repository_interface interface{
	GetUserDocumentByID(id string) (User , error)
	GetUserDocuments() ([]User , error)
	DeleteUserDocument(id string) (error)
	FilterUserDocument(filter map[string]string) ([]User , error)
	UpdateUserPassword(id string , new_hashed_password string) (User , error)
}



