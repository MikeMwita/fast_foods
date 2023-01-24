package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.menu
		var food models.food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		err := menuCollection.Findone(ctx, bson.M{"menu_id": food.Menu.id}).Decode(&menu)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("message was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
}

func GetFoods() {

}

func GetFood() {

}

func UpdateFood() {

}