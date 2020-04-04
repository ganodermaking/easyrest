# easyrest
a easy restful framework

# install
```go
func main() {
	r := easyrest.New()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id([0-9]+)/:name([a-z]+)", func(c *easyrest.Context) {
			c.JSON(http.StatusOK, easyrest.H{
				"id":   c.Query("id"),
				"name": c.Query("name"),
			})
		})

		v1.POST("/users/login", func(c *easyrest.Context) {
			c.JSON(http.StatusOK, easyrest.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":5000")
}
```
