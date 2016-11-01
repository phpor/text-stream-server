package main

import (
	"flag"
	"net/http"
	"github.com/phpor/text-stream-server/ring"
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	port := flag.String("port", "8083", "port to listen")
	flag.Parse()

	ring := ring.New(2048)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		keyword := r.Form.Get("k")
		count, err := strconv.ParseInt(r.Form.Get("cnt"), 10, 64 )
		if err != nil {
			count = 10
		}
		var i int = 0
		for() {
			str := ring.Get(i)
			i += 1
			if (strings.Contains(string(str), keyword)) {
				fmt.Fprintln(w, str)
			}
			if (count != 0 && int(count) > i ) {
				break
			}
		}

	})
	go func() {
		stdin := bufio.NewReaderSize(os.Stdin, 32 * 4096)
		for {
			str, _, err := stdin.ReadLine()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			ring.Set(str)
		}
	}()
	http.ListenAndServe(":" + *port, nil)

}


