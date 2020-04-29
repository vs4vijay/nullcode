package nullcode

import (
	"context"
	"testing"
)

func setup() (Service, context.Context) {
	return NewCodeService(), context.Background()
}

func TestAdd(t *testing.T) {
	svc, ctx := setup()

	c1, err := svc.Add(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// Testing Get
	ok := c1 == "ok"
	if !ok {
		t.Errorf("failed ok")
	}
}
