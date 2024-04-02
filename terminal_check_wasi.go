//go:build wasi
// +build wasi

package log

func isTerminal(fd int) bool {
	return false
}
