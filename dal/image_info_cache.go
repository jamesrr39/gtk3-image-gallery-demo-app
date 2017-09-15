package dal

import (
	"github.com/jamesrr39/gtk3-app/domain"
)

type imageInfoCache struct {
	imageInfos []*domain.ImageInfo
}

func newImageInfoCache() *imageInfoCache {
	return &imageInfoCache{}
}

func (c *imageInfoCache) refresh(imageInfos []*domain.ImageInfo) {
	c.imageInfos = imageInfos
}

func (c *imageInfoCache) getAll() []*domain.ImageInfo {
	return c.imageInfos
}
