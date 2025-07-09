## chango

This repository is for practicing several design patterns in golang that include channel and concurrency.

The name of the package is the name of the pattern. 

All patterns will be available as a command arg. To execute a pattern:

```sh
go run main.go <pattern_name>
```

Some packages will be used as a helper packages and won't have a dedicated arg for execution (logging, error wrap, debug etc.)

No fancy imports - only standard libraries. The repository will include only simple and stupid functionalities and concentrate mainly in the design pattern itself.

### Executions

Execute command by its pattern name

#### hello

```sh
go run main.go -pattern=hello
```

#### messaging

```sh
go run main.go -pattern=messaging -duration=10s
```

#### observer

```sh
go run main.go -pattern=observer -name="etzba/etz" -tag="development" -sha="sha256:111111111111111111111111111111111111"
```

### links

Some inspired by https://dev.to/truongpx396/common-design-patterns-in-golang-5789