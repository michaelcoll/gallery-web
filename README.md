![](doc/logo.png)

# gallery-web ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/michaelcoll/gallery-web/build) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/michaelcoll/gallery-web) ![GitHub](https://img.shields.io/github/license/michaelcoll/gallery-web)
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
```
$ ./gallery-web serve       
   ______
  /\_____\
  \ \__/_/_
  /\ \_____\  gallery web v0.0.0
  \ \/ / / /     -= serve mode =-
   \/_/\/ /
      \/_/

✓ Listening API on 0.0.0.0:8080
✓ Listening daemons on 0.0.0.0:9000
! Registering a new daemon localhost-daemon (v0.0.0) located at localhost:9001... ✓ OK
```