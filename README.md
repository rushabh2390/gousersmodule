### Golang Users module
---
This is module contains User CRUD operations and login api. You can add this users directory directly in any other project to get users api end points ready by just copy and paste user directory in project diretory and add its routes. 
### How to run this.
---   
1. Set up .env file with following keys.Set values as per your requirement
````
SECRET_KEY=xxxx #this is for encryption of token
DATABASE_URL=postgresql://postgres:postgres@localhost/loginmodule
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_HOST=localhost
DATABASE_NAME=loginmodule
DATABASE_PORT=5432
````
2. Install dependenceny   
````
go mod download
````
3. Go to cmd directory and type below commands to run app
````
go run main.go
````
4. You can see api swagger here [http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)

![Api list]("api-swagger.jpg")