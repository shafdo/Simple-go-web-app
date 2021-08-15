
package main

// Libs import
import ("net/http")


func index_handler(w http.ResponseWriter, r *http.Request){
	// send any responses via "w"
	// get any requests via "r"

	http.ServeFile(w, r, "./html/home.html")
}



func main() {

	// Route
	http.HandleFunc("/", index_handler)


	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)

}