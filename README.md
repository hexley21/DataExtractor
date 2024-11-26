# DataExtractor
A CLI that extracts fields you specify from any file you give.

## Install
Go to the [releases](https://github.com/hexley21/DataExtractor/releases) and download any version you want in accordance to your machine.

## Support
For now, only `JSON` and `YAML` files are supported

## Usage
```
datex [file] [flags]
```
- `[file]` - pass a file full path directly or just type a filename and the program will try to find it from the current dir you are in.
- `[flags]` (optional)
  - `i, --indent [int]` - Specify the indent of a resulting file, positive integers are expected, default is 4

Here are the examples of running the program:
- `datex C:\Users\diddy\Documents\DataExtractor\config.yml -i 8`
- `datex config.yml`

The resulting file with all extracted data will be placed at the current directory you are, with `extracted_` prepended to the file's name you passed.
So you get something like `extracted_config.yaml`

Type `datex -h` to see all available options and basic program info.

## Security
DataExtractor follows good practices of security, but 100% security cannot be assured.
DataExtractor is provided **"as is"** without any **warranty**. Use at your own risk.

## License
This project is licensed under the **Apache Software License 2.0**.

See [LICENSE](LICENSE) for more information.

## Acknowledgements
- [golang - go](https://github.com/golang/go) The Go programming language. (BSD-3-Clause license)
- [spf13 - cobra](https://github.com/spf13/cobra) A Commander for modern Go CLI interactions. (Apache-2.0 license)
- [charmbracelet - bubbletea](https://github.com/charmbracelet/bubbletea) A powerful little TUI framework 🏗. (MIT license)
- [charmbracelet - lipgloss](https://github.com/charmbracelet/lipgloss) Style definitions for nice terminal layouts 👄. (MIT license)
- [go-yaml - yaml](https://github.com/go-yaml/yaml) YAML support for the Go language. (MIT license)
