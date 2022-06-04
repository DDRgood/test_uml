package main

import (
	"encoding/csv"
	"fmt"
	"golang/parser_scp"
	"golang/preprocess"
	"golang/scpalgo"
	"golang/scpexpt"
	"golang/scpfunc"
	"golang/supmath"
	_ "net/http/pprof"
	"os"
	"time"
)

var FileNames4 = []string{"./OR/scp41.txt", "./OR/scp42.txt", "./OR/scp43.txt", "./OR/scp44.txt", "./OR/scp45.txt",
	"./OR/scp46.txt", "./OR/scp47.txt", "./OR/scp48.txt", "./OR/scp49.txt", "./OR/scp410.txt"}

var FileNames5 = []string{"./OR/scp51.txt", "./OR/scp52.txt", "./OR/scp53.txt", "./OR/scp54.txt", "./OR/scp55.txt", "./OR/scp56.txt", "./OR/scp57.txt", "./OR/scp58.txt", "./OR/scp59.txt", "./OR/scp510.txt"}
var FileNames6 = []string{"./OR/scp61.txt", "./OR/scp62.txt", "./OR/scp63.txt", "./OR/scp64.txt", "./OR/scp65.txt"}
var FileNamesA = []string{"./OR/scpa1.txt", "./OR/scpa2.txt", "./OR/scpa3.txt", "./OR/scpa4.txt", "./OR/scpa5.txt"}
var FileNamesB = []string{"./OR/scpb1.txt", "./OR/scpb2.txt", "./OR/scpb3.txt", "./OR/scpb4.txt", "./OR/scpb5.txt"}
var FileNamesC = []string{"./OR/scpc1.txt", "./OR/scpc2.txt", "./OR/scpc3.txt", "./OR/scpc4.txt", "./OR/scpc5.txt"}
var FileNamesD = []string{"./OR/scpd1.txt", "./OR/scpd2.txt", "./OR/scpd3.txt", "./OR/scpd4.txt", "./OR/scpd5.txt"}
var FileNamesE = []string{"./OR/scpe1.txt", "./OR/scpe2.txt", "./OR/scpe3.txt", "./OR/scpe4.txt", "./OR/scpe5.txt"}

var FileNamesNRE = []string{"./OR/scpnre1.txt", "./OR/scpnre2.txt", "./OR/scpnre3.txt", "./OR/scpnre4.txt", "./OR/scpnre5.txt"}
var FileNamesNRF = []string{"./OR/scpnrf1.txt", "./OR/scpnrf2.txt", "./OR/scpnrf3.txt", "./OR/scpnrf4.txt", "./OR/scpnrf5.txt"}
var FileNamesNRG = []string{"./OR/scpnrg1.txt", "./OR/scpnrg2.txt", "./OR/scpnrg3.txt", "./OR/scpnrg4.txt", "./OR/scpnrg5.txt"}
var FileNamesNRH = []string{"./OR/scpnrh1.txt", "./OR/scpnrh2.txt", "./OR/scpnrh3.txt", "./OR/scpnrh4.txt", "./OR/scpnrh5.txt"}

var FileNamesRail = []string{"./OR/rail507.txt", "./OR/rail516.txt", "./OR/rail582.txt", "./OR/rail2536.txt", "./OR/rail2586.txt", "./OR/rail4284.txt", "./OR/rail4872.txt"}

var FileNamesAirBus = []string{"./OR/air_bus/aa03.txt", "./OR/air_bus/aa04.txt", "./OR/air_bus/aa05.txt", "./OR/air_bus/aa06.txt",
	"./OR/air_bus/aa11.txt", "./OR/air_bus/aa12.txt", "./OR/air_bus/aa13.txt", "./OR/air_bus/aa14.txt",
	"./OR/air_bus/aa15.txt", "./OR/air_bus/aa16.txt", "./OR/air_bus/aa17.txt", "./OR/air_bus/aa18.txt",
	"./OR/air_bus/aa19.txt", "./OR/air_bus/aa20.txt", "./OR/air_bus/bus1.txt", "./OR/air_bus/bus2.txt"}

