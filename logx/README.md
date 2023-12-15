# xy-essential/logx 日志工具（原 xy-log 日志框架）
logx 是基于zerolog 日志库开发的日志框架
>https://github.com/rs/zerolog
> <br> 专门针对json格式设计，链式日志api，池化日志buffer，日志内容流式写入buffer，减少序列化计算和内存的开销



>结合diode支持线程安全无锁非阻塞的console日志输出
><br>https://github.com/cloudfoundry/go-diodes


>结合trpc-log框架中rollwriter实现了线程安全无锁非阻塞的文件滚动日志输出
><br>https://git.code.oa.com/trpc-go/trpc-go/tree/master/log/rollwriter

>结合zhiyan日志sdk+channel缓冲实现了异步的智研日志输出
> <br>https://git.code.oa.com/zhiyan-log/sdk-go

## 功能
- 基于zerolog提供链式日志api，支持logger与context绑定（filter、middleware中添加固定业务字段）
- 一“键”完成配置，全项目可用，不依赖trpc等其他特定框架，可搭配任意框架使用
- 支持日志内置字段名定义，时间戳字段自定义格式化
- 支持多writer输出，现支持file、console、zhiyan三种输出方式
- 提供异步非阻塞的支持，启用fast输出模式（缓存队列写满时丢弃日志）的情况下可以实现日志对业务逻辑的最小化影响

## 修改记录

### 2021年6月11日整合调整

- xy-log 日志框架现已整合进入 xy-essential 中，作为日志工具模块 logx

### v0.0.8

- 为日志增加全局自增offset功能，便于向智妍等平台上报时确认日志顺序

### v0.0.7

- 调整默认日志时间格式，现在默认打印到微秒单位

### v0.0.6

- 配置方法调整
- 支持多种格式配置
- 支持配置路径
- 修复Close方法的问题
- 添加了设置子logger的帮助方法

### v0.0.3-v0.0.5

- starter问题修复与优化

### v0.0.2

- 添加了starter集成支持
- **包名更改为logx**
### v0.0.1
- 完成框架配置解析
- 三种writer实现

## 使用方式
添加依赖
```shell
go get git.woa.com/xy-vip/utils/xy-go/xy-essential/logx
```

