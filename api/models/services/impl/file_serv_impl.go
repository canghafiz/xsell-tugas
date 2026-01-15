package impl

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"be/models/domains"
)

const (
	MaxImageSize = 50 << 20 // 50 MB
	MaxDocSize   = 10 << 20 // 10 MB
	UploadDir    = "assets"
)

type FileServImpl struct{}

func NewFileServImpl() *FileServImpl {
	return &FileServImpl{}
}

func (serv *FileServImpl) UploadFiles(files []*multipart.FileHeader) ([]domains.File, error) {
	if len(files) == 0 {
		return nil, fmt.Errorf("no files provided")
	}

	if err := os.MkdirAll(UploadDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	var uploaded []domains.File

	allowedExts := map[string]int64{
		".jpg":  MaxImageSize,
		".jpeg": MaxImageSize,
		".png":  MaxImageSize,
		".webp": MaxImageSize,
		".pdf":  MaxDocSize,
		".doc":  MaxDocSize,
		".docx": MaxDocSize,
	}

	for _, fileHeader := range files {
		// --- Validasi ---
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		maxSize, allowed := allowedExts[ext]
		if !allowed {
			return nil, fmt.Errorf("file type not allowed: %s", ext)
		}
		if fileHeader.Size > maxSize {
			return nil, fmt.Errorf("file %s too large (max %d MB)", fileHeader.Filename, maxSize>>20)
		}

		// --- Generate nama aman ---
		baseName := strings.TrimSuffix(fileHeader.Filename, ext)
		safeBaseName := sanitizeFileName(baseName)
		fileName := safeBaseName + ext
		dstPath := filepath.Join(UploadDir, fileName)
		counter := 1
		for {
			if _, err := os.Stat(dstPath); os.IsNotExist(err) {
				break
			}
			fileName = fmt.Sprintf("%s(%d)%s", safeBaseName, counter, ext)
			dstPath = filepath.Join(UploadDir, fileName)
			counter++
		}

		// --- Open file source ---
		src, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %w", fileHeader.Filename, err)
		}

		// --- Open file destiny ---
		dst, errCreate := os.Create(dstPath)
		if errCreate != nil {
			src.Close() // ✅ Tutup segera jika error
			return nil, fmt.Errorf("failed to create file %s: %w", fileName, err)
		}

		// --- Copy file ---
		_, err = io.Copy(dst, src)
		// Close file
		dst.Close()
		src.Close()

		if err != nil {
			// Jika gagal salin, hapus file sementara
			os.Remove(dstPath)
			return nil, fmt.Errorf("failed to save file %s: %w", fileName, err)
		}

		uploaded = append(uploaded, domains.File{
			FileName: fileName,
			FileURL:  fmt.Sprintf("/assets/%s", fileName),
		})
	}

	return uploaded, nil
}

func (serv *FileServImpl) DeleteFile(fileURL string) error {
	if fileURL == "" {
		return fmt.Errorf("file URL is required")
	}

	// Convert URL to filesystem path
	// e.g., "/assets/photo.jpg" → "assets/photo.jpg"
	path := strings.TrimPrefix(fileURL, "/")
	if !strings.HasPrefix(path, UploadDir+"/") {
		return fmt.Errorf("invalid file path")
	}

	fullPath := filepath.Join(".", path)
	if err := os.Remove(fullPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", fileURL)
		}
		return fmt.Errorf("failed to delete file %s: %w", fileURL, err)
	}

	return nil
}

// sanitizeFileName removes unsafe characters for URLs and filesystems
func sanitizeFileName(name string) string {
	// Keep only letters, numbers, underscore, dash, and space
	reg := regexp.MustCompile(`[^a-zA-Z0-9_\- ]+`)
	safe := reg.ReplaceAllString(name, "")

	// Replace spaces with underscores
	safe = strings.ReplaceAll(safe, " ", "_")

	// Ensure not empty
	if safe == "" {
		safe = "file"
	}

	// Limit length
	if len(safe) > 100 {
		safe = safe[:100]
	}

	return safe
}
