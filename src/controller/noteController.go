package controller

import (
	"go-gin-api/src/entity"
	"go-gin-api/src/middleware"
	"go-gin-api/src/model"
	"go-gin-api/src/service"
	"net/http"

	logger "github.com/openscriptsin/go-logger"

	"github.com/gin-gonic/gin"
)

type noteController struct {
	noteServ service.INoteService
	logger   logger.ILogrus
}

func NewNoteController(server *gin.Engine, noteServ service.INoteService, jwtService service.IJWTService, logger logger.ILogrus) {

	controller := &noteController{
		noteServ,
		logger,
	}

	group := server.Group("/notes", middleware.AuthMiddleware(jwtService, logger))

	group.GET("/", func(ctx *gin.Context) {
		notes, err := controller.findAll(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, notes)
		}
	})

	group.GET("/:noteId", func(ctx *gin.Context) {
		noteId := ctx.Param("noteId")
		note, err := controller.findOne(ctx, noteId)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, note)
		}
	})

	group.PUT("/:noteId", func(ctx *gin.Context) {
		noteId := ctx.Param("noteId")
		note, err := controller.updateOne(ctx, noteId)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, note)
		}
	})

	group.POST("/", func(ctx *gin.Context) {
		message, err := controller.save(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": message})
		}

	})

}

// @Summary Add new Note
// @Schemes
// @Description add new note
// @Security ApiKeyAuth
// @Tags Notes
// @Accept json
// @Produce json
// @Success 200 {string} save
// @Router /notes [post]
// @Param Authorization header string true "Bearer Token"
// @Param data body entity.Note true "note payload"
func (c *noteController) save(ctx *gin.Context) (string, error) {
	var obj entity.Note
	err := ctx.ShouldBindJSON(&obj)
	if err != nil {
		return "", err
	}

	_, err = c.noteServ.Save(ctx, &obj)

	if err != nil {
		return "", err
	}

	return "note added successfully", nil
}

// @Summary Fetch notes
// @Schemes
// @Description Fetch all notes for user
// @Security ApiKeyAuth
// @Tags Notes
// @Accept json
// @Produce json
// @Success 200 {object} []model.Note
// @Router /notes [get]
// @Param Authorization header string true "Bearer Token"
func (c *noteController) findAll(ctx *gin.Context) ([]*model.Note, error) {
	notes, err := c.noteServ.FindAllForLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

// @Summary Fetch note
// @Schemes
// @Description Fetch one note for user by note id
// @Security ApiKeyAuth
// @Tags Notes
// @Accept json
// @Produce json
// @Success 200 {object} model.Note
// @Router /notes/{noteId} [get]
// @Param Authorization header string true "Bearer Token"
// @Param noteId path string true "note id"
func (c *noteController) findOne(ctx *gin.Context, noteId string) (*model.Note, error) {
	note, err := c.noteServ.FindOneForLoggedInUser(ctx, noteId)
	if err != nil {
		return nil, err
	}
	return note, nil
}

// @Summary Update Note
// @Schemes
// @Description update a note by its id
// @Security ApiKeyAuth
// @Tags Notes
// @Accept json
// @Produce json
// @Success 200 {object} model.Note
// @Router /notes/{noteId} [put]
// @Param Authorization header string true "Bearer Token"
// @Param noteId path string true "note id"
// @Param data body entity.Note true "note payload"
func (c *noteController) updateOne(ctx *gin.Context, noteId string) (*model.Note, error) {
	var obj entity.Note
	err := ctx.ShouldBindJSON(&obj)
	if err != nil {
		return nil, err
	}

	note, err := c.noteServ.FindOneUpdateForLoggedInUser(ctx, &obj, noteId)

	if err != nil {
		return nil, err
	}

	return note, nil
}
