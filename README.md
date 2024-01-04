# go-webserver

Example websever written in Go. This project uses the echo web library to server a public html. The code has a single point of entry to handle web requests and dynmically parses json data to structs and passes it to the handling function using reflection.

## Usage

1. download source
2. open the main.go file, modify the CONFIG_WEBSERVER_PORT as needed
3. In terminal: go run . 
4. Open your browser to http://localhost:8004  (your port number)
5. To test a post back: http://localhost:8004/public/index.html