func getIterFromTransfer(trans string, filename string, solver scpalgo.ScpSolver) {
	iters := []int{2, 5, 10, 20, 30, 40, 50, 80, 120, 150, 180, 200, 250, 300, 350, 400, 500, 600, 800, 1000}
	numRepeat := 10
	for i := range iters {
		paramsExpt := scpexpt.NewExptParams(20, iters[i], numRepeat,
			supmath.NewBinarizer(supmath.GetTransferByStr(trans), supmath.ElitistDiscrete))
		expt := scpexpt.NewScpExptMaker()
		data, headers := expt.TestSetInstance([]string{filename}, paramsExpt, solver, parser_scp.ParseScp)
		fmt.Println(iters[i])
		expt.Save2File(os.Stdout, data, headers)
	}
}

func main1() {

	paramsExpt := scpexpt.NewExptParams(20, 1000, 15,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.ElitistDiscrete))
	go func() {
		solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 1.0)
		expt := scpexpt.NewScpExptMaker()
		data2, headers2 := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data2, headers2)

		data2, headers2 = expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data2, headers2)

		data2, headers2 = expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data2, headers2)
	}()

	go func() {
		solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 1.0)
		expt := scpexpt.NewScpExptMaker()
		data2, headers2 := expt.TestSetInstance(FileNamesA, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data2, headers2)

		data2, headers2 = expt.TestSetInstance(FileNamesB, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data2, headers2)

		data2, headers2 = expt.TestSetInstance(FileNamesC, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data2, headers2)

		data2, headers2 = expt.TestSetInstance(FileNamesD, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data2, headers2)
	}()

	solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 1.0)
	expt := scpexpt.NewScpExptMaker()
	// v1 elitist

	//data2, headers2 := expt.TestSetInstance(Files, paramsExpt, solver, parser_scp.ParseScp)
	data2, headers2 := expt.TestSetInstance(FileNamesNRF, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)

	data2, headers2 = expt.TestSetInstance(FileNamesNRH, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)

	data2, headers2 = expt.TestSetInstance(FileNamesNRE, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)

	data2, headers2 = expt.TestSetInstance(FileNamesNRG, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)
}

