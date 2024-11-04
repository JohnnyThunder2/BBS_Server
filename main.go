package main

import (
	"fmt"
	"net/http"
	"html/template"
)

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
//					startPage Function 
//		   		handels default request pipleine
//
//
// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

func startPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		var fileName = "index.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("error when parsing file", err)
			return
		}
		t.ExecuteTemplate(w, fileName, nil)
		
		
	case "/page2":
	var fileName = "page2.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("error when parsing file", err)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
	
		case "/page3":
	var fileName = "page3.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("error when parsing file", err)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
	
	//case "*.jpg":
		//return http.FileServer(http.Dir("/")

	case "/johnny":
		fmt.Fprint(w, "Yo, Johnny!")
	default:
		fmt.Fprint(w, "Error 404... file not found.")

	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
//					main Funciton
//
//
// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

func main() {
	
	//This makes sure images will load properly when called by an HTML file
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.HandleFunc("/", startPage)
	
	// HTTP Server Start-up
	//fmt.Println("Starting server at port 80")
	//if err := http.ListenAndServe(":80", nil); err != nil {
		//fmt.Println(err)}

    // HTTPS Server Start-up
    fmt.Println("Starting secure server at port 443")
	if err := http.ListenAndServeTLS(":443","cert.pem","key.pem", nil); err != nil {
		fmt.Println(err)}
	
	// I think this code works to redirect http calls, but it also might be useless...
	if err := http.ListenAndServe(":80", http.HandlerFunc(redirectToTls)); err != nil {
		fmt.Println("ListenAndServe error: %v", err)}
	
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
// http to https redirect function
func redirectToTls(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://:443"+r.RequestURI, http.StatusMovedPermanently)
}
