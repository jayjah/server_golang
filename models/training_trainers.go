// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TrainingTrainer is an object representing the database table.
type TrainingTrainer struct {
	ID         int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	TrainerID  int64 `boil:"trainer_id" json:"trainer_id" toml:"trainer_id" yaml:"trainer_id"`
	TrainingID int64 `boil:"training_id" json:"training_id" toml:"training_id" yaml:"training_id"`

	R *trainingTrainerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L trainingTrainerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TrainingTrainerColumns = struct {
	ID         string
	TrainerID  string
	TrainingID string
}{
	ID:         "id",
	TrainerID:  "trainer_id",
	TrainingID: "training_id",
}

var TrainingTrainerTableColumns = struct {
	ID         string
	TrainerID  string
	TrainingID string
}{
	ID:         "training_trainers.id",
	TrainerID:  "training_trainers.trainer_id",
	TrainingID: "training_trainers.training_id",
}

// Generated where

var TrainingTrainerWhere = struct {
	ID         whereHelperint64
	TrainerID  whereHelperint64
	TrainingID whereHelperint64
}{
	ID:         whereHelperint64{field: "\"training_trainers\".\"id\""},
	TrainerID:  whereHelperint64{field: "\"training_trainers\".\"trainer_id\""},
	TrainingID: whereHelperint64{field: "\"training_trainers\".\"training_id\""},
}

// TrainingTrainerRels is where relationship names are stored.
var TrainingTrainerRels = struct {
	Trainer  string
	Training string
}{
	Trainer:  "Trainer",
	Training: "Training",
}

// trainingTrainerR is where relationships are stored.
type trainingTrainerR struct {
	Trainer  *Trainer  `boil:"Trainer" json:"Trainer" toml:"Trainer" yaml:"Trainer"`
	Training *Training `boil:"Training" json:"Training" toml:"Training" yaml:"Training"`
}

// NewStruct creates a new relationship struct
func (*trainingTrainerR) NewStruct() *trainingTrainerR {
	return &trainingTrainerR{}
}

func (r *trainingTrainerR) GetTrainer() *Trainer {
	if r == nil {
		return nil
	}
	return r.Trainer
}

func (r *trainingTrainerR) GetTraining() *Training {
	if r == nil {
		return nil
	}
	return r.Training
}

// trainingTrainerL is where Load methods for each relationship are stored.
type trainingTrainerL struct{}

var (
	trainingTrainerAllColumns            = []string{"id", "trainer_id", "training_id"}
	trainingTrainerColumnsWithoutDefault = []string{"trainer_id", "training_id"}
	trainingTrainerColumnsWithDefault    = []string{"id"}
	trainingTrainerPrimaryKeyColumns     = []string{"id"}
	trainingTrainerGeneratedColumns      = []string{}
)

type (
	// TrainingTrainerSlice is an alias for a slice of pointers to TrainingTrainer.
	// This should almost always be used instead of []TrainingTrainer.
	TrainingTrainerSlice []*TrainingTrainer
	// TrainingTrainerHook is the signature for custom TrainingTrainer hook methods
	TrainingTrainerHook func(context.Context, boil.ContextExecutor, *TrainingTrainer) error

	trainingTrainerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	trainingTrainerType                 = reflect.TypeOf(&TrainingTrainer{})
	trainingTrainerMapping              = queries.MakeStructMapping(trainingTrainerType)
	trainingTrainerPrimaryKeyMapping, _ = queries.BindMapping(trainingTrainerType, trainingTrainerMapping, trainingTrainerPrimaryKeyColumns)
	trainingTrainerInsertCacheMut       sync.RWMutex
	trainingTrainerInsertCache          = make(map[string]insertCache)
	trainingTrainerUpdateCacheMut       sync.RWMutex
	trainingTrainerUpdateCache          = make(map[string]updateCache)
	trainingTrainerUpsertCacheMut       sync.RWMutex
	trainingTrainerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var trainingTrainerAfterSelectHooks []TrainingTrainerHook

var trainingTrainerBeforeInsertHooks []TrainingTrainerHook
var trainingTrainerAfterInsertHooks []TrainingTrainerHook

var trainingTrainerBeforeUpdateHooks []TrainingTrainerHook
var trainingTrainerAfterUpdateHooks []TrainingTrainerHook

var trainingTrainerBeforeDeleteHooks []TrainingTrainerHook
var trainingTrainerAfterDeleteHooks []TrainingTrainerHook

var trainingTrainerBeforeUpsertHooks []TrainingTrainerHook
var trainingTrainerAfterUpsertHooks []TrainingTrainerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TrainingTrainer) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingTrainerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TrainingTrainer) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingTrainerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TrainingTrainer) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingTrainerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TrainingTrainer) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingTrainerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TrainingTrainer) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingTrainerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TrainingTrainer) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingTrainerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TrainingTrainer) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingTrainerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TrainingTrainer) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingTrainerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TrainingTrainer) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingTrainerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTrainingTrainerHook registers your hook function for all future operations.
func AddTrainingTrainerHook(hookPoint boil.HookPoint, trainingTrainerHook TrainingTrainerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		trainingTrainerAfterSelectHooks = append(trainingTrainerAfterSelectHooks, trainingTrainerHook)
	case boil.BeforeInsertHook:
		trainingTrainerBeforeInsertHooks = append(trainingTrainerBeforeInsertHooks, trainingTrainerHook)
	case boil.AfterInsertHook:
		trainingTrainerAfterInsertHooks = append(trainingTrainerAfterInsertHooks, trainingTrainerHook)
	case boil.BeforeUpdateHook:
		trainingTrainerBeforeUpdateHooks = append(trainingTrainerBeforeUpdateHooks, trainingTrainerHook)
	case boil.AfterUpdateHook:
		trainingTrainerAfterUpdateHooks = append(trainingTrainerAfterUpdateHooks, trainingTrainerHook)
	case boil.BeforeDeleteHook:
		trainingTrainerBeforeDeleteHooks = append(trainingTrainerBeforeDeleteHooks, trainingTrainerHook)
	case boil.AfterDeleteHook:
		trainingTrainerAfterDeleteHooks = append(trainingTrainerAfterDeleteHooks, trainingTrainerHook)
	case boil.BeforeUpsertHook:
		trainingTrainerBeforeUpsertHooks = append(trainingTrainerBeforeUpsertHooks, trainingTrainerHook)
	case boil.AfterUpsertHook:
		trainingTrainerAfterUpsertHooks = append(trainingTrainerAfterUpsertHooks, trainingTrainerHook)
	}
}

