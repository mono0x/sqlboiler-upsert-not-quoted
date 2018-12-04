// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testCharacters(t *testing.T) {
	t.Parallel()

	query := Characters()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCharactersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCharactersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Characters().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCharactersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CharacterSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCharactersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CharacterExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Character exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CharacterExists to return true, but got false.")
	}
}

func testCharactersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	characterFound, err := FindCharacter(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if characterFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCharactersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Characters().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCharactersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Characters().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCharactersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	characterOne := &Character{}
	characterTwo := &Character{}
	if err = randomize.Struct(seed, characterOne, characterDBTypes, false, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}
	if err = randomize.Struct(seed, characterTwo, characterDBTypes, false, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = characterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = characterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Characters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCharactersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	characterOne := &Character{}
	characterTwo := &Character{}
	if err = randomize.Struct(seed, characterOne, characterDBTypes, false, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}
	if err = randomize.Struct(seed, characterTwo, characterDBTypes, false, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = characterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = characterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func characterBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Character) error {
	*o = Character{}
	return nil
}

func characterAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Character) error {
	*o = Character{}
	return nil
}

func characterAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Character) error {
	*o = Character{}
	return nil
}

func characterBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Character) error {
	*o = Character{}
	return nil
}

func characterAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Character) error {
	*o = Character{}
	return nil
}

func characterBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Character) error {
	*o = Character{}
	return nil
}

func characterAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Character) error {
	*o = Character{}
	return nil
}

func characterBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Character) error {
	*o = Character{}
	return nil
}

func characterAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Character) error {
	*o = Character{}
	return nil
}

func testCharactersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Character{}
	o := &Character{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, characterDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Character object: %s", err)
	}

	AddCharacterHook(boil.BeforeInsertHook, characterBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	characterBeforeInsertHooks = []CharacterHook{}

	AddCharacterHook(boil.AfterInsertHook, characterAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	characterAfterInsertHooks = []CharacterHook{}

	AddCharacterHook(boil.AfterSelectHook, characterAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	characterAfterSelectHooks = []CharacterHook{}

	AddCharacterHook(boil.BeforeUpdateHook, characterBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	characterBeforeUpdateHooks = []CharacterHook{}

	AddCharacterHook(boil.AfterUpdateHook, characterAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	characterAfterUpdateHooks = []CharacterHook{}

	AddCharacterHook(boil.BeforeDeleteHook, characterBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	characterBeforeDeleteHooks = []CharacterHook{}

	AddCharacterHook(boil.AfterDeleteHook, characterAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	characterAfterDeleteHooks = []CharacterHook{}

	AddCharacterHook(boil.BeforeUpsertHook, characterBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	characterBeforeUpsertHooks = []CharacterHook{}

	AddCharacterHook(boil.AfterUpsertHook, characterAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	characterAfterUpsertHooks = []CharacterHook{}
}

func testCharactersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCharactersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(characterColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCharactersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCharactersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CharacterSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCharactersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Characters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	characterDBTypes = map[string]string{`ID`: `bigint`, `Name`: `varchar`}
	_                = bytes.MinRead
)

func testCharactersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(characterPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(characterColumns) == len(characterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, characterDBTypes, true, characterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCharactersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(characterColumns) == len(characterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Character{}
	if err = randomize.Struct(seed, o, characterDBTypes, true, characterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, characterDBTypes, true, characterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(characterColumns, characterPrimaryKeyColumns) {
		fields = characterColumns
	} else {
		fields = strmangle.SetComplement(
			characterColumns,
			characterPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := CharacterSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCharactersUpsert(t *testing.T) {
	t.Parallel()

	if len(characterColumns) == len(characterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCharacterUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Character{}
	if err = randomize.Struct(seed, &o, characterDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Character: %s", err)
	}

	count, err := Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, characterDBTypes, false, characterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Character struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Character: %s", err)
	}

	count, err = Characters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}