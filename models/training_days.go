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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TrainingDay is an object representing the database table.
type TrainingDay struct {
	Date             time.Time   `boil:"date" json:"date" toml:"date" yaml:"date"`
	Maxparticipation int         `boil:"maxparticipation" json:"maxparticipation" toml:"maxparticipation" yaml:"maxparticipation"`
	Iscanceled       bool        `boil:"iscanceled" json:"iscanceled" toml:"iscanceled" yaml:"iscanceled"`
	Name             null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Shortdescription null.String `boil:"shortdescription" json:"shortdescription,omitempty" toml:"shortdescription" yaml:"shortdescription,omitempty"`
	ID               int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Text             null.String `boil:"text" json:"text,omitempty" toml:"text" yaml:"text,omitempty"`
	Createdat        time.Time   `boil:"createdat" json:"createdat" toml:"createdat" yaml:"createdat"`
	Updatedat        time.Time   `boil:"updatedat" json:"updatedat" toml:"updatedat" yaml:"updatedat"`
	TrainingID       int64       `boil:"training_id" json:"training_id" toml:"training_id" yaml:"training_id"`

	R *trainingDayR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L trainingDayL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TrainingDayColumns = struct {
	Date             string
	Maxparticipation string
	Iscanceled       string
	Name             string
	Shortdescription string
	ID               string
	Text             string
	Createdat        string
	Updatedat        string
	TrainingID       string
}{
	Date:             "date",
	Maxparticipation: "maxparticipation",
	Iscanceled:       "iscanceled",
	Name:             "name",
	Shortdescription: "shortdescription",
	ID:               "id",
	Text:             "text",
	Createdat:        "createdat",
	Updatedat:        "updatedat",
	TrainingID:       "training_id",
}

var TrainingDayTableColumns = struct {
	Date             string
	Maxparticipation string
	Iscanceled       string
	Name             string
	Shortdescription string
	ID               string
	Text             string
	Createdat        string
	Updatedat        string
	TrainingID       string
}{
	Date:             "training_days.date",
	Maxparticipation: "training_days.maxparticipation",
	Iscanceled:       "training_days.iscanceled",
	Name:             "training_days.name",
	Shortdescription: "training_days.shortdescription",
	ID:               "training_days.id",
	Text:             "training_days.text",
	Createdat:        "training_days.createdat",
	Updatedat:        "training_days.updatedat",
	TrainingID:       "training_days.training_id",
}

// Generated where

var TrainingDayWhere = struct {
	Date             whereHelpertime_Time
	Maxparticipation whereHelperint
	Iscanceled       whereHelperbool
	Name             whereHelpernull_String
	Shortdescription whereHelpernull_String
	ID               whereHelperint64
	Text             whereHelpernull_String
	Createdat        whereHelpertime_Time
	Updatedat        whereHelpertime_Time
	TrainingID       whereHelperint64
}{
	Date:             whereHelpertime_Time{field: "\"training_days\".\"date\""},
	Maxparticipation: whereHelperint{field: "\"training_days\".\"maxparticipation\""},
	Iscanceled:       whereHelperbool{field: "\"training_days\".\"iscanceled\""},
	Name:             whereHelpernull_String{field: "\"training_days\".\"name\""},
	Shortdescription: whereHelpernull_String{field: "\"training_days\".\"shortdescription\""},
	ID:               whereHelperint64{field: "\"training_days\".\"id\""},
	Text:             whereHelpernull_String{field: "\"training_days\".\"text\""},
	Createdat:        whereHelpertime_Time{field: "\"training_days\".\"createdat\""},
	Updatedat:        whereHelpertime_Time{field: "\"training_days\".\"updatedat\""},
	TrainingID:       whereHelperint64{field: "\"training_days\".\"training_id\""},
}

// TrainingDayRels is where relationship names are stored.
var TrainingDayRels = struct {
	Training                  string
	TrainingdateUserTrainings string
}{
	Training:                  "Training",
	TrainingdateUserTrainings: "TrainingdateUserTrainings",
}

// trainingDayR is where relationships are stored.
type trainingDayR struct {
	Training                  *Training         `boil:"Training" json:"Training" toml:"Training" yaml:"Training"`
	TrainingdateUserTrainings UserTrainingSlice `boil:"TrainingdateUserTrainings" json:"TrainingdateUserTrainings" toml:"TrainingdateUserTrainings" yaml:"TrainingdateUserTrainings"`
}

// NewStruct creates a new relationship struct
func (*trainingDayR) NewStruct() *trainingDayR {
	return &trainingDayR{}
}

func (r *trainingDayR) GetTraining() *Training {
	if r == nil {
		return nil
	}
	return r.Training
}

func (r *trainingDayR) GetTrainingdateUserTrainings() UserTrainingSlice {
	if r == nil {
		return nil
	}
	return r.TrainingdateUserTrainings
}

// trainingDayL is where Load methods for each relationship are stored.
type trainingDayL struct{}

var (
	trainingDayAllColumns            = []string{"date", "maxparticipation", "iscanceled", "name", "shortdescription", "id", "text", "createdat", "updatedat", "training_id"}
	trainingDayColumnsWithoutDefault = []string{"date", "training_id"}
	trainingDayColumnsWithDefault    = []string{"maxparticipation", "iscanceled", "name", "shortdescription", "id", "text", "createdat", "updatedat"}
	trainingDayPrimaryKeyColumns     = []string{"id"}
	trainingDayGeneratedColumns      = []string{}
)

type (
	// TrainingDaySlice is an alias for a slice of pointers to TrainingDay.
	// This should almost always be used instead of []TrainingDay.
	TrainingDaySlice []*TrainingDay
	// TrainingDayHook is the signature for custom TrainingDay hook methods
	TrainingDayHook func(context.Context, boil.ContextExecutor, *TrainingDay) error

	trainingDayQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	trainingDayType                 = reflect.TypeOf(&TrainingDay{})
	trainingDayMapping              = queries.MakeStructMapping(trainingDayType)
	trainingDayPrimaryKeyMapping, _ = queries.BindMapping(trainingDayType, trainingDayMapping, trainingDayPrimaryKeyColumns)
	trainingDayInsertCacheMut       sync.RWMutex
	trainingDayInsertCache          = make(map[string]insertCache)
	trainingDayUpdateCacheMut       sync.RWMutex
	trainingDayUpdateCache          = make(map[string]updateCache)
	trainingDayUpsertCacheMut       sync.RWMutex
	trainingDayUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var trainingDayAfterSelectHooks []TrainingDayHook

var trainingDayBeforeInsertHooks []TrainingDayHook
var trainingDayAfterInsertHooks []TrainingDayHook

var trainingDayBeforeUpdateHooks []TrainingDayHook
var trainingDayAfterUpdateHooks []TrainingDayHook

var trainingDayBeforeDeleteHooks []TrainingDayHook
var trainingDayAfterDeleteHooks []TrainingDayHook

var trainingDayBeforeUpsertHooks []TrainingDayHook
var trainingDayAfterUpsertHooks []TrainingDayHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TrainingDay) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingDayAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TrainingDay) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingDayBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TrainingDay) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingDayAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TrainingDay) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingDayBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TrainingDay) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingDayAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TrainingDay) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingDayBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TrainingDay) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingDayAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TrainingDay) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingDayBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TrainingDay) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingDayAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTrainingDayHook registers your hook function for all future operations.
func AddTrainingDayHook(hookPoint boil.HookPoint, trainingDayHook TrainingDayHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		trainingDayAfterSelectHooks = append(trainingDayAfterSelectHooks, trainingDayHook)
	case boil.BeforeInsertHook:
		trainingDayBeforeInsertHooks = append(trainingDayBeforeInsertHooks, trainingDayHook)
	case boil.AfterInsertHook:
		trainingDayAfterInsertHooks = append(trainingDayAfterInsertHooks, trainingDayHook)
	case boil.BeforeUpdateHook:
		trainingDayBeforeUpdateHooks = append(trainingDayBeforeUpdateHooks, trainingDayHook)
	case boil.AfterUpdateHook:
		trainingDayAfterUpdateHooks = append(trainingDayAfterUpdateHooks, trainingDayHook)
	case boil.BeforeDeleteHook:
		trainingDayBeforeDeleteHooks = append(trainingDayBeforeDeleteHooks, trainingDayHook)
	case boil.AfterDeleteHook:
		trainingDayAfterDeleteHooks = append(trainingDayAfterDeleteHooks, trainingDayHook)
	case boil.BeforeUpsertHook:
		trainingDayBeforeUpsertHooks = append(trainingDayBeforeUpsertHooks, trainingDayHook)
	case boil.AfterUpsertHook:
		trainingDayAfterUpsertHooks = append(trainingDayAfterUpsertHooks, trainingDayHook)
	}
}

