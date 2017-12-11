package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	fin, _ := os.Open("./ioteerdr.go")
	defer fin.Close()
	fout, _ := os.Create("./teereader.gz")
	defer fout.Close()
	zip := gzip.NewWriter(fout)
	defer zip.Close()
	sha := sha1.New()
	data := io.TeeReader(fin, sha)
	io.Copy(zip, data)
	fmt.Printf("SHA1 hash %x\n", sha.Sum(nil))
}
