package main

import (
    // "fmt"
    // "html/template"
		"log"
    "net/http"
    // "os"
		// "io"
		"time"
		"io"
		// "encoding/json"
)

func getSearchResult(w http.ResponseWriter, r *http.Request){

  // https://www.nseindia.com/corporates/corpInfo/equities/getBoardMeetings.jsp?Symbol=&Industry=&Period=Latest%20Announced&Purpose=&period=Latest%20Announced&symbol=&industry=&purpose=

	// Create HTTP client with timeout
	client := &http.Client{
	    Timeout: 30 * time.Second,
	}
	query := r.URL.Query().Get("query")
	request, err := http.NewRequest("GET", "https://www.nseindia.com/corporates/common/getCompanyList.jsp?query="+query, nil)
  if err != nil {
      log.Fatal(err)
  }

  // Make request
  response, err := client.Do(request)
  if err != nil {
      log.Fatal(err)
  }
	defer response.Body.Close()
	io.Copy(w, response.Body)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)
	http.HandleFunc("/search", getSearchResult)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
