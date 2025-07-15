# Using Templates in Go

A simple Go web application that demonstrates how to use HTML templates to render dynamic content. This project shows how to combine multiple template files (header, content, and footer) to create a complete web page that displays course information in a table format.

## What This Project Does

This application creates a web server that:
- Serves a single route (`/`) that displays a list of courses
- Uses Go's `html/template` package to render HTML templates
- Combines multiple template files (header, content, footer) into a single page
- Displays course data (name, description, and duration) in an HTML table
- Demonstrates template composition and data binding in Go

## Project Structure

```
.
├── main.go        # Main application with HTTP server and handlers
├── header.html    # HTML header template with DOCTYPE and opening tags
├── content.html   # Main content template with course table
├── footer.html    # HTML footer template with closing tags
├── go.mod         # Go module file
└── README.md      # This file
```

## Features

- **Template Composition**: Combines multiple HTML template files into a single rendered page
- **Data Binding**: Uses Go template syntax to bind course data to HTML elements
- **Clean Structure**: Separates concerns with different template files for different page sections
- **Error Handling**: Includes proper error handling for template rendering

## How to Use

### Prerequisites

- Go 1.16 or later installed on your system
- Basic understanding of Go and HTML

### Running the Application

1. **Clone or download the project** to your local machine

2. **Navigate to the project directory**:
   ```bash
   cd using-templates
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```

4. **Open your web browser** and navigate to:
   ```
   http://localhost:8090
   ```

5. **View the result**: You should see a page titled "Using Templates in Go" with a table displaying two sample courses.

### Expected Output

The web page will display:
- A page title "Using Templates in Go"
- A heading "Cursos" (Courses)
- A table with three columns: Curso, Descrição, Carga Horária
- Two sample courses with their respective information

## Code Structure

### Main Components

- **Course Struct**: Defines the data structure for course information
- **CourseHandler**: HTTP handler function that prepares data and renders templates
- **Template Files**: HTML templates that define the page structure and presentation

### Template Usage

The application demonstrates:
- Loading multiple template files using `template.ParseFiles()`
- Using the `{{range}}` directive to iterate over course data
- Accessing struct fields with `{{.FieldName}}` syntax
- Combining templates to create a complete HTML document

## Customization

### Adding More Courses

To add more courses, modify the `courses` slice in the `CourseHandler` function:

```go
courses := []Course{
    {
        Name:            "Your Course Name",
        Description:     "Your Course Description",
        DurationInHours: "Hours",
    },
    // Add more courses here...
}
```

### Modifying Templates

- **header.html**: Change page title, add CSS, or modify the HTML head section
- **content.html**: Modify the table structure, add styling, or change the layout
- **footer.html**: Add footer content, scripts, or additional closing tags

### Changing the Port

To run the server on a different port, modify the port number in `main.go`:

```go
http.ListenAndServe(":8080", mux) // Change 8090 to your desired port
```

## Learning Objectives

This project demonstrates:
- HTTP server setup in Go
- HTML template parsing and execution
- Template composition with multiple files
- Data binding between Go structs and HTML templates
- Basic web application structure in Go

## Dependencies

This project uses only Go standard library packages:
- `html/template`: For HTML template processing
- `net/http`: For HTTP server functionality

No external dependencies are required.