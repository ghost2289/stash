package autotag

import (
	"github.com/stashapp/stash/pkg/gallery"
	"github.com/stashapp/stash/pkg/image"
	"github.com/stashapp/stash/pkg/models"
	"github.com/stashapp/stash/pkg/scene"
)

func getTagTaggers(p *models.Tag, aliases []string) []tagger {
	ret := []tagger{{
		ID:   p.ID,
		Type: "tag",
		Name: p.Name,
	}}

	for _, a := range aliases {
		ret = append(ret, tagger{
			ID:   p.ID,
			Type: "tag",
			Name: a,
		})
	}

	return ret
}

// TagScenes searches for scenes whose path matches the provided tag name and tags the scene with the tag.
func TagScenes(p *models.Tag, paths []string, aliases []string, rw models.SceneReaderWriter) error {
	t := getTagTaggers(p, aliases)

	for _, tt := range t {
		if err := tt.tagScenes(paths, rw, func(subjectID, otherID int) (bool, error) {
			return scene.AddTag(rw, otherID, subjectID)
		}); err != nil {
			return err
		}
	}
	return nil
}

// TagImages searches for images whose path matches the provided tag name and tags the image with the tag.
func TagImages(p *models.Tag, paths []string, aliases []string, rw models.ImageReaderWriter) error {
	t := getTagTaggers(p, aliases)

	for _, tt := range t {
		if err := tt.tagImages(paths, rw, func(subjectID, otherID int) (bool, error) {
			return image.AddTag(rw, otherID, subjectID)
		}); err != nil {
			return err
		}
	}
	return nil
}

// TagGalleries searches for galleries whose path matches the provided tag name and tags the gallery with the tag.
func TagGalleries(p *models.Tag, paths []string, aliases []string, rw models.GalleryReaderWriter) error {
	t := getTagTaggers(p, aliases)

	for _, tt := range t {
		if err := tt.tagGalleries(paths, rw, func(subjectID, otherID int) (bool, error) {
			return gallery.AddTag(rw, otherID, subjectID)
		}); err != nil {
			return err
		}
	}
	return nil
}
