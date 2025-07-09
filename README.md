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

#### workerpool

```sh
go run main.go -pattern=workerpool -workers=10
```

### Patterns

| Pattern	| Usage  |
|-----------|-------|
| Singleton |	Shared resources (e.g., config, DB) |
| Factory	| Object creation logic |
| Decorator | Adding functionality dynamically |
| Observer |Event-driven systems |
| Strategy | Selecting algorithms dynamically |
| Adapter | Bridging incompatible interfaces |
| Builder	| Building complex objects |
| Chain of Responsibility	| Middleware or request handlers |
| Command	| Queues, undo-redo functionality |
| Options	| Flexible object creation with functional options |
| Error Wrapper | Enhancing errors with context or stack trace |

### links

Some inspired by https://dev.to/truongpx396/common-design-patterns-in-golang-5789