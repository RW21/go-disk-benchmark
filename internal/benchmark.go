package internal

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const BlockSize = 1024

func BenchmarkRead(file *os.File) (time.Duration, error) {
	data := make([]byte, BlockSize)

	start := time.Now()
	for {
		_, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return -1, fmt.Errorf("error reading file: %v", err)
			}
		}
	}
	elapsed := time.Since(start)

	return elapsed, nil
}

func BenchmarkRandomRead(file *os.File) (time.Duration, error) {
	data := make([]byte, BlockSize)
	stat, err := file.Stat()
	if err != nil {
		return -1, fmt.Errorf("error getting file stats: %v", err)
	}

	fileSize := stat.Size()
	numBlocks := fileSize / int64(BlockSize)

	start := time.Now()

	for i := 0; i < 10000; i++ {
		randomBlock := rand.Int63n(numBlocks)
		offset := randomBlock * int64(BlockSize)

		_, err := file.Seek(offset, 0)
		if err != nil {
			return -1, fmt.Errorf("error seeking file: %v", err)
		}

		_, err = file.Read(data)
		if err != nil && err != io.EOF {
			return -1, fmt.Errorf("error reading file: %v", err)
		}
	}

	elapsed := time.Since(start)

	return elapsed, nil
}
