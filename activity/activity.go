package activity

import (
	"context"

	"github.com/cschleiden/go-workflows/internal/activity"
)

type ActivityState = activity.ActivityState

// GetActivityState returns state of currently executing activity.
func State(ctx context.Context) *ActivityState {
	return activity.GetActivityState(ctx)
}
