package hello

import (
    "fmt"
    "net/http"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func Fibonacci(n int) int{  
	FiboResult := 0 
    if n == 1 || n == 2 {  
        FiboResult = 1 
    }else{
		FiboResult = Fibonacci(n-1) + Fibonacci(n-2) 
	} 	
    return FiboResult
}

func FiboTime(fiboIndex int) int64{
    startTime := time.Now() // get current time
    Fibonacci( fiboIndex )
    endTime := time.Now()
	var totalTime int64 = (endTime.Sub(startTime).Nanoseconds())/1000000
	return totalTime
}

type FiboData struct {
    FiboIndex          int
	TimeStamp          time.Time
	Results            int64
}

func dataStore(ctx context.Context, fiboIndex int,results int64) {
	fibodata := &FiboData{
		FiboIndex: fiboIndex,
		TimeStamp:  time.Now(),
		Results:  results,
	}

	key := datastore.NewIncompleteKey(ctx, "FiboData", nil)
	if _, err := datastore.Put(ctx, key, fibodata); err != nil {
		fmt.Println("datastore failed")
	}
}

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	fiboIndex := 40
	for flag := 1; flag <= 30; flag++ {
		totalTime := FiboTime(fiboIndex)
		dataStore(ctx, fiboIndex, totalTime)
		time.Sleep(20 * time.Second)
	}
    fmt.Fprint(w, "Hello, world!")
}
