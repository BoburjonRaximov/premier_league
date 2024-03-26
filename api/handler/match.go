package handler

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"

// 	"strconv"

// 	"github.com/BoburjonRaximov/premier_league/models"
// 	"github.com/BoburjonRaximov/premier_league/pkg/helper"
// 	"github.com/BoburjonRaximov/premier_league/pkg/logger"
// 	response "github.com/BoburjonRaximov/premier_league/responce"

// 	"github.com/gin-gonic/gin"
// )

// // ListAccounts godoc
// // @Router       /match [post]
// // @Summary      create match
// // @Description  api for create match
// // @Tags         matches
// // @Accept       json
// // @Produce      json
// // @Param        match    body     models.CreateMatch  true  "date of Match"
// // @Success      200  {string}   string
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) CreateMatch(c *gin.Context) {
// 	var match models.CreateMatch
// 	err := c.ShouldBindJSON(&match)
// 	if err != nil {
// 		h.log.Error("error while binding:", logger.Error(err))
// 		c.JSON(http.StatusBadRequest, "invalid body")
// 		return
// 	}
// 	fmt.Println(h.strg)
// 	resp, err := h.strg.Match().CreateMatch(c.Request.Context(), match)
// 	if err != nil {
// 		fmt.Println("error Match Create:", err.Error())
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		return
// 	}
// 	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})

// }

// // ListAccounts godoc
// // @Router       /match/{id} [get]
// // @Summary      get match
// // @Description  get matches
// // @Tags         matches
// // @Accept       json
// // @Produce      json
// // @Param        id    path     string  true  "id of match"  Format(uuid)
// // @Success      200  {object}   models.Match
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) GetMatch(c *gin.Context) {
// 	fmt.Println("MethodGet")

// 	id := c.Param("id")

// 	resp, err := h.strg.Match().GetMatch(c.Request.Context(), models.IdRequestMatch{Id: id})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		fmt.Println("error Match Get:", err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp)
// }

// // ListAccounts godoc
// // @Router       /match/{id} [delete]
// // @Summary      delete match
// // @Description   api fot delete matches
// // @Tags         matches
// // @Accept       json
// // @Produce      json
// // @Param        id    path     string  true  "id of match"
// // @Success      200  {strig}   string
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) DeleteMatch(c *gin.Context) {
// 	id := c.Param("id")
// 	if !helper.IsValidUUID(id) {
// 		h.log.Error("error Match Delete:", logger.Error(errors.New("invalid id")))
// 		c.JSON(http.StatusBadRequest, "invalid id")
// 		return
// 	}
// 	resp, err := h.strg.Match().DeleteMatch(c.Request.Context(), models.IdRequestMatch{Id: id})
// 	if err != nil {
// 		h.log.Error("error Match Detete:", logger.Error(err))
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// // ListAccounts godoc
// // @Router       /match [get]
// // @Summary      List matches
// // @Description  get matches
// // @Tags         matches
// // @Accept       json
// // @Produce      json
// // @Param        limit    query     integer  true  "limit for response"  Default(10)
// // @Param        page    query     integer  true  "page of req"  Default(1)
// // @Success      200  {object}   models.GetAllMatch
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) GetAllMatch(c *gin.Context) {
// 	h.log.Info("request GetAll")
// 	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
// 	if err != nil {
// 		h.log.Error("error get page:", logger.Error(err))
// 		c.JSON(http.StatusBadRequest, "invalid page param")
// 		return
// 	}
// 	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
// 	if err != nil {
// 		h.log.Error("error get limit:", logger.Error(err))
// 		c.JSON(http.StatusBadRequest, "invalid page param")
// 		return
// 	}

// 	resp, errs := h.strg.Match().GetAllMatch(c.Request.Context(), models.GetAllMatchRequest{
// 		Page:    page,
// 		Limit:   limit,
// 		Search:  c.Query("search"),
// 	})
// 	if errs != nil {
// 		h.log.Error("error Match GetAll:", logger.Error(err))
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		return
// 	}
// 	h.log.Warn("response to GetAllMatches")
// 	c.JSON(http.StatusOK, resp)
// }
