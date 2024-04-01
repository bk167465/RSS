# RSS in Go (Really Simple Syndication)

• Steps to run this project on localhost: <br />
1. Clone this repository.
2. Make sure you have installed Go and other packages
   - Use Go get <URL> command
   - go mod vendor
   - go mod tidy
4. Change current directory to sql/schema and run up migration command
   - goose postgres URL up
5. Run sqlc genrate command to genrate models file in internal directory
6. Start backend: go build && ./rss
