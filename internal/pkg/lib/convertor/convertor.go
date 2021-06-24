package convertor

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Transform struct {
	OutputObj  interface{}
	InsertObj  interface{}
	TimeFormat string
}

type Tag struct {
	Key       string
	Value     string
	FiledName string
	Args      []string
}

func NewTransform(outObj, inObj interface{}, timeFormat string) *Transform {
	return &Transform{
		OutputObj:  outObj,
		InsertObj:  inObj,
		TimeFormat: timeFormat,
	}
}

func (t *Transform) Transformer() error {

	if GetValueKind(t.OutputObj) != reflect.Ptr {
		return errors.New("输出数据格式必须是指针")
	}

	if GetValueKind(t.InsertObj) == reflect.Map {
		t.transformerMap()
	} else if GetValueKind(t.InsertObj) == reflect.Ptr {
		t.transformerPtr()
	}

	return nil
}

// ptr 类型数据转换
func (t *Transform) transformerPtr() {
	for i := 0; i < GetValueElem(t.OutputObj).NumField(); i++ {
		of := GetValueElemField(t.OutputObj, i)
		otf := GetValueElemTypeField(t.OutputObj, i)
		if !of.CanSet() {
			fmt.Printf("%v:不能被修改 \n", otf.Name)
			continue
		}

		tag := getTag(otf)
		for iI := 0; iI < GetValueElem(t.InsertObj).NumField(); iI++ {
			inf := GetValueElemField(t.InsertObj, iI)
			into := GetValueElemTypeField(t.InsertObj, iI)
			if tag != nil {
				var args []reflect.Value
				startFunc := false
				// 执行自定义方法
				if len(tag.FiledName) < 1 { // 标签只定义了一个参数，则默认第一个参数为 inf.String()
					args = []reflect.Value{reflect.ValueOf(inf.String())} // append args,first arg is inf.string()
					startFunc = into.Name == otf.Name
				} else {
					startFunc = into.Name == tag.FiledName
					args = append(args, GetValueElem(t.InsertObj).FieldByName(tag.FiledName))
					for _, vt := range tag.Args {
						args = append(args, reflect.ValueOf(vt))
					}
				}
				if tag.Key == "Func" && startFunc {
					rs := t.CallOutFunc(tag).Call(args)
					if rs[0].Interface() != nil {
						of.SetString(rs[0].Interface().(string))
						continue
					}
				}

				if inf.Kind() == reflect.Ptr { // for beego orm
					if into.Name == tag.Key {
						relation := inf.Elem().FieldByName(tag.Value)
						setValue(relation, of)
						continue
					}
				} else if inf.Kind() == reflect.Struct { // for gorm
					if into.Name == tag.Key {
						relation := inf.FieldByName(tag.Value)
						setValue(relation, of)
						continue
					}
				}
			}

			time := otf.Tag.Get("time")
			if len(time) > 0 {
				t.TimeFormat = time
			}

			name := otf.Tag.Get("name")
			if into.Name == otf.Name || into.Name == name {
				if into.Type.Name() == "Time" {
					of.SetString(t.setTime(inf, "", t.TimeFormat))
					continue
				}

				if inf.Type() == of.Type() {
					setValue(inf, of)
					continue
				}
			}
			if into.Name == "BaseModel" { // BaseModel for beego orm
				if otf.Name == "ID" {
					of.SetInt(inf.FieldByName("ID").Interface().(int64))
					continue
				} else if otf.Name == "CreatedAt" {
					of.SetString(t.setTime(inf, "CreatedAt", t.TimeFormat))
					continue
				} else if otf.Name == "UpdatedAt" {
					of.SetString(t.setTime(inf, "UpdatedAt", t.TimeFormat))
					continue
				}
			} else if into.Name == "Model" { //Model for gorm
				if otf.Name == "ID" {
					of.SetInt(int64(inf.FieldByName("ID").Interface().(uint)))
					continue
				} else if otf.Name == "CreatedAt" {
					of.SetString(t.setTime(inf, "CreatedAt", t.TimeFormat))
					continue
				} else if otf.Name == "UpdatedAt" {
					of.SetString(t.setTime(inf, "UpdatedAt", t.TimeFormat))
					continue
				}
			}

		}
	}
}

// map 类型数据转换
func (t *Transform) transformerMap() {
	for i := 0; i < GetValueElem(t.OutputObj).NumField(); i++ {
		of := GetValueElemField(t.OutputObj, i)
		otf := GetValueElemTypeField(t.OutputObj, i)
		if !of.CanSet() {
			fmt.Printf("%v:不能被修改 \n", otf.Name)
			continue
		}

		tag := getTag(otf)
		for _, k := range GetMapKeys(t.InsertObj) {
			inf := GetMapValue(t.InsertObj, k)
			keyName := k.String()

			if tag != nil {
				var args []reflect.Value
				startFunc := false
				// 执行自定义方法
				if len(tag.FiledName) < 1 { // 标签只定义了一个参数，则默认第一个参数为 inf.String()
					args = []reflect.Value{reflect.ValueOf(inf.Interface().(string))} // append args,first arg is inf.string()
					startFunc = keyName == otf.Name
				}

				if tag.Key == "Func" && startFunc {
					rs := t.CallOutFunc(tag).Call(args)
					if rs[0].Interface() != nil {
						of.SetString(rs[0].Interface().(string))
						continue
					}
				}

				if keyName == tag.Key {
					relation := inf.Elem().FieldByName(tag.Value)
					t.setMapValue(relation, of)
					continue
				}
			}

			time := otf.Tag.Get("time")
			if len(time) > 0 {
				t.TimeFormat = time
			}

			name := otf.Tag.Get("name")
			if keyName == otf.Name || keyName == name {
				if otf.Type.Name() == "Time" {
					of.SetString(t.setTime(inf, "", t.TimeFormat))
					continue
				}

				t.setMapValue(inf, of)
				continue
			}

		}
	}
}

