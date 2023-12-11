package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          int    `json:"id" uri:"id" binding:"omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func getData() []byte {
	data, _ := os.ReadFile("./data.json")
	return data
}

func addData(todos []Todo) {
	encoded, _ := json.Marshal(todos)
	os.WriteFile("./data.json", encoded, 0644)
}

func Filter(s []Todo, f func(Todo) bool) []Todo {
	filtered := make([]Todo, 0)

	for _, value := range s {
		if f(value) {
			filtered = append(filtered, value)
		}
	}

	return filtered
}

func Map(s []Todo, m func(Todo) Todo) []Todo {
	mapped := make([]Todo, len(s))

	for index, value := range s {
		mapped[index] = m(value)
	}

	return mapped
}

func main() {
	var id int = 1
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		file_content := getData()
		var response []Todo
		json.Unmarshal(file_content, &response)

		c.JSON(200, gin.H{
			"data": response,
		})
	})

	r.POST("/todos", func(c *gin.Context) {
		var todo Todo
		c.ShouldBind(&todo)

		todo.ID = id
		id += 1

		file_content := getData()
		var schema []Todo
		json.Unmarshal(file_content, &schema)

		new_data := append(schema, todo)

		addData(new_data)

		c.JSON(201, gin.H{
			"data": new_data,
		})

	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		_id, _ := strconv.Atoi(id)

		var todo Todo
		c.ShouldBind(&todo)

		file_content := getData()
		var schema []Todo
		json.Unmarshal(file_content, &schema)

		addData(Map(schema, func(_todo Todo) Todo {
			if _todo.ID == _id {
				todo.ID = _todo.ID
				return todo
			}
			return _todo
		}))

		c.JSON(200, gin.H{
			"data": todo,
		})

	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		file_content := getData()
		var schema []Todo
		json.Unmarshal(file_content, &schema)

		addData(Filter(schema, func(todo Todo) bool {
			return todo.ID != id
		}))

		c.JSON(200, gin.H{
			"message": "deleted",
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