// One returns a single trainingDay record from the query.
func (q trainingDayQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TrainingDay, error) {
	o := &TrainingDay{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for training_days")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TrainingDay records from the query.
func (q trainingDayQuery) All(ctx context.Context, exec boil.ContextExecutor) (TrainingDaySlice, error) {
	var o []*TrainingDay

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TrainingDay slice")
	}

	if len(trainingDayAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TrainingDay records in the query.
func (q trainingDayQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count training_days rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q trainingDayQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if training_days exists")
	}

	return count > 0, nil
}

// Training pointed to by the foreign key.
func (o *TrainingDay) Training(mods ...qm.QueryMod) trainingQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TrainingID),
	}

	queryMods = append(queryMods, mods...)

	return Trainings(queryMods...)
}

// TrainingdateUserTrainings retrieves all the user_training's UserTrainings with an executor via trainingdate_id column.
func (o *TrainingDay) TrainingdateUserTrainings(mods ...qm.QueryMod) userTrainingQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_trainings\".\"trainingdate_id\"=?", o.ID),
	)

	return UserTrainings(queryMods...)
}

// LoadTraining allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (trainingDayL) LoadTraining(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTrainingDay interface{}, mods queries.Applicator) error {
	var slice []*TrainingDay
	var object *TrainingDay

	if singular {
		var ok bool
		object, ok = maybeTrainingDay.(*TrainingDay)
		if !ok {
			object = new(TrainingDay)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTrainingDay)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTrainingDay))
			}
		}
	} else {
		s, ok := maybeTrainingDay.(*[]*TrainingDay)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTrainingDay)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTrainingDay))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &trainingDayR{}
		}
		args = append(args, object.TrainingID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &trainingDayR{}
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
		foreign.R.TrainingDays = append(foreign.R.TrainingDays, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TrainingID == foreign.ID {
				local.R.Training = foreign
				if foreign.R == nil {
					foreign.R = &trainingR{}
				}
				foreign.R.TrainingDays = append(foreign.R.TrainingDays, local)
				break
			}
		}
	}

	return nil
}

