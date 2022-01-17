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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Alliance is an object representing the database table.
type Alliance struct {
	ID                    uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name                  string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Ticker                string    `boil:"ticker" json:"ticker" toml:"ticker" yaml:"ticker"`
	CreatorID             uint      `boil:"creator_id" json:"creator_id" toml:"creator_id" yaml:"creator_id"`
	CreatorCorporationID  uint      `boil:"creator_corporation_id" json:"creator_corporation_id" toml:"creator_corporation_id" yaml:"creator_corporation_id"`
	ExecutorCorporationID null.Uint `boil:"executor_corporation_id" json:"executor_corporation_id,omitempty" toml:"executor_corporation_id" yaml:"executor_corporation_id,omitempty"`
	IsClosed              int8      `boil:"is_closed" json:"is_closed" toml:"is_closed" yaml:"is_closed"`
	DateFounded           null.Time `boil:"date_founded" json:"date_founded,omitempty" toml:"date_founded" yaml:"date_founded,omitempty"`
	CreatedAt             time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt             time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *allianceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L allianceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AllianceColumns = struct {
	ID                    string
	Name                  string
	Ticker                string
	CreatorID             string
	CreatorCorporationID  string
	ExecutorCorporationID string
	IsClosed              string
	DateFounded           string
	CreatedAt             string
	UpdatedAt             string
}{
	ID:                    "id",
	Name:                  "name",
	Ticker:                "ticker",
	CreatorID:             "creator_id",
	CreatorCorporationID:  "creator_corporation_id",
	ExecutorCorporationID: "executor_corporation_id",
	IsClosed:              "is_closed",
	DateFounded:           "date_founded",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

var AllianceTableColumns = struct {
	ID                    string
	Name                  string
	Ticker                string
	CreatorID             string
	CreatorCorporationID  string
	ExecutorCorporationID string
	IsClosed              string
	DateFounded           string
	CreatedAt             string
	UpdatedAt             string
}{
	ID:                    "alliances.id",
	Name:                  "alliances.name",
	Ticker:                "alliances.ticker",
	CreatorID:             "alliances.creator_id",
	CreatorCorporationID:  "alliances.creator_corporation_id",
	ExecutorCorporationID: "alliances.executor_corporation_id",
	IsClosed:              "alliances.is_closed",
	DateFounded:           "alliances.date_founded",
	CreatedAt:             "alliances.created_at",
	UpdatedAt:             "alliances.updated_at",
}

// Generated where

type whereHelperuint struct{ field string }

func (w whereHelperuint) EQ(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint) NEQ(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint) LT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint) LTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint) GT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint) GTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint) IN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint) NIN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Uint struct{ field string }

func (w whereHelpernull_Uint) EQ(x null.Uint) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Uint) NEQ(x null.Uint) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Uint) LT(x null.Uint) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Uint) LTE(x null.Uint) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Uint) GT(x null.Uint) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Uint) GTE(x null.Uint) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Uint) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Uint) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperint8 struct{ field string }

