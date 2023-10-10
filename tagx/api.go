package tagx

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/huangyitai/xy-utils/stringx"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// DefaultSeparator ...
const DefaultSeparator = ";"

// DefaultAssignOp ...
const DefaultAssignOp = "="

// DefaultArraySep ...
const DefaultArraySep = ","

// Parse ...
func Parse(tags reflect.StructTag, key string) []string {
	return ParseWithSeparator(tags, key, DefaultSeparator)
}

// ParseWithSeparator ...
func ParseWithSeparator(tags reflect.StructTag, key, sep string) []string {
	if tag, ok := tags.Lookup(key); ok {
		tag = strings.TrimSpace(tag)
		splits := strings.Split(tag, sep)
		for i := range splits {
			splits[i] = strings.TrimSpace(splits[i])
		}
		return splits
	} else {
		return []string{}
	}
}

// TagInfo ...
type TagInfo struct {
	Ignored    bool
	FieldTable map[string]string
}

// JoinTags ...
func JoinTags(tags ...string) reflect.StructTag {
	return reflect.StructTag(strings.Join(tags, " "))
}

// UnmarshalTags ...
func UnmarshalTags(key string, a interface{}, tags ...string) error {
	return UnmarshalTag(key, a, JoinTags(tags...))
}

// UnmarshalTag ...
func UnmarshalTag(key string, a interface{}, tag reflect.StructTag) error {
	return Unmarshal(tag.Get(key), a)
}

// ParseTagInfo ...
func ParseTagInfo(tag string) *TagInfo {
	res := &TagInfo{
		Ignored:    false,
		FieldTable: map[string]string{},
	}

	tag = strings.TrimSpace(tag)
	if tag == "-" {
		res.Ignored = true
		return res
	}

	parts := strings.Split(tag, DefaultSeparator)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		tokens := strings.Split(part, DefaultAssignOp)
		if len(tokens) == 1 {
			res.FieldTable[tokens[0]] = ""
			continue
		}

		val := tokens[len(tokens)-1]

		for i := 0; i < len(tokens)-1; i++ {
			res.FieldTable[tokens[i]] = val
		}
	}

	return res
}

// Unmarshal ...
func Unmarshal(tag string, a interface{}) error {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return nil
	}

	val := reflect.ValueOf(a)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("a must be a pointer to struct")
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("a must be a pointer to struct")
	}

	info := ParseTagInfo(tag)

	type1 := val.Type()
	for i := 0; i < type1.NumField(); i++ {
		fv := val.Field(i)
		ft := type1.Field(i)

		name := strings.TrimSpace(ft.Tag.Get("tagx"))
		if name == "-" {
			continue
		}
		if name == "" {
			name = ft.Name
		}

		str, ok := info.FieldTable[name]
		if !ok {
			if name == "ignored" || name == "Ignored" {
				if info.Ignored {
					str = "true"
				} else {
					str = "false"
				}
			} else {
				continue
			}
		}

		err := setFieldValue(fv, str)
		if err != nil {
			return err
		}
	}
	return nil
}

func setFieldValue(fv reflect.Value, str string) error {
	switch fv.Interface().(type) {
	case string: //string 类型配置读取
		fv.SetString(str)

	case int, int8, int16, int32, int64: //int 读取
		i64, err := cast.ToInt64E(str)
		if err != nil {
			return errors.WithStack(err)
		}
		fv.SetInt(i64)

	case bool: //bool 读取
		if str == "" {
			str = "true"
		}
		b, err := cast.ToBoolE(str)
		if err != nil {
			return errors.WithStack(err)
		}
		fv.SetBool(b)

	case []string:
		if str == "" {
			fv.Set(reflect.ValueOf([]string{}))
		}
		fv.Set(reflect.ValueOf(stringx.SplitAndTrimSpace(str, DefaultArraySep)))

	case []int:
		if str == "" {
			fv.Set(reflect.ValueOf([]int{}))
		}
		slice, err := cast.ToIntSliceE(stringx.SplitAndTrimSpace(str, DefaultArraySep))
		if err != nil {
			return errors.WithStack(err)
		}
		fv.Set(reflect.ValueOf(slice))

	case []bool:
		if str == "" {
			fv.Set(reflect.ValueOf([]bool{}))
		}
		slice, err := cast.ToBoolSliceE(stringx.SplitAndTrimSpace(str, DefaultArraySep))
		if err != nil {
			return errors.WithStack(err)
		}
		fv.Set(reflect.ValueOf(slice))
	default: //其他类型通过反序列化读取
		return fmt.Errorf("unsupported tagx field type %T", fv.Interface())
	}
	return nil
}
