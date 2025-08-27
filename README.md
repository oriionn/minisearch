# MiniSearch
A minimalist search engine that aims to match my personal use of a search engine. It contains different sources of results, such as [Google](https://google.com), [ArchLinux Wiki](https://wiki.archlinux.org) and many others.

## Features
- Remove unwanted sites (allocine.com, pinterest.com etc)
- Various bangs add different features
- Simple calculations
- CLI params

### Bangs
- `!google`: Redirects you to Google when you have more specific needs than those provided by Minisearch.
- `!package`: Searches for packages on the AUR and official Arch repositories.
- `!arch`: Search the Arch Wiki.
- `!wp`: Search on French Wikipedia
- `!wiki`: Search on various wikis (English Wikipedia, French Wikipedia, ArchLinux Wiki etc)

### CLI Params
- `--port <port>` / `-p <port>`: Setting the port of the web server
- `--dev`: Enter to dev mode

## Installation
### Prerequisites
- Go 1.24

### Installation
1. Clone and download modules
```sh
git clone https://github.com/oriionn/minisearch.git
cd minisearch
go mod download
```

2. Compile the project
```sh
make
```

3. Run the project
```sh
./minisearch
```

## License
This project is under a [MIT License](LICENSE)
