package main

import (
	"fmt"
	"io"
	"os"
)

type LimitedReader struct {
	Reader io.Reader
	Limit int64
}

func (r *LimitedReader) Read(b []byte) (n int, err error) {
	if r.Limit <= 0 {
		return 0, io.EOF
	}

	if len(b) > int(r.Limit) {
		b = b[:r.Limit]
	}

	n, err = r.Reader.Read(b)
	r.Limit -= int64(n)
	return
}


func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{
		Reader: r,
		Limit: n,
	}
}


func main() {
	file, err := os.Open("lr")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lr := LimitReader(file, 5)
	buf := make([]byte, 10)
	n, err := lr.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, buf, string(buf))
}
