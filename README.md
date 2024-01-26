
# Go Gin Web Server with User Registration, Authentication, and Notes CRUD

This repository contains a web server implemented in Go using the Gin framework. The server includes functionalities for user registration, user login with JWT token generation, a "Me" endpoint to fetch user details, and CRUD operations for managing notes.

## Getting Started
### Prerequisites
Make sure you have Go installed on your system. You can download it from https://golang.org/dl/.

### Installation
    
1. **Clone the repository:**
    ```bash
    git clone https://github.com/amiransari27/go-gin-poc your_project_name
    cd your_project_name

2. **Install dependencies:**
    ```bash
    go mod download

3. **Configure the database:**

save `example.config.local.yml` as `config.local.yml` and provide your database connection details.

4. **Swagger docs:**

Certainly! Below is an example Swagger documentation snippet based on your provided information.

```
// @Summary Fetch notes
// @Description Fetch all notes for the authenticated user
// @Tags Notes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []model.Note
// @Router /notes [get]
// @Param Authorization header string true "Bearer Token"

// The following comment annotations are for Swagger API documentation.
// They can be added directly above the function definition in your Go code.

// FetchNotesHandler handles the endpoint to fetch all notes for the authenticated user.
// It requires a valid Bearer Token in the Authorization header.
func (c *noteController) findAll(ctx *gin.Context) ([]*model.Note, error) {
    notes, err := c.noteServ.FindAllForLoggedInUser(ctx)
    if err != nil {
        return nil, err
    }
    return notes, nil
}
```

Explanation:

    @Summary: A short summary of what the endpoint does.
    @Schemes: Specifies the communication protocol (e.g., http, https).
    @Description: A more detailed description of the endpoint's purpose.
    @Tags: Specifies the category or group the endpoint belongs to.
    @Accept: Describes the expected request format (JSON in this case).
    @Produce: Describes the expected response format (JSON in this case).
    @Security: Indicates the security scheme used for the endpoint. In this case, it assumes Bearer Token authentication.
    @Success: Describes a successful response with HTTP status code 200 and the expected response format (an array of model.Note objects in this case).
    @Router: Specifies the route of the endpoint.
    @Param: Describes a parameter of the endpoint. In this case, it specifies that the Authorization header is required and should contain a Bearer Token.


5. **Swagger init:**

Certainly! The `swag init` command generates Swagger documentation based on the comments provided in your Go code. Here's how you can run the command:

Make sure you have `swag` installed. If not, install it using:

    go get -u github.com/swaggo/swag/cmd/swag

then Run the following command in the terminal:

    swag init -g ./main.go -o cmd/docs
    
-g: Specifies the entry point of your application (main file).

-o: Specifies the output directory for the generated Swagger files (in this case, cmd/docs).

After running the command, you should see a `docs` folder generated in the specified output directory.


6. **Run the server:**
    ```bash
    go run main.go
The server should now be running at http://localhost:3002.

And Swager url should be http://localhost:3002/swagger/index.html#/