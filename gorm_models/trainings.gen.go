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

func newTraining(db *gorm.DB, opts ...gen.DOOption) training {
	_training := training{}

	_training.trainingDo.UseDB(db, opts...)
	_training.trainingDo.UseModel(&model.Training{})

	tableName := _training.trainingDo.TableName()
	_training.ALL = field.NewAsterisk(tableName)
	_training.Isag = field.NewBool(tableName, "isag")
	_training.Timefrom = field.NewTime(tableName, "timefrom")
	_training.Timetill = field.NewTime(tableName, "timetill")
	_training.Lastcreatedtrainingdates = field.NewTime(tableName, "lastcreatedtrainingdates")
	_training.Weekday = field.NewString(tableName, "weekday")
	_training.Agefrom = field.NewInt32(tableName, "agefrom")
	_training.Agetill = field.NewInt32(tableName, "agetill")
	_training.Color = field.NewString(tableName, "color")
	_training.Deactivated = field.NewBool(tableName, "deactivated")
	_training.ID = field.NewInt64(tableName, "id")
	_training.Name = field.NewString(tableName, "name")
	_training.Shortdescription = field.NewString(tableName, "shortdescription")
	_training.Text = field.NewString(tableName, "text")
	_training.Createdat = field.NewTime(tableName, "createdat")
	_training.Updatedat = field.NewTime(tableName, "updatedat")
	_training.ImageID = field.NewInt64(tableName, "image_id")
	_training.LocationID = field.NewInt64(tableName, "location_id")
	_training.TrainerID = field.NewInt64(tableName, "trainer_id")

	_training.fillFieldMap()

	return _training
}

type training struct {
	trainingDo trainingDo

	ALL                      field.Asterisk
	Isag                     field.Bool
	Timefrom                 field.Time
	Timetill                 field.Time
	Lastcreatedtrainingdates field.Time
	Weekday                  field.String
	Agefrom                  field.Int32
	Agetill                  field.Int32
	Color                    field.String
	Deactivated              field.Bool
	ID                       field.Int64
	Name                     field.String
	Shortdescription         field.String
	Text                     field.String
	Createdat                field.Time
	Updatedat                field.Time
	ImageID                  field.Int64
	LocationID               field.Int64
	TrainerID                field.Int64

	fieldMap map[string]field.Expr
}

func (t training) Table(newTableName string) *training {
	t.trainingDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t training) As(alias string) *training {
	t.trainingDo.DO = *(t.trainingDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *training) updateTableName(table string) *training {
	t.ALL = field.NewAsterisk(table)
	t.Isag = field.NewBool(table, "isag")
	t.Timefrom = field.NewTime(table, "timefrom")
	t.Timetill = field.NewTime(table, "timetill")
	t.Lastcreatedtrainingdates = field.NewTime(table, "lastcreatedtrainingdates")
	t.Weekday = field.NewString(table, "weekday")
	t.Agefrom = field.NewInt32(table, "agefrom")
	t.Agetill = field.NewInt32(table, "agetill")
	t.Color = field.NewString(table, "color")
	t.Deactivated = field.NewBool(table, "deactivated")
	t.ID = field.NewInt64(table, "id")
	t.Name = field.NewString(table, "name")
	t.Shortdescription = field.NewString(table, "shortdescription")
	t.Text = field.NewString(table, "text")
	t.Createdat = field.NewTime(table, "createdat")
	t.Updatedat = field.NewTime(table, "updatedat")
	t.ImageID = field.NewInt64(table, "image_id")
	t.LocationID = field.NewInt64(table, "location_id")
	t.TrainerID = field.NewInt64(table, "trainer_id")

	t.fillFieldMap()

	return t
}

func (t *training) WithContext(ctx context.Context) *trainingDo { return t.trainingDo.WithContext(ctx) }

func (t training) TableName() string { return t.trainingDo.TableName() }

func (t training) Alias() string { return t.trainingDo.Alias() }

func (t *training) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *training) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 18)
	t.fieldMap["isag"] = t.Isag
	t.fieldMap["timefrom"] = t.Timefrom
	t.fieldMap["timetill"] = t.Timetill
	t.fieldMap["lastcreatedtrainingdates"] = t.Lastcreatedtrainingdates
	t.fieldMap["weekday"] = t.Weekday
	t.fieldMap["agefrom"] = t.Agefrom
	t.fieldMap["agetill"] = t.Agetill
	t.fieldMap["color"] = t.Color
	t.fieldMap["deactivated"] = t.Deactivated
	t.fieldMap["id"] = t.ID
	t.fieldMap["name"] = t.Name
	t.fieldMap["shortdescription"] = t.Shortdescription
	t.fieldMap["text"] = t.Text
	t.fieldMap["createdat"] = t.Createdat
	t.fieldMap["updatedat"] = t.Updatedat
	t.fieldMap["image_id"] = t.ImageID
	t.fieldMap["location_id"] = t.LocationID
	t.fieldMap["trainer_id"] = t.TrainerID
}

