package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Testing KOS cli ...")
	fmt.Println("KSPSCRIPT:", os.Getenv("KSPSCRIPT"))
	fmt.Println("KSPSRC:", os.Getenv("KSPSRC"))
	fmt.Println("NOTDEFINED:", os.Getenv("AASDFSADF"))
}
