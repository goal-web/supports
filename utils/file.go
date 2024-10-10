package utils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ExistsPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// FileExists 文件是否存在
func FileExists(filePath string) bool {
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false // 文件不存在
	}
	return err == nil && !stat.IsDir() // 文件存在
}

// CopyFile 复制一个文件
func CopyFile(from, to string, bufferSize int64) error {
	sourceFileStat, statRrr := os.Stat(from)
	if statRrr != nil {
		return statRrr
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", from)
	}

	source, openErr := os.Open(from)
	if openErr != nil {
		return openErr
	}
	defer source.Close()

	_, statRrr = os.Stat(to)
	if statRrr == nil {
		return fmt.Errorf("file %s already exists", to)
	}

	destination, createErr := os.Create(to)
	if createErr != nil {
		return createErr
	}
	defer destination.Close()

	buf := make([]byte, bufferSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err = destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

// AllFiles 获取一个目录下的所有文件
func AllFiles(path string) (results []fs.FileInfo) {
	_ = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			results = append(results, info)
		}
		return nil
	})
	return results
}

// AllDirectories 获取一个目录下的所有目录
func AllDirectories(directory string) (results []string) {
	_ = filepath.WalkDir(directory, func(path string, dir fs.DirEntry, err error) error {
		if err == nil && dir.IsDir() && path != directory {
			results = append(results, strings.ReplaceAll(path, directory, ""))
		}
		return nil
	})
	return results
}
