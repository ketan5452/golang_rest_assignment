package main

// Required imports
import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
)

// Defination of structure of Response
type ResponseStatus struct {
    Status int `json:"Status"`
    Response string `json:"Response"`
}

// Function to show some message on homePage
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnHello(w http.ResponseWriter, r *http.Request){
	
	// Get the query parameter from URL
	param := r.URL.Query().Get("q")

	// Check query parameter value is empty or not
	if param != "" {
		// Check query parameter value is 'hi' or not
		if param == "hi" {
			// Create a valid Response
			returnResponse := ResponseStatus{Status: 200, Response: "Hello"}
			
			// Return a valid Response
			json.NewEncoder(w).Encode(returnResponse)
		} else {
			// Create not valid Response
			returnResponse := ResponseStatus{Status: 400, Response: "Not Found"}
			
			// Return not valid Response
			json.NewEncoder(w).Encode(returnResponse)
		}
	} else {
		// Create not valid Response
		returnResponse := ResponseStatus{Status: 400, Response: "Not Found"}
			
		// Return not valid Response
		json.NewEncoder(w).Encode(returnResponse)
	}
    fmt.Println("Endpoint Hit: returnHello")
}

// Function which handles a incoming request
func handleRequests() {
	
	// If URL match calls to returnHello Function
	http.HandleFunc("/golang", returnHello)
	
    log.Fatal(http.ListenAndServe(":8081", nil))
}

// golang main function
func main() {
	// Call to function which handles a incoming request
    handleRequests()
}