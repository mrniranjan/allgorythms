/* use io.Copy instead of ioutil.ReadAll from loop54.go 
./loop54.go <url1> <url2> <url3>
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
		} else {
			_ , error := io.Copy(os.Stdout, resp.Body)
			if error != nil {
				fmt.Fprintf(os.Stderr, "Error Copying %s: %v\n", url, error)
			}
		}

	}
}
