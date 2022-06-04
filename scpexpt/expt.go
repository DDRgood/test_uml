package scpexpt

import (
	"encoding/csv"
	"fmt"
	"golang/scpalgo"
	"golang/scpfunc"
	"golang/supmath"
	"golang/tableprint"
	"io"
	"time"
)

type Colors int64
type OutputStyle int64

const (
	TableStyle OutputStyle = iota
	CsvStyle
)

const (
	Default Colors = iota
	Green
	Red
)

type Parser func(filename string) (table [][2]int, costs []float64,
	alpha, betta map[int][]int, problem error)

type ExptParams struct {
	PopSize    int
	NumIter    int
	CountExpts int
	Binarizer  *supmath.Binarizer
}

func NewExptParams(popSize, numIter, countExpts int, bin *supmath.Binarizer) *ExptParams {
	return &ExptParams{PopSize: popSize, NumIter: numIter, CountExpts: countExpts, Binarizer: bin}
}

type ScpExptMaker struct {
	resultsHeader []string
	colors        map[Colors]string
}

func NewScpExptMaker() *ScpExptMaker {
	return &ScpExptMaker{resultsHeader: []string{"File", "Min cost", "Mean cost", "Max cost", "Mean size", "Mean time", "Success"},
		colors: map[Colors]string{Default: "\033[0m", Green: "\033[32m", Red: "\033[31m"}}
}

