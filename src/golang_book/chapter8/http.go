package main

import (
	"io"
	"net/http"
)

// hello() will return the embedded html object to the web browser
func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res,
		`<DOCTYPE html>
		<html>
		  <head>
		    <title>Hello, World</title>
		  </head>
		  <body>
		    Hello, World! This is Nick's Go website.
		  </body>
		</html>`)
}

// main() will start a listener service for web browsers on localhost, port 9000.
// since no default file is specified, if the user goes to http://localhost:9000/ it will return a 404 error.
// If the user goes to http://localhost:9000/hello it will run the hello() function above.
// If the user goes to http://localhost:9000/assets (trailing slash is optional), it will display a listing of the root folder (current directory)
func main() {
	http.HandleFunc("/hello", hello)
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir(".")),
		),
	)
	http.ListenAndServe(":9000", nil)
}
