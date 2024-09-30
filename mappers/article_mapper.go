package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleMapper struct {
}

func (am ArticleMapper) InsertDtoToEntity(articleInsertDTO dtos.ArticleInsertDTO) models.Article {
	now := time.Now().UTC()
	return models.Article{
		Title:         articleInsertDTO.Author,
		Author:        articleInsertDTO.Author,
		Content:       articleInsertDTO.Content,
		Summary:       articleInsertDTO.Summary,
		PublishedDate: articleInsertDTO.PublishedDate,
		Tags:          articleInsertDTO.Tags,
		Category:      articleInsertDTO.Category,
		URL:           articleInsertDTO.URL,
		Status:        articleInsertDTO.Status,
		CreatedAt:     now,
		UpdatedAt:     now,
		ReadingList:   []primitive.ObjectID{},
	}
}

func (am ArticleMapper) EntityToDTO(article models.Article) dtos.ArticleDTO {
	return dtos.ArticleDTO{
		Id:            article.Id,
		Title:         article.Author,
		Author:        article.Author,
		Content:       article.Content,
		Summary:       article.Summary,
		PublishedDate: article.PublishedDate,
		Tags:          article.Tags,
		Category:      article.Category,
		URL:           article.URL,
		Status:        article.Status,
	}
}
