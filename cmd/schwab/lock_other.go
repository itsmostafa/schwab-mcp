//go:build !(darwin || dragonfly || freebsd || linux || netbsd || openbsd)

package main

// lockPath is a no-op where syscall.Flock is unavailable.
// ponytail: no Windows lock; add LockFileEx if Windows users hit races.
func lockPath(path string) (unlock func(), err error) {
	return func() {}, nil
}
