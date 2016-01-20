package toolkit

import (
	"reflect"
	"strings"
	"time"
	//"errors"
)

func VariadicToSlice(objs ...interface{}) *[]interface{} {
	result := []interface{}{}
	for _, v := range objs {
		result = append(result, v)
	}
	return &result
}

func MapToSlice(objects map[string]interface{}) []interface{} {
	results := make([]interface{}, 0)
	for _, v := range objects {
		results = append(results, v)
	}
	return results
}

func HasMember(g interface{}, find interface{}) bool {
	found := false
	if IsSlice(g) == false {
		return false
	}

	count := SliceLen(g)
	for i := 0; i < count; i++ {
		v := SliceItem(g, i)
		eq := Compare(v, find, "$eq")
		//Println("L:", v, " F:", find, ", B:", eq, reflect.ValueOf(find).Type().String())
		if eq {
			return true
		}
	}
	return found
}

func ToInterfaceArray(o interface{}) []interface{} {
	if IsSlice(o) == false {
		return []interface{}{}
	}

	//Printf("Slice Data: %s\n", JsonString(o))
	var ret []interface{}
	for i := 0; i < SliceLen(o); i++ {
		sliceItem := SliceItem(o, i)
		//Printf("%d Item: %s\n", i, JsonString(sliceItem))
		ret = append(ret, sliceItem)
	}
	return ret
}

func Compare(v1 interface{}, v2 interface{}, op string) bool {
	vv1 := reflect.Indirect(reflect.ValueOf(v1))
	vv2 := reflect.Indirect(reflect.ValueOf(v2))
	if vv1.Type().String() != vv2.Type().String() {
		return false
	}

	k := strings.ToLower(vv1.Kind().String())
	t := strings.ToLower(vv1.Type().String())
	if strings.Contains(k, "int") || strings.Contains(k, "float") {
		//--- is a number
		// lets convert all to float64 for simplicity
		var vv1o, vv2o float64

		if strings.Contains(k, "int") {
			vv1o = float64(vv1.Int())
			vv2o = float64(vv2.Int())
		} else {
			vv1o = vv1.Float()
			vv2o = vv2.Float()
		}

		//vv1o = ToFloat64(vv1)
		//vv2o = ToFloat64(vv2)
		if op == "$eq" {
			return vv1o == vv2o
		} else if op == "$ne" {
			return vv1o != vv2o
		} else if op == "$lt" {
			return vv1o < vv2o
		} else if op == "$lte" {
			return vv1o <= vv2o
		} else if op == "$gt" {
			return vv1o > vv2o
		} else if op == "$gte" {
			return vv1o >= vv2o
		}
	} else if strings.Contains(t, "time.time") {
		//--- is a time.Time
		vv1o := vv1.Interface().(time.Time)
		vv2o := vv2.Interface().(time.Time)
		if op == "$eq" {
			return vv1o == vv2o
		} else if op == "$ne" {
			return vv1o != vv2o
		} else if op == "$lt" {
			return vv1o.Before(vv2o)
		} else if op == "$lte" {
			return vv1o == vv2o || vv1o.Before(vv2o)
		} else if op == "$gt" {
			return vv1o.After(vv2o)
		} else if op == "$gte" {
			return vv1o == vv2o || vv1o.After(vv2o)
		}

	} else if strings.Contains(t, "bool") {
		vv1o := vv1.Bool()
		vv2o := vv2.Bool()
		if op == "$eq" {
			return vv1o == vv2o
		} else if op == "$ne" {
			return vv1o != vv2o
		}
	} else {
		//--- will be string
		vv1o := vv1.Interface().(string)
		vv2o := vv2.Interface().(string)
		if op == "$eq" {
			return vv1o == vv2o
		} else if op == "$ne" {
			return vv1o != vv2o
		} else if op == "$lt" {
			return vv1o < vv2o
		} else if op == "$lte" {
			return vv1o <= vv2o
		} else if op == "$gt" {
			return vv1o > vv2o
		} else if op == "$gte" {
			return vv1o >= vv2o
		}
	}

	return false
}