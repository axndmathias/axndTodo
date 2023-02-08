package handlers

import (
	"context"
	"time"

	db "github.com/axndmathias/axndTodo/backend/todo/db"
	mod "github.com/axndmathias/axndTodo/backend/todo/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTodo(c *gin.Context) {
	var todo mod.TodoModel
	c.BindJSON(&todo)
	collection := db.MGI.Db.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	todo.ID = primitive.NewObjectID().Hex()
	_, err := collection.InsertOne(ctx, todo)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}

	c.JSON(200, gin.H{"ststus": "success", "message": "Todo added successfully", "todo": "todo"})
}
