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

// News is an object representing the database table.
type News struct {
	Youtubeurl       null.String `boil:"youtubeurl" json:"youtubeurl,omitempty" toml:"youtubeurl" yaml:"youtubeurl,omitempty"`
	Wordpressid      null.String `boil:"wordpressid" json:"wordpressid,omitempty" toml:"wordpressid" yaml:"wordpressid,omitempty"`
	ID               int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name             string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Shortdescription string      `boil:"shortdescription" json:"shortdescription" toml:"shortdescription" yaml:"shortdescription"`
	Text             null.String `boil:"text" json:"text,omitempty" toml:"text" yaml:"text,omitempty"`
	Createdat        time.Time   `boil:"createdat" json:"createdat" toml:"createdat" yaml:"createdat"`
	Updatedat        time.Time   `boil:"updatedat" json:"updatedat" toml:"updatedat" yaml:"updatedat"`
	ImageID          int64       `boil:"image_id" json:"image_id" toml:"image_id" yaml:"image_id"`

	R *newsR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L newsL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NewsColumns = struct {
	Youtubeurl       string
	Wordpressid      string
	ID               string
	Name             string
	Shortdescription string
	Text             string
	Createdat        string
	Updatedat        string
	ImageID          string
}{
	Youtubeurl:       "youtubeurl",
	Wordpressid:      "wordpressid",
	ID:               "id",
	Name:             "name",
	Shortdescription: "shortdescription",
	Text:             "text",
	Createdat:        "createdat",
	Updatedat:        "updatedat",
	ImageID:          "image_id",
}

var NewsTableColumns = struct {
	Youtubeurl       string
	Wordpressid      string
	ID               string
	Name             string
	Shortdescription string
	Text             string
	Createdat        string
	Updatedat        string
	ImageID          string
}{
	Youtubeurl:       "news.youtubeurl",
	Wordpressid:      "news.wordpressid",
	ID:               "news.id",
	Name:             "news.name",
	Shortdescription: "news.shortdescription",
	Text:             "news.text",
	Createdat:        "news.createdat",
	Updatedat:        "news.updatedat",
	ImageID:          "news.image_id",
}

// Generated where

var NewsWhere = struct {
	Youtubeurl       whereHelpernull_String
	Wordpressid      whereHelpernull_String
	ID               whereHelperint64
	Name             whereHelperstring
	Shortdescription whereHelperstring
	Text             whereHelpernull_String
	Createdat        whereHelpertime_Time
	Updatedat        whereHelpertime_Time
	ImageID          whereHelperint64
}{
	Youtubeurl:       whereHelpernull_String{field: "\"news\".\"youtubeurl\""},
	Wordpressid:      whereHelpernull_String{field: "\"news\".\"wordpressid\""},
	ID:               whereHelperint64{field: "\"news\".\"id\""},
	Name:             whereHelperstring{field: "\"news\".\"name\""},
	Shortdescription: whereHelperstring{field: "\"news\".\"shortdescription\""},
	Text:             whereHelpernull_String{field: "\"news\".\"text\""},
	Createdat:        whereHelpertime_Time{field: "\"news\".\"createdat\""},
	Updatedat:        whereHelpertime_Time{field: "\"news\".\"updatedat\""},
	ImageID:          whereHelperint64{field: "\"news\".\"image_id\""},
}

// NewsRels is where relationship names are stored.
var NewsRels = struct {
	Image string
}{
	Image: "Image",
}

// newsR is where relationships are stored.
type newsR struct {
	Image *Image `boil:"Image" json:"Image" toml:"Image" yaml:"Image"`
}

// NewStruct creates a new relationship struct
func (*newsR) NewStruct() *newsR {
	return &newsR{}
}

func (r *newsR) GetImage() *Image {
	if r == nil {
		return nil
	}
	return r.Image
}

// newsL is where Load methods for each relationship are stored.
type newsL struct{}

var (
	newsAllColumns            = []string{"youtubeurl", "wordpressid", "id", "name", "shortdescription", "text", "createdat", "updatedat", "image_id"}
	newsColumnsWithoutDefault = []string{"name", "shortdescription", "image_id"}
	newsColumnsWithDefault    = []string{"youtubeurl", "wordpressid", "id", "text", "createdat", "updatedat"}
	newsPrimaryKeyColumns     = []string{"id"}
	newsGeneratedColumns      = []string{}
)

type (
	// NewsSlice is an alias for a slice of pointers to News.
	// This should almost always be used instead of []News.
	NewsSlice []*News
	// NewsHook is the signature for custom News hook methods
	NewsHook func(context.Context, boil.ContextExecutor, *News) error

	newsQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	newsType                 = reflect.TypeOf(&News{})
	newsMapping              = queries.MakeStructMapping(newsType)
	newsPrimaryKeyMapping, _ = queries.BindMapping(newsType, newsMapping, newsPrimaryKeyColumns)
	newsInsertCacheMut       sync.RWMutex
	newsInsertCache          = make(map[string]insertCache)
	newsUpdateCacheMut       sync.RWMutex
	newsUpdateCache          = make(map[string]updateCache)
	newsUpsertCacheMut       sync.RWMutex
	newsUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var newsAfterSelectHooks []NewsHook

var newsBeforeInsertHooks []NewsHook
var newsAfterInsertHooks []NewsHook

var newsBeforeUpdateHooks []NewsHook
var newsAfterUpdateHooks []NewsHook

var newsBeforeDeleteHooks []NewsHook
var newsAfterDeleteHooks []NewsHook

var newsBeforeUpsertHooks []NewsHook
var newsAfterUpsertHooks []NewsHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *News) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *News) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *News) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *News) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *News) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *News) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *News) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *News) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *News) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNewsHook registers your hook function for all future operations.
func AddNewsHook(hookPoint boil.HookPoint, newsHook NewsHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		newsAfterSelectHooks = append(newsAfterSelectHooks, newsHook)
	case boil.BeforeInsertHook:
		newsBeforeInsertHooks = append(newsBeforeInsertHooks, newsHook)
	case boil.AfterInsertHook:
		newsAfterInsertHooks = append(newsAfterInsertHooks, newsHook)
	case boil.BeforeUpdateHook:
		newsBeforeUpdateHooks = append(newsBeforeUpdateHooks, newsHook)
	case boil.AfterUpdateHook:
		newsAfterUpdateHooks = append(newsAfterUpdateHooks, newsHook)
	case boil.BeforeDeleteHook:
		newsBeforeDeleteHooks = append(newsBeforeDeleteHooks, newsHook)
	case boil.AfterDeleteHook:
		newsAfterDeleteHooks = append(newsAfterDeleteHooks, newsHook)
	case boil.BeforeUpsertHook:
		newsBeforeUpsertHooks = append(newsBeforeUpsertHooks, newsHook)
	case boil.AfterUpsertHook:
		newsAfterUpsertHooks = append(newsAfterUpsertHooks, newsHook)
	}
}

