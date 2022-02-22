package main

import (
	"fmt"
	"todo-app/app/models"
)

func main() {

	fmt.Println(models.Db)

		// u := &models.User{}
		// u.Name = "test2"
		// u.Email = "test2@example.com"
		// u.Password = "testtest"
		// fmt.Println(u)

		// u.CreateUser()


	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "tets2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// fmt.Println(u)
	/*
		user, _ := models.GetUser(2)
		user.CreatedTodo("first todo")
	*/

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreatedTodo("third todo")

	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

		// user2, _ := models.GetUser(3)
		// todos, _ := user2.GetTodosByUser()
		// 		for _, v := range todos {
		// 	fmt.Println(v)
		// }

		t, _ := models.GetTodo(3)
		t.DeleteTodo()
		
}
