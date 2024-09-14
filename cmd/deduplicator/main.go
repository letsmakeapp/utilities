package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"utilities/internal/common/hasher"
	"utilities/internal/deduplicator"
	"utilities/internal/filehasher"
	"utilities/pkg/collectionx"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: <source> <dest>")
	}

	source := os.Args[1]
	dest := os.Args[2]

	lister := deduplicator.NewStdLister()
	sit, err := lister.ListFiles(context.Background(), source)
	if err != nil {
		panic(err)
	}

	store := make(map[string]collectionx.List[deduplicator.File])
	duplicated := collectionx.NewMapSet[string]()

	hasher := filehasher.NewStdFileHasher(hasher.NewSha256Hasher())
	for file := range sit {
		f := file.Unwrap()
		if f.Type == deduplicator.FileTypeDirectory {
			continue
		}

		hash, err := hasher.Hash(context.TODO(), f.Path)
		if err != nil {
			panic(err)
		}

		encoded := hex.EncodeToString(hash)
		fmt.Println(f.Path, encoded)
		if v, found := store[encoded]; found {
			v.Append(f)
			duplicated.Add(encoded)
		} else {
			store[encoded] = collectionx.NewSliceList(f)
		}
	}

	dit, err := lister.ListFiles(context.Background(), dest)
	if err != nil {
		panic(err)
	}

	for file := range dit {
		f := file.Unwrap()
		if f.Type == deduplicator.FileTypeDirectory {
			continue
		}

		hash, err := hasher.Hash(context.TODO(), f.Path)
		if err != nil {
			panic(err)
		}

		encoded := hex.EncodeToString(hash)
		fmt.Println(f.Path, encoded)
		if v, found := store[encoded]; found {
			v.Append(f)
			duplicated.Add(encoded)
		} else {
			store[encoded] = collectionx.NewSliceList(f)
		}
	}

	for _, key := range duplicated.Elements() {
		files := store[key]
		fmt.Println("Duplicated files:")
		for _, file := range files.Iterator() {
			fmt.Printf("\t%s\n", file.Path)
		}
		fmt.Println()
	}
}