// call out func
func (t *Transform) CallOutFunc(tag *Tag) reflect.Value {
	return GetValue(t.OutputObj).MethodByName(tag.Value)
}

// set out map value
func (t *Transform) setMapValue(in, out reflect.Value) {
	switch out.Kind() {
	case reflect.String:
		out.SetString(getMapValueS(in))
	case reflect.Slice:
		reflect.Copy(out, in)
	case reflect.Bool:
		out.SetBool(getMapValueB(in))
	case reflect.Float64, reflect.Float32:
		out.SetFloat(getMapValueF(in))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out.SetInt(getMapValueI(in))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		out.SetUint(getMapValueU(in))
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
	}
}

// time format
func (t *Transform) setTime(inf reflect.Value, fieldName string, timeFormat string) string {

	if inf.IsZero() {
		return ""
	}
	if len(timeFormat) < 1 { // 自定义时间格式
		timeFormat = t.TimeFormat
	}
	args := []reflect.Value{reflect.ValueOf(timeFormat)}
	if len(fieldName) > 0 { // CreatedAt ,UpdatedAt in BaseModel
		inf = inf.FieldByName(fieldName)
	}
	f := inf.MethodByName("Format")
	rs := f.Call(args)

	return rs[0].Interface().(string)
}

// 数据的值
func GetValue(o interface{}) reflect.Value {
	return reflect.ValueOf(o)
}

// 数据的值类型
func GetValueKind(o interface{}) reflect.Kind {
	return reflect.ValueOf(o).Kind()
}

// 数据的值的成员
func GetValueElem(o interface{}) reflect.Value {
	return reflect.ValueOf(o).Elem()
}

// 数据的值的成员类型
func GetValueElemType(o interface{}) reflect.Type {
	return reflect.ValueOf(o).Elem().Type()
}

// 数据的值的成员的值
func GetValueElemField(o interface{}, i int) reflect.Value {
	return reflect.ValueOf(o).Elem().Field(i)
}

// 数据的值的成员类型的值
func GetValueElemTypeField(o interface{}, i int) reflect.StructField {
	return reflect.ValueOf(o).Elem().Type().Field(i)
}

//  map类型数据的 keys
func GetMapKeys(o interface{}) []reflect.Value {
	return reflect.ValueOf(o).MapKeys()
}

// map类型数据的 key 对应数数据
func GetMapValue(o interface{}, key reflect.Value) reflect.Value {
	return reflect.ValueOf(o).MapIndex(key)
}

// set out value
func setValue(in, out reflect.Value) {
	switch out.Kind() {
	case reflect.String:
		out.SetString(in.Interface().(string))
	case reflect.Slice:
		reflect.Copy(out, in)
	case reflect.Bool:
		out.SetBool(in.Bool())
	case reflect.Float64, reflect.Float32:
		out.SetFloat(in.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out.SetInt(in.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		out.SetUint(in.Uint())
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
	}
}

// transform map data
func getMapValueS(in reflect.Value) string {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		return in.Interface().(string)
	case reflect.Bool:
		if in.Bool() {
			return "1"
		} else {
			return "0"
		}
	case reflect.Float64, reflect.Float32:
		return strconv.FormatFloat(in.Float(), 'e', 0, 64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(in.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(in.Uint(), 10)
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return ""
	}

}

// transform map data
func getMapValueB(in reflect.Value) bool {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		if len(in.Interface().(string)) > 0 {
			return true
		} else {
			return false
		}
	case reflect.Bool:
		return in.Interface().(bool)
	case reflect.Float64, reflect.Float32:
		if in.Float() > 0 {
			return true
		} else {
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if in.Int() > 0 {
			return true
		} else {
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if in.Uint() > 0 {
			return true
		} else {
			return false
		}
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return false
	}

}

// transform map data
func getMapValueF(in reflect.Value) float64 {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		f, _ := strconv.ParseFloat(in.Interface().(string), 64)
		return f
	case reflect.Float64, reflect.Float32:
		return in.Float()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(in.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(in.Uint())
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return 0
	}
}

// transform map data
func getMapValueI(in reflect.Value) int64 {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		i, _ := strconv.Atoi(in.Interface().(string))
		return int64(i)
	case reflect.Bool:
		if in.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.Float64, reflect.Float32:
		return int64(in.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return in.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(in.Uint())
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return 0
	}

}

// transform map data
func getMapValueU(in reflect.Value) uint64 {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		i, _ := strconv.ParseUint(in.Interface().(string), 10, 64)
		return i
	case reflect.Bool:
		if in.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.Float64, reflect.Float32:
		return uint64(in.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(in.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return in.Uint()
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return 0
	}

}

// get tag
// first arg is insert object self
func getTag(otf reflect.StructField) *Tag {
	tag := otf.Tag.Get("gtf")
	rtag := strings.Replace(tag, ")", "", 1)
	names := strings.Split(rtag, ".")
	if len(names) <= 1 { // no func name
		return nil
	}

	if !strings.Contains(names[1], "(") { // no args
		return &Tag{Key: names[0], Value: names[1]}
	}

	arg := strings.Split(names[1], "(")
	if len(arg) == 1 { // no args
		return &Tag{Key: names[0], Value: arg[0]}
	}

	args := strings.Split(arg[1], ",") // all args
	if len(args) == 1 {
		return &Tag{Key: names[0], Value: arg[0], FiledName: args[0]}
	}

	return &Tag{Key: names[0], Value: arg[0], FiledName: args[0], Args: args[1:]}

}
