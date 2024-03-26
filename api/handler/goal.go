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
// // @Router       /goal [post]
// // @Summary      create goal
// // @Description  api for create goal
// // @Tags         goals
// // @Accept       json
// // @Produce      json
// // @Param        goal    body     models.CreateGoal  true  "date of goal"
// // @Success      200  {string}   string
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) CreateGoal(c *gin.Context) {
// 	var goal models.CreateGoal
// 	err := c.ShouldBindJSON(&goal)
// 	if err != nil {
// 		h.log.Error("error while binding:", logger.Error(err))
// 		c.JSON(http.StatusBadRequest, "invalid body")
// 		return
// 	}
// 	fmt.Println(h.strg)
// 	resp, err := h.strg.Goal().CreateGoal(c.Request.Context(), goal)
// 	if err != nil {
// 		fmt.Println("error Goal Create:", err.Error())
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		return
// 	}
// 	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})

// }

// // ListAccounts godoc
// // @Router       /goal/{id} [get]
// // @Summary      get goal
// // @Description  get goals
// // @Tags         goals
// // @Accept       json
// // @Produce      json
// // @Param        id    path     string  true  "id of goal"  Format(uuid)
// // @Success      200  {object}   models.Goal
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) GetGoal(c *gin.Context) {
// 	fmt.Println("MethodGet")

// 	id := c.Param("id")

// 	resp, err := h.strg.Goal().GetGoal(c.Request.Context(), models.IdRequestGoal{Id: id})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		fmt.Println("error Goal Get:", err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp)
// }

// // ListAccounts godoc
// // @Router       /goal/{id} [delete]
// // @Summary      delete goal
// // @Description   api fot delete goals
// // @Tags         goals
// // @Accept       json
// // @Produce      json
// // @Param        id    path     string  true  "id of goal"
// // @Success      200  {strig}   string
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) DeleteGoal(c *gin.Context) {
// 	id := c.Param("id")
// 	if !helper.IsValidUUID(id) {
// 		h.log.Error("error Goal Delete:", logger.Error(errors.New("invalid id")))
// 		c.JSON(http.StatusBadRequest, "invalid id")
// 		return
// 	}
// 	resp, err := h.strg.Goal().DeleteGoal(c.Request.Context(), models.IdRequestGoal{Id: id})
// 	if err != nil {
// 		h.log.Error("error Goal Detete:", logger.Error(err))
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// // ListAccounts godoc
// // @Router       /goal [get]
// // @Summary      List goal
// // @Description  get goals
// // @Tags         goals
// // @Accept       json
// // @Produce      json
// // @Param        limit    query     integer  true  "limit for response"  Default(10)
// // @Param        page    query     integer  true  "page of req"  Default(1)
// // @Success      200  {object}   models.GetAllGoal
// // @Failure      400  {object}  response.ErrorResp
// // @Failure      404  {object}  response.ErrorResp
// // @Failure      500  {object}  response.ErrorResp
// func (h *Handler) GetAllGoal(c *gin.Context) {
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

// 	resp, errs := h.strg.Goal().GetAllGoal(c.Request.Context(), models.GetAllGoalRequest{
// 		Page:    page,
// 		Limit:   limit,
// 		Search:  c.Query("search"),
// 	})
// 	if errs != nil {
// 		h.log.Error("error Goal GetAll:", logger.Error(err))
// 		c.JSON(http.StatusInternalServerError, "internal server error")
// 		return
// 	}
// 	h.log.Warn("response to GetAllGoal")
// 	c.JSON(http.StatusOK, resp)
// }
