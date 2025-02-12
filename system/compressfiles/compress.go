package compressfiles

import (
	"archive/tar"
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// addFileToZip adds a single file to the ZIP archive.
func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	// Open the file to be added
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	// Create a ZIP header based on the file information
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filepath.Base(filename)
	header.Method = zip.Deflate

	// Create a writer for the file header
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Copy the file data to the ZIP writer
	_, err = io.Copy(writer, fileToZip)
	return err
}

// addFileToTar adds a single file to the TAR archive.
func AddFileToTar(tarWriter *tar.Writer, filename string) error {
	// Open the file to be added
	fileToTar, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToTar.Close()

	// Get file information
	info, err := fileToTar.Stat()
	if err != nil {
		return err
	}

	// Create a TAR header based on the file information
	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	header.Name = filepath.Base(filename)

	// Write the header to the TAR writer
	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	// Copy the file data to the TAR writer
	_, err = io.Copy(tarWriter, fileToTar)
	return err
}
