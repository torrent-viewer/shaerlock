/*

Shærlock is a media file indexer.

Shærlock is a library walking through directory trees and indexing all of the
media files for which it knows the format. The entry point of the library is a
function called `Investigate`. It takes a path as its only parameter, and will
recurse through it while indexing all of the media files it finds on its way.
The function returns a list of Media files, or a critical error that happened
while indexing the files.

Additionally a shaerlock binary is provided under the package
"github.com/torrent-viewer/shaerlock/cmd/shaerlock". This program takes a list
of folders as arguments and Investigates each of these folders, printing the
found Media structures (and all of their fields) to the standard output.
*/

package shaerlock // import "github.com/torrent-viewer/shaerlock"