// LoadTrainingdateUserTrainings allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (trainingDayL) LoadTrainingdateUserTrainings(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTrainingDay interface{}, mods queries.Applicator) error {
	var slice []*TrainingDay
	var object *TrainingDay

	if singular {
		var ok bool
		object, ok = maybeTrainingDay.(*TrainingDay)
		if !ok {
			object = new(TrainingDay)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTrainingDay)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTrainingDay))
			}
		}
	} else {
		s, ok := maybeTrainingDay.(*[]*TrainingDay)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTrainingDay)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTrainingDay))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &trainingDayR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &trainingDayR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_trainings`),
		qm.WhereIn(`user_trainings.trainingdate_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_trainings")
	}

	var resultSlice []*UserTraining
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_trainings")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_trainings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_trainings")
	}

	if len(userTrainingAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.TrainingdateUserTrainings = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userTrainingR{}
			}
			foreign.R.Trainingdate = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TrainingdateID {
				local.R.TrainingdateUserTrainings = append(local.R.TrainingdateUserTrainings, foreign)
				if foreign.R == nil {
					foreign.R = &userTrainingR{}
				}
				foreign.R.Trainingdate = local
				break
			}
		}
	}

	return nil
}

// SetTraining of the trainingDay to the related item.
// Sets o.R.Training to related.
// Adds o to related.R.TrainingDays.
func (o *TrainingDay) SetTraining(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Training) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"training_days\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"training_id"}),
		strmangle.WhereClause("\"", "\"", 2, trainingDayPrimaryKeyColumns),
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
		o.R = &trainingDayR{
			Training: related,
		}
	} else {
		o.R.Training = related
	}

	if related.R == nil {
		related.R = &trainingR{
			TrainingDays: TrainingDaySlice{o},
		}
	} else {
		related.R.TrainingDays = append(related.R.TrainingDays, o)
	}

	return nil
}

