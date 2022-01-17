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

// CharacterSkill is an object representing the database table.
type CharacterSkill struct {
	CharacterID        uint64    `boil:"character_id" json:"character_id" toml:"character_id" yaml:"character_id"`
	SkillID            uint      `boil:"skill_id" json:"skill_id" toml:"skill_id" yaml:"skill_id"`
	ActiveSkillLevel   uint8     `boil:"active_skill_level" json:"active_skill_level" toml:"active_skill_level" yaml:"active_skill_level"`
	SkillpointsInSkill uint      `boil:"skillpoints_in_skill" json:"skillpoints_in_skill" toml:"skillpoints_in_skill" yaml:"skillpoints_in_skill"`
	TrainedSkillLevel  uint8     `boil:"trained_skill_level" json:"trained_skill_level" toml:"trained_skill_level" yaml:"trained_skill_level"`
	CreatedAt          time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt          time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *characterSkillR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L characterSkillL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CharacterSkillColumns = struct {
	CharacterID        string
	SkillID            string
	ActiveSkillLevel   string
	SkillpointsInSkill string
	TrainedSkillLevel  string
	CreatedAt          string
	UpdatedAt          string
}{
	CharacterID:        "character_id",
	SkillID:            "skill_id",
	ActiveSkillLevel:   "active_skill_level",
	SkillpointsInSkill: "skillpoints_in_skill",
	TrainedSkillLevel:  "trained_skill_level",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

var CharacterSkillTableColumns = struct {
	CharacterID        string
	SkillID            string
	ActiveSkillLevel   string
	SkillpointsInSkill string
	TrainedSkillLevel  string
	CreatedAt          string
	UpdatedAt          string
}{
	CharacterID:        "character_skills.character_id",
	SkillID:            "character_skills.skill_id",
	ActiveSkillLevel:   "character_skills.active_skill_level",
	SkillpointsInSkill: "character_skills.skillpoints_in_skill",
	TrainedSkillLevel:  "character_skills.trained_skill_level",
	CreatedAt:          "character_skills.created_at",
	UpdatedAt:          "character_skills.updated_at",
}

// Generated where

var CharacterSkillWhere = struct {
	CharacterID        whereHelperuint64
	SkillID            whereHelperuint
	ActiveSkillLevel   whereHelperuint8
	SkillpointsInSkill whereHelperuint
	TrainedSkillLevel  whereHelperuint8
	CreatedAt          whereHelpertime_Time
	UpdatedAt          whereHelpertime_Time
}{
	CharacterID:        whereHelperuint64{field: "`character_skills`.`character_id`"},
	SkillID:            whereHelperuint{field: "`character_skills`.`skill_id`"},
	ActiveSkillLevel:   whereHelperuint8{field: "`character_skills`.`active_skill_level`"},
	SkillpointsInSkill: whereHelperuint{field: "`character_skills`.`skillpoints_in_skill`"},
	TrainedSkillLevel:  whereHelperuint8{field: "`character_skills`.`trained_skill_level`"},
	CreatedAt:          whereHelpertime_Time{field: "`character_skills`.`created_at`"},
	UpdatedAt:          whereHelpertime_Time{field: "`character_skills`.`updated_at`"},
}

// CharacterSkillRels is where relationship names are stored.
var CharacterSkillRels = struct {
	Character string
}{
	Character: "Character",
}

// characterSkillR is where relationships are stored.
type characterSkillR struct {
	Character *User `boil:"Character" json:"Character" toml:"Character" yaml:"Character"`
}

// NewStruct creates a new relationship struct
func (*characterSkillR) NewStruct() *characterSkillR {
	return &characterSkillR{}
}

// characterSkillL is where Load methods for each relationship are stored.
type characterSkillL struct{}

var (
	characterSkillAllColumns            = []string{"character_id", "skill_id", "active_skill_level", "skillpoints_in_skill", "trained_skill_level", "created_at", "updated_at"}
	characterSkillColumnsWithoutDefault = []string{"character_id", "skill_id", "active_skill_level", "skillpoints_in_skill", "trained_skill_level", "created_at", "updated_at"}
	characterSkillColumnsWithDefault    = []string{}
	characterSkillPrimaryKeyColumns     = []string{"character_id", "skill_id"}
)

type (
	// CharacterSkillSlice is an alias for a slice of pointers to CharacterSkill.
	// This should almost always be used instead of []CharacterSkill.
	CharacterSkillSlice []*CharacterSkill
	// CharacterSkillHook is the signature for custom CharacterSkill hook methods
	CharacterSkillHook func(context.Context, boil.ContextExecutor, *CharacterSkill) error

	characterSkillQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	characterSkillType                 = reflect.TypeOf(&CharacterSkill{})
	characterSkillMapping              = queries.MakeStructMapping(characterSkillType)
	characterSkillPrimaryKeyMapping, _ = queries.BindMapping(characterSkillType, characterSkillMapping, characterSkillPrimaryKeyColumns)
	characterSkillInsertCacheMut       sync.RWMutex
	characterSkillInsertCache          = make(map[string]insertCache)
	characterSkillUpdateCacheMut       sync.RWMutex
	characterSkillUpdateCache          = make(map[string]updateCache)
	characterSkillUpsertCacheMut       sync.RWMutex
	characterSkillUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var characterSkillBeforeInsertHooks []CharacterSkillHook
var characterSkillBeforeUpdateHooks []CharacterSkillHook
var characterSkillBeforeDeleteHooks []CharacterSkillHook
var characterSkillBeforeUpsertHooks []CharacterSkillHook

var characterSkillAfterInsertHooks []CharacterSkillHook
var characterSkillAfterSelectHooks []CharacterSkillHook
var characterSkillAfterUpdateHooks []CharacterSkillHook
var characterSkillAfterDeleteHooks []CharacterSkillHook
var characterSkillAfterUpsertHooks []CharacterSkillHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CharacterSkill) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterSkillBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CharacterSkill) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterSkillBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CharacterSkill) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterSkillBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CharacterSkill) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterSkillBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CharacterSkill) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterSkillAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CharacterSkill) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterSkillAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CharacterSkill) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterSkillAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CharacterSkill) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterSkillAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CharacterSkill) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterSkillAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCharacterSkillHook registers your hook function for all future operations.
func AddCharacterSkillHook(hookPoint boil.HookPoint, characterSkillHook CharacterSkillHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		characterSkillBeforeInsertHooks = append(characterSkillBeforeInsertHooks, characterSkillHook)
	case boil.BeforeUpdateHook:
		characterSkillBeforeUpdateHooks = append(characterSkillBeforeUpdateHooks, characterSkillHook)
	case boil.BeforeDeleteHook:
		characterSkillBeforeDeleteHooks = append(characterSkillBeforeDeleteHooks, characterSkillHook)
	case boil.BeforeUpsertHook:
		characterSkillBeforeUpsertHooks = append(characterSkillBeforeUpsertHooks, characterSkillHook)
	case boil.AfterInsertHook:
		characterSkillAfterInsertHooks = append(characterSkillAfterInsertHooks, characterSkillHook)
	case boil.AfterSelectHook:
		characterSkillAfterSelectHooks = append(characterSkillAfterSelectHooks, characterSkillHook)
	case boil.AfterUpdateHook:
		characterSkillAfterUpdateHooks = append(characterSkillAfterUpdateHooks, characterSkillHook)
	case boil.AfterDeleteHook:
		characterSkillAfterDeleteHooks = append(characterSkillAfterDeleteHooks, characterSkillHook)
	case boil.AfterUpsertHook:
		characterSkillAfterUpsertHooks = append(characterSkillAfterUpsertHooks, characterSkillHook)
	}
}

