package filesplitter

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func SplitFile(reader io.Reader, filepath string, size int64) (err error) {
	for i := 0; ; i++ {
		var f *os.File
		f, err = os.Create(fmt.Sprintf("%s.part%d", filepath, i))
		if err != nil {
			return err
		}
		_, err = io.CopyN(f, reader, size)
		if err != nil {
			break
		}
	}
	if !errors.Is(err, io.EOF) {
		return
	}
	err = nil

	return
}

func JoinFile(path string, w io.Writer) (err error) {
	basePath, ok := strings.CutSuffix(path, ".part0")
	if !ok {
		err = fmt.Errorf("file name must end with part.0. Actual filename %s", path)
		return
	}
	for i := 0; ; i++ {
		f, err := os.Open(fmt.Sprintf("%s.part%d", basePath, i))
		if os.IsNotExist(err) {
			if i != 0 {
				err = nil
			}
			break
		}
		if err != nil {
			return err
		}
		io.Copy(w, f)
	}
	return
}
