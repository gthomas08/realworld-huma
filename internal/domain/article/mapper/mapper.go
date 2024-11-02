package mapper

import (
	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
	"github.com/gthomas08/realworld-huma/internal/domain/article/dtos"
)

func TagsToTags(tags []model.Tags) []dtos.Tag {
	tagsResponse := make([]dtos.Tag, 0, len(tags))
	for _, tag := range tags {
		tagsResponse = append(tagsResponse, dtos.Tag(tag.Name))
	}
	return tagsResponse
}
