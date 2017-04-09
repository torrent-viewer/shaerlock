package shaerlock

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Media stores all the information we want to associate to a media file (e.g
// its name, path, extension, and various other metadata described in their
// format's intrinsics).
type Media struct {
	Name      string
	Extension string
	FullPath  string
}

// Sleuth is the type of the callbacks used to create a Media structure from a
// file path.
type Sleuth func(path string) (Media, error)

var sleuths map[string]Sleuth

var sleuthsMutex sync.RWMutex

// DuplicationSleuthErr is the error returned when attempting to register a sleuth
// for a format when a sleuth is already registered.
var DuplicateSleuthErr = errors.New("Error registering sleuth: extension is already handled")

// NoSleuthFoundErr is the error returned when trying to examine a file whose
// extension is not found in the sleuths registry.
var NoSleuthFoundErr = errors.New("No sleuth is registered for this file extension")

func init() {
	sleuths = make(map[string]Sleuth)
	RegisterSleuth("mkv", SleuthMatroska)
}

// RegisterSleuth allows third-party to cover additional file formats without
// having to modify/fork this library. It takes a file path extension and the
// function to call when a file matching the given extension is encountered.
func RegisterSleuth(ext string, decoder Sleuth) error {
	sleuthsMutex.Lock()
	defer sleuthsMutex.Unlock()
	if _, ok := sleuths[ext]; ok {
		return DuplicateSleuthErr
	}
	sleuths[ext] = decoder
	return nil
}

// ExamineFile transforms a path into a Media structure, using the most
// appropriate registered Sleuth.
func ExamineFile(path string) (Media, error) {
	sleuthsMutex.RLock()
	defer sleuthsMutex.RUnlock()
	ext := strings.TrimPrefix(filepath.Ext(path), ".")
	if len(ext) == 0 {
		return Media{}, nil
	}
	sleuth, ok := sleuths[ext]
	if !ok {
		return Media{}, NoSleuthFoundErr
	}
	return sleuth(path)
}

// Investigate recursively walks into a directory analyzing all of its
// children and generating Media structures for files whose format is
// known to Shaerlock.
func Investigate(root string) ([]Media, error) {
	var medias []Media
	err := filepath.Walk(root, func(path string, info os.FileInfo, fileErr error) error {
		if fileErr != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		media, err := ExamineFile(path)
		if err == nil {
			medias = append(medias, media)
		}
		return nil
	})
	return medias, err
}
