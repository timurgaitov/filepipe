package main

import (
	"bufio"
	"os"
	"time"
)

func main() {
	pipe, err := os.OpenFile("pipe.txt", os.O_WRONLY, 0)
	if err != nil {
		panic(err)
	}
	defer pipe.Close()

	w := bufio.NewWriter(pipe)
	r := bufio.NewReader(os.Stdin)

	buf := make([]byte, 34*1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			panic(err)
		}
		_, err = w.Write(buf[:n])
		if err != nil {
			panic(err)
		}
		err = w.Flush()
		if err != nil {
			panic(err)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
