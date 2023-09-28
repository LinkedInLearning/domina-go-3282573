package main

import "fmt"


func main() {
	sitioWeb := map[string]interface{}{
		"protocolo": "http",
		"hosts": "localhost",
		"puerto": "8080",
	}
	
	url := fmt.Sprintf("%v://%v:%v", sitioWeb["protocolo"], sitioWeb["hosts"], sitioWeb["puerto"])
	fmt.Println(url)
}
