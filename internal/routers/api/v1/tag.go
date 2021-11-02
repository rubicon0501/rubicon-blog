package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	c.String(http.StatusOK, "Rubicon")
}
func (t Tag) List(c *gin.Context) {
	c.String(http.StatusOK, "Jack")
}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
