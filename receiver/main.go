package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	pipe, err := os.OpenFile("out.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer pipe.Close()

	r := bufio.NewReader(pipe)
	buf := make([]byte, 34*1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				time.Sleep(50 * time.Millisecond)
				continue
			}
			panic(err)
		}
		fmt.Print(string(buf[:n]))
	}
}
