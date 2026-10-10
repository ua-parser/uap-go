package uaparser

import (
	"sync"
	"testing"
)

// TestNewConcurrent_NoRaceOnDefaultRegexDefinitions reproduces the data race
// reported in issue #100. Without the fix, multiple goroutines calling New()
// with the default cached RegexDefinitions racing on the shared []*uaParser
// pointers triggers a race detector failure under -race.
func TestNewConcurrent_NoRaceOnDefaultRegexDefinitions(t *testing.T) {
	const n = 16
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			if _, err := New(); err != nil {
				t.Errorf("New: %v", err)
			}
		}()
	}
	wg.Wait()
}

// TestConcurrentNewWithRegexDefinitions verifies that concurrent New() calls
// passing a shared RegexDefinitions instance are safely isolated and do not race.
func TestConcurrentNewWithRegexDefinitions(t *testing.T) {
	def := defaultRegexesDefinitions()
	const n = 16
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			p, err := New(WithRegexDefinitions(def))
			if err != nil {
				t.Errorf("New: %v", err)
				return
			}
			cli := p.Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
			if cli == nil || cli.UserAgent == nil || cli.UserAgent.Family != "Chrome" {
				t.Errorf("unexpected parse result: %v", cli)
			}
		}()
	}
	wg.Wait()
}

// TestConcurrentParse_SharedParser verifies that multiple goroutines concurrently
// parsing user agent strings using a single shared Parser operate safely without data races.
func TestConcurrentParse_SharedParser(t *testing.T) {
	parser, err := New()
	if err != nil {
		t.Fatalf("failed to create parser: %v", err)
	}

	testCases := []struct {
		ua             string
		expectedFamily string
		expectedOS     string
	}{
		{
			ua:             "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
			expectedFamily: "Chrome",
			expectedOS:     "Windows",
		},
		{
			ua:             "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15",
			expectedFamily: "Safari",
			expectedOS:     "Mac OS X",
		},
		{
			ua:             "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/119.0",
			expectedFamily: "Firefox",
			expectedOS:     "Linux",
		},
		{
			ua:             "Mozilla/5.0 (iPhone; CPU iPhone OS 17_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
			expectedFamily: "Mobile Safari UI/WKWebView",
			expectedOS:     "iOS",
		},
		{
			ua:             "UnknownCustomBot/1.0",
			expectedFamily: "UnknownCustomBot",
			expectedOS:     "Other",
		},
	}

	const goroutines = 16
	const iterations = 50

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(gid int) {
			defer wg.Done()
			for iter := 0; iter < iterations; iter++ {
				tc := testCases[(gid+iter)%len(testCases)]
				cli := parser.Parse(tc.ua)
				if cli == nil {
					t.Errorf("nil client returned for UA %q", tc.ua)
					continue
				}
				if cli.UserAgent == nil || cli.UserAgent.Family != tc.expectedFamily {
					t.Errorf("UA family mismatch: got %v, want %q", cli.UserAgent, tc.expectedFamily)
				}
				if cli.Os == nil || cli.Os.Family != tc.expectedOS {
					t.Errorf("OS family mismatch: got %v, want %q", cli.Os, tc.expectedOS)
				}
				if cli.Device == nil {
					t.Errorf("nil device for UA %q", tc.ua)
				}
			}
		}(g)
	}

	wg.Wait()
}

// TestConcurrentSubParsers_SharedParser verifies that ParseUserAgent, ParseOs,
// and ParseDevice can be called concurrently across goroutines on a shared Parser.
func TestConcurrentSubParsers_SharedParser(t *testing.T) {
	parser, err := New()
	if err != nil {
		t.Fatalf("failed to create parser: %v", err)
	}

	sampleUA := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"

	const goroutines = 16
	const iterations = 50

	var wg sync.WaitGroup
	wg.Add(goroutines * 3)

	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				ua := parser.ParseUserAgent(sampleUA)
				if ua == nil || ua.Family != "Chrome" {
					t.Errorf("unexpected UA: %v", ua)
				}
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				os := parser.ParseOs(sampleUA)
				if os == nil || os.Family != "Windows" {
					t.Errorf("unexpected OS: %v", os)
				}
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				dvc := parser.ParseDevice(sampleUA)
				if dvc == nil {
					t.Errorf("nil device returned")
				}
			}
		}()
	}

	wg.Wait()
}

// TestConcurrentParseWithSort_SharedParser tests that concurrent parsing does not race
// with periodic sorting triggered by checkAndSort when WithSort is enabled.
func TestConcurrentParseWithSort_SharedParser(t *testing.T) {
	parser, err := New(
		WithSort(true),
		WithCacheSize(32),
		WithMatchIdxNotOk(0),
	)
	if err != nil {
		t.Fatalf("failed to create parser: %v", err)
	}

	// Lower threshold to trigger checkAndSort repeatedly during test execution.
	parser.config.MissesThreshold = 5

	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15",
		"Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/119.0",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 17_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
		"CustomCrawler/2.0 (Testing bot)",
		"YetAnotherBot/3.0",
	}

	const goroutines = 16
	const iterations = 50

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(gid int) {
			defer wg.Done()
			for iter := 0; iter < iterations; iter++ {
				ua := userAgents[(gid+iter)%len(userAgents)]
				cli := parser.Parse(ua)
				if cli == nil || cli.UserAgent == nil || cli.Os == nil || cli.Device == nil {
					t.Errorf("incomplete client parsed for %q", ua)
				}
			}
		}(g)
	}

	wg.Wait()
}
