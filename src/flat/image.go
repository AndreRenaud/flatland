package flat

import (
	"bytes"
	"flatland/src/asset"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Image struct {
	Path string
	img  *ebiten.Image
}

var _ = asset.RegisterAsset(Image{})

type ImageContainer interface {
	GetImage() *ebiten.Image
}

func (i *Image) PostLoad() {
	fmt.Printf("Post load %#v\n", i)
	content, err := asset.ReadFile(i.Path)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	i.img = ebiten.NewImageFromImage(img)
}

func (i *Image) GetImage() *ebiten.Image {
	return i.img
}
