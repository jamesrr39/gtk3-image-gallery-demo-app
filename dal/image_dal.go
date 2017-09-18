package dal

import (
	"errors"
	"image"
	_ "image/gif"  // Decode
	_ "image/jpeg" // Decode
	_ "image/png"  // Decode
	"os"
	"path/filepath"
	"strings"

	"github.com/jamesrr39/goutil/dirtraversal"
	"github.com/jamesrr39/gtk3-image-gallery-demo-app/domain"
	"github.com/spf13/afero"
)

type ImageDAL struct {
	rootPath       string
	fs             afero.Fs
	imageInfoCache *imageInfoCache
}

func NewImageDAL(rootPath string) *ImageDAL {
	return &ImageDAL{rootPath, afero.NewOsFs(), newImageInfoCache()}
}

func (dal *ImageDAL) ScanRootPath() error {
	var imageInfos []*domain.ImageInfo

	err := afero.Walk(dal.fs, dal.rootPath, func(path string, fileInfo os.FileInfo, err error) error {
		if nil != err {
			return err
		}

		if fileInfo.IsDir() {
			return nil
		}

		relativeFilePath := strings.TrimPrefix(strings.TrimPrefix(path, dal.rootPath), string(filepath.Separator))

		imageInfos = append(imageInfos, domain.NewImageInfo(relativeFilePath))

		return nil
	})

	if nil != err {
		return err
	}

	dal.imageInfoCache.refresh(imageInfos)

	return nil
}

func (dal *ImageDAL) GetAll() []*domain.ImageInfo {
	return dal.imageInfoCache.getAll()
}

func (dal *ImageDAL) GetImage(relativePath string) (image.Image, error) {
	isTryingToTraverse := dirtraversal.IsTryingToTraverseUp(relativePath)
	if isTryingToTraverse {
		return nil, errors.New("Illegal traverse up the tree")
	}

	file, err := os.Open(filepath.Join(dal.rootPath, relativePath))
	if nil != err {
		return nil, err
	}
	defer file.Close()

	picture, _, err := image.Decode(file)
	if nil != err {
		return nil, err
	}

	return picture, nil
}
