# go-flavor-parser

See the [Main project github page](https://github.com/bunniesandbeatings/go-flavor)

## Build

**Must use go 1.5+**

```
export GO15VENDOREXPERIMENT=1
go build github.com/bunniesandbeatings/go-flavor-parser

```

## Test

**Must use go 1.5+**

```
export GO15VENDOREXPERIMENT=1
go test github.com/bunniesandbeatings/go-flavor-parser/...
```

## Run

Binary is called `baduk`

`./scripts/sample` has sample usage (assuming the usual $GOPATH and $PATH settings).

## Contributing

Notes about working on this poject:
  * Uses GO15VENDOREXPERIMENT and [govendor](https://github.com/kardianos/govendor) to vendor libs.

## Assumptions

  * Code is buildable
  * Golang enforces no cycles in imports