func (w whereHelperint8) EQ(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint8) NEQ(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint8) LT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint8) LTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint8) GT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint8) GTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint8) IN(slice []int8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint8) NIN(slice []int8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AllianceWhere = struct {
	ID                    whereHelperuint
	Name                  whereHelperstring
	Ticker                whereHelperstring
	CreatorID             whereHelperuint
	CreatorCorporationID  whereHelperuint
	ExecutorCorporationID whereHelpernull_Uint
	IsClosed              whereHelperint8
	DateFounded           whereHelpernull_Time
	CreatedAt             whereHelpertime_Time
	UpdatedAt             whereHelpertime_Time
}{
	ID:                    whereHelperuint{field: "`alliances`.`id`"},
	Name:                  whereHelperstring{field: "`alliances`.`name`"},
	Ticker:                whereHelperstring{field: "`alliances`.`ticker`"},
	CreatorID:             whereHelperuint{field: "`alliances`.`creator_id`"},
	CreatorCorporationID:  whereHelperuint{field: "`alliances`.`creator_corporation_id`"},
	ExecutorCorporationID: whereHelpernull_Uint{field: "`alliances`.`executor_corporation_id`"},
	IsClosed:              whereHelperint8{field: "`alliances`.`is_closed`"},
	DateFounded:           whereHelpernull_Time{field: "`alliances`.`date_founded`"},
	CreatedAt:             whereHelpertime_Time{field: "`alliances`.`created_at`"},
	UpdatedAt:             whereHelpertime_Time{field: "`alliances`.`updated_at`"},
}

// AllianceRels is where relationship names are stored.
var AllianceRels = struct {
}{}

// allianceR is where relationships are stored.
type allianceR struct {
}

// NewStruct creates a new relationship struct
func (*allianceR) NewStruct() *allianceR {
	return &allianceR{}
}

// allianceL is where Load methods for each relationship are stored.
type allianceL struct{}

var (
	allianceAllColumns            = []string{"id", "name", "ticker", "creator_id", "creator_corporation_id", "executor_corporation_id", "is_closed", "date_founded", "created_at", "updated_at"}
	allianceColumnsWithoutDefault = []string{"id", "name", "ticker", "creator_id", "creator_corporation_id", "executor_corporation_id", "date_founded", "created_at", "updated_at"}
	allianceColumnsWithDefault    = []string{"is_closed"}
	alliancePrimaryKeyColumns     = []string{"id"}
)

type (
	// AllianceSlice is an alias for a slice of pointers to Alliance.
	// This should almost always be used instead of []Alliance.
	AllianceSlice []*Alliance
	// AllianceHook is the signature for custom Alliance hook methods
	AllianceHook func(context.Context, boil.ContextExecutor, *Alliance) error

	allianceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	allianceType                 = reflect.TypeOf(&Alliance{})
	allianceMapping              = queries.MakeStructMapping(allianceType)
	alliancePrimaryKeyMapping, _ = queries.BindMapping(allianceType, allianceMapping, alliancePrimaryKeyColumns)
	allianceInsertCacheMut       sync.RWMutex
	allianceInsertCache          = make(map[string]insertCache)
	allianceUpdateCacheMut       sync.RWMutex
	allianceUpdateCache          = make(map[string]updateCache)
	allianceUpsertCacheMut       sync.RWMutex
	allianceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var allianceBeforeInsertHooks []AllianceHook
var allianceBeforeUpdateHooks []AllianceHook
var allianceBeforeDeleteHooks []AllianceHook
var allianceBeforeUpsertHooks []AllianceHook

var allianceAfterInsertHooks []AllianceHook
var allianceAfterSelectHooks []AllianceHook
var allianceAfterUpdateHooks []AllianceHook
var allianceAfterDeleteHooks []AllianceHook
var allianceAfterUpsertHooks []AllianceHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Alliance) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allianceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Alliance) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allianceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Alliance) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allianceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Alliance) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allianceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Alliance) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allianceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Alliance) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allianceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Alliance) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allianceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Alliance) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allianceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Alliance) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allianceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAllianceHook registers your hook function for all future operations.
func AddAllianceHook(hookPoint boil.HookPoint, allianceHook AllianceHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		allianceBeforeInsertHooks = append(allianceBeforeInsertHooks, allianceHook)
	case boil.BeforeUpdateHook:
		allianceBeforeUpdateHooks = append(allianceBeforeUpdateHooks, allianceHook)
	case boil.BeforeDeleteHook:
		allianceBeforeDeleteHooks = append(allianceBeforeDeleteHooks, allianceHook)
	case boil.BeforeUpsertHook:
		allianceBeforeUpsertHooks = append(allianceBeforeUpsertHooks, allianceHook)
	case boil.AfterInsertHook:
		allianceAfterInsertHooks = append(allianceAfterInsertHooks, allianceHook)
	case boil.AfterSelectHook:
		allianceAfterSelectHooks = append(allianceAfterSelectHooks, allianceHook)
	case boil.AfterUpdateHook:
		allianceAfterUpdateHooks = append(allianceAfterUpdateHooks, allianceHook)
	case boil.AfterDeleteHook:
		allianceAfterDeleteHooks = append(allianceAfterDeleteHooks, allianceHook)
	case boil.AfterUpsertHook:
		allianceAfterUpsertHooks = append(allianceAfterUpsertHooks, allianceHook)
	}
}

