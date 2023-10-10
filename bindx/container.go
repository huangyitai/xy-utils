package bindx

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"sync/atomic"

	"github.com/huangyitai/xy-utils/deferx"
	"github.com/huangyitai/xy-utils/dox"
	"github.com/huangyitai/xy-utils/tagx"
	"github.com/rs/zerolog/log"
)

type binding struct {
	name   string
	binder Binder
	value  *reflect.Value
	slots  []Slot
	bound  bool
	watch  bool
}

type orderedBinder struct {
	Binder
	order int
}

// Order ...
func (o *orderedBinder) Order() int {
	return o.order
}

type orderedWatchBinder struct {
	WatchBinder
	order int
}

// Order ...
func (o *orderedWatchBinder) Order() int {
	return o.order
}

// Container ...
type Container struct {
	bindingTable map[string]map[reflect.Type]*binding
	typeConv     map[string]map[reflect.Type]reflect.Type
	binderSet    map[Binder]bool

	binders  []OrderedBinder
	slots    []Slot
	bindings []*binding
}

// AddOne ...
func (c *Container) AddOne(name string, i interface{}, tags ...string) error {
	tag := tagx.JoinTags(tags...)
	val := reflect.ValueOf(i)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("i must be a valid pointer")
	}

	info := TagInfo{}
	err := tagx.UnmarshalTag(TagKey, &info, tag)
	if err != nil {
		return err
	}

	//任何情况下都会使用name，bindx.Name只是用于结构体字段初始化时指定和字段名不同的name
	if info.Name != "" && info.Name != name {
		log.Warn().Str("sName", name).Str("sTagName", info.Name).
			Msg("[AddOne] conflict between name and tag name")
	}

	if info.Ignored {
		return nil
	}

	c.slots = append(c.slots, &slotImpl{
		name:  name,
		value: val.Elem(),
		order: info.Order,
		tags:  tag,
	})
	return nil
}

// AddCombo ...
func (c *Container) AddCombo(a ...interface{}) error {
	for _, p := range a {
		val := reflect.ValueOf(p)
		if val.Kind() != reflect.Ptr {
			return fmt.Errorf("combo must be a pointer to struct")
		}

		v := val.Elem()
		if v.Kind() != reflect.Struct {
			return fmt.Errorf("combo must be a pointer to struct")
		}
	}

	for _, p := range a {
		v := reflect.ValueOf(p).Elem()
		t := v.Type()

		for i := 0; i < t.NumField(); i++ {
			fv := v.Field(i)
			ft := t.Field(i)

			//跳过embed字段
			//if ft.Anonymous {
			//	continue
			//}

			name := ft.Name

			info := TagInfo{}
			err := tagx.UnmarshalTag(TagKey, &info, ft.Tag)
			if err != nil {
				return err
			}

			//如果bindx的值为"-"，则表示跳过该字段
			if info.Ignored {
				continue
			}

			if info.Name != "" {
				name = info.Name
			}

			c.slots = append(c.slots, &slotImpl{
				name:  name,
				value: fv,
				order: info.Order,
				tags:  ft.Tag,
			})
		}
	}
	return nil
}

// AddBinder ...
func (c *Container) AddBinder(binder Binder) {
	if binder == nil {
		return
	}

	//binder简单去重
	if c.binderSet[binder] {
		return
	}
	c.binderSet[binder] = true

	if i, ok := binder.(OrderedBinder); ok {
		c.binders = append(c.binders, i)
	} else {
		if w, ok := binder.(WatchBinder); ok {
			c.binders = append(c.binders, &orderedWatchBinder{
				WatchBinder: w,
				order:       0,
			})
		} else {
			c.binders = append(c.binders, &orderedBinder{
				Binder: binder,
				order:  0,
			})
		}
	}
}

// AddBinderWithOrder ...
func (c *Container) AddBinderWithOrder(binder Binder, order int) {
	c.binders = append(c.binders, &orderedBinder{
		Binder: binder,
		order:  order,
	})
}

