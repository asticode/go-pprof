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

// Profile profiles your app
func Profile() (Closer, error) {
	// Init profiler
	var p = NewProfiler()

	// CPU
	var err error
	if *profileCPU {
		p.profilingCPU = true
		if p.fCPU, err = os.Create("./profile.cpu"); err != nil {
			return p, err
		}
		pprof.StartCPUProfile(p.fCPU)
	}

	// Memory
	if *profileMem {
		p.profilingMem = true
		if p.fMem, err = os.Create("./profile.mem"); err != nil {
			return p, err
		}
	}
	return p, nil
}

// Profiler represents a profiler
type Profiler struct {
	fCPU, fMem                 *os.File
	profilingCPU, profilingMem bool
}

// NewProfiler creates a new profiler
func NewProfiler() Profiler {
	return Profiler{}
}

// Close allows profiler to implement the Closer interface
func (p Profiler) Close() {
	// Stop profiling
	if p.profilingCPU {
		pprof.StopCPUProfile()
	}
	if p.profilingMem {
		pprof.WriteHeapProfile(p.fMem)
	}

	// Close files
	if p.fCPU != nil {
		p.fCPU.Close()
	}
	if p.fMem != nil {
		p.fMem.Close()
	}
}
