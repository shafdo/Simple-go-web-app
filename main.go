
package main

// Libs import
import (
	"log"
	"net/http")


func index_handler(w http.ResponseWriter, r *http.Request){
	// send any responses via "w"
	// get any requests via "r"

	http.ServeFile(w, r, "./html/home.html")
}


func main() {
	
	var	PORT string = ":8080" // Change to desired port 

	// Route
	http.HandleFunc("/", index_handler)

	log.Println("Server Started.")
	log.Println("[+] Visit http://localhost"+PORT)

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(PORT, nil))

}