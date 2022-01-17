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

// Bloodline is an object representing the database table.
type Bloodline struct {
	ID            uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name          string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	RaceID        uint      `boil:"race_id" json:"race_id" toml:"race_id" yaml:"race_id"`
	CorporationID uint      `boil:"corporation_id" json:"corporation_id" toml:"corporation_id" yaml:"corporation_id"`
	ShipTypeID    uint      `boil:"ship_type_id" json:"ship_type_id" toml:"ship_type_id" yaml:"ship_type_id"`
	Charisma      uint      `boil:"charisma" json:"charisma" toml:"charisma" yaml:"charisma"`
	Intelligence  uint      `boil:"intelligence" json:"intelligence" toml:"intelligence" yaml:"intelligence"`
	Memory        uint      `boil:"memory" json:"memory" toml:"memory" yaml:"memory"`
	Perception    uint      `boil:"perception" json:"perception" toml:"perception" yaml:"perception"`
	Willpower     uint      `boil:"willpower" json:"willpower" toml:"willpower" yaml:"willpower"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *bloodlineR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bloodlineL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BloodlineColumns = struct {
	ID            string
	Name          string
	RaceID        string
	CorporationID string
	ShipTypeID    string
	Charisma      string
	Intelligence  string
	Memory        string
	Perception    string
	Willpower     string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	Name:          "name",
	RaceID:        "race_id",
	CorporationID: "corporation_id",
	ShipTypeID:    "ship_type_id",
	Charisma:      "charisma",
	Intelligence:  "intelligence",
	Memory:        "memory",
	Perception:    "perception",
	Willpower:     "willpower",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var BloodlineTableColumns = struct {
	ID            string
	Name          string
	RaceID        string
	CorporationID string
	ShipTypeID    string
	Charisma      string
	Intelligence  string
	Memory        string
	Perception    string
	Willpower     string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "bloodlines.id",
	Name:          "bloodlines.name",
	RaceID:        "bloodlines.race_id",
	CorporationID: "bloodlines.corporation_id",
	ShipTypeID:    "bloodlines.ship_type_id",
	Charisma:      "bloodlines.charisma",
	Intelligence:  "bloodlines.intelligence",
	Memory:        "bloodlines.memory",
	Perception:    "bloodlines.perception",
	Willpower:     "bloodlines.willpower",
	CreatedAt:     "bloodlines.created_at",
	UpdatedAt:     "bloodlines.updated_at",
}

// Generated where

var BloodlineWhere = struct {
	ID            whereHelperuint
	Name          whereHelperstring
	RaceID        whereHelperuint
	CorporationID whereHelperuint
	ShipTypeID    whereHelperuint
	Charisma      whereHelperuint
	Intelligence  whereHelperuint
	Memory        whereHelperuint
	Perception    whereHelperuint
	Willpower     whereHelperuint
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
}{
	ID:            whereHelperuint{field: "`bloodlines`.`id`"},
	Name:          whereHelperstring{field: "`bloodlines`.`name`"},
	RaceID:        whereHelperuint{field: "`bloodlines`.`race_id`"},
	CorporationID: whereHelperuint{field: "`bloodlines`.`corporation_id`"},
	ShipTypeID:    whereHelperuint{field: "`bloodlines`.`ship_type_id`"},
	Charisma:      whereHelperuint{field: "`bloodlines`.`charisma`"},
	Intelligence:  whereHelperuint{field: "`bloodlines`.`intelligence`"},
	Memory:        whereHelperuint{field: "`bloodlines`.`memory`"},
	Perception:    whereHelperuint{field: "`bloodlines`.`perception`"},
	Willpower:     whereHelperuint{field: "`bloodlines`.`willpower`"},
	CreatedAt:     whereHelpertime_Time{field: "`bloodlines`.`created_at`"},
	UpdatedAt:     whereHelpertime_Time{field: "`bloodlines`.`updated_at`"},
}

// BloodlineRels is where relationship names are stored.
var BloodlineRels = struct {
}{}

// bloodlineR is where relationships are stored.
type bloodlineR struct {
}

// NewStruct creates a new relationship struct
func (*bloodlineR) NewStruct() *bloodlineR {
	return &bloodlineR{}
}

// bloodlineL is where Load methods for each relationship are stored.
type bloodlineL struct{}

var (
	bloodlineAllColumns            = []string{"id", "name", "race_id", "corporation_id", "ship_type_id", "charisma", "intelligence", "memory", "perception", "willpower", "created_at", "updated_at"}
	bloodlineColumnsWithoutDefault = []string{"id", "name", "race_id", "corporation_id", "ship_type_id", "charisma", "intelligence", "memory", "perception", "willpower", "created_at", "updated_at"}
	bloodlineColumnsWithDefault    = []string{}
	bloodlinePrimaryKeyColumns     = []string{"id"}
)

type (
	// BloodlineSlice is an alias for a slice of pointers to Bloodline.
	// This should almost always be used instead of []Bloodline.
	BloodlineSlice []*Bloodline
	// BloodlineHook is the signature for custom Bloodline hook methods
	BloodlineHook func(context.Context, boil.ContextExecutor, *Bloodline) error

	bloodlineQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bloodlineType                 = reflect.TypeOf(&Bloodline{})
	bloodlineMapping              = queries.MakeStructMapping(bloodlineType)
	bloodlinePrimaryKeyMapping, _ = queries.BindMapping(bloodlineType, bloodlineMapping, bloodlinePrimaryKeyColumns)
	bloodlineInsertCacheMut       sync.RWMutex
	bloodlineInsertCache          = make(map[string]insertCache)
	bloodlineUpdateCacheMut       sync.RWMutex
	bloodlineUpdateCache          = make(map[string]updateCache)
	bloodlineUpsertCacheMut       sync.RWMutex
	bloodlineUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bloodlineBeforeInsertHooks []BloodlineHook
var bloodlineBeforeUpdateHooks []BloodlineHook
var bloodlineBeforeDeleteHooks []BloodlineHook
var bloodlineBeforeUpsertHooks []BloodlineHook

var bloodlineAfterInsertHooks []BloodlineHook
var bloodlineAfterSelectHooks []BloodlineHook
var bloodlineAfterUpdateHooks []BloodlineHook
var bloodlineAfterDeleteHooks []BloodlineHook
var bloodlineAfterUpsertHooks []BloodlineHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Bloodline) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bloodlineBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Bloodline) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bloodlineBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Bloodline) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bloodlineBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Bloodline) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bloodlineBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Bloodline) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bloodlineAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Bloodline) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bloodlineAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Bloodline) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bloodlineAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Bloodline) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bloodlineAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Bloodline) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bloodlineAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBloodlineHook registers your hook function for all future operations.
func AddBloodlineHook(hookPoint boil.HookPoint, bloodlineHook BloodlineHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		bloodlineBeforeInsertHooks = append(bloodlineBeforeInsertHooks, bloodlineHook)
	case boil.BeforeUpdateHook:
		bloodlineBeforeUpdateHooks = append(bloodlineBeforeUpdateHooks, bloodlineHook)
	case boil.BeforeDeleteHook:
		bloodlineBeforeDeleteHooks = append(bloodlineBeforeDeleteHooks, bloodlineHook)
	case boil.BeforeUpsertHook:
		bloodlineBeforeUpsertHooks = append(bloodlineBeforeUpsertHooks, bloodlineHook)
	case boil.AfterInsertHook:
		bloodlineAfterInsertHooks = append(bloodlineAfterInsertHooks, bloodlineHook)
	case boil.AfterSelectHook:
		bloodlineAfterSelectHooks = append(bloodlineAfterSelectHooks, bloodlineHook)
	case boil.AfterUpdateHook:
		bloodlineAfterUpdateHooks = append(bloodlineAfterUpdateHooks, bloodlineHook)
	case boil.AfterDeleteHook:
		bloodlineAfterDeleteHooks = append(bloodlineAfterDeleteHooks, bloodlineHook)
	case boil.AfterUpsertHook:
		bloodlineAfterUpsertHooks = append(bloodlineAfterUpsertHooks, bloodlineHook)
	}
}

// OneG returns a single bloodline record from the query using the global executor.
func (q bloodlineQuery) OneG(ctx context.Context) (*Bloodline, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single bloodline record from the query.
func (q bloodlineQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Bloodline, error) {
	o := &Bloodline{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for bloodlines")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Bloodline records from the query using the global executor.
func (q bloodlineQuery) AllG(ctx context.Context) (BloodlineSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Bloodline records from the query.
func (q bloodlineQuery) All(ctx context.Context, exec boil.ContextExecutor) (BloodlineSlice, error) {
	var o []*Bloodline

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to Bloodline slice")
	}

	if len(bloodlineAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Bloodline records in the query, and panics on error.
func (q bloodlineQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Bloodline records in the query.
func (q bloodlineQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count bloodlines rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q bloodlineQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q bloodlineQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if bloodlines exists")
	}

	return count > 0, nil
}

// Bloodlines retrieves all the records using an executor.
func Bloodlines(mods ...qm.QueryMod) bloodlineQuery {
	mods = append(mods, qm.From("`bloodlines`"))
	return bloodlineQuery{NewQuery(mods...)}
}

// FindBloodlineG retrieves a single record by ID.
func FindBloodlineG(ctx context.Context, iD uint, selectCols ...string) (*Bloodline, error) {
	return FindBloodline(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindBloodline retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBloodline(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*Bloodline, error) {
	bloodlineObj := &Bloodline{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `bloodlines` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bloodlineObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from bloodlines")
	}

	if err = bloodlineObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bloodlineObj, err
	}

	return bloodlineObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Bloodline) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Bloodline) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no bloodlines provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(bloodlineColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bloodlineInsertCacheMut.RLock()
	cache, cached := bloodlineInsertCache[key]
	bloodlineInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bloodlineAllColumns,
			bloodlineColumnsWithDefault,
			bloodlineColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bloodlineType, bloodlineMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bloodlineType, bloodlineMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `bloodlines` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `bloodlines` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `bloodlines` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, bloodlinePrimaryKeyColumns))
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
		return errors.Wrap(err, "boiler: unable to insert into bloodlines")
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
		return errors.Wrap(err, "boiler: unable to populate default values for bloodlines")
	}

CacheNoHooks:
	if !cached {
		bloodlineInsertCacheMut.Lock()
		bloodlineInsertCache[key] = cache
		bloodlineInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Bloodline record using the global executor.
// See Update for more documentation.
func (o *Bloodline) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Bloodline.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Bloodline) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bloodlineUpdateCacheMut.RLock()
	cache, cached := bloodlineUpdateCache[key]
	bloodlineUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bloodlineAllColumns,
			bloodlinePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update bloodlines, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `bloodlines` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, bloodlinePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bloodlineType, bloodlineMapping, append(wl, bloodlinePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update bloodlines row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for bloodlines")
	}

	if !cached {
		bloodlineUpdateCacheMut.Lock()
		bloodlineUpdateCache[key] = cache
		bloodlineUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q bloodlineQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q bloodlineQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for bloodlines")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for bloodlines")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BloodlineSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BloodlineSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bloodlinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `bloodlines` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bloodlinePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in bloodline slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all bloodline")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Bloodline) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLBloodlineUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Bloodline) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no bloodlines provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(bloodlineColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBloodlineUniqueColumns, o)

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

	bloodlineUpsertCacheMut.RLock()
	cache, cached := bloodlineUpsertCache[key]
	bloodlineUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bloodlineAllColumns,
			bloodlineColumnsWithDefault,
			bloodlineColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bloodlineAllColumns,
			bloodlinePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("boiler: unable to upsert bloodlines, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`bloodlines`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `bloodlines` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(bloodlineType, bloodlineMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bloodlineType, bloodlineMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert for bloodlines")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(bloodlineType, bloodlineMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for bloodlines")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for bloodlines")
	}

CacheNoHooks:
	if !cached {
		bloodlineUpsertCacheMut.Lock()
		bloodlineUpsertCache[key] = cache
		bloodlineUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Bloodline record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Bloodline) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Bloodline record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Bloodline) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no Bloodline provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bloodlinePrimaryKeyMapping)
	sql := "DELETE FROM `bloodlines` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from bloodlines")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for bloodlines")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q bloodlineQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q bloodlineQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no bloodlineQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from bloodlines")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for bloodlines")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o BloodlineSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BloodlineSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bloodlineBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bloodlinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `bloodlines` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bloodlinePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from bloodline slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for bloodlines")
	}

	if len(bloodlineAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Bloodline) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("boiler: no Bloodline provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Bloodline) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBloodline(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BloodlineSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("boiler: empty BloodlineSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BloodlineSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BloodlineSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bloodlinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `bloodlines`.* FROM `bloodlines` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bloodlinePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in BloodlineSlice")
	}

	*o = slice

	return nil
}

// BloodlineExistsG checks if the Bloodline row exists.
func BloodlineExistsG(ctx context.Context, iD uint) (bool, error) {
	return BloodlineExists(ctx, boil.GetContextDB(), iD)
}

// BloodlineExists checks if the Bloodline row exists.
func BloodlineExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `bloodlines` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if bloodlines exists")
	}

	return exists, nil
}
