# snowball

Snowball is a variant implementation of Twitter's [Snowflake](https://en.wikipedia.org/wiki/Snowflake_ID) unique ID
generation algorithm. This implementation specifically has some modifications to allow for both increased scalability
over time, and an extended period of valid IDs generated based on a given epoch. The implementation is also intentionally
made configurable for, say, production deployments to a cluster/data center.

This Golang module provides the following functionality:
- A basic UID generator using the "Snowball" algorithm
- Methods to parse Snowball IDs
- Methods to encode/decode Snowball IDs in binary, Base16 (hex), Base32, Base62, and Base64

## Snowball ID Structure

The ID format used by Snowball is very similar to Twitter's original design for Snowflake IDs. However, there are some
key changes made to provide the following benefits:
- Providing a larger time range where valid IDs can be generated
- Providing a larger range for server IDs

Snowball's ID format is as follows:
- The ID fits fully within an unsigned 64-bit integer (uint64)
- 42 bits, including the "sign" bit, are used to store a millisecond-precision timestamp, based on a custom epoch provided
  as an environment variable
- 11 bits are used to store a server ID, doubling the number of valid server IDs in a Snowflake ID (2048 vs 1024)
- 11 bits are used to store a sequence number, halving the possible number vs. Snowflake (2048 vs 4096)

This format carries some downsides, namely being able to generate fewer unique IDs in a given millisecond, but nonetheless
is very performant in a distributed environment. Some estimates for this performance are:
- Up to 2048 individual servers that can produce IDs independently
- Each server can produce up to 2048 IDs per millisecond, meaning all servers can produce up to approx. 4.2 million IDs
  in a given millisecond, or over 4 billion IDs per second
- The 42-bit wide timestamp section allows for unique ID generation well into the future - IDs are guaranteed to be unique
  over 139 years following the pre-set epoch

## Getting Started

### Installation

Installation is as simple as running:
```bash
go get github.com/MrM21632/snowball
```

And importing the module into your code is also simple:
```go
import (
	"github.com/MrM21632/snowball"
)
```

### Usage

Snowball handles Server IDs in one of two ways, depending on how you configure it:
1. Developers can pre-assign IDs to each server the service will run on. If so, you will want to define the variable
   `SNOWBALL_NODE_ID` in your system environment.
2. (Development ongoing) Snowball can also derive server IDs from the IP address of the machine, pod, container, etc. that
   it's running on. In this case, you'll need to define the variable `SERVER_IP_ADDRESS` with the machine IP address in
   your system environment.

In addition, you'll also need to provide the variable `SNOWBALL_EPOCH_MS` in your system environment to define what epoch
Snowball uses. By default, if none is provided, it will use the same epoch Twitter uses for Snowflake - 1288834974657 (i.e.,
Thursday, November 4, 2010 01:42:54 UTC).

From here, simply import the package as normal, then initialize the server, and you're ready to start generating IDs!
```go
package main

import (
    "fmt"
    "github.com/MrM21632/snowball"
)

func main() {
    // Set to false if you're providing pre-assigned IDs to servers
    node, err := snowball.InitNode(true)
    if err != nil {
        fmt.Println(err)
        return
    }

    id := node.GenerateID()
    // ... do what you need with the ID afterwards.
}
```

Snowball also comes with a default executable service, leveraging [Gin](https://pkg.go.dev/github.com/gin-gonic/gin) for
HTTP requests and [Prometheus](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus) for metrics collection.

In addition, there are also custom Docker images available for both amd64 and arm64 platforms. Simply run
```bash
docker pull mrm21632/snowball:v1.2.0  # or whatever the latest version happens to be
```

and deploy to your heart's content.

### Testing

Executing unit tests:
```bash
go test -v ./...
```

Executing benchmarks:
```bash
go test -bench=.
```
