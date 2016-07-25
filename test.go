package main

import (
	"bufio"
	"fmt"
	"github.com/streamrail/uap-go/uaparser"
	"os"
	"strconv"
	"time"
	"sync"
	netUtil "github.com/streamrail/common/util/net"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s [old|new|both] [concurrency level]\n", os.Args[0])
		return
	}
	var wg sync.WaitGroup
	cLevel, _ := strconv.Atoi(os.Args[2])
	switch (os.Args[1]) {
		case "new":
			fmt.Printf("Running new version of uap...")
			uaParser, _ := uaparser.NewWithOptions("/etc/regexes.yaml", (uaparser.EOsLookUpMode | uaparser.EUserAgentLookUpMode), 100, 20, true, true)
			for i := 0; i < cLevel; i++ {
				wg.Add(1)
				go runTest(uaParser, i, &wg)
			}
			wg.Wait()
			return
		case "old":
			fmt.Printf("Running old version of uap...")
			uaParser, _ := uaparser.New("/etc/regexes.yaml")
			for i := 0; i < cLevel; i++ {
				wg.Add(1)
				go runTest(uaParser, i, &wg)
			}
			wg.Wait()
			return
		case "both":
			fmt.Printf("Running new version of uap...")
			uaParser, _ := uaparser.NewWithOptions("/etc/regexes.yaml", (uaparser.EOsLookUpMode | uaparser.EUserAgentLookUpMode), 100, 20, true, true)
			for i := 0; i < cLevel; i++ {
				wg.Add(1)
				runTest(uaParser, i, &wg)
			}
			fmt.Printf("Running old version of uap...")
			uaParser, _ = uaparser.New("/etc/regexes.yaml")
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
	for i:= 0; i < id * 2; i++ {
		delim += "\t"
	}
	totalLines := countLines()
	platforms := map[string]int{"mobile": 0, "desktop": 0}
	var totalTime time.Duration
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		line++
//		if line == 5000 {
//			break
//		}
		start := time.Now()
		if netUtil.IsMobileUA(str, uaParser.Parse(str)) {
			platforms["mobile"]++
		} else {
			platforms["desktop"]++
		}
		elapsed := time.Since(start)
		totalTime += elapsed
		fmt.Printf("\r\t\t\t\t%s%.2f%% completed|", delim, float64(line * 100)/float64(totalLines))
	}
	fmt.Printf("\nProcessed lines: %d. Test took %s\nResult: %+v\n", line, totalTime, platforms)
	wg.Done()
}

func countLines() (int) {
	file, _ := os.Open("./uas")
	fileScanner := bufio.NewScanner(file)
	defer file.Close()
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
//		if lineCount == 5000 {
//			break
//		}
	}
	return lineCount
}
