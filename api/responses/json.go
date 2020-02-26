package responses

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// JSON - for successful response
func JSON(w http.ResponseWriter, statusCode int, data interface{}){
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "%s", data )
	}
}

// ERROR for error-responses
func ERROR(w http.ResponseWriter, statusCode int, err error){
	w.WriteHeader(statusCode) 
	if err := json.NewEncoder(w).Encode( struct{ Error string `json:"error"` }{ err.Error() } ); err != nil{
		fmt.Fprintf(w, "%s", err.Error() )	
	}
}