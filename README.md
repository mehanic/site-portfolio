# site-portfolio

It looks like PostgreSQL is using **peer authentication**, which means it tries to match the system user with the database user. Since your system user **"mehanic"** doesn’t match **"postgres"**, the authentication is failing.

---

### ✅ **Fix: Switch to Password Authentication**
1. **Open the PostgreSQL authentication config file** (`pg_hba.conf`):  

   ```sh
   sudo nano /etc/postgresql/*/main/pg_hba.conf
   ```

2. **Find the line that looks like this** (near the top):

   ```
   local   all   postgres   peer
   ```

   **Change `peer` to `md5`** so it looks like this:

   ```
   local   all   postgres   md5
   ```

3. **Save and exit**:  
   - Press `CTRL + X`
   - Press `Y` to save
   - Press `Enter`

4. **Restart PostgreSQL**:

   ```sh
   sudo systemctl restart postgresql
   ```

---

### ✅ **Set a Password for "postgres"**
Now, log in as the **postgres system user**:

```sh
sudo -i -u postgres
```

Then, enter PostgreSQL:

```sh
psql
```

Set a password for **postgres**:

```sql
ALTER USER postgres PASSWORD 'your-secure-password';
```

Exit PostgreSQL:

```sql
\q
```

Now, try logging in again:

```sh
psql -U postgres -W
```

It should now **prompt for a password**, and you can use the one you just set.

---

### ✅ **Try Your Original Command Again**
Now, run:

```sh
psql -U postgres
```

It should work! 🎉

---

////


To add a third language, like Dutch (`nl`), you can follow these steps:

### 1. **Update the `getLanguage` function:**
Add support for the new language `nl`:

```go
func getLanguage(c *gin.Context) string {
	lang := c.Query("lang") // Get `lang` from URL query

	// Allow English, German, and Dutch. Default to English.
	if lang != "de" && lang != "en" && lang != "nl" {
		lang = "en"
	}

	log.Printf("Selected language: %s", lang) // Log selected language
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

### 3. **Create the new Dutch templates:**
In your `templates` folder, add:

- `index_nl.html`
- `services_nl.html`
- `contact_nl.html`
- ... (other pages for Dutch)

### 4. **Modify the navigation to switch between languages:**
In your HTML templates, update the language switch links. For example, in `index_en.html`:

```html
<a href="/?lang=en">English</a>
<a href="/?lang=de">Deutsch</a>
<a href="/?lang=nl">Nederlands</a>
```

### 5. **Verify the `routes` are correct:**
Your existing route definitions should already handle this:

```go
router.GET("/", HandleHome)
router.GET("/contact", HandleContact)
router.GET("/services", HandleServices)
router.GET("/exclusives", HandleExclusives)
router.GET("/support", HandleSupport)
router.GET("/courses", HandleCourses)
router.GET("/photos", HandlePhotos)
router.GET("/videos", HandleVideos)
router.GET("/author", HandleAuthor)
```

---

### ✅ **Test the application:**
- `/` for English (default)
- `/?lang=de` for German
- `/?lang=nl` for Dutch

---

Let me know if you'd like me to show the folder structure or add enhancements!