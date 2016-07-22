# Golang User Agent parser

This is the Go implementation of the [ua-parser](https://github.com/tobie/ua-parser)

# Usage

## Install

    $ go get github.com/ua-parser/uap-go/uaparser

## Updating

  `uap-core` definitions are, by default, compiled and included in the Go portion of this package. To regenerate these definitions you can run the `build.sh` script.

## Testing

    $ cd uaparser
    $ go test -v -cover

When adding new feature you can check for data race with `go test -race`, although this command
will take more than 2 minutes to execute since the `-race` flag adds some instrumentation on all regex matches

So a quicker way to test data race on the main `*parser.Parse` method can be:

    $ go test -race -run=Concurrency  # filter to execute `TestGenericParseMethodConcurrency` only

## Benching

If needed, you can run benchmark on your latest feature to be compared (using `benchcmp`) against the current performance

    $ cd uaparser
    $ git checkout master
    $ go test -bench=.  -run=none > ~/old.benchmark
    $ git checkout my_latest_feature
    $ go test -bench=.  -run=none > ~/new.benchmark
    $ benchcmp ~/old.benchmark ~/new.benchmark

# Example

```go
package main

import (
  "fmt"
  "log"

  "github.com/ua-parser/uap-go/uaparser"
)

func main() {
  uagent := "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us; Silk/1.1.0-80) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16 Silk-Accelerated=true"

  parser, err := uaparser.New("./regexes.yaml")
  if err != nil {
    log.Fatal(err)
  }

  client := parser.Parse(uagent)

  fmt.Println(client.UserAgent.Family)  // "Amazon Silk"
  fmt.Println(client.UserAgent.Major)   // "1"
  fmt.Println(client.UserAgent.Minor)   // "1"
  fmt.Println(client.UserAgent.Patch)   // "0-80"
  fmt.Println(client.Os.Family)         // "Android"
  fmt.Println(client.Os.Major)          // ""
  fmt.Println(client.Os.Minor)          // ""
  fmt.Println(client.Os.Patch)          // ""
  fmt.Println(client.Os.PatchMinor)     // ""
  fmt.Println(client.Device.Family)     // "Kindle Fire"
}

```

# Authors

* Yihuan Zhou

(Based on the Java implementation by Steve Jiang and using agent data from BrowserScope)
