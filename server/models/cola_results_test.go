// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testColaResults(t *testing.T) {
	t.Parallel()

	query := ColaResults()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testColaResultsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
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

	count, err := ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testColaResultsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ColaResults().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testColaResultsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ColaResultSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testColaResultsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ColaResultExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ColaResult exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ColaResultExists to return true, but got false.")
	}
}

func testColaResultsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	colaResultFound, err := FindColaResult(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if colaResultFound == nil {
		t.Error("want a record, got nil")
	}
}

func testColaResultsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ColaResults().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testColaResultsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ColaResults().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testColaResultsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	colaResultOne := &ColaResult{}
	colaResultTwo := &ColaResult{}
	if err = randomize.Struct(seed, colaResultOne, colaResultDBTypes, false, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}
	if err = randomize.Struct(seed, colaResultTwo, colaResultDBTypes, false, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = colaResultOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = colaResultTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ColaResults().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testColaResultsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	colaResultOne := &ColaResult{}
	colaResultTwo := &ColaResult{}
	if err = randomize.Struct(seed, colaResultOne, colaResultDBTypes, false, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}
	if err = randomize.Struct(seed, colaResultTwo, colaResultDBTypes, false, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = colaResultOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = colaResultTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func colaResultBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ColaResult) error {
	*o = ColaResult{}
	return nil
}

func colaResultAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ColaResult) error {
	*o = ColaResult{}
	return nil
}

func colaResultAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ColaResult) error {
	*o = ColaResult{}
	return nil
}

func colaResultBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ColaResult) error {
	*o = ColaResult{}
	return nil
}

func colaResultAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ColaResult) error {
	*o = ColaResult{}
	return nil
}

func colaResultBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ColaResult) error {
	*o = ColaResult{}
	return nil
}

func colaResultAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ColaResult) error {
	*o = ColaResult{}
	return nil
}

func colaResultBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ColaResult) error {
	*o = ColaResult{}
	return nil
}

func colaResultAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ColaResult) error {
	*o = ColaResult{}
	return nil
}

func testColaResultsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ColaResult{}
	o := &ColaResult{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, colaResultDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ColaResult object: %s", err)
	}

	AddColaResultHook(boil.BeforeInsertHook, colaResultBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	colaResultBeforeInsertHooks = []ColaResultHook{}

	AddColaResultHook(boil.AfterInsertHook, colaResultAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	colaResultAfterInsertHooks = []ColaResultHook{}

	AddColaResultHook(boil.AfterSelectHook, colaResultAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	colaResultAfterSelectHooks = []ColaResultHook{}

	AddColaResultHook(boil.BeforeUpdateHook, colaResultBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	colaResultBeforeUpdateHooks = []ColaResultHook{}

	AddColaResultHook(boil.AfterUpdateHook, colaResultAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	colaResultAfterUpdateHooks = []ColaResultHook{}

	AddColaResultHook(boil.BeforeDeleteHook, colaResultBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	colaResultBeforeDeleteHooks = []ColaResultHook{}

	AddColaResultHook(boil.AfterDeleteHook, colaResultAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	colaResultAfterDeleteHooks = []ColaResultHook{}

	AddColaResultHook(boil.BeforeUpsertHook, colaResultBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	colaResultBeforeUpsertHooks = []ColaResultHook{}

	AddColaResultHook(boil.AfterUpsertHook, colaResultAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	colaResultAfterUpsertHooks = []ColaResultHook{}
}

func testColaResultsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testColaResultsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(colaResultColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testColaResultToOneColaTypeUsingCola(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ColaResult
	var foreign ColaType

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, colaResultDBTypes, false, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, colaTypeDBTypes, false, colaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaType struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ColaID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Cola().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ColaResultSlice{&local}
	if err = local.L.LoadCola(ctx, tx, false, (*[]*ColaResult)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Cola == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Cola = nil
	if err = local.L.LoadCola(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Cola == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testColaResultToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ColaResult
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, colaResultDBTypes, false, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ColaResultSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*ColaResult)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testColaResultToOneSetOpColaTypeUsingCola(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ColaResult
	var b, c ColaType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colaResultDBTypes, false, strmangle.SetComplement(colaResultPrimaryKeyColumns, colaResultColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, colaTypeDBTypes, false, strmangle.SetComplement(colaTypePrimaryKeyColumns, colaTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, colaTypeDBTypes, false, strmangle.SetComplement(colaTypePrimaryKeyColumns, colaTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ColaType{&b, &c} {
		err = a.SetCola(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Cola != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ColaColaResults[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ColaID != x.ID {
			t.Error("foreign key was wrong value", a.ColaID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ColaID))
		reflect.Indirect(reflect.ValueOf(&a.ColaID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ColaID != x.ID {
			t.Error("foreign key was wrong value", a.ColaID, x.ID)
		}
	}
}
func testColaResultToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ColaResult
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colaResultDBTypes, false, strmangle.SetComplement(colaResultPrimaryKeyColumns, colaResultColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ColaResults[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testColaResultsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
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

func testColaResultsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ColaResultSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testColaResultsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ColaResults().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	colaResultDBTypes = map[string]string{`ID`: `integer`, `UserID`: `integer`, `ColaID`: `integer`, `ResultDate`: `timestamp without time zone`, `Note`: `character varying`}
	_                 = bytes.MinRead
)

func testColaResultsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(colaResultPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(colaResultAllColumns) == len(colaResultPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testColaResultsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(colaResultAllColumns) == len(colaResultPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ColaResult{}
	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, colaResultDBTypes, true, colaResultPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(colaResultAllColumns, colaResultPrimaryKeyColumns) {
		fields = colaResultAllColumns
	} else {
		fields = strmangle.SetComplement(
			colaResultAllColumns,
			colaResultPrimaryKeyColumns,
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

	slice := ColaResultSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testColaResultsUpsert(t *testing.T) {
	t.Parallel()

	if len(colaResultAllColumns) == len(colaResultPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ColaResult{}
	if err = randomize.Struct(seed, &o, colaResultDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ColaResult: %s", err)
	}

	count, err := ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, colaResultDBTypes, false, colaResultPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ColaResult struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ColaResult: %s", err)
	}

	count, err = ColaResults().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
