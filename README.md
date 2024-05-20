# Course Management API

This is a simple Course Management API written in Go using the Gorilla Mux router. It provides basic CRUD operations for managing courses. The API is designed to run as a web server and handle HTTP requests for creating, reading, updating, and deleting courses.

## Features

- **Get all courses**: Retrieve a list of all courses.
- **Get one course**: Retrieve details of a single course by its ID.
- **Create a new course**: Add a new course to the list.
- **Update a course**: Modify the details of an existing course.
- **Delete a course**: Remove a course from the list.

## Endpoints

- **GET /**: Home page.
- **GET /courses**: Get a list of all courses.
- **GET /course/{id}**: Get details of a single course by ID.
- **POST /course**: Create a new course.
- **PUT /course/{id}**: Update an existing course by ID.
- **DELETE /course/{id}**: Delete a course by ID.

## Running the Application

### Prerequisites

- Go (latest version recommended)
- Gorilla Mux package

### Setup

1. **Clone the repository**:
    ```bash
    git clone <repository_url>
    cd <repository_directory>
    ```

2. **Install dependencies**:
    ```bash
    go get -u github.com/gorilla/mux
    ```

3. **Run the application**:
    ```bash
    go run main.go
    ```

4. **The server will start on port 8000**. You can access it at `http://localhost:8000`.

## Example Requests

### Get All Courses

```bash
curl -X GET http://localhost:8000/courses
```

### Get One Course

```bash
curl -X GET http://localhost:8000/course/1
```

### Create a New Course

```bash
curl -X POST http://localhost:8000/course -H "Content-Type: application/json" -d '{
    "coursename": "Go Programming",
    "courseprice": 300,
    "author": {
        "fullname": "Alice Smith",
        "website": "www.alicesmith.com"
    }
}'
```

### Update a Course

```bash
curl -X PUT http://localhost:8000/course/1 -H "Content-Type: application/json" -d '{
    "coursename": "Advanced Go Programming",
    "courseprice": 350,
    "author": {
        "fullname": "Alice Smith",
        "website": "www.alicesmith.com"
    }
}'
```

### Delete a Course

```bash
curl -X DELETE http://localhost:8000/course/1
```

## Code Structure

- **Model**:
  - `Course`: Represents a course with ID, name, price, and author.
  - `Author`: Represents the author of a course with full name and website.

- **Fake Database**:
  - `courses`: A slice to store the list of courses.

- **Middleware**:
  - `IsEmpty()`: Checks if a course object is empty.

- **Controller Functions**:
  - `HomePage(w http.ResponseWriter, r *http.Request)`: Handles requests to the home page.
  - `GetCourses(w http.ResponseWriter, r *http.Request)`: Retrieves all courses.
  - `GetOneCourse(w http.ResponseWriter, r *http.Request)`: Retrieves a single course by ID.
  - `CreateCourse(w http.ResponseWriter, r *http.Request)`: Creates a new course.
  - `UpdateCourse(w http.ResponseWriter, r *http.Request)`: Updates an existing course by ID.
  - `DeleteCourse(w http.ResponseWriter, r *http.Request)`: Deletes a course by ID.

- **Main Function**:
  - Initializes the server, sets up routes, and starts listening on port 8000.



If you have any questions or suggestions, please feel free to reach out.

---

Happy coding!