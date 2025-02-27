package mergi

import (
	"image"
)

// Importer uses to plug any kind of image.Image importer
type Importer interface {
	Import() (image.Image, error)
}

// Exporter uses to plug any kind of image.Image exporter
type Exporter interface {
	Export() error
}

// Import uses to import image.Image from different sources
// Multiple io implementation can be find in io pkg
//
// for more Import usages https://github.com/dmitridb/mergi/examples
func Import(importer Importer) (image.Image, error) {
	return importer.Import()
}

// Export uses to export output do different sources
// Multiple exporter implementation can be find in io pkg
//
// for more Import usages https://github.com/dmitridb/mergi/examples
func Export(exporter Exporter) error {
	return exporter.Export()
}
