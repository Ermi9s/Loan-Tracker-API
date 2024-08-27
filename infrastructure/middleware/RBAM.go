package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	domain "github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RoleBasedAuth(protected bool , user_collection database.CollectionInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get token with claims
		auth := c.GetHeader("Authorization")

		var claims = domain.UserClaims{}
		authSplit := strings.Split(auth, " ")
		_, err := jwt.ParseWithClaims(authSplit[1], &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("accessblahblah"), nil
		})

		if err != nil {
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		user := domain.UserClaims{
			ID: claims.ID,}
		
		
		duser := GetUser(user.ID , user_collection)	
		
		if duser.Is_Admin {
			c.Set("filter", bson.M{})
		} else {
			if protected {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "you need to be an admin"})
				c.Abort()
				return
			}

			path := c.Request.URL.Path
			idx := c.Param("id")
			objid, _ := primitive.ObjectIDFromHex(idx)
        	if strings.Contains(path, "user") && idx != "" && objid != user.ID {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
			c.Set("filter", bson.M{"user._id": claims.ID})
		}

		c.Next()
	}

}


func GetUser(id primitive.ObjectID , user_c database.CollectionInterface) domain.User {
	user := domain.User{}
	filter := bson.M{"_id": id}
	user_c.FindOne(context.TODO() , filter).Decode(&user)
	return user
}