func (expt *ScpExptMaker) Save2File(writer io.Writer, data [][]string, headers []string) {
	table := tableprint.NewWriter(writer)
	table.SetHeader(headers)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func (expt *ScpExptMaker) Save2Csv(writer *csv.Writer, data [][]string, headers []string) {
	err := writer.Write(headers)
	if err != nil {
		fmt.Println("Problem with saving in csv file, ", err.Error())
		return
	}
	writer.WriteAll(data)
}

func (expt *ScpExptMaker) TestSetInstance(instances []string, params *ExptParams,
	solver scpalgo.ScpSolver, parser Parser) ([][]string, []string) {

	data := make([][]string, len(instances))
	for i := range data {
		data[i] = make([]string, len(expt.resultsHeader))
	}

	for index := range instances {
		costsSlice, sizesSlice, timesSlice, ok := TestOneInstance(instances[index], params, solver, parser)

		data[index][0] = string(instances[index])
		data[index][1] = fmt.Sprint(supmath.MinFloat64(costsSlice))
		data[index][2] = fmt.Sprint(supmath.MeanFloat64(costsSlice))
		data[index][3] = fmt.Sprint(supmath.MaxFloat64(costsSlice))
		data[index][4] = fmt.Sprint(supmath.MeanFloat64(sizesSlice))
		data[index][5] = fmt.Sprintf("%v", time.Duration(supmath.MeanInt64(timesSlice)))
		if ok == true {
			data[index][6] = "OK"
		} else {
			data[index][6] = "FAIl"
		}
	}
	return data, expt.resultsHeader
}

//TestOneInstance: do experiments on one instance
func TestOneInstance(instance string, params *ExptParams,
	solver scpalgo.ScpSolver, parser Parser) ([]float64, []float64, []int64, bool) {

	_, costs, alpha, betta, err := parser(instance)

	if err != nil {
		fmt.Println("Problem in instance: ", instance, err.Error())
	}

	// expt info
	costsSlice := make([]float64, params.CountExpts, params.CountExpts)
	sizesSlice := make([]float64, params.CountExpts, params.CountExpts)
	timeSlice := make([]int64, params.CountExpts, params.CountExpts)
	okNormal := true

	for i := 0; i < params.CountExpts; i++ {
		startTime := time.Now()
		repairer := scpfunc.NewSolutionRepairer(alpha, betta, costs)

		// call function
		value, optimum := solver.Solve(params.PopSize, params.NumIter, costs, repairer, params.Binarizer)

		elapsedTime := time.Since(startTime)
		costsSlice[i] = value
		timeSlice[i] = int64(elapsedTime)

		// calc num of columns
		sizesSlice[i] = 0
		for _, e := range optimum {
			if e > 0 {
				sizesSlice[i] += 1.0
			}
		}

		//check solution
		ok, okWhat := repairer.CheckSolution(optimum)
		if !ok {
			fmt.Println("-----Problem with file", instance, okWhat)
			okNormal = false
		}
	}

	return costsSlice, sizesSlice, timeSlice, okNormal

}

/*

DONT USE IT DONT USE IT DELETE


*/

func (expt *ScpExptMaker) TestSetInstanceBHA(limits []float64, instances []string, params *ExptParams,
	solver *scpalgo.BHASolver, parser Parser) ([][]string, []string) {

	data := make([][]string, len(instances))
	for i := range data {
		data[i] = make([]string, len(expt.resultsHeader)+2)
	}

	for index := range instances {
		costsSlice, sizesSlice, timesSlice, ok, stepsSlice, countsSlice := TestOneInstanceBHA(limits[index], instances[index], params, solver, parser)

		data[index][0] = string(instances[index])
		data[index][1] = fmt.Sprint(supmath.MinFloat64(costsSlice))
		data[index][2] = fmt.Sprint(supmath.MeanFloat64(costsSlice))
		data[index][3] = fmt.Sprint(supmath.MaxFloat64(costsSlice))
		data[index][4] = fmt.Sprint(supmath.MeanFloat64(sizesSlice))
		//data[index][5] = fmt.Sprintf("%v", time.Duration(supmath.MeanInt64(timesSlice)))
		data[index][5] = fmt.Sprintf("%v", time.Duration(supmath.CalcMedianInt(timesSlice)))
		if ok == true {
			data[index][6] = "OK"
		} else {
			data[index][6] = "FAIl"
		}
		data[index][7] = fmt.Sprint(supmath.CalcMedian(stepsSlice))
		data[index][8] = fmt.Sprint(supmath.CalcMedian(countsSlice))
	}
	head := make([]string, 0)
	head = append(head, expt.resultsHeader...)
	head = append(head, "steps")
	head = append(head, "count math")
	return data, head
}

//TestOneInstance: do experiments on one instance
func TestOneInstanceBHA(limit float64, instance string, params *ExptParams,
	solver *scpalgo.BHASolver, parser Parser) ([]float64, []float64, []int64, bool, []float64, []float64) {

	_, costs, alpha, betta, err := parser(instance)

	if err != nil {
		fmt.Println("Problem in instance: ", instance, err.Error())
	}

	// expt info
	costsSlice := make([]float64, params.CountExpts, params.CountExpts)
	sizesSlice := make([]float64, params.CountExpts, params.CountExpts)
	timeSlice := make([]int64, params.CountExpts, params.CountExpts)

	stepsSlice := make([]float64, params.CountExpts, params.CountExpts)
	countSlice := make([]float64, params.CountExpts, params.CountExpts)
	okNormal := true

	for i := 0; i < params.CountExpts; i++ {
		startTime := time.Now()
		repairer := scpfunc.NewSolutionRepairer(alpha, betta, costs)

		// call function
		value, optimum, steps, counts := solver.Solve1(limit, params.PopSize, params.NumIter, costs, repairer, params.Binarizer)

		elapsedTime := time.Since(startTime)
		costsSlice[i] = value
		timeSlice[i] = int64(elapsedTime)

		stepsSlice[i] = steps
		countSlice[i] = counts

		// calc num of columns
		sizesSlice[i] = 0
		for _, e := range optimum {
			if e > 0 {
				sizesSlice[i] += 1.0
			}
		}

		//check solution
		ok, okWhat := repairer.CheckSolution(optimum)
		if !ok {
			fmt.Println("-----Problem with file", instance, okWhat)
			okNormal = false
		}
	}

	return costsSlice, sizesSlice, timeSlice, okNormal, stepsSlice, countSlice

}

//// FFFA

func (expt *ScpExptMaker) TestSetInstanceFFA(limits []float64, instances []string, params *ExptParams,
	solver *scpalgo.FFASolver, parser Parser) ([][]string, []string) {

	data := make([][]string, len(instances))
	for i := range data {
		data[i] = make([]string, len(expt.resultsHeader)+2)
	}

	for index := range instances {
		costsSlice, sizesSlice, timesSlice, ok, stepsSlice, countsSlice := TestOneInstanceFFA(limits[index], instances[index], params, solver, parser)

		data[index][0] = string(instances[index])
		data[index][1] = fmt.Sprint(supmath.MinFloat64(costsSlice))
		data[index][2] = fmt.Sprint(supmath.MeanFloat64(costsSlice))
		data[index][3] = fmt.Sprint(supmath.MaxFloat64(costsSlice))
		data[index][4] = fmt.Sprint(supmath.MeanFloat64(sizesSlice))
		data[index][5] = fmt.Sprintf("%v", time.Duration(supmath.CalcMedianInt(timesSlice)))
		if ok == true {
			data[index][6] = "OK"
		} else {
			data[index][6] = "FAIl"
		}
		data[index][7] = fmt.Sprint(supmath.CalcMedian(stepsSlice))
		data[index][8] = fmt.Sprint(supmath.CalcMedian(countsSlice))
	}
	head := make([]string, 0)
	head = append(head, expt.resultsHeader...)
	head = append(head, "steps")
	head = append(head, "count math")
	return data, head
}

//TestOneInstance: do experiments on one instance
func TestOneInstanceFFA(limit float64, instance string, params *ExptParams,
	solver *scpalgo.FFASolver, parser Parser) ([]float64, []float64, []int64, bool, []float64, []float64) {

	_, costs, alpha, betta, err := parser(instance)

	if err != nil {
		fmt.Println("Problem in instance: ", instance, err.Error())
	}

	// expt info
	costsSlice := make([]float64, params.CountExpts, params.CountExpts)
	sizesSlice := make([]float64, params.CountExpts, params.CountExpts)
	timeSlice := make([]int64, params.CountExpts, params.CountExpts)

	stepsSlice := make([]float64, params.CountExpts, params.CountExpts)
	countSlice := make([]float64, params.CountExpts, params.CountExpts)
	okNormal := true

	for i := 0; i < params.CountExpts; i++ {
		startTime := time.Now()
		repairer := scpfunc.NewSolutionRepairer(alpha, betta, costs)

		// call function
		value, optimum, steps, counts := solver.Solve1(limit, params.PopSize, params.NumIter, costs, repairer, params.Binarizer)

		elapsedTime := time.Since(startTime)
		costsSlice[i] = value
		timeSlice[i] = int64(elapsedTime)

		stepsSlice[i] = steps
		countSlice[i] = counts

		// calc num of columns
		sizesSlice[i] = 0
		for _, e := range optimum {
			if e > 0 {
				sizesSlice[i] += 1.0
			}
		}

		//check solution
		ok, okWhat := repairer.CheckSolution(optimum)
		if !ok {
			fmt.Println("-----Problem with file", instance, okWhat)
			okNormal = false
		}
	}

	return costsSlice, sizesSlice, timeSlice, okNormal, stepsSlice, countSlice

}
