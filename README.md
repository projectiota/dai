# Dai
Dai is simple imperative build system for C projects written in Go.

Dai is:
- Imperative - it demands certain directory structure for the sake of simplicity
- C only - other languages are not supported. Cross-compilation is supported.
- Linux only - you are developer, right?

### Documentation
Dai is similar to `go` tool in the sense that it inposes certain directory structure, for several reasons:
- Good practice is not an option - ti is mandatory ("There is only one way to do it")
- Keep the tool simple

Dai thus demands that all:
- Source files are kept in `src` directory
- Header files are kept in `inc` directory
- Libraries are kept in the `lib` directory

If this is respected then all is needed is to run `dai` and it will will run with default configuration in the current directory.

Other wise you can tell point it to your project directory:
```bash
dai <path_to_project>
```

Dai takes a config from the special file called `dai.toml` placed in the root of your project directory.

Here is the example of `dait.toml`:
```toml
### Target
target = "hello"

### Toolchain
[toolchain]

prefix = ""

# Compiler
compiler = "gcc"
cflags = [""]

# Linker
linker = "ld"
ldflags = [""]
```

### Author
Dai is engineered by [@drasko](https://github.com/drasko)

### License
[Apache License, version 2.0](LICENSE)
