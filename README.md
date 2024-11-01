# GoBuilder

`GoBuilder` is a Go-based CLI tool and library that compiles Go applications for multiple operating systems and architectures with ease. It automatically reads the module name from the `go.mod` file to name the output binaries, making it a convenient solution for cross-platform builds.
![GoBuilder](/img.png)


## Features

- **Multi-platform builds**: Automatically compiles your Go application for multiple OS/architecture pairs.
- **Dynamic naming**: Uses the module name from `go.mod` to name binaries, ensuring output files are correctly labeled by target platform.
- **Single target or full build**: Supports building for all targets or a specific OS/architecture.


**Installation**:
You can install GoBuilder using the following command:
   ```bash
   go install github.com/9dl/gobuilder@latest
   ```

## Usage

To use GoBuilder, you can build your Go applications for specific operating systems and architectures. Here’s how to do it:

1. **Build for a specific target**:
   To build your Go application for a specific OS and architecture, run:
   ```bash
   gobuilder <GOOS> <GOARCH>
   ```
   Replace `<GOOS>` with your desired operating system (e.g., `linux`, `windows`, `darwin`) and `<GOARCH>` with the desired architecture (e.g., `amd64`, `386`, `arm`).

2. **Build for all targets**:
   To build your Go application for all predefined targets, simply run:
   ```bash
   gobuilder
    ```

### Example

Here’s an example of building a Go application for Linux AMD64:
```bash
gobuilder linux amd64
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.