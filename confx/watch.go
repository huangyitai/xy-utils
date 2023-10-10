package confx

import (
	"bytes"
	"reflect"
	"sync"

	"github.com/huangyitai/xy-utils/dox"
	"github.com/huangyitai/xy-utils/xxx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type mapChannelBlock struct {
	ch       chan map[string][]byte
	val      map[string][]byte
	init     bool
	end      bool
	lock     sync.Locker
	loadCond *sync.Cond
}

// WatchMapChannel ...
func WatchMapChannel(ch chan map[string][]byte) WatchFunc {
	var lock sync.Mutex
	b := mapChannelBlock{
		ch:       ch,
		val:      nil,
		lock:     &lock,
		loadCond: sync.NewCond(&lock),
	}

	dox.Do(func() {
		for m := range ch {
			b.lock.Lock()
			b.val = m
			b.init = true
			b.lock.Unlock()
			b.loadCond.Broadcast()
		}
		b.lock.Lock()
		b.end = true
		b.loadCond.Broadcast()
		b.lock.Unlock()
	}).Daemon()

	return func(s string, f func([]byte) error) error {
		b.loadCond.L.Lock()
		defer b.loadCond.L.Unlock()

		//判断是否已经有初值
		if !b.init {
			//没有初值则开始等待（如果在watch之前没有传入初值，会进入死等）
			b.loadCond.Wait()
			end := b.end
			//等到ch被关闭，直接返回
			if end {
				return nil
			}
		}

		//已经有初值，进行首次初始化
		old := b.val[s]
		err := f(old)
		if err != nil {
			return err
		}

		//开始监听
		dox.Do(func() {
			for {
				end := false
				func() {
					b.loadCond.L.Lock()
					defer b.loadCond.L.Unlock()
					b.loadCond.Wait()
					end = b.end
					if end {
						return
					}

					cur := b.val[s]
					if !bytes.Equal(old, cur) {
						err = f(cur)
						if err != nil {
							log.Err(err).Send()
						}
						old = cur
					}
				}()
				if end {
					break
				}
			}
		}).Daemon()
		return nil
	}
}

// NewMapChannelWatch ...
func NewMapChannelWatch() (WatchFunc, chan map[string][]byte) {
	ch := make(chan map[string][]byte, 2)
	return WatchMapChannel(ch), ch
}

var watchFuncTable = map[string]WatchFunc{}

// RegisterWatch ...
func RegisterWatch(name string, watch WatchFunc) {
	watchFuncTable[name] = watch
}

// WatchFunc ...
type WatchFunc func(string, func([]byte) error) error

// WatchValue ...
func (w WatchFunc) WatchValue(key string, watchType reflect.Type, callback func(value reflect.Value) error,
	format string, opts ...viper.DecoderConfigOption) error {
	sign := xxx.NewSignStr().WithPath("confx", "WatchValue").WithProp("key", key)
	return w(key, func(bs []byte) error {
		var v reflect.Value
		if watchType.Kind() == reflect.Ptr {
			v = reflect.New(watchType.Elem())
		} else {
			v = reflect.New(watchType).Elem()
		}

		err := ReadValue(bs, v, format, opts...)
		if err != nil {
			log.Err(err).
				Str("sKey", key).Bytes("sBytes", bs).
				Msgf("%s ReadValueWithPath fail", sign)
			return err
		}

		err = callback(v)
		if err != nil {
			log.Err(err).Msgf("%s callback fail", sign)
			return err
		}
		log.Debug().Msgf("%s updated", sign)
		log.Trace().Str("sKey", key).Bytes("sBytes", bs).Msgf("%s updated detail", sign)
		return nil
	})
}

// WatchValueWithPath ...
func (w WatchFunc) WatchValueWithPath(key string, watchType reflect.Type, callback func(value reflect.Value) error,
	format, path string, opts ...viper.DecoderConfigOption) error {
	sign := xxx.NewSignStr().WithPath("confx", "WatchValue").WithProp("key", key)
	return w(key, func(bs []byte) error {
		var v reflect.Value
		if watchType.Kind() == reflect.Ptr {
			v = reflect.New(watchType.Elem())
		} else {
			v = reflect.New(watchType).Elem()
		}

		err := ReadValueWithPath(bs, v, format, path, opts...)
		if err != nil {
			log.Err(err).
				Str("sKey", key).Bytes("sBytes", bs).
				Msgf("%s ReadValueWithPath fail", sign)
			return err
		}

		err = callback(v)
		if err != nil {
			log.Err(err).Msgf("%s callback fail", sign)
			return err
		}
		log.Debug().Msgf("%s updated", sign)
		log.Trace().Str("sKey", key).Bytes("sBytes", bs).Msgf("%s updated detail", sign)
		return nil
	})
}
