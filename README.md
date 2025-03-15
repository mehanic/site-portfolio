```go
func getLanguage(c *gin.Context) string {
	lang := c.Query("lang") 

	// Allow English, German, and Dutch. Default to English.
	if lang != "de" && lang != "en" && lang != "nl" {
		lang = "en"
	}

	log.Printf("Selected language: %s", lang) 
	return lang
}
```

### 2. **Update the `renderPage` function:**
Ensure that `Lang` is passed as a dynamic variable so the template can handle the links:

```go
func renderPage(c *gin.Context, page string) {
	lang := getLanguage(c)                         // Get the selected language
	template := page + "_" + lang + ".html"        // Generate the template name
	log.Printf("Rendering template: %s", template) // Log the template being used
	c.HTML(http.StatusOK, template, gin.H{"Lang": lang}) // Pass the language variable
}
```


### 3. **Provision portfollio:**
# first variant but need postgresql as systemd
```
go run cmd/main.go
```
# second variand but need podman and podman-compose
  ```
  podman-compose build
  podman-compose up
  podman-compose down
  podman images 
  podman rmi portfolio 
 ```
 # third variant but need terraform
  ```
  terraform init
  terraform plan
  terraform apply 
  terraform destroy
```