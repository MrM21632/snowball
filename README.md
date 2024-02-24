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

First, ensure `SNOWBALL_EPOCH_MS` and `SNOWBALL_NODE_ID` are defined in the system environment variables. From there,
utilizing the module is very straightforward. Import the package normally, initialize a new server node using `InitNode()`,
then call `GenerateID()` to create and retrieve a new Snowball ID.
```go
package main

import (
    "fmt"
    "github.com/MrM21632/snowball"
)

func main() {
    node, err := snowball.InitNode()
    if err != nil {
        fmt.Println(err)
        return
    }

    id := node.GenerateID()
    // ... do what you need with the ID afterwards.
}
```

### Testing

Executing unit tests:
```bash
go test -v ./...
```

Executing benchmarks:
```bash
go test -bench=.
```
