package internal

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateTestFile(loc *string) (*os.File, error) {
	fmt.Println("Creating file...")
	fp := filepath.Join(*loc, "test.txt")

	// Enable synchnous option to avoid caching
	f, err := os.OpenFile(fp, os.O_RDWR|os.O_SYNC, 0644)

	return f, err
}

func FillFileRandom(f *os.File, bytes int64) error {
	fmt.Println("Filling file...")

	_, err := io.CopyN(f, rand.Reader, bytes)

	return err
}
