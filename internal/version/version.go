package version

import "fmt"

var (
	// Commit is populated during build using Go's support for ldflags. This should refer to
	// the git commit SHA1 for HEAD at the time of build.
	Commit string

	// Release is populated during build using Go's support for ldflags. This should
	// match the release tag.
	Release string
)

// FullVersionString returns a string in the format "Release (Commit)" for use in
// displaying version information
func FullVersionString() string {
	if Release == "" {
		Release = "dev"
	}
	if Commit == "" {
		Commit = "unknown"
	}
	return fmt.Sprintf("%s (%s)", Release, Commit)
}
