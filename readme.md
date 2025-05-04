# lazyPm

![Made for Lazy Devs](https://img.shields.io/badge/made%20for-lazy%20devs-yellow?style=flat-square&logo=go)

**Too lazy to remember whether it's `npm run`, `go run`, or `yarn something`?**  
**lazyPm** is your new best friend — a friendly CLI wrapper for package managers, so you can spend more time sipping coffee and less time memorizing commands.

---

## Overview

**lazyPm** is a CLI-agnostic wrapper for package managers, designed to simplify and streamline your development workflow. Whether you're managing Go modules, npm packages, or other dependencies, lazyPm provides a unified interface to handle common tasks like listing scripts, installing packages, and running commands.

## About me

Just a heads-up — **I'm primarily a TypeScript developer**, and **this is my first project in Go**! 😅  
I’m still learning the ins and outs of Go, so if you spot any weird practices or quirks in the code, feel free to point them out!  
Please be kind and indulgent — I’m here to learn and improve. 🚀

If you're a Go expert, your feedback is **more than welcome**. Let’s make lazyPm even better, together!

## Inspiration

lazyPm draws inspiration from [lazynpm](https://github.com/jesseduffield/lazynpm), a terminal UI for npm commands. While lazyPm currently operates as a command-line interface (CLI), there's a vision to evolve it into a full-fledged terminal user interface (TUI) in the future.

## Features

- **Package Manager Detection**: Automatically detects the package manager used in your project directory.
- **Unified Commands**: Use consistent commands across different package managers.
- **Extensible**: Easily add support for more package managers as needed.

## Installation

### From Source

1. Clone the repository:

```bash
git clone https://github.com/Tatsuyasan/lazyPm.git
cd lazyPm
```

2. Build the application:

```bash
go build -o lpm
```

3. Optionally, move the binary to a directory in your PATH:

```bash
mv lpm /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/Tatsuyasan/lazyPm@latest
```

## Usage

lazyPm provides several commands to interact with your project's package manager:

### `lpm list scripts`

Lists all available scripts defined in your project's package manager configuration.

```bash
lpm list scripts
```

### `lpm list deps`

Lists all dependencies of your project.

```bash
lpm list deps
```

### `lpm install`

Installs the dependencies for your project.

```bash
lpm install
```

### `lpm run <script>`

Runs a specified script defined in your project's package manager configuration.

```bash
lpm run build
```

## Future Plans

There's an ongoing effort to develop a terminal user interface (TUI) for lazyPm, inspired by the [lazynpm project](https://github.com/jesseduffield/lazynpm). This TUI will provide an interactive and user-friendly interface to manage your project's package manager tasks more efficiently.

## Contributing

Contributions are welcome! If you have suggestions, improvements, or bug fixes, please open an issue or submit a pull request.

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

---

_Made with 💻 + ☕ by lazy devs, for lazy devs._

## Contributing

Contributions are welcome! If you have suggestions, improvements, or bug fixes, please open an issue or submit a pull request.

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

---

_Made with 💻 + ☕ by lazy devs, for lazy devs._
