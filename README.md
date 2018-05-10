# go-coinmarketcap

[![GoDoc](https://godoc.org/github.com/xkitpro/go-coinmarketcap?status.svg)](https://godoc.org/github.com/xkitpro/go-coinmarketcap)

## Usage

```go
import "github.com/xkitpro/go-coinmarketcap"
```

Construct a new CoinMarketCap client:
```go
cmc := coinmarketcap.NewClient()
```

### Examples

```go
package main

import cmc "github.com/xkitpro/go-coinmarketcap"

func main() {
    cli := cmc.NewClient()
    
    global, _, err := cli.GetGlobal(context.Background(), &cmc.GetGlobalOptions{
        Convert: "EUR",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(global)
}

```
For complete usage of go-coinmarketcap, see the full [package docs](https://godoc.org/github.com/xkitpro/go-coinmarketcap).