// One returns a single trainingTrainer record from the query.
func (q trainingTrainerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TrainingTrainer, error) {
	o := &TrainingTrainer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for training_trainers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TrainingTrainer records from the query.
func (q trainingTrainerQuery) All(ctx context.Context, exec boil.ContextExecutor) (TrainingTrainerSlice, error) {
	var o []*TrainingTrainer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TrainingTrainer slice")
	}

	if len(trainingTrainerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TrainingTrainer records in the query.
func (q trainingTrainerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count training_trainers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q trainingTrainerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if training_trainers exists")
	}

	return count > 0, nil
}

// Trainer pointed to by the foreign key.
func (o *TrainingTrainer) Trainer(mods ...qm.QueryMod) trainerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TrainerID),
	}

	queryMods = append(queryMods, mods...)

	return Trainers(queryMods...)
}

// Training pointed to by the foreign key.
func (o *TrainingTrainer) Training(mods ...qm.QueryMod) trainingQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TrainingID),
	}

	queryMods = append(queryMods, mods...)

	return Trainings(queryMods...)
}

// LoadTrainer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (trainingTrainerL) LoadTrainer(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTrainingTrainer interface{}, mods queries.Applicator) error {
	var slice []*TrainingTrainer
	var object *TrainingTrainer

	if singular {
		var ok bool
		object, ok = maybeTrainingTrainer.(*TrainingTrainer)
		if !ok {
			object = new(TrainingTrainer)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTrainingTrainer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTrainingTrainer))
			}
		}
	} else {
		s, ok := maybeTrainingTrainer.(*[]*TrainingTrainer)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTrainingTrainer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTrainingTrainer))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &trainingTrainerR{}
		}
		args = append(args, object.TrainerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &trainingTrainerR{}
			}

			for _, a := range args {
				if a == obj.TrainerID {
					continue Outer
				}
			}

			args = append(args, obj.TrainerID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`trainers`),
		qm.WhereIn(`trainers.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Trainer")
	}

	var resultSlice []*Trainer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Trainer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for trainers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for trainers")
	}

	if len(trainerAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Trainer = foreign
		if foreign.R == nil {
			foreign.R = &trainerR{}
		}
		foreign.R.TrainingTrainers = append(foreign.R.TrainingTrainers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TrainerID == foreign.ID {
				local.R.Trainer = foreign
				if foreign.R == nil {
					foreign.R = &trainerR{}
				}
				foreign.R.TrainingTrainers = append(foreign.R.TrainingTrainers, local)
				break
			}
		}
	}

	return nil
}

// LoadTraining allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (trainingTrainerL) LoadTraining(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTrainingTrainer interface{}, mods queries.Applicator) error {
	var slice []*TrainingTrainer
	var object *TrainingTrainer

	if singular {
		var ok bool
		object, ok = maybeTrainingTrainer.(*TrainingTrainer)
		if !ok {
			object = new(TrainingTrainer)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTrainingTrainer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTrainingTrainer))
			}
		}
	} else {
		s, ok := maybeTrainingTrainer.(*[]*TrainingTrainer)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTrainingTrainer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTrainingTrainer))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &trainingTrainerR{}
		}
		args = append(args, object.TrainingID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &trainingTrainerR{}
			}

			for _, a := range args {
				if a == obj.TrainingID {
					continue Outer
				}
			}

			args = append(args, obj.TrainingID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`trainings`),
		qm.WhereIn(`trainings.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Training")
	}

	var resultSlice []*Training
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Training")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for trainings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for trainings")
	}

	if len(trainingAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Training = foreign
		if foreign.R == nil {
			foreign.R = &trainingR{}
		}
		foreign.R.TrainingTrainers = append(foreign.R.TrainingTrainers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TrainingID == foreign.ID {
				local.R.Training = foreign
				if foreign.R == nil {
					foreign.R = &trainingR{}
				}
				foreign.R.TrainingTrainers = append(foreign.R.TrainingTrainers, local)
				break
			}
		}
	}

	return nil
}

