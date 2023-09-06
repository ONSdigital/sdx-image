package schema

import (
	"fmt"
	"sdxImage/pkg/test"
	"sort"
	"testing"
)

func TestReadAbs(t *testing.T) {
	test.SetCwdToRoot()
	result, err := Read("abs_1802")
	if err != nil {
		t.Errorf("failed to read abs_1802 with error: %q", err.Error())
	}
	fmt.Println(result)
}

func TestReadMbs(t *testing.T) {
	test.SetCwdToRoot()
	result, err := Read("mbs_0106")
	if err != nil {
		t.Errorf("failed to read mbs_0106 with error: %q", err.Error())
	}
	fmt.Println(result)
}

func TestReadBerd(t *testing.T) {
	test.SetCwdToRoot()
	result, err := Read("berd_0001")
	if err != nil {
		t.Errorf("failed to read berd_0001 with error: %q", err.Error())
	}
	fmt.Println(result)
}

func TestReadBerdShort(t *testing.T) {
	test.SetCwdToRoot()
	result, err := Read("berd_0006")
	if err != nil {
		t.Errorf("failed to read berd_0006 with error: %q", err.Error())
	}
	fmt.Println(result)
}

func TestReadUkis(t *testing.T) {
	test.SetCwdToRoot()
	result, err := Read("ukis_0002")
	if err != nil {
		t.Errorf("failed to read ukis_0002 with error: %q", err.Error())
	}
	fmt.Println(result)
}

func TestReadStocks(t *testing.T) {
	test.SetCwdToRoot()
	set := map[string]struct{}{}
	for x := 1; x < 10; x++ {
		name := fmt.Sprintf("stocks_000%v", x)
		readStock(name, set)
	}
	more := []int{10, 11, 12, 13, 14, 33, 34, 51, 52, 57, 58, 61, 70}
	for _, y := range more {
		name := fmt.Sprintf("stocks_00%v", y)
		readStock(name, set)
	}
	list := make([]string, len(set))
	i := 0
	for k := range set {
		list[i] = k
		i++
	}
	sort.Strings(list)
	for _, z := range list {
		fmt.Println("\"" + z + "\":\"#" + z + "\",")
	}
}

func readStock(name string, set map[string]struct{}) {
	result, err := Read(name)
	if err != nil {
		fmt.Printf("failed to read %v with error: %q", name, err.Error())
	}
	for _, a := range result.Answers.AnswerMap {
		set[a.Qcode] = struct{}{}
	}
}
