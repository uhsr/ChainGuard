// cmd/chainguard/main.go
package main

import (
"flag"
"log"
"os"

"chainguard/internal/chainguard"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := chainguard.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
