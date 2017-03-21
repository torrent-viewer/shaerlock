package main

import (
	"fmt"
	"os"

	"github.com/torrent-viewer/shaerlock"
)

func printMedia(media shaerlock.Media) {
	fmt.Printf("\n%v\n\tExt: %v\n\tFullPath: %v\n", media.Name, media.Extension, media.FullPath)
}

func main() {
	var args []string
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %v DIRS...", os.Args[0])
		os.Exit(1)
	}
	args = os.Args[1:]
	for _, path := range args {
		medias, err := shaerlock.Investigate(path)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		for _, media := range medias {
			printMedia(media)
		}
	}
}
