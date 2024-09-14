package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"utilities/internal/common/hasher"
	"utilities/internal/filehasher"
)

func main() {
	if len(os.Args) != 2 {
		panic("missing path argument")
	}

	path := os.Args[1]
	sha256 := hasher.NewSha256Hasher()
	fileHasher := filehasher.NewStdFileHasher(sha256)

	checksum, err := fileHasher.Hash(context.TODO(), path)
	if err != nil {
		panic(err)
	}

	checkSumHex := hex.EncodeToString(checksum)
	fmt.Println("Lowers: ", checkSumHex)
	fmt.Println("Uppers: ", strings.ToUpper(checkSumHex))
}
