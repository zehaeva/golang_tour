package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(a []byte) (int, error) {
	n, err := r.r.Read(a)
	for i := 0; i < n; i++ {
		if 65 <= a[i] && a[i] < 78 {
			a[i] += 13
		} else if a[i] >= 78 && a[i] < 91 {
			a[i] -= 13
		} else if a[i] >= 97 && a[i] < 110 {
			a[i] += 13
		} else if a[i] >= 110 && a[i] < 123 {
			a[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
