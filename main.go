package main

import (
	dp "dsa/DP"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(dp.BestSumRec(9, []int{3, 2, 4}))

	log.Println("Execution Time: ", time.Since(start))
}