// Bind ...
// 理一下顺序
// 1.首先应该收集所有能绑定的类型，建立类型和binder的单射表binders，并根据最具类型原则简化binders
// 2.顺序遍历所有需要绑定的对象Slot，根据binders选择binder并建立binding表bindings
// 3.遍历bindings，对binding进行绑定
// 1.1 遍历Slot，将其中实现了Binder接口的对象注册成为Binder
//
//	for _, slot := range c.slots {
//		if slot.Value().Type().AssignableTo(binderType) {
//			c.AddBinder(slot.Value().Type(), reflect.New(slot.Value().Type()).Elem().Interface().(Binder))
//		}
//	}
func (c *Container) Bind(ctx context.Context) (err error) {
	defer deferx.PanicToError(&err)

	//1 sort
	c.sortBinderAndSlotByOrder()

	//2.1 建立binding表
	err = c.initBinding()
	if err != nil {
		return err
	}

	//3
	for i := 0; i < len(c.bindings); i++ {
		b := c.bindings[i]
		if len(b.slots) == 0 {
			continue
		}

		if b.binder.Cacheable() {
			if b.bound {
				continue
			} else {
				slot := b.slots[0]
				if b.watch {
					watchBinder, ok := b.binder.(WatchBinder)
					if !ok {
						return fmt.Errorf("illegal state, binding.watch = true, binding.binder is not WatchBinder")
					}

					atom := &atomic.Value{}
					for _, s := range b.slots {
						//构造watch型的生成函数
						v := makeWatchFunc(s.Type(), atom)
						s.SetValue(v)
					}

					err = watchBinder.Watch(ctx, c, slot, dox.NoopRun, atom)
					if err != nil {
						return err
					}

					vAtom := reflect.ValueOf(atom)
					b.value = &vAtom
				} else {
					err = b.binder.Bind(ctx, c, slot)
					if err != nil {
						return err
					}

					v := slot.Value()
					//其余slot值注入
					for i := 1; i < len(b.slots); i++ {
						b.slots[i].Value().Set(v)
					}
					b.value = &v

				}
				b.bound = true
			}
		} else {
			if b.watch {
				watchBinder, ok := b.binder.(WatchBinder)
				if !ok {
					return fmt.Errorf("illegal state, binding.watch = true, binding.binder is not WatchBinder")
				}

				for _, slot := range b.slots {
					atom := &atomic.Value{}
					slot.SetValue(makeWatchFunc(slot.Type(), atom))

					err = watchBinder.Watch(ctx, c, slot, dox.NoopRun, atom)
					if err != nil {
						return err
					}
				}
			} else {
				for _, slot := range b.slots {
					err = b.binder.Bind(ctx, c, slot)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (c *Container) sortBinderAndSlotByOrder() {
	//1.1 对binders进行排序
	sort.SliceStable(c.binders, func(i, j int) bool {
		return c.binders[i].Order() < c.binders[j].Order()
	})

	//1.2 对slots进行排序
	sort.SliceStable(c.slots, func(i, j int) bool {
		if c.slots[i].Order() != c.slots[j].Order() {
			return c.slots[i].Order() < c.slots[j].Order()
		}
		if c.slots[i].Name() != c.slots[j].Name() || c.slots[i].Type() == c.slots[j].Type() {
			return false
		}
		return c.slots[i].Type().AssignableTo(c.slots[j].Type())
	})
}

func (c *Container) initBinding() error {
	for i, slot := range c.slots {
		log.Trace().
			Str("sName", slot.Name()).Str("sTag", string(slot.Tag())).Str("sType", slot.Type().String()).
			Int("iIndex", i).
			Msg("[bindx]add binding")
		err := c.addBinding(slot)
		if err != nil {
			return err
		}
	}
	return nil
}

// BindOne ...
func (c *Container) BindOne(ctx context.Context, name string, i interface{}, tags ...string) error {
	return c.WatchOne(ctx, name, i, dox.NoopRun, tags...)
}

// WatchOne ...
func (c *Container) WatchOne(ctx context.Context, name string, i interface{}, callback dox.Run, tags ...string) error {
	tag := tagx.JoinTags(tags...)
	val := reflect.ValueOf(i)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("slot must be a pointer")
	}

	info := TagInfo{}
	err := tagx.UnmarshalTag(TagKey, &info, tag)
	if err != nil {
		return err
	}

	//该处验证实际上意义不大，实际上任何情况都应该使用name，bindx.Name只是用于初始化绑定时指定name
	if info.Name != "" && info.Name != name {
		log.Warn().Str("sName", name).Str("sTagName", info.Name).
			Msg("[WatchOne] conflict between name and tag name")
	}

	if info.Ignored {
		return nil
	}

	val = val.Elem()

	var slot Slot = &slotImpl{
		name:  name,
		value: val,
		order: 0,
		trace: false,
		tags:  tag,
	}

	b, err := c.getBindings(slot, false)
	if err != nil {
		return err
	}

	if b.binder.Cacheable() {
		err = c.watchCacheable(ctx, b, slot, val, callback)
		if err != nil {
			return err
		}
	} else {
		if b.watch {
			watchBinder, ok := b.binder.(WatchBinder)
			if !ok {
				return fmt.Errorf("illegal state, binding.watch = true, binding.binder is not WatchBinder")
			}

			//先设置监听函数值，再调用Watch
			atom := &atomic.Value{}
			slot.SetValue(makeWatchFunc(slot.Type(), atom))

			err = watchBinder.Watch(ctx, c, slot, callback, atom)
			if err != nil {
				return err
			}
		} else {
			err = b.binder.Bind(ctx, c, slot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Container) watchCacheable(ctx context.Context, b *binding, slot Slot, val reflect.Value,
	callback dox.Run) error {

	//判断单例是否已初始化
	if b.bound {
		if b.watch {
			atom, ok := b.value.Interface().(*atomic.Value)
			if !ok {
				return fmt.Errorf("illegal state, binding.watch = true, binding.value is not *atomic.Value")
			}

			//先设置函数体，然后应该立刻调用一次callback函数
			v := makeWatchFunc(slot.Type(), atom)
			val.Set(v)

			//手动调用一次callback，以完成后续初始化
			err := callback()
			if err != nil {
				return err
			}
		} else {
			val.Set(*b.value)
		}
		return nil
	} else {
		//单例未初始化

		if len(b.slots) != 0 {
			//binding中有外部slot

			slot = b.slots[0]
		}

		if b.watch {
			watchBinder, ok := b.binder.(WatchBinder)
			if !ok {
				return fmt.Errorf("illegal state, binding.watch = true, binding.binder is not WatchBinder")
			}

			//先构造函数体，并设置函数值，然后进行Watch，尽量避免空函数问题
			atom := &atomic.Value{}
			for _, s := range b.slots {
				//构造watch型的生成函数，每个slot的函数类型可能不同，因此需要每个单独生成
				v := makeWatchFunc(s.Type(), atom)
				s.SetValue(v)
			}
			val.Set(makeWatchFunc(val.Type(), atom))

			err := watchBinder.Watch(ctx, c, slot, callback, atom)
			if err != nil {
				return err
			}

			vAtom := reflect.ValueOf(atom)
			b.value = &vAtom

		} else {
			if len(b.slots) != 0 { //binding中有外部slot
				slot = b.slots[0]
			}

			err := b.binder.Bind(ctx, c, slot)
			if err != nil {
				return err
			}

			v := slot.Value()
			//其余slot值注入
			for i := 1; i < len(b.slots); i++ {
				b.slots[i].Value().Set(v)
			}
			b.value = &v

			if len(b.slots) != 0 {
				val.Set(v)
			}
		}

		b.bound = true
	}
	return nil
}

// GetValue ...
func (c *Container) GetValue(ctx context.Context, name string, t reflect.Type, tags ...string) (reflect.Value, error) {
	p := reflect.New(t).Interface()
	err := c.BindOne(ctx, name, p, tags...)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(p).Elem(), nil
}

func makeWatchFunc(t reflect.Type, val *atomic.Value) reflect.Value {
	return reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
		i := val.Load()
		if i == nil {
			return []reflect.Value{reflect.Zero(t.Out(0))}
		}
		return []reflect.Value{reflect.ValueOf(i)}
	})
}

func (c *Container) addBinding(slot Slot) error {
	_, err := c.getBindings(slot, true)
	return err
}

func (c *Container) getWatchBinder(slot Slot) (Binder, error) {
	if !isWatchFuncType(slot.Type()) {
		return nil, fmt.Errorf("[%s] is not watch func type", slot.Type())
	}

	rt := slot.Type().Out(0)

	tempSlot := slotImpl{
		name:  slot.Name(),
		value: reflect.Zero(rt),
		order: slot.Order(),
		trace: false,
		tags:  slot.Tag(),
	}

	for _, binder := range c.binders {
		//首先筛选出支持Watch的Binder
		if _, ok := binder.(WatchBinder); !ok {
			continue
		}

		if binder.Bindable(&tempSlot) {
			return binder, nil
		}
	}
	return noopBinder{}, nil
}

func (c *Container) getConcreteType(name string, t reflect.Type) (reflect.Type, error) {
	t2t, ok := c.typeConv[name]
	if !ok {
		t2t = map[reflect.Type]reflect.Type{}
		c.typeConv[name] = t2t
	}

	ct, ok := t2t[t]
	if ok {
		return ct, nil
	}

	ct = t
	for t2 := range t2t {
		if t2 != t && t2.AssignableTo(t) { //t2是t的子类型
			if t2.AssignableTo(ct) { //找到比当前dt更具体的t2
				ct = t2t[t2]
			} else {
				if ct.AssignableTo(t2) { //找到了当前dt的父类t2
					continue
				} else { //找到了当前dt的兄弟类型，无法选择
					return nil, fmt.Errorf("mutiple choice of %s: [%s], [%s]", t, ct, t2)
				}
			}
		}
	}

	//对于函数类型:没有输入参数，只返回一个结果，进行特殊子类型判断
	if ct == t && isWatchFuncType(t) {
		rt := t.Out(0)
		crt := ct.Out(0)
		for t2 := range t2t {
			if !isWatchFuncType(t2) {
				continue
			}
			rt2 := t2.Out(0)
			if t2 != t && rt2.AssignableTo(rt) {
				//t2的返回类型是 t返回类型的 子类型

				if rt2.AssignableTo(crt) {
					//t2的返回类型是 ct返回类型的子类型

					ct = t2t[t2]
					crt = ct.Out(0)
				} else {
					if crt.AssignableTo(rt2) {
						//t2的返回类型是 ct返回类型的父类型
						continue
					} else {
						//t2的返回类型是 ct返回类型的兄弟类型，无法选择
						return nil, fmt.Errorf("multiple choice of watch %s: [%s], [%s]", t.Out(0), crt, rt2)
					}
				}
			}
		}
	}

	t2t[t] = ct

	return ct, nil
}

func isWatchFuncType(t reflect.Type) bool {
	return t.Kind() == reflect.Func && t.NumIn() == 0 && t.NumOut() == 1
}

func (c *Container) getBinder(slot Slot) (Binder, error) {
	for _, binder := range c.binders {
		if binder.Bindable(slot) {
			return binder, nil
		}
	}
	return noopBinder{}, nil
}

func (c *Container) getBindings(slot Slot, appendSlot bool) (*binding, error) {
	t2b, ok := c.bindingTable[slot.Name()]
	if !ok {
		t2b = map[reflect.Type]*binding{}
		c.bindingTable[slot.Name()] = t2b
	}

	ct, err := c.getConcreteType(slot.Name(), slot.Type())
	if err != nil {
		return nil, err
	}

	b, ok := t2b[ct]
	if ok {
		if appendSlot {
			b.slots = append(b.slots, slot)
		}
		return b, nil
	}

	binder, err := c.getBinder(slot)
	if err != nil {
		return nil, err
	}

	watch := false

	//如果没有找到合适的binder，判断是否为watch函数类型
	if _, ok := binder.(noopBinder); ok && isWatchFuncType(slot.Type()) {
		//寻找watch函数的binder
		binder, err = c.getWatchBinder(slot)
		if err != nil {
			return nil, err
		}
		watch = true
	}

	b = &binding{
		name:   slot.Name(),
		binder: binder,
		value:  nil,
		slots:  []Slot{},
		bound:  false,
		watch:  watch,
	}

	if appendSlot {
		b.slots = []Slot{slot}
	}

	t2b[ct] = b
	c.bindings = append(c.bindings, b)

	return b, nil
}
