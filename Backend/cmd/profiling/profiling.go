package profiling

// To generate a CPU and memory profile:
// execute 'go run main.go -cpuprofile=cpu.prof -memprofile=mem.prof'

// To show CPU and memory profile:
// install the Graphviz visualization tool. https://graphviz.org/download/
// execute 'go tool pprof -http=:8080 cpu.prof' to show CPU profile
// execute 'go tool pprof -http=:8080 mem.prof' to show memory profile

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	indexer "github.com/JuanPLoaizaC/mail_search/tree/main/Backend/pkg/Indexer"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func RunProfiling() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	indexer.NewIndexer().Run()
	end := time.Now()
	fmt.Println("execution time " + end.Sub(start).String())

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
