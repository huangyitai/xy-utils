package logx

import (
	"fmt"
	"testing"
	"time"

	"github.com/huangyitai/xy-utils/confx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_callerMarshal(t *testing.T) {
	fmt.Println(callerMarshal(0, "C:/Users/shawnstang/GolandProjects/xy-go/xy-log/example/sample.go", 32))
	fmt.Println(callerMarshal(0, "/opt/xy-go/xy-log/example/sample.go", 32))
	fmt.Println(callerMarshal(0, "sample.go", 32))
	fmt.Println(callerMarshal(0, "/sample.go", 32))
	fmt.Println(callerMarshal(0, "example/sample.go", 32))
}

func TestConfig(t *testing.T) {
	Convey("配置日志", t, func() {
		Convey("读取整个文件", func() {
			err := SetupReadConfig(confx.ReadFile, "test/config.yaml", "yaml")
			So(err, ShouldBeNil)
			log.Info().
				Int("age", 23).Str("name", "Bob").
				Msg("say hello!")
			log.Debug().
				Int("age", 23).Str("name", "Bob").
				Msg("say world!")
		})
		Convey("读取文件配置路径", func() {
			err := SetupReadConfigWithPath(confx.ReadFile, "test/config_path.yaml", "yaml", "xy-go.log")
			So(err, ShouldBeNil)
			log.Info().
				Int("age", 25).Str("name", "Tom").Bytes("bbb", []byte("中文测试😀😀")).EmbedObject(Any("first_name", 123.22)).
				Msg("say hello!")
			log.Debug().
				Int("age", 25).Str("name", "Tom").
				Msg("say world!")
		})
	})
	time.Sleep(time.Millisecond * 500)
	CloseAndWait()
}

func Benchmark_LogToJSONStr(b *testing.B) {
	payload := Config{}
	log.Info().Str("sPayload", ToJSONStr(payload)).Msg("sample")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	for i := 0; i < b.N; i++ {
		log.Info().Str("sPayload", ToJSONStr(payload)).Msg("sample")
	}
}

func Benchmark_LogJSONStr(b *testing.B) {
	payload := Config{}
	log.Info().Stringer("sPayload", JSONStr(payload)).Msg("sample")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	for i := 0; i < b.N; i++ {
		log.Info().Stringer("sPayload", JSONStr(payload)).Msg("")
	}
}
