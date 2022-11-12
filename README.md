<!--suppress HtmlDeprecatedAttribute -->
<p align="center">
    <img src="https://raw.githubusercontent.com/michaelcoll/gallery-web/main/doc/logo.png" alt="Logo" /><br />
</p>
<p align="center">
    A simple decentralized photo gallery
</p>
<p align="center">
    <img src="https://img.shields.io/github/workflow/status/michaelcoll/gallery-web/build" alt="Logo" />
    <img src="https://img.shields.io/github/v/release/michaelcoll/gallery-web" alt="Logo" />
    <img src="https://img.shields.io/github/license/michaelcoll/gallery-web" alt="Logo" />
</p>

# gallery-web
The main web app of the gallery app

## Usage
```
$ ./gallery-web serve --help
Starts the serve

Usage:
  gallery-web serve [flags]

Flags:
  -h, --help   help for serve

Global Flags:
      --verbose   Verbose display
```

## Exemple
![](doc/gallery-web.webp)

## Docker image usage
```
docker run -ti --rm \
-p 9000:9000 \
-p 8080:8080 \
ghcr.io/michaelcoll/gallery-web:latest
```

## Build
### Requirements

- Node >= 18.10
- Go >= 1.19
- make command

### Commands
```
make prepare
```
```
make
```

## Run from code
```
make run
```