func PSOS3() {
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.ElitistDiscrete))
	solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	expt := scpexpt.NewScpExptMaker()
	go func() {

		data1, headers1 := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

	}()

	go func() {
		data1, headers1 := expt.TestSetInstance(FileNamesA, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNamesB, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNamesC, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

	}()

	data1, headers1 := expt.TestSetInstance(FileNamesNRE, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRH, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRG, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

}

func PSO() {
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		//supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
		supmath.NewBinarizer(supmath.GetTransferByStr("s1"), supmath.ElitistDiscrete))
	solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	expt := scpexpt.NewScpExptMaker()
	go func() {

		data1, headers1 := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

	}()

	go func() {
		data1, headers1 := expt.TestSetInstance(FileNamesA, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNamesB, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNamesC, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

	}()

	data1, headers1 := expt.TestSetInstance(FileNamesNRE, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRH, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRG, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRF, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

}

func main11() {
	//PSO()
	//fmt.Println("Was s1 elist")
	//PSOS3()
	//fmt.Println("BHA s1 elist")
	//BHA()
	//BHA_iters()
	///FFA_iters()
	//FFA()
	//smallBHA()

}

// nolimit rand
func main2() {
	//solver := scpalgo.NewFFASolver([]float64{0.1}, 1.0, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	//solver := scpalgo.NewBHASolver(scpalgo.NoneNorm, scpalgo.RandCollapse, 1.0)
	///solver := scpalgo.NewPSOSolver([]float64{0.1}, 1.0, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	///getIterFromTransfer("v2", FileNames6[0], solver)
	//testPSO()
	///	fmt.Println("PSO")

	//solver := scpalgo.NewRSMSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.ChangeBest)
	//solver := scpalgo.NewFFASolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	//solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.00002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	//solver.MoveOne = true

	/*
		Тестируем FFA для разных случаев (потом еще для BestMove тоже самое было бы неплохо
		а в целом, тенденция если логична, то пофиг
	*/

	var Files = append(FileNames4, FileNames5...)
	Files = append(Files, FileNames6...)
	Files = append(Files, FileNamesA...)
	Files = append(Files, FileNamesB...)
	Files = append(Files, FileNamesC...)
	Files = append(Files, FileNamesD...)
	Files = append(Files, FileNamesE...)
	Files = append(Files, FileNamesNRE...)
	Files = append(Files, FileNamesNRF...)
	Files = append(Files, FileNamesNRG...)
	Files = append(Files, FileNamesNRH...)
	/*
		go func() {
			paramsExpt := scpexpt.NewExptParams(20, 1000, 30,
				supmath.NewBinarizer(supmath.GetTransferByStr("v1"), supmath.StandardDiscrete))
			solver := scpalgo.NewFFASolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
			expt := scpexpt.NewScpExptMaker()
			data1, headers1 := expt.TestSetInstance(Files, paramsExpt, solver, parser_scp.ParseScp)
			fmt.Println("FFA v1 stand")
			headers1[0] += "v1-st"
			expt.Save2File(os.Stdout, data1, headers1)

		}()
	*/
	//solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.BestFFMove, scpalgo.NoChange)
	solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 0.4)
	expt := scpexpt.NewScpExptMaker()
	// v1 elitist
	paramsExpt := scpexpt.NewExptParams(25, 1000, 1,
		supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
	//data2, headers2 := expt.TestSetInstance(Files, paramsExpt, solver, parser_scp.ParseScp)
	data2, headers2 := expt.TestSetInstance([]string{Files[60]}, paramsExpt, solver, parser_scp.ParseScp)
	fmt.Println("FFA v1 elitist")
	headers2[0] += "v1-e"

	expt.Save2File(os.Stdout, data2, headers2)

	//RAIL
	/*
		go func() {
			testFFARailSmall(solver, "pso", []string{FileNamesRail[3]})
		}()
		go func() {
			testFFARailSmall(solver, "pso", []string{FileNamesRail[4]})
		}()
		data1, headers1 := expt.TestSetInstance([]string{FileNamesRail[3]}, paramsExpt, solver, parser_scp.ParseRail)
		expt.Save2File(os.Stdout, data1, headers1)

	*/
	//getIterFromTransfer("v4", FileNames4[0], solver)

}

func t1() {
	var Files = append(FileNames4, FileNames5...)
	Files = append(Files, FileNames6...)
	go func() {
		paramsExpt := scpexpt.NewExptParams(20, 1000, 30,
			supmath.NewBinarizer(supmath.GetTransferByStr("v1"), supmath.StandardDiscrete))
		solver := scpalgo.NewFFASolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
		expt := scpexpt.NewScpExptMaker()
		data1, headers1 := expt.TestSetInstance(Files, paramsExpt, solver, parser_scp.ParseScp)
		fmt.Println("FFA v1 stand")
		headers1[0] += "v1-st"
		expt.Save2File(os.Stdout, data1, headers1)

		file, _ := os.Create("ffaV1Stand.csv")
		defer file.Close()
		writer := csv.NewWriter(file)
		expt.Save2Csv(writer, data1, headers1)

		// v1 elitist
		paramsExpt = scpexpt.NewExptParams(20, 1000, 30,
			supmath.NewBinarizer(supmath.GetTransferByStr("v1"), supmath.ElitistDiscrete))
		data2, headers2 := expt.TestSetInstance(Files, paramsExpt, solver, parser_scp.ParseScp)
		fmt.Println("FFA v1 elitist")
		headers2[0] += "v1-e"

		expt.Save2File(os.Stdout, data2, headers2)

		file1, _ := os.Create("ffaV1Elist.csv")
		defer file1.Close()
		writer1 := csv.NewWriter(file1)
		expt.Save2Csv(writer1, data2, headers2)

	}()

	paramsExpt := scpexpt.NewExptParams(20, 1000, 30,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	solver := scpalgo.NewFFASolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	expt := scpexpt.NewScpExptMaker()
	data1, headers1 := expt.TestSetInstance(Files, paramsExpt, solver, parser_scp.ParseScp)
	fmt.Println("FFA s3 stand")
	headers1[0] += "s3-st"
	expt.Save2File(os.Stdout, data1, headers1)

	file, _ := os.Create("ffaS3Stand.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	expt.Save2Csv(writer, data1, headers1)

	// v1 elitist
	paramsExpt = scpexpt.NewExptParams(20, 1000, 30,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.ElitistDiscrete))
	data2, headers2 := expt.TestSetInstance(Files, paramsExpt, solver, parser_scp.ParseScp)
	fmt.Println("FFA s3 elitist")
	headers2[0] += "s3-e"

	expt.Save2File(os.Stdout, data2, headers2)

	file1, _ := os.Create("ffaS3Elist.csv")
	defer file1.Close()
	writer1 := csv.NewWriter(file1)
	expt.Save2Csv(writer1, data2, headers2)
}

func testFFAMoveStandGamma() {
	fmt.Println("FFA MoveStand, Gamma = 0.00002, NoChangeBest firefly")
	solver := scpalgo.NewFFASolver([]float64{0.1}, 0.00002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	expt := scpexpt.NewScpExptMaker()
	go func() {
		data, headers := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	data1, headers1 := expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data2, headers2 := expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)
}

func testFFAMoveBestGamma() {
	fmt.Println("FFA MoveBest, Gamma = 0.00002, NoChangeBest firefly")
	solver := scpalgo.NewFFASolver([]float64{0.1}, 0.00002, 1.0, 2, scpalgo.BestFFMove, scpalgo.NoChange)
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	expt := scpexpt.NewScpExptMaker()
	go func() {
		data, headers := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	data1, headers1 := expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data2, headers2 := expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)
}

func testPSONRE() {
	fmt.Println("PSO")
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.ChangeBest)
	expt := scpexpt.NewScpExptMaker()
	go func() {
		data, headers := expt.TestSetInstance(FileNamesNRG, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	data1, headers1 := expt.TestSetInstance(FileNamesA, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data2, headers2 := expt.TestSetInstance(FileNamesD, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)
}

func testBHRandVersion() {
	fmt.Println("Stand BH with Rand collapse")
	solver := scpalgo.NewBHASolver(scpalgo.NoneNorm, scpalgo.RandCollapse, 0.8)
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	expt := scpexpt.NewScpExptMaker()
	go func() {
		data, headers := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	go func() {
		data, headers := expt.TestSetInstance(FileNamesD, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	data1, headers1 := expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data2, headers2 := expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)
}

func testBHModVersion() {
	fmt.Println("Stand BH with Mod collapse Max Norme")
	solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 0.8)
	paramsExpt := scpexpt.NewExptParams(20, 1000, 20,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	expt := scpexpt.NewScpExptMaker()
	go func() {
		data, headers := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	go func() {
		data, headers := expt.TestSetInstance(FileNamesD, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	data1, headers1 := expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data2, headers2 := expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)
}

func testFFAMoveStandAlpha() {
	fmt.Println("testFFAMoveStandAlpha")

}

func testPSO() {
	fmt.Println("PSO")
	paramsExpt := scpexpt.NewExptParams(20, 1000, 20,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.ChangeBest)
	expt := scpexpt.NewScpExptMaker()
	go func() {
		data, headers := expt.TestSetInstance(FileNamesC, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	go func() {
		data, headers := expt.TestSetInstance(FileNamesB, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	data1, headers1 := expt.TestSetInstance(FileNamesD, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data2, headers2 := expt.TestSetInstance(FileNamesA, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)
}

func test1() {
	fmt.Println("Improve")
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	solver := scpalgo.NewImproveBHASolver(scpalgo.NoneNorm, scpalgo.RandCollapse, 0.2)
	expt := scpexpt.NewScpExptMaker()
	go func() {
		data, headers := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	data1, headers1 := expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data2, headers2 := expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)

}

func test2() {
	fmt.Println("PSO")
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.ChangeBest)
	expt := scpexpt.NewScpExptMaker()
	go func() {
		data, headers := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	data1, headers1 := expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data2, headers2 := expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)

}

func test3() {
	fmt.Println("Rand ")
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
	solver := scpalgo.NewRSMSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.BestFFMove, scpalgo.ChangeBest)
	expt := scpexpt.NewScpExptMaker()
	go func() {
		data, headers := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data, headers)
	}()
	data1, headers1 := expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data2, headers2 := expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data2, headers2)

}

func testFFARailSmall(solver scpalgo.ScpSolver, str string, filenames []string) {
	data := make([][]string, 10, 10)
	for i := range data {
		data[i] = make([]string, 4)
	}
	for i, filename := range filenames {
		table, costs, alpha, betta, err := parser_scp.ParseRail(filename)
		if err != nil {
			fmt.Println("Problem with file ", filename)
			continue
		}
		startTime := time.Now()

		reducer := preprocess.NewReduceTable(table, costs, alpha, betta)
		sol := reducer.Reduce()
		elapsedTime := time.Since(startTime)
		alphaNew, bettaNew := reducer.GetAlpha(), reducer.GetBetta()
		fmt.Println("After alpha", len(alphaNew), "betta", len(bettaNew))
		data[i][2] = elapsedTime.String()

		repair := scpfunc.NewSolutionRepairer(alphaNew, bettaNew, costs)
		startTime = time.Now()
		value, optimum := solver.Solve(20, 1000, costs, repair,
			supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.StandardDiscrete))
		elapsedTime = time.Since(startTime)
		s, _ := scpfunc.CalcFitness(sol, costs)
		data[i][0] = fmt.Sprint(value + s)
		data[i][1] = elapsedTime.String()
		_, data[i][3] = repair.CheckSolution(optimum)
	}
	expt := scpexpt.NewScpExptMaker()
	expt.Save2File(os.Stdout, data, []string{"Value" + str, "Time work", "Time reduce", "Check"})
}

/*

1
2
3
3
4
5

6
7
7



*/

func smallBHA() {
	paramsExpt := scpexpt.NewExptParams(20, 1000, 3,
		//supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
		supmath.NewBinarizer(supmath.GetTransferByStr("s3"), supmath.ElitistDiscrete))
	solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 0.3)
	expt := scpexpt.NewScpExptMaker()
	go func() {

		data1, headers1 := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

	}()

	go func() {
		data1, headers1 := expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)
		data2, headers2 := expt.TestSetInstance(FileNamesNRH, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data2, headers2)
	}()

	data1, headers1 := expt.TestSetInstance(FileNamesNRG, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

}

func BHA() {
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		//supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
		supmath.NewBinarizer(supmath.GetTransferByStr("s1"), supmath.ElitistDiscrete))
	//solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 0.6)
	expt := scpexpt.NewScpExptMaker()
	go func() {

		data1, headers1 := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

	}()

	go func() {
		data1, headers1 := expt.TestSetInstance(FileNamesA, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNamesB, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNamesC, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

	}()

	data1, headers1 := expt.TestSetInstance(FileNamesNRE, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRH, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRG, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRF, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

}

func FFA() {
	paramsExpt := scpexpt.NewExptParams(20, 1000, 5,
		//supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
		supmath.NewBinarizer(supmath.GetTransferByStr("v3"), supmath.ElitistDiscrete))
	solver := scpalgo.NewFFASolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.BestFFMove, scpalgo.NoChange)
	//solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 0.6)
	expt := scpexpt.NewScpExptMaker()
	fmt.Println("s1 stand")
	go func() {

		data1, headers1 := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNames5, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNames6, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

	}()

	go func() {
		data1, headers1 := expt.TestSetInstance(FileNamesA, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNamesB, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

		data1, headers1 = expt.TestSetInstance(FileNamesC, paramsExpt, solver, parser_scp.ParseScp)
		expt.Save2File(os.Stdout, data1, headers1)

	}()

	data1, headers1 := expt.TestSetInstance(FileNamesNRE, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRH, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRG, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

	data1, headers1 = expt.TestSetInstance(FileNamesNRF, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)

}

func AirLine() {

	expt := scpexpt.NewScpExptMaker()

	//FFA
	fmt.Println("FFA Results===========================")
	paramsExpt := scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
	solver := scpalgo.NewFFASolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.BestFFMove, scpalgo.NoChange)
	data1, headers1 := expt.TestSetInstance(FileNamesAirBus, paramsExpt, solver, parser_scp.ParseAirline)
	expt.Save2File(os.Stdout, data1, headers1)
	fmt.Println("END FFA +==============================")

	//PSO
	fmt.Println("PSO Results===========================")
	paramsExpt = scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
	solver1 := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	data2, headers2 := expt.TestSetInstance(FileNamesAirBus, paramsExpt, solver1, parser_scp.ParseAirline)
	expt.Save2File(os.Stdout, data2, headers2)
	fmt.Println("END PSO +==============================")

	//BHA
	fmt.Println("BHA Results===========================")
	paramsExpt = scpexpt.NewExptParams(20, 1000, 10,
		supmath.NewBinarizer(supmath.GetTransferByStr("s12"), supmath.ElitistDiscrete))
	solver2 := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 0.35)
	data3, headers3 := expt.TestSetInstance(FileNamesAirBus, paramsExpt, solver2, parser_scp.ParseAirline)
	expt.Save2File(os.Stdout, data3, headers3)
	fmt.Println("END BHA +==============================")

}

func BHA_iters() {
	paramsExpt := scpexpt.NewExptParams(20, 30, 10,
		//supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
		supmath.NewBinarizer(supmath.GetTransferByStr("s12"), supmath.ElitistDiscrete))
	//solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 1.0)
	expt := scpexpt.NewScpExptMaker()

	data1, headers1 := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)
}

func FFA_iters() {
	paramsExpt := scpexpt.NewExptParams(20, 30, 1,
		//supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
		supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
	solver := scpalgo.NewFFASolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)

	expt := scpexpt.NewScpExptMaker()

	data1, headers1 := expt.TestSetInstance(FileNames4, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)
}

var scp4Limits = []float64{430, 512, 516, 494, 514, 560, 430, 492, 644, 515}

var scp5Limits = []float64{254, 305, 228, 243, 211, 213, 293, 288, 279, 265}

var scp6Limits = []float64{138, 146, 145, 131, 161}

var scpALimits = []float64{254, 253, 234, 235, 237}

var scpBLimits = []float64{70, 77, 81, 79, 72}

var scpCLimits = []float64{227, 220, 244, 220, 217}

var scpNRGLimits = []float64{179, 158, 170, 171, 170}

var scpDLimits = []float64{60, 66, 74, 62, 62}
var scpNRELimits = []float64{29, 30, 27, 28, 28}
var scpNRFLimits = []float64{14, 15, 15, 14, 14}
var scpNRHLimits = []float64{65, 65, 61, 59, 55}

var scpAIRLimitsBHA = []float64{33155, 34573, 31623, 37464, 35492, 30825,
	33211, 33219, 34485, 32886, 31678, 36799, 32317, 34912, 27983, 68283}

var scpAIRLimitsFFA = []float64{33155, 34731, 31623, 37468, 35502, 31063, 33523,
	33451, 34651, 32983, 31955, 37134, 32544, 35477, 28064, 68060}

func BHA_ITER_NEW_TEST(files []string, limits []float64) {
	paramsExpt := scpexpt.NewExptParams(20, 1000, 30,
		//supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
		supmath.NewBinarizer(supmath.GetTransferByStr("s12"), supmath.ElitistDiscrete))
	//solver := scpalgo.NewPSOSolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)
	solver := scpalgo.NewBHASolver(scpalgo.MaxNorm, scpalgo.StandCollapse, 0.15)
	expt := scpexpt.NewScpExptMaker()

	data1, headers1 := expt.TestSetInstanceBHA(limits, files, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)
}

func FFA_ITER_NEW_TEST(files []string, limits []float64) {
	paramsExpt := scpexpt.NewExptParams(20, 1000, 30,
		//supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
		supmath.NewBinarizer(supmath.GetTransferByStr("v2"), supmath.ElitistDiscrete))
	solver := scpalgo.NewFFASolver([]float64{0.1}, 0.0002, 1.0, 2, scpalgo.StandardMove, scpalgo.NoChange)

	expt := scpexpt.NewScpExptMaker()

	data1, headers1 := expt.TestSetInstanceFFA(limits, files, paramsExpt, solver, parser_scp.ParseScp)
	expt.Save2File(os.Stdout, data1, headers1)
}

func main() {
	//PSO()
	//fmt.Println("Was s1 elist")
	//PSOS3()
	//fmt.Println("BHA s1 elist")
	//BHA()
	//BHA_iters()
	///FFA_iters()
	//FFA()
	//smallBHA()

	//FFA_ITER_NEW_TEST(FileNames4, scp4Limits)
	//FFA_ITER_NEW_TEST(FileNames5, scp5Limits)
	//FFA_ITER_NEW_TEST(FileNames6, scp6Limits)
	//FFA_ITER_NEW_TEST(FileNamesA, scpALimits)
	/*	FFA_ITER_NEW_TEST(FileNamesB, scpBLimits)
		FFA_ITER_NEW_TEST(FileNamesC, scpCLimits)
		FFA_ITER_NEW_TEST(FileNamesD, scpDLimits)
		FFA_ITER_NEW_TEST(FileNamesNRE, scpNRELimits)
		FFA_ITER_NEW_TEST(FileNamesNRF, scpNRFLimits)
		FFA_ITER_NEW_TEST(FileNamesNRH, scpNRHLimits) */
	//FFA_ITER_NEW_TEST(FileNamesAirBus, scpAIRLimitsFFA)
	//FFA_ITER_NEW_TEST(FileNamesNRG, scpNRGLimits)
	fmt.Println("BHA BHA BHA")
	//BHA_ITER_NEW_TEST(FileNames4, scp4Limits)
	//BHA_ITER_NEW_TEST(FileNames5, scp5Limits)
	//BHA_ITER_NEW_TEST(FileNames6, scp6Limits)
	//BHA_ITER_NEW_TEST(FileNamesA, scpALimits)*/
	//BHA_ITER_NEW_TEST(FileNamesNRG, scpNRGLimits)
	BHA_ITER_NEW_TEST(FileNamesB, scpBLimits)
	BHA_ITER_NEW_TEST(FileNamesC, scpCLimits)
	BHA_ITER_NEW_TEST(FileNamesD, scpDLimits)
	BHA_ITER_NEW_TEST(FileNamesNRE, scpNRELimits)
	BHA_ITER_NEW_TEST(FileNamesNRF, scpNRFLimits)
	BHA_ITER_NEW_TEST(FileNamesNRH, scpNRHLimits)
	//BHA_ITER_NEW_TEST(FileNamesAirBus, scpAIRLimitsBHA)
}
