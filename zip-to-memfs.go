package zipmemfs

import (
	"archive/zip"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"io"
)

func New(file string) (billy.Filesystem, error) {
	// Create an empty memfs
	memfs := memfs.New()

	z, err := zip.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer z.Close()

	// Read every file and store it in memfs
	for _, f := range z.File {
		err := copy(f, memfs)
		if err != nil {
			return nil, err
		}
	}

	// We are done
	return memfs, nil
}

func copy(file *zip.File, fs billy.Filesystem) error {
	handle, err := file.Open()
	if err != nil {
		return err
	}
	defer handle.Close()

	mem, err := fs.Create(file.Name)
	if err != nil {
		return err
	}
	if _, err = io.Copy(mem, handle); err != nil {
		return err
	}
	return nil
}
