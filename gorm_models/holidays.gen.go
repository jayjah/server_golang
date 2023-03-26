// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gorm_models

import (
	"context"
	"golang_server/gorm_models/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newHoliday(db *gorm.DB, opts ...gen.DOOption) holiday {
	_holiday := holiday{}

	_holiday.holidayDo.UseDB(db, opts...)
	_holiday.holidayDo.UseModel(&model.Holiday{})

	tableName := _holiday.holidayDo.TableName()
	_holiday.ALL = field.NewAsterisk(tableName)
	_holiday.ID = field.NewInt64(tableName, "id")
	_holiday.Name = field.NewString(tableName, "name")
	_holiday.Description = field.NewString(tableName, "description")
	_holiday.Date = field.NewTime(tableName, "date")

	_holiday.fillFieldMap()

	return _holiday
}

type holiday struct {
	holidayDo holidayDo

	ALL         field.Asterisk
	ID          field.Int64
	Name        field.String
	Description field.String
	Date        field.Time

	fieldMap map[string]field.Expr
}

func (h holiday) Table(newTableName string) *holiday {
	h.holidayDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h holiday) As(alias string) *holiday {
	h.holidayDo.DO = *(h.holidayDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *holiday) updateTableName(table string) *holiday {
	h.ALL = field.NewAsterisk(table)
	h.ID = field.NewInt64(table, "id")
	h.Name = field.NewString(table, "name")
	h.Description = field.NewString(table, "description")
	h.Date = field.NewTime(table, "date")

	h.fillFieldMap()

	return h
}

func (h *holiday) WithContext(ctx context.Context) *holidayDo { return h.holidayDo.WithContext(ctx) }

func (h holiday) TableName() string { return h.holidayDo.TableName() }

func (h holiday) Alias() string { return h.holidayDo.Alias() }

func (h *holiday) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *holiday) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 4)
	h.fieldMap["id"] = h.ID
	h.fieldMap["name"] = h.Name
	h.fieldMap["description"] = h.Description
	h.fieldMap["date"] = h.Date
}

func (h holiday) clone(db *gorm.DB) holiday {
	h.holidayDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h holiday) replaceDB(db *gorm.DB) holiday {
	h.holidayDo.ReplaceDB(db)
	return h
}

type holidayDo struct{ gen.DO }

func (h holidayDo) Debug() *holidayDo {
	return h.withDO(h.DO.Debug())
}

func (h holidayDo) WithContext(ctx context.Context) *holidayDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h holidayDo) ReadDB() *holidayDo {
	return h.Clauses(dbresolver.Read)
}

func (h holidayDo) WriteDB() *holidayDo {
	return h.Clauses(dbresolver.Write)
}

func (h holidayDo) Session(config *gorm.Session) *holidayDo {
	return h.withDO(h.DO.Session(config))
}

func (h holidayDo) Clauses(conds ...clause.Expression) *holidayDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h holidayDo) Returning(value interface{}, columns ...string) *holidayDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h holidayDo) Not(conds ...gen.Condition) *holidayDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h holidayDo) Or(conds ...gen.Condition) *holidayDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h holidayDo) Select(conds ...field.Expr) *holidayDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h holidayDo) Where(conds ...gen.Condition) *holidayDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h holidayDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *holidayDo {
	return h.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (h holidayDo) Order(conds ...field.Expr) *holidayDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h holidayDo) Distinct(cols ...field.Expr) *holidayDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h holidayDo) Omit(cols ...field.Expr) *holidayDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h holidayDo) Join(table schema.Tabler, on ...field.Expr) *holidayDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h holidayDo) LeftJoin(table schema.Tabler, on ...field.Expr) *holidayDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h holidayDo) RightJoin(table schema.Tabler, on ...field.Expr) *holidayDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h holidayDo) Group(cols ...field.Expr) *holidayDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h holidayDo) Having(conds ...gen.Condition) *holidayDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h holidayDo) Limit(limit int) *holidayDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h holidayDo) Offset(offset int) *holidayDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h holidayDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *holidayDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h holidayDo) Unscoped() *holidayDo {
	return h.withDO(h.DO.Unscoped())
}

func (h holidayDo) Create(values ...*model.Holiday) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h holidayDo) CreateInBatches(values []*model.Holiday, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h holidayDo) Save(values ...*model.Holiday) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h holidayDo) First() (*model.Holiday, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Holiday), nil
	}
}

func (h holidayDo) Take() (*model.Holiday, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Holiday), nil
	}
}

func (h holidayDo) Last() (*model.Holiday, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Holiday), nil
	}
}

func (h holidayDo) Find() ([]*model.Holiday, error) {
	result, err := h.DO.Find()
	return result.([]*model.Holiday), err
}

func (h holidayDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Holiday, err error) {
	buf := make([]*model.Holiday, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h holidayDo) FindInBatches(result *[]*model.Holiday, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h holidayDo) Attrs(attrs ...field.AssignExpr) *holidayDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h holidayDo) Assign(attrs ...field.AssignExpr) *holidayDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h holidayDo) Joins(fields ...field.RelationField) *holidayDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h holidayDo) Preload(fields ...field.RelationField) *holidayDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h holidayDo) FirstOrInit() (*model.Holiday, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Holiday), nil
	}
}

func (h holidayDo) FirstOrCreate() (*model.Holiday, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Holiday), nil
	}
}

func (h holidayDo) FindByPage(offset int, limit int) (result []*model.Holiday, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h holidayDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h holidayDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h holidayDo) Delete(models ...*model.Holiday) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *holidayDo) withDO(do gen.Dao) *holidayDo {
	h.DO = *do.(*gen.DO)
	return h
}