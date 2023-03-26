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

func newUserTraining(db *gorm.DB, opts ...gen.DOOption) userTraining {
	_userTraining := userTraining{}

	_userTraining.userTrainingDo.UseDB(db, opts...)
	_userTraining.userTrainingDo.UseModel(&model.UserTraining{})

	tableName := _userTraining.userTrainingDo.TableName()
	_userTraining.ALL = field.NewAsterisk(tableName)
	_userTraining.ID = field.NewInt64(tableName, "id")
	_userTraining.Validatedfromuser = field.NewBool(tableName, "validatedfromuser")
	_userTraining.UserID = field.NewInt64(tableName, "user_id")
	_userTraining.TrainingdateID = field.NewInt64(tableName, "trainingdate_id")

	_userTraining.fillFieldMap()

	return _userTraining
}

type userTraining struct {
	userTrainingDo userTrainingDo

	ALL               field.Asterisk
	ID                field.Int64
	Validatedfromuser field.Bool
	UserID            field.Int64
	TrainingdateID    field.Int64

	fieldMap map[string]field.Expr
}

func (u userTraining) Table(newTableName string) *userTraining {
	u.userTrainingDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userTraining) As(alias string) *userTraining {
	u.userTrainingDo.DO = *(u.userTrainingDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userTraining) updateTableName(table string) *userTraining {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.Validatedfromuser = field.NewBool(table, "validatedfromuser")
	u.UserID = field.NewInt64(table, "user_id")
	u.TrainingdateID = field.NewInt64(table, "trainingdate_id")

	u.fillFieldMap()

	return u
}

func (u *userTraining) WithContext(ctx context.Context) *userTrainingDo {
	return u.userTrainingDo.WithContext(ctx)
}

func (u userTraining) TableName() string { return u.userTrainingDo.TableName() }

func (u userTraining) Alias() string { return u.userTrainingDo.Alias() }

func (u *userTraining) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userTraining) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["id"] = u.ID
	u.fieldMap["validatedfromuser"] = u.Validatedfromuser
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["trainingdate_id"] = u.TrainingdateID
}

func (u userTraining) clone(db *gorm.DB) userTraining {
	u.userTrainingDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userTraining) replaceDB(db *gorm.DB) userTraining {
	u.userTrainingDo.ReplaceDB(db)
	return u
}

type userTrainingDo struct{ gen.DO }

func (u userTrainingDo) Debug() *userTrainingDo {
	return u.withDO(u.DO.Debug())
}

func (u userTrainingDo) WithContext(ctx context.Context) *userTrainingDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userTrainingDo) ReadDB() *userTrainingDo {
	return u.Clauses(dbresolver.Read)
}

func (u userTrainingDo) WriteDB() *userTrainingDo {
	return u.Clauses(dbresolver.Write)
}

func (u userTrainingDo) Session(config *gorm.Session) *userTrainingDo {
	return u.withDO(u.DO.Session(config))
}

func (u userTrainingDo) Clauses(conds ...clause.Expression) *userTrainingDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userTrainingDo) Returning(value interface{}, columns ...string) *userTrainingDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userTrainingDo) Not(conds ...gen.Condition) *userTrainingDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userTrainingDo) Or(conds ...gen.Condition) *userTrainingDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userTrainingDo) Select(conds ...field.Expr) *userTrainingDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userTrainingDo) Where(conds ...gen.Condition) *userTrainingDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userTrainingDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userTrainingDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userTrainingDo) Order(conds ...field.Expr) *userTrainingDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userTrainingDo) Distinct(cols ...field.Expr) *userTrainingDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userTrainingDo) Omit(cols ...field.Expr) *userTrainingDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userTrainingDo) Join(table schema.Tabler, on ...field.Expr) *userTrainingDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userTrainingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userTrainingDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userTrainingDo) RightJoin(table schema.Tabler, on ...field.Expr) *userTrainingDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userTrainingDo) Group(cols ...field.Expr) *userTrainingDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userTrainingDo) Having(conds ...gen.Condition) *userTrainingDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userTrainingDo) Limit(limit int) *userTrainingDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userTrainingDo) Offset(offset int) *userTrainingDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userTrainingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userTrainingDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userTrainingDo) Unscoped() *userTrainingDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userTrainingDo) Create(values ...*model.UserTraining) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userTrainingDo) CreateInBatches(values []*model.UserTraining, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userTrainingDo) Save(values ...*model.UserTraining) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userTrainingDo) First() (*model.UserTraining, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTraining), nil
	}
}

func (u userTrainingDo) Take() (*model.UserTraining, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTraining), nil
	}
}

func (u userTrainingDo) Last() (*model.UserTraining, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTraining), nil
	}
}

func (u userTrainingDo) Find() ([]*model.UserTraining, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserTraining), err
}

func (u userTrainingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserTraining, err error) {
	buf := make([]*model.UserTraining, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userTrainingDo) FindInBatches(result *[]*model.UserTraining, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userTrainingDo) Attrs(attrs ...field.AssignExpr) *userTrainingDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userTrainingDo) Assign(attrs ...field.AssignExpr) *userTrainingDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userTrainingDo) Joins(fields ...field.RelationField) *userTrainingDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userTrainingDo) Preload(fields ...field.RelationField) *userTrainingDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userTrainingDo) FirstOrInit() (*model.UserTraining, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTraining), nil
	}
}

func (u userTrainingDo) FirstOrCreate() (*model.UserTraining, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTraining), nil
	}
}

func (u userTrainingDo) FindByPage(offset int, limit int) (result []*model.UserTraining, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userTrainingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userTrainingDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userTrainingDo) Delete(models ...*model.UserTraining) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userTrainingDo) withDO(do gen.Dao) *userTrainingDo {
	u.DO = *do.(*gen.DO)
	return u
}