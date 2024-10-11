package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingListMapper struct {
}

func (rlm ReadingListMapper) InsertDtoToEntity(insertDTO dtos.ReadingListInsertDTO) models.ReadingsList {
	return models.ReadingsList{
		Id:          primitive.NewObjectID(),
		Name:        insertDTO.Name,
		ReadingIds:  []primitive.ObjectID{},
		Description: insertDTO.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

func (rlm ReadingListMapper) EntityToDTO(readingList models.ReadingsList) dtos.ReadingListDTO {
	return dtos.ReadingListDTO{
		Id:          primitive.NewObjectID(),
		Name:        readingList.Name,
		ReadingIds:  readingList.ReadingIds,
		Description: readingList.Description,
	}
}

func (rlm ReadingListMapper) InsertDtoToUpdateEntity(insertDTO dtos.ReadingListInsertDTO) models.ReadingsList {
	return models.ReadingsList{
		Name:        insertDTO.Name,
		Description: insertDTO.Description,
		UpdatedAt:   time.Now().UTC(),
	}
}
