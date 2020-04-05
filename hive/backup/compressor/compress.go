package compressor

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Compress(source string, target string, database string){

	sdir, err := os.Open(source)
	if err != nil {
		panic(err)
	}

	defer sdir.Close()

	files, err := sdir.Readdir(0)
	if err != nil{
		panic(err)
	}

	tarfile, err := os.Create(target)

	defer tarfile.Close()
	var fileWriter io.WriteCloser = tarfile

	if strings.HasSuffix(target, ".gz") {
		fileWriter = gzip.NewWriter(tarfile)
		defer fileWriter.Close()
	}

	tarfileWriter := tar.NewWriter(fileWriter)
	defer tarfileWriter.Close()

	for _, fileInfo := range files {

		if fileInfo.IsDir() {
			continue
		}

		file, err := os.Open(sdir.Name() + string(filepath.Separator) + fileInfo.Name())
		if err != nil{
			panic(err)
		}

		defer file.Close()

		header := new(tar.Header)
		header.Name = file.Name()
		header.Size = fileInfo.Size()
		header.Mode = int64(fileInfo.Mode())
		header.ModTime = fileInfo.ModTime()

		err = tarfileWriter.WriteHeader(header)

		if err != nil {
			panic(err)
		}

		_, err = io.Copy(tarfileWriter, file)

		if err != nil {
			panic(err)
		}
	}
}
