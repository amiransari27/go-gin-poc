package service

import (
	"go-gin-api/src/dao"
	"go-gin-api/src/entity"
	"go-gin-api/src/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type INoteService interface {
	Save(*gin.Context, *entity.Note) (interface{}, error)
	FindAllForLoggedInUser(*gin.Context) ([]*model.Note, error)
	FindOneForLoggedInUser(*gin.Context, string) (*model.Note, error)
	FindOneUpdateForLoggedInUser(*gin.Context, *entity.Note, string) (*model.Note, error)
}

type noteService struct {
	noteDao dao.INoteDao
	userDao dao.IUserDao
}

func NewNoteService(nd dao.INoteDao, ud dao.IUserDao) INoteService {
	return &noteService{
		noteDao: nd,
		userDao: ud,
	}
}

func (s *noteService) Save(ctx *gin.Context, note *entity.Note) (interface{}, error) {

	// fetch user data from mongo
	userObj, err := getLoggedInUser(ctx, s.userDao)
	if err != nil {
		return nil, err
	}

	noteId, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	newNote := &model.Note{
		UserId:     userObj.UserId,
		NoteId:     noteId.String(),
		Title:      note.Title,
		Content:    note.Content,
		SharedWith: []string{},
	}

	inserted, err := s.noteDao.Save(newNote)

	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (s *noteService) FindAllForLoggedInUser(ctx *gin.Context) ([]*model.Note, error) {

	// fetch user data from mongo
	userObj, err := getLoggedInUser(ctx, s.userDao)
	if err != nil {
		return nil, err
	}

	notes, err := s.noteDao.Find(bson.M{"userId": userObj.UserId})
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *noteService) FindOneForLoggedInUser(ctx *gin.Context, noteId string) (*model.Note, error) {

	// fetch user data from mongo
	userObj, err := getLoggedInUser(ctx, s.userDao)
	if err != nil {
		return nil, err
	}

	note, err := s.noteDao.FindOne(bson.M{"userId": userObj.UserId, "noteId": noteId})
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (s *noteService) FindOneUpdateForLoggedInUser(ctx *gin.Context, note *entity.Note, noteId string) (*model.Note, error) {

	// fetch user data from mongo
	userObj, err := getLoggedInUser(ctx, s.userDao)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"userId": userObj.UserId,
		"noteId": noteId,
	}

	updateNote := bson.M{
		"$set": note,
	}

	updatedNote, err := s.noteDao.FindOneAndUpdate(filter, updateNote)

	if err != nil {
		return nil, err
	}
	return updatedNote, nil
}
