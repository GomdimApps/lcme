package compressfiles

import (
	"archive/tar"
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// AddFileToZip adiciona um único arquivo ao arquivo ZIP.
// Se baseFolder não for vazio, é calculado o caminho relativo para preservar a estrutura.
func AddFileToZip(zipWriter *zip.Writer, filename, baseFolder string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	var relPath string
	if baseFolder != "" {
		relPath, err = filepath.Rel(baseFolder, filename)
		if err != nil {
			return err
		}
	} else {
		relPath = filepath.Base(filename)
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = relPath
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}

// AddFileToTar adiciona um único arquivo ao arquivo TAR.
// Se baseFolder não for vazio, é calculado o caminho relativo para preservar a estrutura.
func AddFileToTar(tarWriter *tar.Writer, filename, baseFolder string) error {
	fileToTar, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToTar.Close()

	info, err := fileToTar.Stat()
	if err != nil {
		return err
	}

	var relPath string
	if baseFolder != "" {
		relPath, err = filepath.Rel(baseFolder, filename)
		if err != nil {
			return err
		}
	} else {
		relPath = filepath.Base(filename)
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	header.Name = relPath

	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, fileToTar)
	return err
}
