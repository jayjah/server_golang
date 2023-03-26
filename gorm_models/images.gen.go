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

func newImage(db *gorm.DB, opts ...gen.DOOption) image {
	_image := image{}

	_image.imageDo.UseDB(db, opts...)
	_image.imageDo.UseModel(&model.Image{})

	tableName := _image.imageDo.TableName()
	_image.ALL = field.NewAsterisk(tableName)
	_image.Shortdescription = field.NewString(tableName, "shortdescription")
	_image.Extracontent = field.NewString(tableName, "extracontent")
	_image.Filepath = field.NewString(tableName, "filepath")
	_image.Type = field.NewString(tableName, "type")
	_image.ID = field.NewInt64(tableName, "id")
	_image.Name = field.NewString(tableName, "name")
	_image.Text = field.NewString(tableName, "text")
	_image.Createdat = field.NewTime(tableName, "createdat")
	_image.Updatedat = field.NewTime(tableName, "updatedat")

	_image.fillFieldMap()

	return _image
}

type image struct {
	imageDo imageDo

	ALL              field.Asterisk
	Shortdescription field.String
	Extracontent     field.String
	Filepath         field.String
	Type             field.String
	ID               field.Int64
	Name             field.String
	Text             field.String
	Createdat        field.Time
	Updatedat        field.Time

	fieldMap map[string]field.Expr
}

func (i image) Table(newTableName string) *image {
	i.imageDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i image) As(alias string) *image {
	i.imageDo.DO = *(i.imageDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *image) updateTableName(table string) *image {
	i.ALL = field.NewAsterisk(table)
	i.Shortdescription = field.NewString(table, "shortdescription")
	i.Extracontent = field.NewString(table, "extracontent")
	i.Filepath = field.NewString(table, "filepath")
	i.Type = field.NewString(table, "type")
	i.ID = field.NewInt64(table, "id")
	i.Name = field.NewString(table, "name")
	i.Text = field.NewString(table, "text")
	i.Createdat = field.NewTime(table, "createdat")
	i.Updatedat = field.NewTime(table, "updatedat")

	i.fillFieldMap()

	return i
}

func (i *image) WithContext(ctx context.Context) *imageDo { return i.imageDo.WithContext(ctx) }

func (i image) TableName() string { return i.imageDo.TableName() }

func (i image) Alias() string { return i.imageDo.Alias() }

func (i *image) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *image) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 9)
	i.fieldMap["shortdescription"] = i.Shortdescription
	i.fieldMap["extracontent"] = i.Extracontent
	i.fieldMap["filepath"] = i.Filepath
	i.fieldMap["type"] = i.Type
	i.fieldMap["id"] = i.ID
	i.fieldMap["name"] = i.Name
	i.fieldMap["text"] = i.Text
	i.fieldMap["createdat"] = i.Createdat
	i.fieldMap["updatedat"] = i.Updatedat
}

func (i image) clone(db *gorm.DB) image {
	i.imageDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i image) replaceDB(db *gorm.DB) image {
	i.imageDo.ReplaceDB(db)
	return i
}

type imageDo struct{ gen.DO }

func (i imageDo) Debug() *imageDo {
	return i.withDO(i.DO.Debug())
}

func (i imageDo) WithContext(ctx context.Context) *imageDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imageDo) ReadDB() *imageDo {
	return i.Clauses(dbresolver.Read)
}

func (i imageDo) WriteDB() *imageDo {
	return i.Clauses(dbresolver.Write)
}

func (i imageDo) Session(config *gorm.Session) *imageDo {
	return i.withDO(i.DO.Session(config))
}

func (i imageDo) Clauses(conds ...clause.Expression) *imageDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imageDo) Returning(value interface{}, columns ...string) *imageDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imageDo) Not(conds ...gen.Condition) *imageDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imageDo) Or(conds ...gen.Condition) *imageDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imageDo) Select(conds ...field.Expr) *imageDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imageDo) Where(conds ...gen.Condition) *imageDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imageDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imageDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imageDo) Order(conds ...field.Expr) *imageDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imageDo) Distinct(cols ...field.Expr) *imageDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imageDo) Omit(cols ...field.Expr) *imageDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imageDo) Join(table schema.Tabler, on ...field.Expr) *imageDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imageDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imageDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imageDo) RightJoin(table schema.Tabler, on ...field.Expr) *imageDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imageDo) Group(cols ...field.Expr) *imageDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imageDo) Having(conds ...gen.Condition) *imageDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imageDo) Limit(limit int) *imageDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imageDo) Offset(offset int) *imageDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imageDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imageDo) Unscoped() *imageDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imageDo) Create(values ...*model.Image) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imageDo) CreateInBatches(values []*model.Image, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imageDo) Save(values ...*model.Image) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imageDo) First() (*model.Image, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) Take() (*model.Image, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) Last() (*model.Image, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) Find() ([]*model.Image, error) {
	result, err := i.DO.Find()
	return result.([]*model.Image), err
}

func (i imageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Image, err error) {
	buf := make([]*model.Image, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imageDo) FindInBatches(result *[]*model.Image, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imageDo) Attrs(attrs ...field.AssignExpr) *imageDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imageDo) Assign(attrs ...field.AssignExpr) *imageDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imageDo) Joins(fields ...field.RelationField) *imageDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imageDo) Preload(fields ...field.RelationField) *imageDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imageDo) FirstOrInit() (*model.Image, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) FirstOrCreate() (*model.Image, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) FindByPage(offset int, limit int) (result []*model.Image, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i imageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imageDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imageDo) Delete(models ...*model.Image) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imageDo) withDO(do gen.Dao) *imageDo {
	i.DO = *do.(*gen.DO)
	return i
}