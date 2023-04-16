package helloworld

/*
	This api presents a simple Hello World response to the client.
	The two response formats are xml and json. The client can use
	the Accept Header in its request to specify the required response
	format.
*/

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strings"
	"time"
)

type helloworldXml struct {
	XMLName xml.Name `xml:"Response"`
	Message string   `xml:"message"`
	Time    string   `xml:"time"`
}

type helloworldJson struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	if strings.ContainsAny(r.Header.Get("Accept"), strings.ToLower("xml")) {
		rw.Header().Set("Content-Type", "application/xml")
		hw := helloworldXml{
			Message: "Hello World",
			Time:    time.Now().Format("2006-01-02-15:04:05"),
		}
		err := xml.NewEncoder(rw).Encode(hw)
		if err != nil {
			log.Print(err)
		}
	} else {
		rw.Header().Set("Content-Type", "application/json")
		hw := helloworldJson{
			Message: "Hello World",
			Time:    time.Now().Format("2006-01-02-15:04:05"),
		}
		err := json.NewEncoder(rw).Encode(hw)
		if err != nil {
			log.Print(err)
		}
	}

}