// One returns a single news record from the query.
func (q newsQuery) One(ctx context.Context, exec boil.ContextExecutor) (*News, error) {
	o := &News{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for news")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all News records from the query.
func (q newsQuery) All(ctx context.Context, exec boil.ContextExecutor) (NewsSlice, error) {
	var o []*News

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to News slice")
	}

	if len(newsAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all News records in the query.
func (q newsQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count news rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q newsQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if news exists")
	}

	return count > 0, nil
}

// Image pointed to by the foreign key.
func (o *News) Image(mods ...qm.QueryMod) imageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ImageID),
	}

	queryMods = append(queryMods, mods...)

	return Images(queryMods...)
}

// LoadImage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (newsL) LoadImage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNews interface{}, mods queries.Applicator) error {
	var slice []*News
	var object *News

	if singular {
		var ok bool
		object, ok = maybeNews.(*News)
		if !ok {
			object = new(News)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeNews)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeNews))
			}
		}
	} else {
		s, ok := maybeNews.(*[]*News)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeNews)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeNews))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &newsR{}
		}
		args = append(args, object.ImageID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &newsR{}
			}

			for _, a := range args {
				if a == obj.ImageID {
					continue Outer
				}
			}

			args = append(args, obj.ImageID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`images`),
		qm.WhereIn(`images.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Image")
	}

	var resultSlice []*Image
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Image")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for images")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for images")
	}

	if len(imageAfterSelectHooks) != 0 {
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
		object.R.Image = foreign
		if foreign.R == nil {
			foreign.R = &imageR{}
		}
		foreign.R.News = append(foreign.R.News, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ImageID == foreign.ID {
				local.R.Image = foreign
				if foreign.R == nil {
					foreign.R = &imageR{}
				}
				foreign.R.News = append(foreign.R.News, local)
				break
			}
		}
	}

	return nil
}

// SetImage of the news to the related item.
// Sets o.R.Image to related.
// Adds o to related.R.News.
func (o *News) SetImage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Image) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"news\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"image_id"}),
		strmangle.WhereClause("\"", "\"", 2, newsPrimaryKeyColumns),
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

	o.ImageID = related.ID
	if o.R == nil {
		o.R = &newsR{
			Image: related,
		}
	} else {
		o.R.Image = related
	}

	if related.R == nil {
		related.R = &imageR{
			News: NewsSlice{o},
		}
	} else {
		related.R.News = append(related.R.News, o)
	}

	return nil
}

// News retrieves all the records using an executor.
func AllNews(mods ...qm.QueryMod) newsQuery {
	mods = append(mods, qm.From("\"news\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"news\".*"})
	}

	return newsQuery{q}
}

// FindNews retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNews(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*News, error) {
	newsObj := &News{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"news\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, newsObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from news")
	}

	if err = newsObj.doAfterSelectHooks(ctx, exec); err != nil {
		return newsObj, err
	}

	return newsObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *News) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no news provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(newsColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	newsInsertCacheMut.RLock()
	cache, cached := newsInsertCache[key]
	newsInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			newsAllColumns,
			newsColumnsWithDefault,
			newsColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(newsType, newsMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(newsType, newsMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"news\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"news\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into news")
	}

	if !cached {
		newsInsertCacheMut.Lock()
		newsInsertCache[key] = cache
		newsInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the News.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *News) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	newsUpdateCacheMut.RLock()
	cache, cached := newsUpdateCache[key]
	newsUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			newsAllColumns,
			newsPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update news, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"news\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, newsPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(newsType, newsMapping, append(wl, newsPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update news row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for news")
	}

	if !cached {
		newsUpdateCacheMut.Lock()
		newsUpdateCache[key] = cache
		newsUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q newsQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for news")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for news")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NewsSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), newsPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"news\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, newsPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in news slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all news")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *News) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no news provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(newsColumnsWithDefault, o)

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

	newsUpsertCacheMut.RLock()
	cache, cached := newsUpsertCache[key]
	newsUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			newsAllColumns,
			newsColumnsWithDefault,
			newsColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			newsAllColumns,
			newsPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert news, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(newsPrimaryKeyColumns))
			copy(conflict, newsPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"news\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(newsType, newsMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(newsType, newsMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert news")
	}

	if !cached {
		newsUpsertCacheMut.Lock()
		newsUpsertCache[key] = cache
		newsUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single News record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *News) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no News provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), newsPrimaryKeyMapping)
	sql := "DELETE FROM \"news\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from news")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for news")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q newsQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no newsQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from news")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for news")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NewsSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(newsBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), newsPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"news\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, newsPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from news slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for news")
	}

	if len(newsAfterDeleteHooks) != 0 {
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
func (o *News) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNews(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NewsSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NewsSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), newsPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"news\".* FROM \"news\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, newsPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NewsSlice")
	}

	*o = slice

	return nil
}

// NewsExists checks if the News row exists.
func NewsExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"news\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if news exists")
	}

	return exists, nil
}

// Exists checks if the News row exists.
func (o *News) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return NewsExists(ctx, exec, o.ID)
}
