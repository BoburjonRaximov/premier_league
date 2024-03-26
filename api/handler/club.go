package handler

import (
	"errors"
	"fmt"
	"net/http"

	"strconv"

	"github.com/BoburjonRaximov/premier_league/models"
	"github.com/BoburjonRaximov/premier_league/pkg/helper"
	"github.com/BoburjonRaximov/premier_league/pkg/logger"
	response "github.com/BoburjonRaximov/premier_league/responce"

	"github.com/gin-gonic/gin"
)

// ListAccounts godoc
// @Router       /club [post]
// @Summary      create club
// @Description  api for create club
// @Tags         clubs
// @Accept       json
// @Produce      json
// @Param        club    body     models.CreateClub  true  "date of club"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateClub(c *gin.Context) {
	var club models.CreateClub
	err := c.ShouldBindJSON(&club)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	fmt.Println(h.strg)
	resp, err := h.strg.Club().CreateClub(c.Request.Context(), club)
	if err != nil {
		fmt.Println("error Club Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})

}

// ListAccounts godoc
// @Router       /club/{id} [get]
// @Summary      get club
// @Description  get clubs
// @Tags         clubs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of club"  Format(uuid)
// @Success      200  {object}   models.Club
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetClub(c *gin.Context) {
	fmt.Println("MethodGet")

	id := c.Param("id")

	resp, err := h.strg.Club().GetClub(c.Request.Context(), models.IdRequestClub{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Club Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /club/{id} [put]
// @Summary      updateda club
// @Description   api fot update clubs
// @Tags         clubs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of club"
// @Param        club    body     models.CreateClub  true  "id of club"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateClub(c *gin.Context) {
	var club models.Club
	err := c.ShouldBind(&club)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	club.Id = c.Param("id")
	resp, err := h.strg.Club().UpdateClub(c.Request.Context(),club)
	if err != nil {
		fmt.Println("error Club Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)

}

// ListAccounts godoc
// @Router       /club/{id} [delete]
// @Summary      delete club
// @Description   api fot delete clubes
// @Tags         clubs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of club"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteClub(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Club Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.Club().DeleteClub(c.Request.Context(), models.IdRequestClub{Id: id})
	if err != nil {
		h.log.Error("error Club Detete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /club [get]
// @Summary      List clubs
// @Description  get clubs
// @Tags         clubs
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}   models.GetAllClub
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllClub(c *gin.Context) {
	h.log.Info("request GetAll")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	resp, errs := h.strg.Club().GetAllClub(c.Request.Context(), models.GetAllClubRequest{
		Page:    page,
		Limit:   limit,
		Search:  c.Query("search"),
	})
	if errs != nil {
		h.log.Error("error Club GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllClubs")
	c.JSON(http.StatusOK, resp)
}
