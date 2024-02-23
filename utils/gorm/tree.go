package gorm

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type TreeOptions struct {
	DB      *gorm.DB
	Table   string
	AppId   string
	Id      string
	Preload []string
	Lazy    bool
}

func Tree[T interface{}](options TreeOptions) ([]*T, error) {
	var (
		row  T
		list []*T
		sql  *gorm.DB
	)

	sql = options.DB.Table(options.Table)
	for _, field := range options.Preload {
		sql.Preload(field)
	}
	sql.Where("app_id = ?", options.AppId)
	if options.Id != "" {
		sql.Where("id = ?", options.Id)
		if res := sql.First(&row); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return list, res.Error
		}
	}

	if !options.Lazy {
		var parentId string
		if options.Id != "" {
			parentId = fmt.Sprintf(" AND parent_id = '%s'", options.Id)
		}

		statement := fmt.Sprintf(`WITH RECURSIVE tree AS (SELECT * FROM %s WHERE app_id = '%s' %s AND deleted_at IS NULL UNION ALL SELECT t.* FROM %s t INNER JOIN tree ON tree.id = t.parent_id) SELECT * FROM tree;`, options.Table, options.AppId, parentId, options.Table)
		sql = options.DB.Raw(statement)
		for _, field := range options.Preload {
			sql.Preload(field)
		}
	}
	sql.Find(&list)

	if len(options.Id) != 0 {
		list = append(list, &row)
	}

	return list, nil
}
