package main

import (
	"drive-speed-test/internal"
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	fileSizeMegaByte = flag.Int64("mb", 1000, "the file size for the benchmark in Megabytes")
	mode             = flag.String("mode", "both", "benchmark mode: options are 'sequential', 'random', 'both'")
	loc              = flag.String("loc", ".", "location of the file")
)

func main() {
	flag.Parse()

	f, err := internal.CreateTestFile(loc)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	internal.FillFileRandom(f, 1024*1024**fileSizeMegaByte)

	fileInfo, _ := f.Stat()
	fileSizeMegaByte := fileInfo.Size() / 1024 / 1024

	var elapsed time.Duration

	switch *mode {
	case "sequential":
		elapsed, err = internal.BenchmarkRead(f)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sequential read speed: %.2f MB/s\n", calculateSpeed(fileSizeMegaByte, elapsed))

	case "random":
		elapsed, err = internal.BenchmarkRandomRead(f)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Random read speed: %.2f MB/s\n", calculateSpeed(fileSizeMegaByte, elapsed))
	default:
		elapsed, err = internal.BenchmarkRead(f)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sequential read speed: %.2f MB/s\n", calculateSpeed(fileSizeMegaByte, elapsed))

		elapsed, err = internal.BenchmarkRandomRead(f)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Random read speed: %.2f MB/s\n", calculateSpeed(fileSizeMegaByte, elapsed))

	}
}

func calculateSpeed(fileSize int64, elapsed time.Duration) float64 {
	return float64(fileSize) / float64(elapsed.Seconds())
}