func (t training) clone(db *gorm.DB) training {
	t.trainingDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t training) replaceDB(db *gorm.DB) training {
	t.trainingDo.ReplaceDB(db)
	return t
}

type trainingDo struct{ gen.DO }

func (t trainingDo) Debug() *trainingDo {
	return t.withDO(t.DO.Debug())
}

func (t trainingDo) WithContext(ctx context.Context) *trainingDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t trainingDo) ReadDB() *trainingDo {
	return t.Clauses(dbresolver.Read)
}

func (t trainingDo) WriteDB() *trainingDo {
	return t.Clauses(dbresolver.Write)
}

func (t trainingDo) Session(config *gorm.Session) *trainingDo {
	return t.withDO(t.DO.Session(config))
}

func (t trainingDo) Clauses(conds ...clause.Expression) *trainingDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t trainingDo) Returning(value interface{}, columns ...string) *trainingDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t trainingDo) Not(conds ...gen.Condition) *trainingDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t trainingDo) Or(conds ...gen.Condition) *trainingDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t trainingDo) Select(conds ...field.Expr) *trainingDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t trainingDo) Where(conds ...gen.Condition) *trainingDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t trainingDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *trainingDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t trainingDo) Order(conds ...field.Expr) *trainingDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t trainingDo) Distinct(cols ...field.Expr) *trainingDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t trainingDo) Omit(cols ...field.Expr) *trainingDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t trainingDo) Join(table schema.Tabler, on ...field.Expr) *trainingDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t trainingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *trainingDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t trainingDo) RightJoin(table schema.Tabler, on ...field.Expr) *trainingDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t trainingDo) Group(cols ...field.Expr) *trainingDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t trainingDo) Having(conds ...gen.Condition) *trainingDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t trainingDo) Limit(limit int) *trainingDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t trainingDo) Offset(offset int) *trainingDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t trainingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *trainingDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t trainingDo) Unscoped() *trainingDo {
	return t.withDO(t.DO.Unscoped())
}

func (t trainingDo) Create(values ...*model.Training) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t trainingDo) CreateInBatches(values []*model.Training, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t trainingDo) Save(values ...*model.Training) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t trainingDo) First() (*model.Training, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Training), nil
	}
}

func (t trainingDo) Take() (*model.Training, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Training), nil
	}
}

func (t trainingDo) Last() (*model.Training, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Training), nil
	}
}

func (t trainingDo) Find() ([]*model.Training, error) {
	result, err := t.DO.Find()
	return result.([]*model.Training), err
}

func (t trainingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Training, err error) {
	buf := make([]*model.Training, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t trainingDo) FindInBatches(result *[]*model.Training, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t trainingDo) Attrs(attrs ...field.AssignExpr) *trainingDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t trainingDo) Assign(attrs ...field.AssignExpr) *trainingDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t trainingDo) Joins(fields ...field.RelationField) *trainingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t trainingDo) Preload(fields ...field.RelationField) *trainingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t trainingDo) FirstOrInit() (*model.Training, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Training), nil
	}
}

func (t trainingDo) FirstOrCreate() (*model.Training, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Training), nil
	}
}

func (t trainingDo) FindByPage(offset int, limit int) (result []*model.Training, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t trainingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t trainingDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t trainingDo) Delete(models ...*model.Training) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *trainingDo) withDO(do gen.Dao) *trainingDo {
	t.DO = *do.(*gen.DO)
	return t
}
