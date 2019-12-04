package test

import (
	"os"
	"os/signal"
	"testing"
)

// CleanupAll cleans all created contexts
func CleanupAll() {
	Cleanup(contexts...)
}

// Cleanup for all given contexts
func Cleanup(contexts ...*Context) {
	for _, ctx := range contexts {
		ctx.Cleanup()
	}
}

// Cleanup iterates through the list of registered CleanupFunc functions and calls them
func (ctx *Context) Cleanup() {
	for _, f := range ctx.CleanupList {
		doCleanup(ctx.T, f)
	}
}

// AddToCleanup adds the cleanup function as the first function to the cleanup list,
// we want to delete the last thing first
func (ctx *Context) AddToCleanup(f CleanupFunc) {
	ctx.CleanupList = append([]CleanupFunc{f}, ctx.CleanupList...)
}

// CleanupOnInterrupt will execute the function cleanup if an interrupt signal is caught
func CleanupOnInterrupt(t *testing.T, cleanup CleanupFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			t.Logf("Test interrupted, cleaning up.")
			doCleanup(t, cleanup)
			os.Exit(2)
		}
	}()
}

func doCleanup(t *testing.T, cleanup CleanupFunc) {
	if !PerformCleanup(t) {
		t.Log("Skipping cleanup!")
		return
	}
	err := cleanup()
	if err != nil {
		t.Error(err)
	}
}
