// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiler

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

// TypeGroup is an object representing the database table.
type TypeGroup struct {
	ID         uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name       string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Published  int8      `boil:"published" json:"published" toml:"published" yaml:"published"`
	CategoryID uint      `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *typeGroupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L typeGroupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TypeGroupColumns = struct {
	ID         string
	Name       string
	Published  string
	CategoryID string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	Name:       "name",
	Published:  "published",
	CategoryID: "category_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var TypeGroupTableColumns = struct {
	ID         string
	Name       string
	Published  string
	CategoryID string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "type_groups.id",
	Name:       "type_groups.name",
	Published:  "type_groups.published",
	CategoryID: "type_groups.category_id",
	CreatedAt:  "type_groups.created_at",
	UpdatedAt:  "type_groups.updated_at",
}

// Generated where

var TypeGroupWhere = struct {
	ID         whereHelperuint
	Name       whereHelperstring
	Published  whereHelperint8
	CategoryID whereHelperuint
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	ID:         whereHelperuint{field: "`type_groups`.`id`"},
	Name:       whereHelperstring{field: "`type_groups`.`name`"},
	Published:  whereHelperint8{field: "`type_groups`.`published`"},
	CategoryID: whereHelperuint{field: "`type_groups`.`category_id`"},
	CreatedAt:  whereHelpertime_Time{field: "`type_groups`.`created_at`"},
	UpdatedAt:  whereHelpertime_Time{field: "`type_groups`.`updated_at`"},
}

// TypeGroupRels is where relationship names are stored.
var TypeGroupRels = struct {
}{}

// typeGroupR is where relationships are stored.
type typeGroupR struct {
}

// NewStruct creates a new relationship struct
func (*typeGroupR) NewStruct() *typeGroupR {
	return &typeGroupR{}
}

// typeGroupL is where Load methods for each relationship are stored.
type typeGroupL struct{}

var (
	typeGroupAllColumns            = []string{"id", "name", "published", "category_id", "created_at", "updated_at"}
	typeGroupColumnsWithoutDefault = []string{"id", "name", "published", "category_id", "created_at", "updated_at"}
	typeGroupColumnsWithDefault    = []string{}
	typeGroupPrimaryKeyColumns     = []string{"id"}
)

type (
	// TypeGroupSlice is an alias for a slice of pointers to TypeGroup.
	// This should almost always be used instead of []TypeGroup.
	TypeGroupSlice []*TypeGroup
	// TypeGroupHook is the signature for custom TypeGroup hook methods
	TypeGroupHook func(context.Context, boil.ContextExecutor, *TypeGroup) error

	typeGroupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	typeGroupType                 = reflect.TypeOf(&TypeGroup{})
	typeGroupMapping              = queries.MakeStructMapping(typeGroupType)
	typeGroupPrimaryKeyMapping, _ = queries.BindMapping(typeGroupType, typeGroupMapping, typeGroupPrimaryKeyColumns)
	typeGroupInsertCacheMut       sync.RWMutex
	typeGroupInsertCache          = make(map[string]insertCache)
	typeGroupUpdateCacheMut       sync.RWMutex
	typeGroupUpdateCache          = make(map[string]updateCache)
	typeGroupUpsertCacheMut       sync.RWMutex
	typeGroupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var typeGroupBeforeInsertHooks []TypeGroupHook
var typeGroupBeforeUpdateHooks []TypeGroupHook
var typeGroupBeforeDeleteHooks []TypeGroupHook
var typeGroupBeforeUpsertHooks []TypeGroupHook

var typeGroupAfterInsertHooks []TypeGroupHook
var typeGroupAfterSelectHooks []TypeGroupHook
var typeGroupAfterUpdateHooks []TypeGroupHook
var typeGroupAfterDeleteHooks []TypeGroupHook
var typeGroupAfterUpsertHooks []TypeGroupHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TypeGroup) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range typeGroupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TypeGroup) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range typeGroupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TypeGroup) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range typeGroupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TypeGroup) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range typeGroupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TypeGroup) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range typeGroupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TypeGroup) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range typeGroupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TypeGroup) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range typeGroupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TypeGroup) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range typeGroupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TypeGroup) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range typeGroupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTypeGroupHook registers your hook function for all future operations.
func AddTypeGroupHook(hookPoint boil.HookPoint, typeGroupHook TypeGroupHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		typeGroupBeforeInsertHooks = append(typeGroupBeforeInsertHooks, typeGroupHook)
	case boil.BeforeUpdateHook:
		typeGroupBeforeUpdateHooks = append(typeGroupBeforeUpdateHooks, typeGroupHook)
	case boil.BeforeDeleteHook:
		typeGroupBeforeDeleteHooks = append(typeGroupBeforeDeleteHooks, typeGroupHook)
	case boil.BeforeUpsertHook:
		typeGroupBeforeUpsertHooks = append(typeGroupBeforeUpsertHooks, typeGroupHook)
	case boil.AfterInsertHook:
		typeGroupAfterInsertHooks = append(typeGroupAfterInsertHooks, typeGroupHook)
	case boil.AfterSelectHook:
		typeGroupAfterSelectHooks = append(typeGroupAfterSelectHooks, typeGroupHook)
	case boil.AfterUpdateHook:
		typeGroupAfterUpdateHooks = append(typeGroupAfterUpdateHooks, typeGroupHook)
	case boil.AfterDeleteHook:
		typeGroupAfterDeleteHooks = append(typeGroupAfterDeleteHooks, typeGroupHook)
	case boil.AfterUpsertHook:
		typeGroupAfterUpsertHooks = append(typeGroupAfterUpsertHooks, typeGroupHook)
	}
}

// OneG returns a single typeGroup record from the query using the global executor.
func (q typeGroupQuery) OneG(ctx context.Context) (*TypeGroup, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single typeGroup record from the query.
func (q typeGroupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TypeGroup, error) {
	o := &TypeGroup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for type_groups")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TypeGroup records from the query using the global executor.
func (q typeGroupQuery) AllG(ctx context.Context) (TypeGroupSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TypeGroup records from the query.
func (q typeGroupQuery) All(ctx context.Context, exec boil.ContextExecutor) (TypeGroupSlice, error) {
	var o []*TypeGroup

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to TypeGroup slice")
	}

	if len(typeGroupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TypeGroup records in the query, and panics on error.
func (q typeGroupQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TypeGroup records in the query.
func (q typeGroupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count type_groups rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q typeGroupQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q typeGroupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if type_groups exists")
	}

	return count > 0, nil
}

// TypeGroups retrieves all the records using an executor.
func TypeGroups(mods ...qm.QueryMod) typeGroupQuery {
	mods = append(mods, qm.From("`type_groups`"))
	return typeGroupQuery{NewQuery(mods...)}
}

// FindTypeGroupG retrieves a single record by ID.
func FindTypeGroupG(ctx context.Context, iD uint, selectCols ...string) (*TypeGroup, error) {
	return FindTypeGroup(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTypeGroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTypeGroup(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*TypeGroup, error) {
	typeGroupObj := &TypeGroup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `type_groups` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, typeGroupObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from type_groups")
	}

	if err = typeGroupObj.doAfterSelectHooks(ctx, exec); err != nil {
		return typeGroupObj, err
	}

	return typeGroupObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TypeGroup) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TypeGroup) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no type_groups provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(typeGroupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	typeGroupInsertCacheMut.RLock()
	cache, cached := typeGroupInsertCache[key]
	typeGroupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			typeGroupAllColumns,
			typeGroupColumnsWithDefault,
			typeGroupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(typeGroupType, typeGroupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(typeGroupType, typeGroupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `type_groups` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `type_groups` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `type_groups` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, typeGroupPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "boiler: unable to insert into type_groups")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for type_groups")
	}

CacheNoHooks:
	if !cached {
		typeGroupInsertCacheMut.Lock()
		typeGroupInsertCache[key] = cache
		typeGroupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TypeGroup record using the global executor.
// See Update for more documentation.
func (o *TypeGroup) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TypeGroup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TypeGroup) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	typeGroupUpdateCacheMut.RLock()
	cache, cached := typeGroupUpdateCache[key]
	typeGroupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			typeGroupAllColumns,
			typeGroupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update type_groups, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `type_groups` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, typeGroupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(typeGroupType, typeGroupMapping, append(wl, typeGroupPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update type_groups row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for type_groups")
	}

	if !cached {
		typeGroupUpdateCacheMut.Lock()
		typeGroupUpdateCache[key] = cache
		typeGroupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q typeGroupQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q typeGroupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for type_groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for type_groups")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TypeGroupSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TypeGroupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("boiler: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), typeGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `type_groups` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, typeGroupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in typeGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all typeGroup")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TypeGroup) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLTypeGroupUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TypeGroup) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no type_groups provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(typeGroupColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTypeGroupUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	typeGroupUpsertCacheMut.RLock()
	cache, cached := typeGroupUpsertCache[key]
	typeGroupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			typeGroupAllColumns,
			typeGroupColumnsWithDefault,
			typeGroupColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			typeGroupAllColumns,
			typeGroupPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("boiler: unable to upsert type_groups, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`type_groups`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `type_groups` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(typeGroupType, typeGroupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(typeGroupType, typeGroupMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "boiler: unable to upsert for type_groups")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(typeGroupType, typeGroupMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for type_groups")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for type_groups")
	}

CacheNoHooks:
	if !cached {
		typeGroupUpsertCacheMut.Lock()
		typeGroupUpsertCache[key] = cache
		typeGroupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TypeGroup record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TypeGroup) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TypeGroup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TypeGroup) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no TypeGroup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), typeGroupPrimaryKeyMapping)
	sql := "DELETE FROM `type_groups` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from type_groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for type_groups")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q typeGroupQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q typeGroupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no typeGroupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from type_groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for type_groups")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TypeGroupSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TypeGroupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(typeGroupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), typeGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `type_groups` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, typeGroupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from typeGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for type_groups")
	}

	if len(typeGroupAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TypeGroup) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("boiler: no TypeGroup provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TypeGroup) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTypeGroup(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TypeGroupSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("boiler: empty TypeGroupSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TypeGroupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TypeGroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), typeGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `type_groups`.* FROM `type_groups` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, typeGroupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in TypeGroupSlice")
	}

	*o = slice

	return nil
}

// TypeGroupExistsG checks if the TypeGroup row exists.
func TypeGroupExistsG(ctx context.Context, iD uint) (bool, error) {
	return TypeGroupExists(ctx, boil.GetContextDB(), iD)
}

// TypeGroupExists checks if the TypeGroup row exists.
func TypeGroupExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `type_groups` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if type_groups exists")
	}

	return exists, nil
}
