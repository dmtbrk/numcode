# Numcode

Numcode is a Go module and a CLI tool for encoding and decoding English text with the following rules:
```
a -> 1
e -> 2
i -> 3
o -> 4
u -> 5
```
So that `Hello` becomes `H2ll4`.

It is in no means production-ready.

## Usage

As a tool:

```
go run ./cmd/numcode -e "Hello"
H2ll4
```
```
go run ./cmd/numcode -d "H2ll4"
Hello
```

As a package:
```
import "github.com/ortymid/numcode/pkg/numcode"
```
