package data

import (
	"math"
	"strings"

	"github.com/muradab/gomovies/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func ValidateFilters(v *validator.Validator, filters Filters) {
	v.Check(filters.Page > 0, "page", "must be greater than 0")
	v.Check(filters.Page <= 10_000_000, "page", "must be a maximum of 10 million")

	v.Check(filters.PageSize > 0, "page_size", "must be greater than 0")
	v.Check(filters.PageSize <= 100, "page_size", "must be a maximum of 100")

	v.Check(validator.PermittedValue(filters.Sort, filters.SortSafelist...), "sort", "invalid sort value")
}

func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}
func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}

func (f Filters) limit() int {
	return f.PageSize
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {

	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}

}
