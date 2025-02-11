// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"
	"workout-training-api/internal/graphql"
	"workout-training-api/internal/graphql/graph/model"


)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _SignInResp_token(ctx context.Context, field graphql.CollectedField, obj *model.SignInResp) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SignInResp_token(ctx, field)
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
		return obj.Token, nil
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

func (ec *executionContext) fieldContext_SignInResp_token(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SignInResp",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SignUpResp_token(ctx context.Context, field graphql.CollectedField, obj *model.SignUpResp) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SignUpResp_token(ctx, field)
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
		return obj.Token, nil
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

func (ec *executionContext) fieldContext_SignUpResp_token(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SignUpResp",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var signInRespImplementors = []string{"SignInResp"}

func (ec *executionContext) _SignInResp(ctx context.Context, sel ast.SelectionSet, obj *model.SignInResp) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, signInRespImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("SignInResp")
		case "token":
			out.Values[i] = ec._SignInResp_token(ctx, field, obj)
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

var signUpRespImplementors = []string{"SignUpResp"}

func (ec *executionContext) _SignUpResp(ctx context.Context, sel ast.SelectionSet, obj *model.SignUpResp) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, signUpRespImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("SignUpResp")
		case "token":
			out.Values[i] = ec._SignUpResp_token(ctx, field, obj)
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

func (ec *executionContext) marshalNSignInResp2workoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐSignInResp(ctx context.Context, sel ast.SelectionSet, v model.SignInResp) graphql.Marshaler {
	return ec._SignInResp(ctx, sel, &v)
}

func (ec *executionContext) marshalNSignInResp2ᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐSignInResp(ctx context.Context, sel ast.SelectionSet, v *model.SignInResp) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._SignInResp(ctx, sel, v)
}

func (ec *executionContext) marshalNSignUpResp2workoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐSignUpResp(ctx context.Context, sel ast.SelectionSet, v model.SignUpResp) graphql.Marshaler {
	return ec._SignUpResp(ctx, sel, &v)
}

func (ec *executionContext) marshalNSignUpResp2ᚖworkoutᚑtrainingᚑapiᚋinternalᚋgraphqlᚋgraphᚋmodelᚐSignUpResp(ctx context.Context, sel ast.SelectionSet, v *model.SignUpResp) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._SignUpResp(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
