package domain

type ImageInfo struct {
	RelativePath string
}

func NewImageInfo(relativeFilePath string) *ImageInfo {
	return &ImageInfo{relativeFilePath}
}
