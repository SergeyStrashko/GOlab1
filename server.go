package main
import (
	"fmt";
	"net/http";
	"encoding/json";
	"time";
)

func main() {
	http.HandleFunc("/time", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")
		jsonTime, err := json.Marshal(map[string]string{ "time": time.Now().Format(time.RFC3339) });

		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError);
			return;
		}

		resp.Write(jsonTime)
	})

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "text/html")
		resp.Write([]byte("Lab1"))
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8795", nil)
}