package schema

import (
	"fmt"
	"sdxImage/pkg/test"
	"testing"
)

func TestCache(t *testing.T) {
	test.SetCwdToRoot()
	schemaCache := NewCache()
	schemaName := "abs_1802"
	schema, err := schemaCache.GetSchema(schemaName)
	if err != nil {
		t.Errorf("failed to get schema: %q with error: %q", schemaName, err.Error())
	}

	titles := schema.ListTitles()
	fmt.Println(titles)
	titles[0] = "Mutated Title!!!"

	schema, err = schemaCache.GetSchema(schemaName)
	if err != nil {
		t.Errorf("failed to get schema: %q with error: %q", schemaName, err.Error())
	}

	if schema.ListTitles()[0] == "Mutated Title!!!" {
		t.Errorf("Able to mutate instruments!")
	}
}
