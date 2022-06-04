package scpalgo

import (
	"golang/scpfunc"
	"golang/supmath"
)

type ScpSolver interface {
	Solve(popSize int, numIter int, costs []float64,
		repair *scpfunc.SolutionRepairer, binarizer *supmath.Binarizer) (float64, []float64)
}
