package mapper

import "test/cmd/model"

func ToCommonParam(offset, limit int, query, orderBy, sortBy string) model.CommonParam {
	return model.CommonParam{
		QueryBy: query,
		OrderBy: orderBy,
		SortBy:  sortBy,
		Offset:  offset,
		Limit:   limit,
	}
}
