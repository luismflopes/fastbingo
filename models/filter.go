package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Filter entity
type Filters struct {
	Page         int64
	PerPage      int64
	Limit        int64
	Offset       int64
	FilterFields string
	FilterValues []interface{}
}

func ParseFilters(allowedFilters map[string]string, c *gin.Context) *Filters {
	// order_by
	// order_direction
	f := Filters{Page: 1, PerPage: 10}

	fields := []string{}
	for param, v := range c.Request.URL.Query() {
		// fmt.Println(k, v)

		if param == "page" {
			if p, err := strconv.ParseInt(v[0], 0, 64); err == nil {
				f.Page = p
			}
			continue
		}

		if param == "per_page" {
			if p, err := strconv.ParseInt(v[0], 0, 64); err == nil && p <= 100 {
				f.PerPage = p
			}
			continue
		}

		for k, operator := range allowedFilters {
			if k == param && v[0] != "" {
				fields = append(fields, fmt.Sprintf("%s %s ?", param, operator))
				f.FilterValues = append(f.FilterValues, fmt.Sprintf("%%%s%%", v[0]))
				break
			}
		}
	}

	f.FilterFields = strings.Join(fields, ", ")
	f.Limit = f.PerPage
	f.Offset = (f.Page - 1) * f.PerPage

	return &f
}
