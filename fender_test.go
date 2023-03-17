package fender_test

import (
	"context"
	"testing"

	"github.com/foomo/fender"
	"github.com/foomo/fender/fend"
	"github.com/foomo/fender/rule"
	"github.com/stretchr/testify/assert"
)

type Gender string

const (
	GenderMale    Gender = "male"
	GenderFemmale Gender = "female"
)

func (g Gender) String() string {
	return string(g)
}

func (g Gender) Valid() bool {
	return g == GenderMale || g == GenderFemmale
}

var (
	email = fend.NewRules(
		rule.StringMin(6),
		rule.StringMax(30),
		rule.Email,
	)
	title = fend.NewRules(
		rule.StringMin(2),
		rule.NotEqual("dr."),
	)
	gender = fend.NewRules(
		rule.Valid[Gender],
	)
	dynamic = fend.DynamicVar(
		func(ctx context.Context) error {
			return fend.NewRuleError("dynamic", "foo")
		},
	)
	firstName = fend.NewRules(
		rule.StringMin(6),
		rule.StringMax(30),
		rule.Alpha,
	)
	streetNumber = fend.NewRules(
		rule.NumberMin(1),
		rule.NumberMax(10),
	)
)

func TestAll(t *testing.T) {
	tests := []struct {
		name            string
		fends           fend.Fends
		wantAllErr      string
		wantFirstErr    string
		wantAllFirstErr string
	}{
		{
			name: "success values",
			fends: fend.Fends{
				title.Var("dr"),
				gender.Var(GenderMale),
				firstName.Var("foobar"),
				streetNumber.Var(5),
				email.Var("foo@bar.com"),
			},
		},
		{
			name: "failure values",
			fends: fend.Fends{
				title.Var("dr."),
				gender.Var("divers"),
				firstName.Var("foo"),
				streetNumber.Var(15),
				email.Var("foo"),
				dynamic,
			},
			wantAllErr:      "equal=dr.;valid;min=6;min=10;min=6:email=parse;dynamic:foo",
			wantFirstErr:    "equal=dr.",
			wantAllFirstErr: "equal=dr.",
		},
		{
			name: "success fields",
			fends: fend.Fends{
				title.Field("title", "dr"),
				gender.Field("gender", GenderMale),
				firstName.Field("firstName", "foobar"),
				streetNumber.Field("firstName", 5),
				email.Field("email", "foo@bar.com"),
			},
		},
		{
			name: "failure fields",
			fends: fend.Fends{
				title.Field("title", "dr."),
				gender.Field("gender", "divers"),
				firstName.Field("firstName", "foo"),
				streetNumber.Field("firstName", 15),
				email.Field("email", "foo"),
				dynamic,
			},
			wantAllErr:      "equal=dr.;valid;min=6;min=10;min=6:email=parse;dynamic:foo",
			wantFirstErr:    "equal=dr.",
			wantAllFirstErr: "equal=dr.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := fender.All(context.TODO(), tt.fends...); tt.wantAllErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantAllErr)
			}

			if err := fender.All(context.TODO(), tt.fends...); tt.wantAllErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantAllErr)
			}

			if err := fender.First(context.TODO(), tt.fends...); tt.wantFirstErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantFirstErr)
			}

			if err := fender.AllFirst(context.TODO(), tt.fends...); tt.wantAllFirstErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantAllFirstErr)
			}
		})
	}
}
