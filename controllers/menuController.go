package controllers

import (
	"context"
	"net/http"
	"time"
)

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("menu_id")
		var food models.Menu

		err := foodCollection.FindOne(ctx, bson.M{"menu_id": menuid}).Decode(&menu)
		defer cancel
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured while fetching the menu"}
			)
		}
		c.JSON(http.StatusOK,menu)
	}
}
