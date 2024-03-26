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
// // @Router       /result [post]
// // @Summary      create result
// // @Description  api for create result
// // @Tags         results
// // @Accept       json
// // @Produce      json
// // @Param        club    body     models.CreateResult  true  "date of result"
// // @Success      200  {string}   string
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) CreateResult(c *gin.Context) {
// 	var result models.CreateResult
// 	err := c.ShouldBindJSON(&result)
// 	if err != nil {
// 		h.log.Error("error while binding:", logger.Error(err))
// 		c.JSON(http.StatusBadRequest, "invalid body")
// 		return
// 	}
// 	fmt.Println(h.strg)
// 	resp, err := h.strg.Result().CreateResult(c.Request.Context(), result)
// 	if err != nil {
// 		fmt.Println("error Result Create:", err.Error())
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		return
// 	}
// 	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})

// }

// // ListAccounts godoc
// // @Router       /result/{id} [get]
// // @Summary      get result
// // @Description  get results
// // @Tags         results
// // @Accept       json
// // @Produce      json
// // @Param        id    path     string  true  "id of result"  Format(uuid)
// // @Success      200  {object}   models.Result
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) GetResult(c *gin.Context) {
// 	fmt.Println("MethodGet")

// 	id := c.Param("id")

// 	resp, err := h.strg.Result().GetResult(c.Request.Context(), models.IdRequestResult{Id: id})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		fmt.Println("error Result Get:", err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp)
// }

// // ListAccounts godoc
// // @Router       /result/{id} [delete]
// // @Summary      delete result
// // @Description   api fot delete results
// // @Tags         results
// // @Accept       json
// // @Produce      json
// // @Param        id    path     string  true  "id of result"
// // @Success      200  {strig}   string
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) DeleteResult(c *gin.Context) {
// 	id := c.Param("id")
// 	if !helper.IsValidUUID(id) {
// 		h.log.Error("error Club Delete:", logger.Error(errors.New("invalid id")))
// 		c.JSON(http.StatusBadRequest, "invalid id")
// 		return
// 	}
// 	resp, err := h.strg.Result().DeleteResult(c.Request.Context(), models.IdRequestResult{Id: id})
// 	if err != nil {
// 		h.log.Error("error Result Detete:", logger.Error(err))
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// // ListAccounts godoc
// // @Router       /result [get]
// // @Summary      List results
// // @Description  get results
// // @Tags         results
// // @Accept       json
// // @Produce      json
// // @Param        limit    query     integer  true  "limit for response"  Default(10)
// // @Param        page    query     integer  true  "page of req"  Default(1)
// // @Success      200  {object}   models.GetAllResult
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) GetAllResult(c *gin.Context) {
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

// 	resp, errs := h.strg.Result().GetAllResult(c.Request.Context(), models.GetAllResultRequest{
// 		Page:    page,
// 		Limit:   limit,
// 		Search:  c.Query("search"),
// 	})
// 	if errs != nil {
// 		h.log.Error("error Result GetAll:", logger.Error(err))
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		return
// 	}
// 	h.log.Warn("response to GetAllResult")
// 	c.JSON(http.StatusOK, resp)
// }
