package common

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_PAGE = 0
	DEFAULT_LIMIT = 10
)

type FilterParam struct {
	Page	int		`json:"page"`
	Limit	int		`json:"limit"`
	Offset	int		`json:"offset"`
	Order	int		`json:"order"`
}

func GetPagination(c *gin.Context) (int, int) {
	defaultPage := fmt.Sprintf("%d", DEFAULT_PAGE)
	page, err := strconv.Atoi(c.DefaultQuery("page", defaultPage))

	if err != nil {
		page = DEFAULT_PAGE
	}
	
	defaultLimit := fmt.Sprintf("%d", DEFAULT_LIMIT)
	limit, err := strconv.Atoi(c.DefaultQuery("limit", defaultLimit))

	if err != nil {
		page = DEFAULT_LIMIT
	}
	
	return page, limit
}

// Get skip value for mysql
func GetSkipValue(page int, limit int) int {
	if page < 1 {
		page = 1
	}

	return (page - 1) * limit
}