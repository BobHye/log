//go:build wasip1
// +build wasip1

package log

func isTerminal(fd int) bool {
	return false
}
