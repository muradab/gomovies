package data

import "github.com/muradab/gomovies/internal/validator"

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
