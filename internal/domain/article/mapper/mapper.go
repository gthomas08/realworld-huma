package mapper

import (
	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
	"github.com/gthomas08/realworld-huma/internal/domain/article/dtos"
)

func TagsToTags(tags []model.Tags) []dtos.Tag {
	var tagsResponse []dtos.Tag
	for _, tag := range tags {
		tagsResponse = append(tagsResponse, dtos.Tag(tag.Tag))
	}
	return tagsResponse
}
