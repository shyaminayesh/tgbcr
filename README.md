# tgbcr ( Telegram Bot Command Router )

This module is providing capability to handle and route request to appropriate functions by scanning the Telegram bot update messages. It's just scan the text and then provide implementation to forward the request to predefined handler function. This can be also used to any other purpose that need this capability.


## Usage

```go
package main

import (
    "github.com/shyaminayesh/tgbcr"
)

func Hello(message) {
    fmt.Println(message)
}

func main() {
    Router := tgbcr.New()
    Router.Handle("/hello", Hello)
    Router.Dispatch()
}
```