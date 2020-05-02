/*Using net/http module: 

To run:
go build loop54.go
./loop54  <url1> <url2> <url3>
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		/*http.Get function makes an HTTP request
                and if there is no error returns the result in
                struct resp. The Body field contains server response
                as readeable stream*/
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
			os.Exit(1)
		}
		/*ioutil reads the entire response , the result is stored in b */
		b, err := ioutil.ReadAll(resp.Body)
		/* we are using defer to make sure that in case any errors we still close
                the Body stream*/
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading%s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
		

}
