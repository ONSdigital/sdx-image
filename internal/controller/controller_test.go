package controller

import (
	"fmt"
	"io"
	"os"
	"sdxImage/internal/test"
	"testing"
)

func getSubmission(filename string) ([]byte, error) {
	jsonFile, err := os.Open("examples/submissions/" + filename + ".json")
	if err != nil {
		return nil, err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)
	bytes, _ := io.ReadAll(jsonFile)
	return bytes, nil
}

func runFromFile(filename string, t *testing.T) {
	test.SetCwdToRoot()
	submission, err := getSubmission(filename)
	if err != nil {
		t.Errorf("failed to read submission: %q", err)
	}
	result, e := Run(submission)
	if e != nil {
		t.Errorf("failed with error: %q", err)
	}
	err = test.SaveJPG("examples/images/"+filename+".jpg", result, 100)
	if err != nil {
		t.Errorf("failed to create image for %s with error: %q", filename, err.Error())
	}
}

func TestAbs(t *testing.T) {
	runFromFile("abs_1802", t)
}

func TestMbs(t *testing.T) {
	runFromFile("mbs_0106", t)
}

func TestMcg(t *testing.T) {
	runFromFile("mcg_0002", t)
}

func TestMbsV1(t *testing.T) {
	runFromFile("v1/mbs_0106", t)
}

func TestV1Abs1808(t *testing.T) {
	runFromFile("v1/abs_1808", t)
}

func TestV1Abs1862(t *testing.T) {
	runFromFile("v1/abs_1862", t)
}

func TestV1Abs1874(t *testing.T) {
	runFromFile("v1/abs_1874", t)
}

func TestBricksV1(t *testing.T) {
	runFromFile("v1/bricks", t)
}

func TestBerd0001V1(t *testing.T) {
	runFromFile("v1/berd_0001", t)
}

func TestBerd0001(t *testing.T) {
	runFromFile("berd_0001", t)
}

func TestBerd0006V1(t *testing.T) {
	runFromFile("v1/berd_0006", t)
}

func TestUkis0001V1(t *testing.T) {
	runFromFile("v1/ukis_0001", t)
}

func TestUkisFail(t *testing.T) {
	runFromFile("v1/ukis_fail", t)
}

func TestSandLand(t *testing.T) {
	runFromFile("v1/sand_land", t)
}

func TestSandMarine(t *testing.T) {
	runFromFile("v1/sand_marine", t)
}

func TestCreditGrantors(t *testing.T) {
	runFromFile("v1/mgc_0001", t)
}

func TestStocks0001(t *testing.T) {
	runFromFile("stocks_0001", t)
}

func TestStocks0033(t *testing.T) {
	runFromFile("stocks_0033", t)
}

func TestStocks0052(t *testing.T) {
	runFromFile("stocks_0052", t)
}

func TestQcas(t *testing.T) {
	runFromFile("qcas_0018", t)
}

func TestAbs1801(t *testing.T) {
	runFromFile("abs_1801", t)
}

func TestAbs1809(t *testing.T) {
	runFromFile("abs_1809", t)
}

func TestAbs1819(t *testing.T) {
	runFromFile("abs_1819", t)
}

func TestAbs1869(t *testing.T) {
	runFromFile("abs_1869", t)
}

func TestMbs0167(t *testing.T) {
	runFromFile("mbs_0167", t)
}

func TestMbs0201(t *testing.T) {
	runFromFile("mbs_0201", t)
}

func TestMbs0255(t *testing.T) {
	runFromFile("mbs_0255", t)
}

func TestAcas0002(t *testing.T) {
	runFromFile("acas_0002", t)
}

func TestBlocks0001(t *testing.T) {
	runFromFile("blocks_0001", t)
}

func TestBricks0001(t *testing.T) {
	runFromFile("bricks_0001", t)
}

func TestConstruction0001(t *testing.T) {
	runFromFile("construction_0001", t)
}

func TestDes0001(t *testing.T) {
	runFromFile("des_0001", t)
}

func TestMwss0001(t *testing.T) {
	runFromFile("mwss_0005", t)
}

func TestQbs0001(t *testing.T) {
	runFromFile("qbs_0001", t)
}

func TestQpses160(t *testing.T) {
	runFromFile("qpses160_0002", t)
}

func TestQpses165(t *testing.T) {
	runFromFile("qpses165_0002", t)
}

func TestQpses169(t *testing.T) {
	runFromFile("qpses169_0003", t)
}

func TestUkis_0001(t *testing.T) {
	runFromFile("ukis_0001", t)
}

func TestVacancies(t *testing.T) {
	runFromFile("vacancies_0006", t)
}

func TestEpe(t *testing.T) {
	runFromFile("epe_0003", t)
}

func TestLcre(t *testing.T) {
	runFromFile("lcre_0009", t)
}

func TestRsi(t *testing.T) {
	runFromFile("rsi_0102", t)
}

func TestFuels(t *testing.T) {
	runFromFile("fuels_0002", t)
}

func TestRailways(t *testing.T) {
	runFromFile("railways_0001", t)
}

func TestTiles(t *testing.T) {
	runFromFile("qrt_0001", t)
}

func TestSlate(t *testing.T) {
	runFromFile("qs_0001", t)
}

func TestSand(t *testing.T) {
	runFromFile("qsl_0002", t)
}

func TestMarine(t *testing.T) {
	runFromFile("qsm_0002", t)
}