// SetTrainer of the trainingTrainer to the related item.
// Sets o.R.Trainer to related.
// Adds o to related.R.TrainingTrainers.
func (o *TrainingTrainer) SetTrainer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Trainer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"training_trainers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"trainer_id"}),
		strmangle.WhereClause("\"", "\"", 2, trainingTrainerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TrainerID = related.ID
	if o.R == nil {
		o.R = &trainingTrainerR{
			Trainer: related,
		}
	} else {
		o.R.Trainer = related
	}

	if related.R == nil {
		related.R = &trainerR{
			TrainingTrainers: TrainingTrainerSlice{o},
		}
	} else {
		related.R.TrainingTrainers = append(related.R.TrainingTrainers, o)
	}

	return nil
}

// SetTraining of the trainingTrainer to the related item.
// Sets o.R.Training to related.
// Adds o to related.R.TrainingTrainers.
func (o *TrainingTrainer) SetTraining(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Training) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"training_trainers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"training_id"}),
		strmangle.WhereClause("\"", "\"", 2, trainingTrainerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TrainingID = related.ID
	if o.R == nil {
		o.R = &trainingTrainerR{
			Training: related,
		}
	} else {
		o.R.Training = related
	}

	if related.R == nil {
		related.R = &trainingR{
			TrainingTrainers: TrainingTrainerSlice{o},
		}
	} else {
		related.R.TrainingTrainers = append(related.R.TrainingTrainers, o)
	}

	return nil
}

