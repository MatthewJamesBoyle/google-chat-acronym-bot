package persist

// GetInMemoryDbForTest is a trick to expose a private var so I can test the concrete implementation.
// This should not be called outside of a _test file.
func GetInMemoryDbForTest(i InMemoryPersister) map[string]string {
	return i.mem
}
