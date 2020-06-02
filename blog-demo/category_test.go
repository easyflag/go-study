package db

import (
	"testing"
)

func init() {
	err := Init()
	if err != nil {
		panic(err)
	}
}

func TestGetCategoryList(t *testing.T) {
	list, err := GetCategoryList()
	if err != nil {
		t.Errorf("get category list failed,error:%v\n", err)
	}

	for _, v := range list {
		t.Logf("%#v\n", *v)
	}
}

func TestSelectCategory(t *testing.T) {
	var IDs [5]uint8
	for i := range IDs {
		IDs[i] = uint8(i)
	}

	categories, err := SelectCategory(IDs[:], "class")
	if err != nil {
		t.Errorf("get category info failed,error:%v\n", err)
	}

	t.Logf("%#v\n", categories[0])
}
