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

func newTrainingTrainer(db *gorm.DB, opts ...gen.DOOption) trainingTrainer {
	_trainingTrainer := trainingTrainer{}

	_trainingTrainer.trainingTrainerDo.UseDB(db, opts...)
	_trainingTrainer.trainingTrainerDo.UseModel(&model.TrainingTrainer{})

	tableName := _trainingTrainer.trainingTrainerDo.TableName()
	_trainingTrainer.ALL = field.NewAsterisk(tableName)
	_trainingTrainer.ID = field.NewInt64(tableName, "id")
	_trainingTrainer.TrainerID = field.NewInt64(tableName, "trainer_id")
	_trainingTrainer.TrainingID = field.NewInt64(tableName, "training_id")

	_trainingTrainer.fillFieldMap()

	return _trainingTrainer
}

type trainingTrainer struct {
	trainingTrainerDo trainingTrainerDo

	ALL        field.Asterisk
	ID         field.Int64
	TrainerID  field.Int64
	TrainingID field.Int64

	fieldMap map[string]field.Expr
}

func (t trainingTrainer) Table(newTableName string) *trainingTrainer {
	t.trainingTrainerDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t trainingTrainer) As(alias string) *trainingTrainer {
	t.trainingTrainerDo.DO = *(t.trainingTrainerDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *trainingTrainer) updateTableName(table string) *trainingTrainer {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.TrainerID = field.NewInt64(table, "trainer_id")
	t.TrainingID = field.NewInt64(table, "training_id")

	t.fillFieldMap()

	return t
}

func (t *trainingTrainer) WithContext(ctx context.Context) *trainingTrainerDo {
	return t.trainingTrainerDo.WithContext(ctx)
}

func (t trainingTrainer) TableName() string { return t.trainingTrainerDo.TableName() }

func (t trainingTrainer) Alias() string { return t.trainingTrainerDo.Alias() }

func (t *trainingTrainer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *trainingTrainer) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["id"] = t.ID
	t.fieldMap["trainer_id"] = t.TrainerID
	t.fieldMap["training_id"] = t.TrainingID
}

func (t trainingTrainer) clone(db *gorm.DB) trainingTrainer {
	t.trainingTrainerDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t trainingTrainer) replaceDB(db *gorm.DB) trainingTrainer {
	t.trainingTrainerDo.ReplaceDB(db)
	return t
}

type trainingTrainerDo struct{ gen.DO }

func (t trainingTrainerDo) Debug() *trainingTrainerDo {
	return t.withDO(t.DO.Debug())
}

func (t trainingTrainerDo) WithContext(ctx context.Context) *trainingTrainerDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t trainingTrainerDo) ReadDB() *trainingTrainerDo {
	return t.Clauses(dbresolver.Read)
}

func (t trainingTrainerDo) WriteDB() *trainingTrainerDo {
	return t.Clauses(dbresolver.Write)
}

func (t trainingTrainerDo) Session(config *gorm.Session) *trainingTrainerDo {
	return t.withDO(t.DO.Session(config))
}

func (t trainingTrainerDo) Clauses(conds ...clause.Expression) *trainingTrainerDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t trainingTrainerDo) Returning(value interface{}, columns ...string) *trainingTrainerDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t trainingTrainerDo) Not(conds ...gen.Condition) *trainingTrainerDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t trainingTrainerDo) Or(conds ...gen.Condition) *trainingTrainerDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t trainingTrainerDo) Select(conds ...field.Expr) *trainingTrainerDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t trainingTrainerDo) Where(conds ...gen.Condition) *trainingTrainerDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t trainingTrainerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *trainingTrainerDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t trainingTrainerDo) Order(conds ...field.Expr) *trainingTrainerDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t trainingTrainerDo) Distinct(cols ...field.Expr) *trainingTrainerDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t trainingTrainerDo) Omit(cols ...field.Expr) *trainingTrainerDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t trainingTrainerDo) Join(table schema.Tabler, on ...field.Expr) *trainingTrainerDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t trainingTrainerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *trainingTrainerDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t trainingTrainerDo) RightJoin(table schema.Tabler, on ...field.Expr) *trainingTrainerDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t trainingTrainerDo) Group(cols ...field.Expr) *trainingTrainerDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t trainingTrainerDo) Having(conds ...gen.Condition) *trainingTrainerDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t trainingTrainerDo) Limit(limit int) *trainingTrainerDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t trainingTrainerDo) Offset(offset int) *trainingTrainerDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t trainingTrainerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *trainingTrainerDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t trainingTrainerDo) Unscoped() *trainingTrainerDo {
	return t.withDO(t.DO.Unscoped())
}

func (t trainingTrainerDo) Create(values ...*model.TrainingTrainer) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t trainingTrainerDo) CreateInBatches(values []*model.TrainingTrainer, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t trainingTrainerDo) Save(values ...*model.TrainingTrainer) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t trainingTrainerDo) First() (*model.TrainingTrainer, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TrainingTrainer), nil
	}
}

func (t trainingTrainerDo) Take() (*model.TrainingTrainer, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TrainingTrainer), nil
	}
}

func (t trainingTrainerDo) Last() (*model.TrainingTrainer, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TrainingTrainer), nil
	}
}

func (t trainingTrainerDo) Find() ([]*model.TrainingTrainer, error) {
	result, err := t.DO.Find()
	return result.([]*model.TrainingTrainer), err
}

func (t trainingTrainerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TrainingTrainer, err error) {
	buf := make([]*model.TrainingTrainer, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t trainingTrainerDo) FindInBatches(result *[]*model.TrainingTrainer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t trainingTrainerDo) Attrs(attrs ...field.AssignExpr) *trainingTrainerDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t trainingTrainerDo) Assign(attrs ...field.AssignExpr) *trainingTrainerDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t trainingTrainerDo) Joins(fields ...field.RelationField) *trainingTrainerDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t trainingTrainerDo) Preload(fields ...field.RelationField) *trainingTrainerDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t trainingTrainerDo) FirstOrInit() (*model.TrainingTrainer, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TrainingTrainer), nil
	}
}

func (t trainingTrainerDo) FirstOrCreate() (*model.TrainingTrainer, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TrainingTrainer), nil
	}
}

func (t trainingTrainerDo) FindByPage(offset int, limit int) (result []*model.TrainingTrainer, count int64, err error) {
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

func (t trainingTrainerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t trainingTrainerDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t trainingTrainerDo) Delete(models ...*model.TrainingTrainer) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *trainingTrainerDo) withDO(do gen.Dao) *trainingTrainerDo {
	t.DO = *do.(*gen.DO)
	return t
}
