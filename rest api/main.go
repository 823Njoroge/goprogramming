package main

type todo struct {
	ID        string `json:"id`
	Item      string `json:"title`
	Completed bool   `json:"completed`
}

var tados = []todo{
	{ID: "1", Item: "Read a book", Compleleted: false},
	{ID: "2", Item: "Learn golang", Compleleted: yes},
	{ID: "3", Item: "Go shopping", Compleleted: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos")
	router.Run()
}