### 配置说明
简化配置信息
```yaml
global:
  fields:
    "@podid": $POD_IP
writer:
  - type: console
    level: debug
  - type: file
    level: info
    log_path: /usr/local/xy/log
    filename: xy.log
    roll_type: time
    max_age: 7
    time_unit: day
    max_size_mb: 10
  - type: zhiyan
    level: debug
    report_topic: fb_xxxxxx
```
全部配置项默认值及说明
```yaml
#全局日志配置
global:
  #关闭自动打印caller信息，设置为true时关闭打印
  disable_caller: false
  #关闭自动打印错误栈信息，设置为true时关闭打印，需要error中包含调用栈信息
  disable_error_stack: false
  #关闭自动打印时间信息，设置为true时关闭打印
  disable_timestamp: false
  #自定义默认打印字段，打印格式均为string
  fields:
    #样例，yaml语法： @开头的字符串key或者值均需要用""包装
    "@podip": $POD_IP
    env_name: production
    #支持用$VAR的语法插入单个环境变量VAR，使用${VAR}的形式将环境变量VAR的值插入字符串中
    app_name: sample${ENV}server

#内置字段key格式
format:
  #时间戳字段key
  time_key: "@time"
  #时间戳format格式，""表示采用unix时间戳，UNIXMS表示采用unix毫秒，UNIXMICRO表示采用unix微秒
  time_field_format: 2006-01-02 15:04:05.000
  #日志等级字段key
  level_key: "@level"
  #caller字段key
  caller_key: "@caller"
  #错误调用栈字段key
  stacktrace_key: "@stacktrace"
  #message字段key
  message_key: "@message"
  #error字段key
  error_key: "@error"

#输出writer配置
writer:
    #writer类型，必填 值域：file-文件 console-命令行 zhiyan-智研日志汇
  - type: console
    #输出日志等级，必填 值域：trace debug info warn error fatal panic ""-表示不输出
    level: debug
    #console专用，true表示按key=value方式美化命令行输出
    pretty: false
    #console专用，true表示同步打印日志，将会进行加锁防止竞争，但多个console输出同步打印可能会出现竞争问题
    #false表示异步打印日志，并且会在队列满时丢弃日志
    write_sync: false
    #异步日志缓冲队列大小
    queue_size: 10000
    #console专用，从缓冲队列中取消息写入命令行的间隔时间，大于0将采用轮询队列输出模式，设置为0表示采用sync.Cond使消费者等待队列可用
    write_interval_ms: 0

  - type: file
    level: info
    #file专用，表示文件输出的路径
    log_path: /usr/local/xy/log
    #file专用，表示输出的文件名称（前缀）
    filename: xy.log
    #file专用，表示输出是否进行压缩
    compress: false
    #file专用，表示输出方式，值域： sync-同步写文件 async-异步写文件 fast-异步快速模式，队列满时丢弃日志
    write_mode: async
    queue_size: 10000
    #file专用，写文件缓冲，单位byte，缓冲大小超过此值时进行写文件
    buffer_size: 4096
    #file专用，表示写文件的间隔时间，可选值 >0
    write_interval_ms: 100
    #file专用，表示文件滚动方式 值域： time-按时间滚动 size-按文件大小滚动
    roll_type: time
    #file专用，保留额外备份文件的最久时间
    max_age: 0
    #file专用，表示滚动或者备份文件的时间单位 可选 minute hour day month year
    time_unit: day
    #file专用，表示保留备份文件的最大数量
    max_backups: 0
    #file专用，最大单个日志文件大小，（可能会稍微超过此限制）
    max_size_mb: 100

  - type: zhiyan
    level: debug
    #zhiyan专用，必填，智研日志汇创建sdk接入点时自动生成的topic
    report_topic: fb-fc41e4afecf298e
    #zhiyan专用，上报协议，tcp或udp
    report_proto: tcp
    #zhiyan专用，上报服务器地址，""表示自动解析北极星获取地址
    report_addr: ""
    #zhiyan专用，上报日志时主机标识，为空会自动填入服务所在机器ip
    report_host: ""
    #zhiyan专用，启动日志客户端数量，当出现日志输出不及时，丢日志时可以适当调高
    zhiyan_client_num: 3
    #zhiyan专用，表示writer写入队列时是否等待队列空闲，true时可能会阻塞业务，false时可能会丢日志
    wait_queue: false
    queue_size: 10000
```


【推荐】使用方式
```go
import "git.woa.com/xy-vip/utils/xy-go/xy-essential/logx"
import "github.com/rs/zerolog/log"

//---配置部分---

//从xy_log.yaml文件读取日志配置（使用confx包）
logx.SetupReadConfig(confx.ReadFile, "xy_log.yaml", confx.YAML)
//或不使用confx包
logx.SetupReadConfig(ioutil.ReadFile, "xy_log.yaml", "yaml")

//从文件读取日志配置，带key路径（从文件中指定key路径下读取配置）
logx.SetupReadConfigWithPath(ioutil.ReadFile, "trpc_go.yaml", "yaml", "xy-go.log")

//从字符串配置
logx.SetupFromString(str, "yaml")
logx.SetupFromStringWithPath(str, "yaml", "xy-go.log")

//从结构体配置
cfg := new(logx.Config)
//……
logx.SetupFromConfig(cfg)

//---使用部分---
log.Info().Str("user_id", "xxxxxx").Int("user_level", 2).Msg("login!")
log.Info().Str("user_id", "xxxxxx").Int("user_level", 2).Err(err).Send()
log.Ctx(ctx).Info().Msg("end")
//...... 其他方法详见https://github.com/rs/zerolog
```
基本注意点是链式API通过Msg、Msgf、Send方法结束才会将日志输出，否则不会输出

### 性能Tips
zerolog是结构化日志框架，因此输出字段或者参数信息时请尽量使用类型特定的方法，如Str、Int、Ints、Float64、Time等，对于非结构体和map大量使用Interface（使用反射序列化）会带来不必要的性能开销。

注意采用了链式方法输出各种字段信息后，可以尽量避免Msgf这种消息输出方式，减少fmt.Sprintf带来的性能影响

