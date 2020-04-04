# easyrest
a easy restful framework

# install
```go
func main() {
	r := goish.New()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id([0-9]+)/:name([a-z]+)", func(c *goish.Context) {
			c.JSON(http.StatusOK, goish.H{
				"id":   c.Query("id"),
				"name": c.Query("name"),
			})
		})

		v1.POST("/users/login", func(c *goish.Context) {
			c.JSON(http.StatusOK, goish.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":5000")
}
```
