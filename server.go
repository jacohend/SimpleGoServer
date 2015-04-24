/* Jacob Henderson 
 * 
 * Instructions: 
 * sudo apt-get install golang-go
 * go build <filename.go>
 * ./filename
 * open a browser and navigate to localhost:8080/example/time
 *
 */

package main

import (
    "fmt"
    "net/http"
    "time"
    "io/ioutil"
)

//length of our site directory. 
//we use this to know how many characters to go forward in the URL string until we hit the web directory name
const lenPath = len("/example/")

//global variables
var retrieval_time string

//HTTP response handler. This gets called when someone visits the page 
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]   //get the web directory name from the URL 
    if title == string("time"){   //if the web directory in the url is "time"
        body := "The Time is"        //assign string variable
        retrieval_time = time.Now().Format("Jan 2, 2006 at 3:04pm (MST)")  //fill an example time format with current time
        fmt.Fprintf(w, "<h1><b>%s</b></h1><div>%s</div>", retrieval_time, body)  //return an http string to user's browser
    }else{
        contents,_ := ioutil.ReadFile(title)
        result := string(contents)
        fmt.Fprintf(w, "%s", result)
    }
}

func main() {
    http.HandleFunc("/example/", viewHandler)   //viewHandler gets called when someone visits our page
    http.ListenAndServe(":8080", nil)           //starts web server on port 8080
}
