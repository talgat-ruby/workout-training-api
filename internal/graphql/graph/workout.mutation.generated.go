// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"workout-training-api/internal/graphql/graph/model"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Exercise_id(ctx context.Context, field graphql.CollectedField, obj *model.Exercise) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Exercise_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Exercise_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Exercise",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Exercise_name(ctx context.Context, field graphql.CollectedField, obj *model.Exercise) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Exercise_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Exercise_name(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Exercise",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Exercise_category(ctx context.Context, field graphql.CollectedField, obj *model.Exercise) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Exercise_category(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Category, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Exercise_category(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Exercise",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Exercise_repetitions(ctx context.Context, field graphql.CollectedField, obj *model.Exercise) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Exercise_repetitions(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Repetitions, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int32)
	fc.Result = res
	return ec.marshalNInt2int32(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Exercise_repetitions(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Exercise",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Exercise_sets(ctx context.Context, field graphql.CollectedField, obj *model.Exercise) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Exercise_sets(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Sets, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int32)
	fc.Result = res
	return ec.marshalNInt2int32(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Exercise_sets(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Exercise",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Exercise_weight(ctx context.Context, field graphql.CollectedField, obj *model.Exercise) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Exercise_weight(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Weight, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalNFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Exercise_weight(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Exercise",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Float does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Workout_id(ctx context.Context, field graphql.CollectedField, obj *model.Workout) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Workout_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Workout_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Workout",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Workout_name(ctx context.Context, field graphql.CollectedField, obj *model.Workout) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Workout_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Workout_name(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Workout",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Workout_exercises(ctx context.Context, field graphql.CollectedField, obj *model.Workout) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Workout_exercises(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Exercises, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.Exercise)
	fc.Result = res
	return ec.marshalNExercise2ᚕᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐExercise(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Workout_exercises(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Workout",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Exercise_id(ctx, field)
			case "name":
				return ec.fieldContext_Exercise_name(ctx, field)
			case "category":
				return ec.fieldContext_Exercise_category(ctx, field)
			case "repetitions":
				return ec.fieldContext_Exercise_repetitions(ctx, field)
			case "sets":
				return ec.fieldContext_Exercise_sets(ctx, field)
			case "weight":
				return ec.fieldContext_Exercise_weight(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Exercise", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Workout_description(ctx context.Context, field graphql.CollectedField, obj *model.Workout) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Workout_description(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Workout_description(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Workout",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Workout_scheduledTimes(ctx context.Context, field graphql.CollectedField, obj *model.Workout) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Workout_scheduledTimes(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ScheduledTimes, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int32)
	fc.Result = res
	return ec.marshalNInt2int32(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Workout_scheduledTimes(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Workout",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputExerciseInput(ctx context.Context, obj any) (model.ExerciseInput, error) {
	var it model.ExerciseInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "name", "category", "repetitions", "sets", "weight"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNID2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "name":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "category":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("category"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Category = data
		case "repetitions":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("repetitions"))
			data, err := ec.unmarshalNInt2int32(ctx, v)
			if err != nil {
				return it, err
			}
			it.Repetitions = data
		case "sets":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("sets"))
			data, err := ec.unmarshalNInt2int32(ctx, v)
			if err != nil {
				return it, err
			}
			it.Sets = data
		case "weight":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("weight"))
			data, err := ec.unmarshalNFloat2float64(ctx, v)
			if err != nil {
				return it, err
			}
			it.Weight = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputWorkoutInput(ctx context.Context, obj any) (model.WorkoutInput, error) {
	var it model.WorkoutInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "name", "exercises", "description", "scheduledTimes"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNID2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "name":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "exercises":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("exercises"))
			data, err := ec.unmarshalNExerciseInput2ᚕᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐExerciseInput(ctx, v)
			if err != nil {
				return it, err
			}
			it.Exercises = data
		case "description":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("description"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Description = data
		case "scheduledTimes":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("scheduledTimes"))
			data, err := ec.unmarshalNInt2int32(ctx, v)
			if err != nil {
				return it, err
			}
			it.ScheduledTimes = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var exerciseImplementors = []string{"Exercise"}

func (ec *executionContext) _Exercise(ctx context.Context, sel ast.SelectionSet, obj *model.Exercise) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, exerciseImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Exercise")
		case "id":
			out.Values[i] = ec._Exercise_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "name":
			out.Values[i] = ec._Exercise_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "category":
			out.Values[i] = ec._Exercise_category(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "repetitions":
			out.Values[i] = ec._Exercise_repetitions(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "sets":
			out.Values[i] = ec._Exercise_sets(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "weight":
			out.Values[i] = ec._Exercise_weight(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var workoutImplementors = []string{"Workout"}

func (ec *executionContext) _Workout(ctx context.Context, sel ast.SelectionSet, obj *model.Workout) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, workoutImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Workout")
		case "id":
			out.Values[i] = ec._Workout_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "name":
			out.Values[i] = ec._Workout_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "exercises":
			out.Values[i] = ec._Workout_exercises(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "description":
			out.Values[i] = ec._Workout_description(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "scheduledTimes":
			out.Values[i] = ec._Workout_scheduledTimes(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNExercise2ᚕᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐExercise(ctx context.Context, sel ast.SelectionSet, v []*model.Exercise) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOExercise2ᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐExercise(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	return ret
}

func (ec *executionContext) unmarshalNExerciseInput2ᚕᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐExerciseInput(ctx context.Context, v any) ([]*model.ExerciseInput, error) {
	var vSlice []any
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*model.ExerciseInput, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalOExerciseInput2ᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐExerciseInput(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNWorkout2ᚕᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐWorkoutᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.Workout) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNWorkout2ᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐWorkout(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNWorkout2ᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐWorkout(ctx context.Context, sel ast.SelectionSet, v *model.Workout) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Workout(ctx, sel, v)
}

func (ec *executionContext) unmarshalNWorkoutInput2workoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐWorkoutInput(ctx context.Context, v any) (model.WorkoutInput, error) {
	res, err := ec.unmarshalInputWorkoutInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOExercise2ᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐExercise(ctx context.Context, sel ast.SelectionSet, v *model.Exercise) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Exercise(ctx, sel, v)
}

func (ec *executionContext) unmarshalOExerciseInput2ᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐExerciseInput(ctx context.Context, v any) (*model.ExerciseInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputExerciseInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