// OneG returns a single characterSkill record from the query using the global executor.
func (q characterSkillQuery) OneG(ctx context.Context) (*CharacterSkill, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single characterSkill record from the query.
func (q characterSkillQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CharacterSkill, error) {
	o := &CharacterSkill{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for character_skills")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all CharacterSkill records from the query using the global executor.
func (q characterSkillQuery) AllG(ctx context.Context) (CharacterSkillSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all CharacterSkill records from the query.
func (q characterSkillQuery) All(ctx context.Context, exec boil.ContextExecutor) (CharacterSkillSlice, error) {
	var o []*CharacterSkill

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to CharacterSkill slice")
	}

	if len(characterSkillAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all CharacterSkill records in the query, and panics on error.
func (q characterSkillQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all CharacterSkill records in the query.
func (q characterSkillQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count character_skills rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q characterSkillQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q characterSkillQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if character_skills exists")
	}

	return count > 0, nil
}

// Character pointed to by the foreign key.
func (o *CharacterSkill) Character(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`character_id` = ?", o.CharacterID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// LoadCharacter allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (characterSkillL) LoadCharacter(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCharacterSkill interface{}, mods queries.Applicator) error {
	var slice []*CharacterSkill
	var object *CharacterSkill

	if singular {
		object = maybeCharacterSkill.(*CharacterSkill)
	} else {
		slice = *maybeCharacterSkill.(*[]*CharacterSkill)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &characterSkillR{}
		}
		args = append(args, object.CharacterID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &characterSkillR{}
			}

			for _, a := range args {
				if a == obj.CharacterID {
					continue Outer
				}
			}

			args = append(args, obj.CharacterID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.character_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(characterSkillAfterSelectHooks) != 0 {
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
		object.R.Character = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.CharacterCharacterSkills = append(foreign.R.CharacterCharacterSkills, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CharacterID == foreign.CharacterID {
				local.R.Character = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.CharacterCharacterSkills = append(foreign.R.CharacterCharacterSkills, local)
				break
			}
		}
	}

	return nil
}

// SetCharacterG of the characterSkill to the related item.
// Sets o.R.Character to related.
// Adds o to related.R.CharacterCharacterSkills.
// Uses the global database handle.
func (o *CharacterSkill) SetCharacterG(ctx context.Context, insert bool, related *User) error {
	return o.SetCharacter(ctx, boil.GetContextDB(), insert, related)
}

// SetCharacter of the characterSkill to the related item.
// Sets o.R.Character to related.
// Adds o to related.R.CharacterCharacterSkills.
func (o *CharacterSkill) SetCharacter(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `character_skills` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"character_id"}),
		strmangle.WhereClause("`", "`", 0, characterSkillPrimaryKeyColumns),
	)
	values := []interface{}{related.CharacterID, o.CharacterID, o.SkillID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CharacterID = related.CharacterID
	if o.R == nil {
		o.R = &characterSkillR{
			Character: related,
		}
	} else {
		o.R.Character = related
	}

	if related.R == nil {
		related.R = &userR{
			CharacterCharacterSkills: CharacterSkillSlice{o},
		}
	} else {
		related.R.CharacterCharacterSkills = append(related.R.CharacterCharacterSkills, o)
	}

	return nil
}

// CharacterSkills retrieves all the records using an executor.
func CharacterSkills(mods ...qm.QueryMod) characterSkillQuery {
	mods = append(mods, qm.From("`character_skills`"))
	return characterSkillQuery{NewQuery(mods...)}
}

// FindCharacterSkillG retrieves a single record by ID.
func FindCharacterSkillG(ctx context.Context, characterID uint64, skillID uint, selectCols ...string) (*CharacterSkill, error) {
	return FindCharacterSkill(ctx, boil.GetContextDB(), characterID, skillID, selectCols...)
}

// FindCharacterSkill retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCharacterSkill(ctx context.Context, exec boil.ContextExecutor, characterID uint64, skillID uint, selectCols ...string) (*CharacterSkill, error) {
	characterSkillObj := &CharacterSkill{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `character_skills` where `character_id`=? AND `skill_id`=?", sel,
	)

	q := queries.Raw(query, characterID, skillID)

	err := q.Bind(ctx, exec, characterSkillObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from character_skills")
	}

	if err = characterSkillObj.doAfterSelectHooks(ctx, exec); err != nil {
		return characterSkillObj, err
	}

	return characterSkillObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CharacterSkill) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CharacterSkill) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no character_skills provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(characterSkillColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	characterSkillInsertCacheMut.RLock()
	cache, cached := characterSkillInsertCache[key]
	characterSkillInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			characterSkillAllColumns,
			characterSkillColumnsWithDefault,
			characterSkillColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(characterSkillType, characterSkillMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(characterSkillType, characterSkillMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `character_skills` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `character_skills` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `character_skills` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, characterSkillPrimaryKeyColumns))
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
		return errors.Wrap(err, "boiler: unable to insert into character_skills")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.CharacterID,
		o.SkillID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for character_skills")
	}

CacheNoHooks:
	if !cached {
		characterSkillInsertCacheMut.Lock()
		characterSkillInsertCache[key] = cache
		characterSkillInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single CharacterSkill record using the global executor.
// See Update for more documentation.
func (o *CharacterSkill) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the CharacterSkill.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CharacterSkill) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	characterSkillUpdateCacheMut.RLock()
	cache, cached := characterSkillUpdateCache[key]
	characterSkillUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			characterSkillAllColumns,
			characterSkillPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update character_skills, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `character_skills` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, characterSkillPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(characterSkillType, characterSkillMapping, append(wl, characterSkillPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update character_skills row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for character_skills")
	}

	if !cached {
		characterSkillUpdateCacheMut.Lock()
		characterSkillUpdateCache[key] = cache
		characterSkillUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q characterSkillQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q characterSkillQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for character_skills")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for character_skills")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CharacterSkillSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CharacterSkillSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterSkillPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `character_skills` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterSkillPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in characterSkill slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all characterSkill")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CharacterSkill) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLCharacterSkillUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CharacterSkill) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no character_skills provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(characterSkillColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCharacterSkillUniqueColumns, o)

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

	characterSkillUpsertCacheMut.RLock()
	cache, cached := characterSkillUpsertCache[key]
	characterSkillUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			characterSkillAllColumns,
			characterSkillColumnsWithDefault,
			characterSkillColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			characterSkillAllColumns,
			characterSkillPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("boiler: unable to upsert character_skills, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`character_skills`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `character_skills` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(characterSkillType, characterSkillMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(characterSkillType, characterSkillMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert for character_skills")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(characterSkillType, characterSkillMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to retrieve unique values for character_skills")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to populate default values for character_skills")
	}

CacheNoHooks:
	if !cached {
		characterSkillUpsertCacheMut.Lock()
		characterSkillUpsertCache[key] = cache
		characterSkillUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single CharacterSkill record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CharacterSkill) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single CharacterSkill record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CharacterSkill) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no CharacterSkill provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), characterSkillPrimaryKeyMapping)
	sql := "DELETE FROM `character_skills` WHERE `character_id`=? AND `skill_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from character_skills")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for character_skills")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q characterSkillQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q characterSkillQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no characterSkillQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from character_skills")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for character_skills")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o CharacterSkillSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CharacterSkillSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(characterSkillBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterSkillPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `character_skills` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterSkillPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from characterSkill slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for character_skills")
	}

	if len(characterSkillAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CharacterSkill) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("boiler: no CharacterSkill provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CharacterSkill) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCharacterSkill(ctx, exec, o.CharacterID, o.SkillID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CharacterSkillSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("boiler: empty CharacterSkillSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CharacterSkillSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CharacterSkillSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterSkillPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `character_skills`.* FROM `character_skills` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterSkillPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in CharacterSkillSlice")
	}

	*o = slice

	return nil
}

// CharacterSkillExistsG checks if the CharacterSkill row exists.
func CharacterSkillExistsG(ctx context.Context, characterID uint64, skillID uint) (bool, error) {
	return CharacterSkillExists(ctx, boil.GetContextDB(), characterID, skillID)
}

// CharacterSkillExists checks if the CharacterSkill row exists.
func CharacterSkillExists(ctx context.Context, exec boil.ContextExecutor, characterID uint64, skillID uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `character_skills` where `character_id`=? AND `skill_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, characterID, skillID)
	}
	row := exec.QueryRowContext(ctx, sql, characterID, skillID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if character_skills exists")
	}

	return exists, nil
}
