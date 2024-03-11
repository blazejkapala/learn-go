# Go + HTMX Hello World Application: Proof of Concept

This proof of concept (PoC) demonstrates a simple web application that uses Go as the backend server and HTMX for dynamic frontend interactions, specifically showcasing a toggle feature. The application allows users to click a button to display a "Hello, World!" message. Clicking the button again hides the message, with the toggle state managed server-side in Go, leveraging sessions for persistence across page reloads.

## Objectives

- Demonstrate Go and HTMX Integration: Showcase how Go can serve both static files and handle API requests, with HTMX enabling dynamic content updates without full page reloads.
- Session Management: Utilize server-side session management in Go to remember the toggle state of the "Hello, World!" message for each user session.
- Minimal Dependency: The application aims to use minimal external libraries, focusing on Go's standard library and HTMX for frontend dynamics.

## Prerequisites

- Go (1.15 or later recommended)
- A modern web browser that supports HTMX

## Getting Started

To get the application running on your local machine, follow these steps:

### 1. Clone the Repository

First, clone this repository to your local machine using Git:

```
git clone https://github.com/blazejkapala/learn-go.git
cd project
```


### 2. Run the Go Server

Navigate to the backend directory and start the Go server:

```
cd backend
go run main.go
```


This command starts the Go server on `http://localhost:8888`. The server handles both serving the frontend static files and API requests.

### 3. Access the Application

Open your web browser and go to `http://localhost:8888`. You should see the application's main page with a "Say Hello!" button.

### 4. Interact with the Application

Click the "Say Hello!" button to display the "Hello, World!" message. Clicking the button again will hide the message. This toggle behavior is managed by the server, demonstrating a simple interaction between HTMX on the frontend and Go on the backend.

## Project Structure

- `/backend`: Contains the Go server code.
  - `main.go`: The main server file that sets up the routes and session management.
- `/frontend`: Contains the static files served by the Go server.
  - `index.html`: The main HTML file for the frontend.
  - `htmx.min.js`: The HTMX library (if you're hosting it locally).

## Features

- **Go Backend**: A simple HTTP server that serves static files and handles API requests.
- **HTMX Frontend**: Dynamic frontend interactions without full page reloads, enhancing user experience.
- **Session Management**: Server-side toggle state management using sessions.

## Contributing

Contributions to this project are welcome! Please fork the repository and submit a pull request with your changes or improvements.

## License

This project is licensed under the [MIT License](LICENSE). See the LICENSE file in this repository for more information.
