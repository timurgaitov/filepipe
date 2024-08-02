package main

import (
	"bufio"
	"io"
	"os"
	"time"
)

func main() {
	in, err := os.OpenFile("in.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.OpenFile("out.txt", os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
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
		_, err = w.Write(buf[:n])
		if err != nil {
			panic(err)
		}
		err = w.Flush()
		if err != nil {
			panic(err)
		}
	}
}
