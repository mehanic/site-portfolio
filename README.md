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


------------


The error message `FATAL: password authentication failed for user "mehanic"` means that PostgreSQL is rejecting the password. Here’s how to fix it:

---

### **1. Reset the Password for User "mehanic"**
Since you are logged in as `postgres`, you can manually reset the password.

1. In the PostgreSQL prompt (`postgres=#`), run:

   ```sql
   ALTER USER mehanic WITH PASSWORD 'newpassword';
   ```

   **Replace `'newpassword'` with a strong password**.

2. Exit the PostgreSQL prompt:

   ```sql
   \q
   ```

---

### **2. Verify `pg_hba.conf` (PostgreSQL Authentication Configuration)**

Since your database is running on Ubuntu, check if the authentication method is correct:

1. Open `pg_hba.conf` with a text editor:

   ```bash
   sudo nano /etc/postgresql/16/main/pg_hba.conf
   ```

2. Look for a line like this:

   ```plaintext
   local   all             all                                     peer
   ```

   **Change it to:**

   ```plaintext
   local   all             all                                     md5
   ```

   This makes sure PostgreSQL asks for a password instead of using system user authentication.

3. **Save and Exit** (`CTRL + X`, then `Y`, then `Enter`).

---

### **3. Restart PostgreSQL**
After modifying `pg_hba.conf`, restart PostgreSQL to apply changes:

```bash
sudo systemctl restart postgresql
```

---

### **4. Try Logging In Again**
Now, try logging in with the new password:

```bash
sudo -u postgres psql
```

Then, switch to the database:

```sql
\c portfolio mehanic
```

Enter the new password and check if the connection is successful.

---

### **5. Update Connection String in Your Application**
Now that the password is reset, update your **Go application’s** database connection string:

```go
DBConnStr = "user=mehanic password=newpassword dbname=portfolio host=127.0.0.1 port=5432 sslmode=disable"
```

**Then, restart your application:**

```bash
go run cmd/main.go
```

---

### **Summary:**
1. **Reset the password** for the `mehanic` user in PostgreSQL.
2. **Change `pg_hba.conf`** to use `md5` authentication.
3. **Restart PostgreSQL** to apply the changes.
4. **Test login again** using `psql`.
5. **Update your application’s database connection string**.

This should resolve your authentication issue. Let me know if you need more help! 🚀


The error **"role 'mehanic' does not exist"** means the PostgreSQL user (role) you're trying to use doesn't exist in the database. You need to create it first.

---

### **1. Create the User (Role) "mehanic"**
Since you are logged in as `postgres`, run the following command in the PostgreSQL prompt:

```sql
CREATE ROLE mehanic WITH LOGIN PASSWORD 'newpassword';
```

> Replace `'newpassword'` with a strong password.

Then, **grant privileges** to the user:

```sql
ALTER ROLE mehanic CREATEDB;
ALTER ROLE mehanic CREATEROLE;
ALTER ROLE mehanic SUPERUSER;
```

---

### **2. Verify If the Role Exists**
Check if the user `mehanic` is created:

```sql
\du
```

It should display a list of database roles, including `mehanic`.

---

### **3. Create the Database (If It Doesn't Exist)**
If your database `portfolio` does not exist, create it:

```sql
CREATE DATABASE portfolio OWNER mehanic;
```

Then, **grant all privileges**:

```sql
GRANT ALL PRIVILEGES ON DATABASE portfolio TO mehanic;
```

---

### **4. Update `pg_hba.conf` (If Needed)**
If you're still getting authentication errors:

1. Open the PostgreSQL authentication file:

   ```bash
   sudo nano /etc/postgresql/16/main/pg_hba.conf
   ```

2. Find this line:

   ```plaintext
   local   all             all                                     peer
   ```

   **Change it to:**

   ```plaintext
   local   all             all                                     md5
   ```

3. Save the file and exit (`CTRL + X`, then `Y`, then `Enter`).

4. Restart PostgreSQL:

   ```bash
   sudo systemctl restart postgresql
   ```

---

### **5. Test the Connection**
Now, try logging in as `mehanic`:

```bash
psql -U mehanic -d portfolio -h 127.0.0.1 -W
```

It should prompt for a password—enter `newpassword`.

---

### **6. Update Your Application**
Modify your **Go application's** database connection string in `config.go`:

```go
DBConnStr = "user=mehanic password=newpassword dbname=portfolio host=127.0.0.1 port=5432 sslmode=disable"
```

Restart your application:

```bash
go run cmd/main.go
```

---

### **Summary**
✅ **Created the user (`mehanic`)**  
✅ **Created the database (`portfolio`)**  
✅ **Granted permissions**  
✅ **Updated authentication settings (`pg_hba.conf`)**  
✅ **Restarted PostgreSQL**  
✅ **Updated the app's connection string**  

Now, your app should connect successfully! 🚀 Let me know if you need more help.


sudo -u ubuntu /usr/local/bin/myservice
