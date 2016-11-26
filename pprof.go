package pprof

import (
	"flag"
	"os"
	"runtime/pprof"
)

// Flags
var (
	profileCPU = flag.Bool("profile-cpu", false, "If yes, CPU will be profiled")
	profileMem = flag.Bool("profile-mem", false, "If yes, memory will be profiled")
)

// Closer represents an object capable of closing
type Closer interface {
	Close()
}

// closer represents the profile closer
type closer struct {
	fCPU, fMem *os.File
}

// Close allows closer to implement the Closer interface
func (c closer) Close() {
	c.fCPU.Close()
	c.fMem.Close()
	pprof.StopCPUProfile()
}

// Profile profiles your app
func Profile() (Closer, error) {
	// CPU
	var c closer
	var err error
	if *profileCPU {
		if c.fCPU, err = os.Create("./profile.cpu"); err != nil {
			return c, err
		}
		pprof.StartCPUProfile(c.fCPU)
	}

	// Memory
	if *profileMem {
		if c.fMem, err = os.Create("./profile.mem"); err != nil {
			return c, err
		}
		pprof.WriteHeapProfile(c.fMem)
	}
	return c, nil
}
