package main

type todo struct {
	ID        string `json:"id`
	Item      string `json:"title`
	Completed bool   `json:"completed`
}

var tados = []todo{
	{ID: "1", Item: "Read a book", Completed: false},
	{ID: "2", Item: "Learn golang", Completed: yes},
	{ID: "3", Item: "Go shopping", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos")
	router.Run()
}