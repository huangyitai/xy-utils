package xxx

// NoopSink 用于接收任意类型，任意数量参数，不进行任何操作，用于解决编写examples时演示参数未被使用，报错的问题
func NoopSink(objs ...interface{}) {}
