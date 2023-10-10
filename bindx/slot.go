package bindx

import "reflect"

type slotImpl struct {
	name  string
	value reflect.Value

	order int
	trace bool

	tags reflect.StructTag
}

// Order ...
func (s *slotImpl) Order() int {
	return s.order
}

// Tag ...
func (s *slotImpl) Tag() reflect.StructTag {
	return s.tags
}

// Name ...
func (s *slotImpl) Name() string {
	return s.name
}

// Value ...
func (s *slotImpl) Value() reflect.Value {
	return s.value
}

// Type ...
func (s *slotImpl) Type() reflect.Type {
	return s.value.Type()
}

// WatchType ...
func (s *slotImpl) WatchType() reflect.Type {
	if !isWatchFuncType(s.value.Type()) {
		return nil
	}
	return s.value.Type().Out(0)
}

// Get ...
func (s *slotImpl) Get() interface{} {
	return s.value.Interface()
}

// Set ...
func (s *slotImpl) Set(i interface{}) {
	s.value.Set(reflect.ValueOf(i))
}

// Pointer ...
func (s *slotImpl) Pointer() interface{} {
	return s.value.Addr().Interface()
}

// SetPointer ...
func (s *slotImpl) SetPointer(i interface{}) {
	s.value.Set(reflect.ValueOf(i).Elem())
}

// SetValue ...
func (s *slotImpl) SetValue(v reflect.Value) {
	s.value.Set(v)
}

// SetIndirectValue ...
func (s *slotImpl) SetIndirectValue(v reflect.Value) {
	s.value.Set(reflect.Indirect(v))
}
