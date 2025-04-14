package test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/zuadi/tecamino-dbm.git/handlers"
	"github.com/zuadi/tecamino-dbm.git/models"
	"github.com/zuadi/tecamino-dbm.git/server"
	"github.com/zuadi/tecamino-dbm.git/utils"
)

func TestCreateDps(t *testing.T) {
	dmaHandler, err := handlers.NewDbmHandler(".", "test")
	if err != nil {
		t.Fatal(err)
	}

	rand.NewSource(time.Now().UnixNano())

	ndps := utils.ListofA2ZZ()
	l := len(ndps)
	s := time.Now()
	for _, dp := range ndps {
		for i := 0; i < 100; i++ {
			err = dmaHandler.ImportDatapoints(&models.Datapoint{
				Path:  fmt.Sprintf("Test:%s:%03d", dp, i),
				Type:  RandomType(),
				Value: rand.Int31(),
			})
			if err != nil {
				t.Fatal(err)
			}
		}
	}
	fmt.Printf("time used to create %d datapoints: %v\n", l*100, time.Since(s))

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Allocated: %.2f MB\n", float64(m.Alloc)/1024/1024)
	fmt.Printf("Total Allocated (ever): %.2f MB\n", float64(m.TotalAlloc)/1024/1024)
	fmt.Printf("System Memory Obtained: %.2f MB\n", float64(m.Sys)/1024/1024)
	fmt.Printf("Heap In Use: %.2f MB\n", float64(m.HeapInuse)/1024/1024)
	fmt.Printf("GC Runs: %d\n", m.NumGC)

	err = dmaHandler.SaveDb()
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuery(t *testing.T) {
	dmaHandler, err := handlers.NewDbmHandler(".", "test")
	if err != nil {
		panic(err)
	}

	// for i, o := range dmaHandler.QueryDatapoints(".*002.*") {
	// 	fmt.Println(600, i, o)
	// }

	for i, o := range dmaHandler.QueryDatapoints("Test:A:000") {
		fmt.Println(600, i, o)
	}

}

func TestUpdateDps(t *testing.T) {
	dmaHandler, err := handlers.NewDbmHandler(".", "test")
	if err != nil {
		t.Fatal(err)
	}

	rand.NewSource(time.Now().UnixNano())

	ndps := utils.ListofA2ZZ()
	l := len(ndps)
	s := time.Now()
	for j, dp := range ndps {
		if j > 2 {
			break
		}
		for i := 0; i < 100; i++ {
			err = dmaHandler.UpdateDatapointValue(fmt.Sprintf("Test:%s:%03d", dp, i), rand.Int31())
			if err != nil {
				t.Fatal(err)
			}
		}
	}
	fmt.Printf("time used to update %d datapoints: %v\n", l*100, time.Since(s))

	time.Sleep(5 * time.Second)

	fmt.Println("save data")
	err = dmaHandler.SaveDb()
	if err != nil {
		t.Fatal(err)
	}
}

func TestServer(t *testing.T) {
	fmt.Println("start")
	server := server.NewServer()
	t.Fatal(server.Serve(8100))

}