// AddTrainingdateUserTrainings adds the given related objects to the existing relationships
// of the training_day, optionally inserting them as new records.
// Appends related to o.R.TrainingdateUserTrainings.
// Sets related.R.Trainingdate appropriately.
func (o *TrainingDay) AddTrainingdateUserTrainings(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserTraining) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TrainingdateID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_trainings\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"trainingdate_id"}),
				strmangle.WhereClause("\"", "\"", 2, userTrainingPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TrainingdateID = o.ID
		}
	}

	if o.R == nil {
		o.R = &trainingDayR{
			TrainingdateUserTrainings: related,
		}
	} else {
		o.R.TrainingdateUserTrainings = append(o.R.TrainingdateUserTrainings, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userTrainingR{
				Trainingdate: o,
			}
		} else {
			rel.R.Trainingdate = o
		}
	}
	return nil
}

// TrainingDays retrieves all the records using an executor.
func TrainingDays(mods ...qm.QueryMod) trainingDayQuery {
	mods = append(mods, qm.From("\"training_days\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"training_days\".*"})
	}

	return trainingDayQuery{q}
}

// FindTrainingDay retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTrainingDay(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TrainingDay, error) {
	trainingDayObj := &TrainingDay{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"training_days\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, trainingDayObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from training_days")
	}

	if err = trainingDayObj.doAfterSelectHooks(ctx, exec); err != nil {
		return trainingDayObj, err
	}

	return trainingDayObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TrainingDay) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no training_days provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(trainingDayColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	trainingDayInsertCacheMut.RLock()
	cache, cached := trainingDayInsertCache[key]
	trainingDayInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			trainingDayAllColumns,
			trainingDayColumnsWithDefault,
			trainingDayColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(trainingDayType, trainingDayMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(trainingDayType, trainingDayMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"training_days\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"training_days\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into training_days")
	}

	if !cached {
		trainingDayInsertCacheMut.Lock()
		trainingDayInsertCache[key] = cache
		trainingDayInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TrainingDay.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TrainingDay) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	trainingDayUpdateCacheMut.RLock()
	cache, cached := trainingDayUpdateCache[key]
	trainingDayUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			trainingDayAllColumns,
			trainingDayPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update training_days, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"training_days\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, trainingDayPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(trainingDayType, trainingDayMapping, append(wl, trainingDayPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update training_days row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for training_days")
	}

	if !cached {
		trainingDayUpdateCacheMut.Lock()
		trainingDayUpdateCache[key] = cache
		trainingDayUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q trainingDayQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for training_days")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for training_days")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TrainingDaySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), trainingDayPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"training_days\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, trainingDayPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in trainingDay slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all trainingDay")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TrainingDay) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no training_days provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(trainingDayColumnsWithDefault, o)

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

	trainingDayUpsertCacheMut.RLock()
	cache, cached := trainingDayUpsertCache[key]
	trainingDayUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			trainingDayAllColumns,
			trainingDayColumnsWithDefault,
			trainingDayColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			trainingDayAllColumns,
			trainingDayPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert training_days, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(trainingDayPrimaryKeyColumns))
			copy(conflict, trainingDayPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"training_days\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(trainingDayType, trainingDayMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(trainingDayType, trainingDayMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert training_days")
	}

	if !cached {
		trainingDayUpsertCacheMut.Lock()
		trainingDayUpsertCache[key] = cache
		trainingDayUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TrainingDay record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TrainingDay) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TrainingDay provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), trainingDayPrimaryKeyMapping)
	sql := "DELETE FROM \"training_days\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from training_days")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for training_days")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q trainingDayQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no trainingDayQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from training_days")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for training_days")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TrainingDaySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(trainingDayBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), trainingDayPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"training_days\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, trainingDayPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from trainingDay slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for training_days")
	}

	if len(trainingDayAfterDeleteHooks) != 0 {
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
func (o *TrainingDay) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTrainingDay(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TrainingDaySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TrainingDaySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), trainingDayPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"training_days\".* FROM \"training_days\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, trainingDayPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TrainingDaySlice")
	}

	*o = slice

	return nil
}

// TrainingDayExists checks if the TrainingDay row exists.
func TrainingDayExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"training_days\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if training_days exists")
	}

	return exists, nil
}

// Exists checks if the TrainingDay row exists.
func (o *TrainingDay) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TrainingDayExists(ctx, exec, o.ID)
}
