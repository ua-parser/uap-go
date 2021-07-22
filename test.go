package main

import (
	"bufio"
	"fmt"
	"github.com/Servicewall/uap-go/uaparser"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s [old|new|both] [concurrency level]\n", os.Args[0])
		return
	}
	var wg sync.WaitGroup
	cLevel, _ := strconv.Atoi(os.Args[2])
	switch os.Args[1] {
	case "new":
		fmt.Println("Running new version of uap...")
		uaParser, _ := uaparser.NewWithOptions("./uap-core/regexes.yaml", (uaparser.EOsLookUpMode | uaparser.EUserAgentLookUpMode), 100, 20, true, true)
		for i := 0; i < cLevel; i++ {
			wg.Add(1)
			go runTest(uaParser, i, &wg)
		}
		wg.Wait()
		return
	case "old":
		fmt.Println("Running old version of uap...")
		uaParser, _ := uaparser.New("./uap-core/regexes.yaml")
		for i := 0; i < cLevel; i++ {
			wg.Add(1)
			go runTest(uaParser, i, &wg)
		}
		wg.Wait()
		return
	case "both":
		fmt.Println("Running new version of uap...")
		uaParser, _ := uaparser.NewWithOptions("./uap-core/regexes.yaml", (uaparser.EOsLookUpMode | uaparser.EUserAgentLookUpMode), 100, 20, true, true)
		for i := 0; i < cLevel; i++ {
			wg.Add(1)
			runTest(uaParser, i, &wg)
		}
		fmt.Println("Running old version of uap...")
		uaParser, _ = uaparser.New("./uap-core/regexes.yaml")
		for i := 0; i < cLevel; i++ {
			wg.Add(1)
			runTest(uaParser, i, &wg)
		}
		wg.Wait()
		return
	default:
		fmt.Printf("Usage: %s [old|new|both]\n", os.Args[0])
		return
	}
}

func runTest(uaParser *uaparser.Parser, id int, wg *sync.WaitGroup) {
	file, err := os.Open("./uas")
	if err != nil {
		fmt.Printf("Failed to open ./wrappers file. Error: %s\n", err.Error())
		return
	}
	defer file.Close()
	line := 0
	delim := ""
	for i := 0; i < id*2; i++ {
		delim += "\t"
	}
	totalLines := countLines()
	var totalTime time.Duration
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		start := time.Now()
		_ = uaParser.Parse(str)
		elapsed := time.Since(start)
		totalTime += elapsed
		line++
		fmt.Printf("\r\t\t\t\t%s%.2f%% completed|", delim, float64(line*100)/float64(totalLines))
	}
	fmt.Printf("\nProcessed lines: %d. Test took %s, average took %.3f ms\n", line, totalTime, float64(totalTime.Nanoseconds()/1e6)/(float64(line)))
	wg.Done()
}

func countLines() int {
	file, _ := os.Open("./uas")
	fileScanner := bufio.NewScanner(file)
	defer file.Close()
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}
