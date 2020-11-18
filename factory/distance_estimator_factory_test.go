package factory

import (
	"github.com/sebastian-sz/GotwockAppServer/distance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_initializeHaversteinDistanceEstimator(t *testing.T) {
	haversteinDistanceEstimator := initializeHaversteinDistanceEstimator()

	_, ok := haversteinDistanceEstimator.(distance.Estimator)

	assert.True(t, ok)
}
