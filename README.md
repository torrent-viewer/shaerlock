# Shærlock [![Build Status](https://travis-ci.org/torrent-viewer/shaerlock.svg?branch=master)](https://travis-ci.org/torrent-viewer/shaerlock)

## Media file tracker

Shærlock is a library walking through directory trees and indexing all of the
media files for which it knows the format. The entry point of the library is a
function called `Investigate`. It takes a path as its only parameter, and will
recurse through it while indexing all of the media files it finds on its way.
The function returns a list of Media files, or a critical error that happened
while indexing the files.

```go
package main

import (
    "fmt"
    "os"

    "github.com/torrent-viewer/shaerlock"
)

func main() {
    medias, err := shaerlock.Investigate(".")
    if err != nil {
        fmt.Fprintf(os.Stderr, "shaerlock: %v", err.Error())
        os.Exit(1)
    }
    for _, media := range medias {
        fmt.Printf("%s\n", media.Name)
    }
}
```
