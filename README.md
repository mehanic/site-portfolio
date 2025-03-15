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
 # third variant but need terraform (experimental)
  ```
  terraform init
  terraform plan
  terraform apply 
  terraform destroy
```

After the containers are up, you can check if the contact_messages table has been created by connecting to the PostgreSQL container and running a query:

```
docker exec -it site-portfolio-db psql -U user -d portfolio
```
Then run the following SQL query to check if the table exists:

```
SELECT * FROM contact_messages;
```
If the table exists, the query should return the columns, or it will return an empty set if there are no entries yet.

The first part postgres-data:/var/lib/postgresql/data ensures data persistence.
The second part ./init.sql:/docker-entrypoint-initdb.d/init.sql ensures that the init.sql file is mounted and can be executed during the initialization of the PostgreSQL container.


Where is the Database Data Stored?:
The database data (including the messages you inserted) is stored in the volume mounted at /var/lib/postgresql/data inside the container. In your Docker Compose configuration, you have:


volumes:
  - postgres-data:/var/lib/postgresql/data
The postgres-data volume holds the database files. You can find this volume on your local machine (outside the container) by inspecting your Docker/Podman volumes.

To list your Docker/Podman volumes:

```
podman volume ls
```
This will show you the volumes available on your local system. The volume postgres-data is the one where your database files are stored, which includes the data for contact_messages.

3. How to Access the PostgreSQL Database Files:
If you want to directly access the files stored by PostgreSQL, you can do so by inspecting the volume:

Find the volume path: On Docker or Podman, you can use the inspect command to get the path where the volume is stored:

```
podman volume inspect postgres-data
```
This will give you the file path of the volume on your local machine.

Navigate to the volume path: Once you know the volume path, you can go to the file system location and find the PostgreSQL data files.

4. How to Query the Database from Inside the Container:
If you want to query the database directly from inside the container (other than psql), you can:

Get a shell inside the container:

```
podman exec -it site-portfolio-db bash
```
Use psql from within the container: Once inside, you can run the psql command to access the database, like:

```
psql -U user -d portfolio
```
5. Where Are the Logs and Data Saved?:
The logs related to PostgreSQL are typically stored in the container's logs. You can access the logs via Podman by running:

```
podman logs site-portfolio-db
```

Steps to Remove Messages from PostgreSQL Database:
Log into the database: You're already inside the PostgreSQL container with this command:

```
podman exec -it site-portfolio-db psql -U user -d portfolio
```
Delete all records in the contact_messages table: Once logged into the PostgreSQL shell, execute the following command to remove all messages:

```
DELETE FROM contact_messages;
```
This will delete all records from the contact_messages table, but the table structure will remain intact.

Verify the deletion: After deleting the records, you can verify that the table is empty by running:

```
SELECT * FROM contact_messages;
```
This should return an empty result set if the deletion was successful.

If You Want to Remove the Table Completely:
If you want to drop the contact_messages table entirely (which means removing the table structure as well), you can run:

```
DROP TABLE IF EXISTS contact_messages;
```
This will delete both the data and the table itself.

Deleting the Database Volume:
If you're asking how to completely wipe the data stored in the PostgreSQL volume (including everything like tables, data, and configurations), you can delete the volume. You already found the volume associated with your container:

Volume name: site-portfolio_postgres-data
To remove the volume completely (thus wiping all PostgreSQL data), you can run:

```
podman volume rm site-portfolio_postgres-data
```
This will remove the volume, and all data inside PostgreSQL (including your tables and messages) will be wiped.


To remove all messages from the contact_messages table, use:

DELETE FROM contact_messages;

To remove the table itself, use:
```
DROP TABLE IF EXISTS contact_messages;
```
To remove the entire PostgreSQL volume (and all data), use:
```
podman volume rm site-portfolio_postgres-data
```