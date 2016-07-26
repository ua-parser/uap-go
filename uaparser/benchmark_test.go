package uaparser

import (
	"log"
	"testing"
)

var benchedParser *Parser
var benchedParserWithOptions *Parser

func init() {
	var err error
	benchedParser, err = New("../uap-core/regexes.yaml")
	if err != nil {
		log.Fatal(err)
	}
	benchedParserWithOptions, err = NewWithOptions("../uap-core/regexes.yaml", (EOsLookUpMode | EUserAgentLookUpMode), 100, 20, true, true)
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ua := range uas {
			benchedParser.Parse(ua)
		}
	}
}

func BenchmarkParserWithOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ua := range uas {
			benchedParserWithOptions.Parse(ua)
		}
	}
}

var uas = []string{
	"Mozilla/5.0 (Windows NT 6.0; rv:2.0b6pre) Gecko/20100907 Firefox/4.0b6pre",
	"Mozilla/5.0 (Windows NT 5.2; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
	"Mozilla/5.0 (Windows NT 6.0; rv:2.0b6pre) Gecko/20100907 Firefox/4.0b6pre",
	"Mozilla/5.0 (Windows NT 5.2; rv:2.1.1) Gecko/ Firefox/5.0.1",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; fr; rv:1.9.2.8) Gecko/20100722 Firefox 3.6.8 GTB7.1",
	"Mozilla/5.0 (Windows NT 5.2; rv:2.0.1) Gecko Firefox/5.0.1",
	"Mozilla/5.0 (Android; Linux armv7l; rv:2.0b6pre) Gecko/20100907 Firefox/4.0b6pre Fennec/2.0b1pre",
	"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.1.8) Gecko/20100408 Thunderbird/3.0.3",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.5; rv:2.0.1) Gecko/20100101 Firefox/4.0.1 SeaMonkey/2.1.1",
	"ACS-NF	3.0	ACS-NF/3.0 NEC-c616/001.00",
	"BlackBerry7730	3.7.1	BlackBerry7730/3.7.1 UP.Link/5.1.2.5",
	"Galeon	1.3.18	Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.3) Gecko/20041020 Galeon/1.3.18",
	"Opera/9.80 (J2ME/MIDP; Opera Mini/7.1.32052/35.5706; U; id) Presto/2.8.119 Version/11.10",
}