// TrainingTrainers retrieves all the records using an executor.
func TrainingTrainers(mods ...qm.QueryMod) trainingTrainerQuery {
	mods = append(mods, qm.From("\"training_trainers\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"training_trainers\".*"})
	}

	return trainingTrainerQuery{q}
}

// FindTrainingTrainer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTrainingTrainer(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TrainingTrainer, error) {
	trainingTrainerObj := &TrainingTrainer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"training_trainers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, trainingTrainerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from training_trainers")
	}

	if err = trainingTrainerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return trainingTrainerObj, err
	}

	return trainingTrainerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TrainingTrainer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no training_trainers provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(trainingTrainerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	trainingTrainerInsertCacheMut.RLock()
	cache, cached := trainingTrainerInsertCache[key]
	trainingTrainerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			trainingTrainerAllColumns,
			trainingTrainerColumnsWithDefault,
			trainingTrainerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(trainingTrainerType, trainingTrainerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(trainingTrainerType, trainingTrainerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"training_trainers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"training_trainers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into training_trainers")
	}

	if !cached {
		trainingTrainerInsertCacheMut.Lock()
		trainingTrainerInsertCache[key] = cache
		trainingTrainerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TrainingTrainer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TrainingTrainer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	trainingTrainerUpdateCacheMut.RLock()
	cache, cached := trainingTrainerUpdateCache[key]
	trainingTrainerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			trainingTrainerAllColumns,
			trainingTrainerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update training_trainers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"training_trainers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, trainingTrainerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(trainingTrainerType, trainingTrainerMapping, append(wl, trainingTrainerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update training_trainers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for training_trainers")
	}

	if !cached {
		trainingTrainerUpdateCacheMut.Lock()
		trainingTrainerUpdateCache[key] = cache
		trainingTrainerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q trainingTrainerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for training_trainers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for training_trainers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TrainingTrainerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), trainingTrainerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"training_trainers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, trainingTrainerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in trainingTrainer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all trainingTrainer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TrainingTrainer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no training_trainers provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(trainingTrainerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	trainingTrainerUpsertCacheMut.RLock()
	cache, cached := trainingTrainerUpsertCache[key]
	trainingTrainerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			trainingTrainerAllColumns,
			trainingTrainerColumnsWithDefault,
			trainingTrainerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			trainingTrainerAllColumns,
			trainingTrainerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert training_trainers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(trainingTrainerPrimaryKeyColumns))
			copy(conflict, trainingTrainerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"training_trainers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(trainingTrainerType, trainingTrainerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(trainingTrainerType, trainingTrainerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert training_trainers")
	}

	if !cached {
		trainingTrainerUpsertCacheMut.Lock()
		trainingTrainerUpsertCache[key] = cache
		trainingTrainerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TrainingTrainer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TrainingTrainer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TrainingTrainer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), trainingTrainerPrimaryKeyMapping)
	sql := "DELETE FROM \"training_trainers\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from training_trainers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for training_trainers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q trainingTrainerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no trainingTrainerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from training_trainers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for training_trainers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TrainingTrainerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(trainingTrainerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), trainingTrainerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"training_trainers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, trainingTrainerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from trainingTrainer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for training_trainers")
	}

	if len(trainingTrainerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TrainingTrainer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTrainingTrainer(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TrainingTrainerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TrainingTrainerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), trainingTrainerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"training_trainers\".* FROM \"training_trainers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, trainingTrainerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TrainingTrainerSlice")
	}

	*o = slice

	return nil
}

// TrainingTrainerExists checks if the TrainingTrainer row exists.
func TrainingTrainerExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"training_trainers\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if training_trainers exists")
	}

	return exists, nil
}

// Exists checks if the TrainingTrainer row exists.
func (o *TrainingTrainer) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TrainingTrainerExists(ctx, exec, o.ID)
}
