package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleService interface {
	GetArticleId(ArticleId primitive.ObjectID) (*dtos.ArticleDTO, error)
	CreateArticle(ArticleInsertDTO dtos.ArticleInsertDTO) error
	UpdateArticle(ArticleId primitive.ObjectID, ArticleInsertDTO dtos.ArticleInsertDTO) error
	DeleteArticle(ArticleId primitive.ObjectID) error
}

type ArticleServiceImpl struct {
	ArticleMapper    mappers.ArticleMapper
	commonRepository repository.Repository[models.Article]
}

func NewArticleService(commonRepository repository.Repository[models.Article]) ArticleService {
	return &ArticleServiceImpl{
		commonRepository: commonRepository,
	}

}

func (bs ArticleServiceImpl) GetArticleId(articleId primitive.ObjectID) (*dtos.ArticleDTO, error) {
	article, err := bs.commonRepository.GetByID(context.TODO(), articleId)
	if err != nil {
		return nil, err
	}

	ArticleDTO := bs.ArticleMapper.EntityToDTO(*article)
	return &ArticleDTO, nil
}

func (bs ArticleServiceImpl) CreateArticle(articleInsertDTO dtos.ArticleInsertDTO) error {
	newArticle := bs.ArticleMapper.InsertDtoToEntity(articleInsertDTO)

	if _, err := bs.commonRepository.Create(context.TODO(), &newArticle); err != nil {
		return err
	}

	return nil
}

func (bs ArticleServiceImpl) UpdateArticle(articleId primitive.ObjectID, ArticleInsertDTO dtos.ArticleInsertDTO) error {
	articleUpdated := bs.ArticleMapper.InsertDtoToEntity(ArticleInsertDTO)

	if _, err := bs.commonRepository.Update(context.TODO(), articleId, &articleUpdated); err != nil {
		return err
	}

	return nil
}

func (bs ArticleServiceImpl) DeleteArticle(articleId primitive.ObjectID) error {
	if _, err := bs.commonRepository.DeleteById(context.TODO(), articleId); err != nil {
		return err
	}
	return nil
}