// OneG returns a single alliance record from the query using the global executor.
func (q allianceQuery) OneG(ctx context.Context) (*Alliance, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single alliance record from the query.
func (q allianceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Alliance, error) {
	o := &Alliance{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for alliances")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Alliance records from the query using the global executor.
func (q allianceQuery) AllG(ctx context.Context) (AllianceSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Alliance records from the query.
func (q allianceQuery) All(ctx context.Context, exec boil.ContextExecutor) (AllianceSlice, error) {
	var o []*Alliance

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to Alliance slice")
	}

	if len(allianceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Alliance records in the query, and panics on error.
func (q allianceQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Alliance records in the query.
func (q allianceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count alliances rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q allianceQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q allianceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if alliances exists")
	}

	return count > 0, nil
}

// Alliances retrieves all the records using an executor.
func Alliances(mods ...qm.QueryMod) allianceQuery {
	mods = append(mods, qm.From("`alliances`"))
	return allianceQuery{NewQuery(mods...)}
}

// FindAllianceG retrieves a single record by ID.
func FindAllianceG(ctx context.Context, iD uint, selectCols ...string) (*Alliance, error) {
	return FindAlliance(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindAlliance retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAlliance(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*Alliance, error) {
	allianceObj := &Alliance{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `alliances` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, allianceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from alliances")
	}

	if err = allianceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return allianceObj, err
	}

	return allianceObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Alliance) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Alliance) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no alliances provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(allianceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	allianceInsertCacheMut.RLock()
	cache, cached := allianceInsertCache[key]
	allianceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			allianceAllColumns,
			allianceColumnsWithDefault,
			allianceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(allianceType, allianceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(allianceType, allianceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `alliances` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `alliances` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `alliances` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, alliancePrimaryKeyColumns))
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
		return errors.Wrap(err, "boiler: unable to insert into alliances")
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
		return errors.Wrap(err, "boiler: unable to populate default values for alliances")
	}

CacheNoHooks:
	if !cached {
		allianceInsertCacheMut.Lock()
		allianceInsertCache[key] = cache
		allianceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Alliance record using the global executor.
// See Update for more documentation.
func (o *Alliance) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Alliance.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Alliance) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	allianceUpdateCacheMut.RLock()
	cache, cached := allianceUpdateCache[key]
	allianceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			allianceAllColumns,
			alliancePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update alliances, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `alliances` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, alliancePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(allianceType, allianceMapping, append(wl, alliancePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update alliances row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for alliances")
	}

	if !cached {
		allianceUpdateCacheMut.Lock()
		allianceUpdateCache[key] = cache
		allianceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q allianceQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q allianceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for alliances")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for alliances")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AllianceSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AllianceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), alliancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `alliances` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, alliancePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in alliance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all alliance")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Alliance) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLAllianceUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Alliance) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no alliances provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(allianceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAllianceUniqueColumns, o)

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

	allianceUpsertCacheMut.RLock()
	cache, cached := allianceUpsertCache[key]
	allianceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			allianceAllColumns,
			allianceColumnsWithDefault,
			allianceColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			allianceAllColumns,
			alliancePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("boiler: unable to upsert alliances, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`alliances`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `alliances` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(allianceType, allianceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(allianceType, allianceMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert for alliances")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(allianceType, allianceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for alliances")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for alliances")
	}

CacheNoHooks:
	if !cached {
		allianceUpsertCacheMut.Lock()
		allianceUpsertCache[key] = cache
		allianceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Alliance record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Alliance) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Alliance record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Alliance) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no Alliance provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), alliancePrimaryKeyMapping)
	sql := "DELETE FROM `alliances` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from alliances")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for alliances")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q allianceQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q allianceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no allianceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from alliances")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for alliances")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AllianceSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AllianceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(allianceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), alliancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `alliances` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, alliancePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from alliance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for alliances")
	}

	if len(allianceAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Alliance) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("boiler: no Alliance provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Alliance) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAlliance(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AllianceSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("boiler: empty AllianceSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AllianceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AllianceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), alliancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `alliances`.* FROM `alliances` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, alliancePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in AllianceSlice")
	}

	*o = slice

	return nil
}

// AllianceExistsG checks if the Alliance row exists.
func AllianceExistsG(ctx context.Context, iD uint) (bool, error) {
	return AllianceExists(ctx, boil.GetContextDB(), iD)
}

// AllianceExists checks if the Alliance row exists.
func AllianceExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `alliances` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if alliances exists")
	}

	return exists, nil
}
