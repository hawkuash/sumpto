## Requirements

- [Libvips](https://www.libvips.org/install.html) 8.10 or newer

If you're gonna build it yourself:
- C compatible compiler such as gcc 4.6 or clang 3.0 or newer
- Go 1.24 or newer

## Installation

1. Just follow [libvips installation guide](https://www.libvips.org/install.html)
On Windows also add next environmental variables:
- path: *path_to_libvips*\bin 
- PKG_CONFIG_PATH: *path_to_libvips*\lib\pkgconfig

2. If you got executable, do the usual stuff. Else, while in root directory of cloned repo, read [this](https://go.dev/doc/tutorial/compile-install) and run
```Go
go install
```
3. Run it

## Overview

Resize and convert as much images as you have with one command. That's basically it.
Sumpto is currently in development, so only PNG and JPEG for now. 
Other image and video format support as well as new features will be there.