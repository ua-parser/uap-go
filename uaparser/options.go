package uaparser

type Option func(*Parser)

func WithMode(mode LookupMode) Option {
	return func(s *Parser) {
		s.config.Mode = mode
	}
}

func WithSort(useSort bool) Option {
	return func(s *Parser) {
		s.config.UseSort = useSort
	}
}

func WithDebug(debug bool) Option {
	return func(s *Parser) {
		s.config.DebugMode = debug
	}
}

func WithCacheSize(size int) Option {
	return func(s *Parser) {
		s.config.CacheSize = size
	}
}

func WithMissesThreshold(threshold uint64) Option {
	return func(s *Parser) {
		s.config.MissesThreshold = threshold
	}
}

func WithMatchIdxNotOk(idx int) Option {
	return func(s *Parser) {
		s.config.MatchIdxNotOk = idx
	}
}

func WithRegexesDefinitions(def RegexesDefinitions) Option {
	return func(s *Parser) {
		s.RegexesDefinitions = &def
	}
}
