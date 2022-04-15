package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/rickyseezy/block/internal/domain/models"
	"log"
	"net/http"
)

// groups godoc
// @Summary      List groups
// @Description  List groups of index ids
// @Tags         Groups
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Group
// @Failure      400  {object}  models.AppResponse
// @Router       /groups [get]
func (s *Server) groups(c *gin.Context) {
	groups, err := s.blockIndexApp.ListGroups(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models.AppResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// group godoc
// @Summary      Get group
// @Description  Show a group of index ids
// @Tags         Groups
// @Accept       json
// @Produce      json
// @Param        id path int true "group ID"
// @Success      200  {object}  models.Group
// @Failure      400  {object}  models.AppResponse
// @Router       /groups/{id} [get]
func (s *Server) group(c *gin.Context) {
	group, err := s.blockIndexApp.GetGroup(c, c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models.AppResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, group)

}

// index godoc
// @Summary      Get index
// @Description  Show an index
// @Tags         Indexes
// @Accept       json
// @Produce      json
// @Param        id   path int true "index ID"
// @Success      200  {object}  models.Index
// @Failure      400  {object}  models.AppResponse
// @Router       /indexes/{id} [get]
func (s *Server) index(c *gin.Context) {
	groups, err := s.blockIndexApp.GetIndex(c, c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models.AppResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// block godoc
// @Summary      Get block
// @Description  Get block information
// @Tags         Blocks
// @Accept       json
// @Produce      json
// @Param        search path string true "search param can have the following values BlockNumber | BlockHash | 'latest'"
// @Success      200  {object}  models.Block
// @Failure      400  {object}  models.AppResponse
// @Router       /blocks/{search} [get]
func (s *Server) block(c *gin.Context) {
	block, err := s.blockIndexApp.GetBlock(c, c.Param("search"))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models.AppResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, block)
}
