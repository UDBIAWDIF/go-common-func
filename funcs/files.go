package funcs

import (
	"archive/zip"
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// 保存 image.Image 数据到图片文件, 将用 go 官方的图片数据解析器自动判断数据格式, 也会用文件扩展名自动存成相应的图片类型
func SaveImageDataToFileAutoType(imageFile string, imageData image.Image) error {
	jpgQuality := 80
	gifNumColors := 256
	var filePerm fs.FileMode = 0666

	saveImgFuncsMap := map[string]func(*bufio.Writer, image.Image) error{
		".jpg": func(imageFileWriter *bufio.Writer, imageData image.Image) error {
			return jpeg.Encode(imageFileWriter, imageData, &jpeg.Options{Quality: jpgQuality})
		},
		".png": func(imageFileWriter *bufio.Writer, imageData image.Image) error {
			return png.Encode(imageFileWriter, imageData)
		},
		".gif": func(imageFileWriter *bufio.Writer, imageData image.Image) error {
			return gif.Encode(imageFileWriter, imageData, &gif.Options{NumColors: gifNumColors})
		},
	}

	imageFileHandle, err := os.OpenFile(imageFile, os.O_SYNC|os.O_RDWR|os.O_CREATE, filePerm)
	if err != nil {
		return err
	}
	defer imageFileHandle.Close()

	imageFileWriter := bufio.NewWriter(imageFileHandle)
	ext := filepath.Ext(imageFile)
	saveImgFunc, funcExists := saveImgFuncsMap[ext]
	if funcExists {
		err = saveImgFunc(imageFileWriter, imageData)
		imageFileWriter.Flush()
	} else {
		err = fmt.Errorf("截图无法保存, 不支持保存为此图片格式: %s", ext)
	}

	return err
}

/**
 * 复制文件
 */
func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

/**
 * 带通配符删除文件
 */
func RemoveFilesWildCard(filesWildCard string) {
	files, err := filepath.Glob(filesWildCard)
	if err != nil {
		return
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			continue
		}
	}
}

func CheckFileIsExist(filename string) bool {
	var exist = false
	_, err := os.Stat(filename)
	exist = err == nil

	if err != nil {
		exist = !os.IsNotExist(err)
	}

	return exist
}

func RootPath() string {
	var rootPath, _ = os.Getwd()
	return rootPath + "/"
}

func ReadFileAll(filename string) (content []byte, err error) {
	runnable := true
	var hFile *os.File

	hFile, err = os.Open(filename)
	if err != nil {
		runnable = false
	}

	if runnable {
		content, err = io.ReadAll(hFile)
		if err != nil {
			runnable = false
		}
	}

	if runnable {
		hFile.Close()
	}

	return
}

/**
 * 生成多级目录, 目录属性默认设为 os.ModePerm
 *
 * @param path 目录, 可支持多级目录
 *
 */
func MakeDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

/**
 * 以覆盖的方式写文件, 文件不存在时自动创建属性默认设为 0644
 *
 * @param fileName 文件名
 * @param content 文件内容
 *
 * @return 错误信息
 */
func WriteFile(fileFullName string, content []byte) error {
	dir, _ := filepath.Split(fileFullName)
	MakeDirs(dir)
	return os.WriteFile(fileFullName, content, 0644)
}

// 将指定目录下的所有文件压缩成一个zip文件
func ZipDirFiles(zipFile, dir string) error {
	zipFileHandle, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer zipFileHandle.Close()

	archiveWriter := zip.NewWriter(zipFileHandle)
	defer archiveWriter.Close()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(path)

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archiveWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			io.Copy(writer, file)
		}
		return err
	})

	return err
}

// 将指定目录下的所有文件压缩成一个zip文件
func ZipContent(zipFile string, fileName string, content []byte) error {
	zipFileHandle, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer zipFileHandle.Close()

	archiveWriter := zip.NewWriter(zipFileHandle)
	defer archiveWriter.Close()

	fileWriter, _ := archiveWriter.Create(fileName)
	fileWriter.Write(content)

	return err
}

// 解压zip文件里第一个文件(用于只压缩了一个文件的zip包, 经常会碰到, 避免再输入文件名)
func UnZipFirstFile(zipFile string) ([]byte, error) {
	var content []byte
	archive, err := zip.OpenReader(zipFile)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, fileInArchive := range archive.File {
		var fileRead io.ReadCloser
		fileRead, err = fileInArchive.Open()
		if err != nil {
			break
		}

		content = make([]byte, fileInArchive.UncompressedSize64)
		_, err = fileRead.Read(content)
		fileRead.Close()
		break
	}

	return content, err
}
