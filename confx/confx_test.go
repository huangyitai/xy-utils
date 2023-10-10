package confx

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/huangyitai/xy-utils/bindx"
	"github.com/huangyitai/xy-utils/testx"
	"github.com/mitchellh/mapstructure"
	. "github.com/smartystreets/goconvey/convey"
)

// Person ...
type Person struct {
	Name   string
	Age    int
	OpenID int `mapstructure:"open_id"`
	EMail  string
}

func (p *Person) Hello() string {
	return "Hello " + p.Name
}

type HelloAware interface {
	Hello() string
}

func TestReadConfig(t *testing.T) {

	type Person struct {
		Name   string
		Age    int
		OpenID int `mapstructure:"open_id"`
		EMail  string
	}

	type Config struct {
		P1 Person                 `confx:"name=test/person.json"`
		P2 Person                 `confx:"name=test/person.yaml;format=yaml"`
		P3 Person                 `confx:"name=test/person_path.json;path=AAA.BBB"`
		P4 Person                 `confx:"name=test/person_path.yaml;format=yaml;path=AAA.BBB"`
		P5 map[string]interface{} `confx:"name=test/person_path.yaml;format=yaml;path=AAA.BBB; binding"`
	}

	Convey("读取配置", t, func() {
		Convey("不带路径 json", func() {
			p := Person{}
			err := ReadFile.Unmarshal("test/person.json", &p, "json")
			Printf("%+v\n", p)

			So(err, ShouldBeNil)
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")

			pm := map[string]interface{}{}
			err = ReadFile.Unmarshal("test/person.json", &pm, "json")
			fmt.Println("PM")
			testx.PrintJSONPretty(pm)
		})

		Convey("不带路径 yaml", func() {
			p := Person{}
			err := ReadFile.Unmarshal("test/person.yaml", &p, "yaml")
			Printf("%+v\n", p)

			So(err, ShouldBeNil)
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")
		})

		Convey("带路径 json", func() {
			p := Person{}
			err := ReadFile.UnmarshalWithPath("test/person_path.json", &p, "json", "AAA.BBB")
			Printf("%+v\n", p)

			So(err, ShouldBeNil)
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")
		})

		Convey("带路径 yaml", func() {
			p := Person{}
			err := ReadFile.UnmarshalWithPath("test/person_path.yaml", &p, "yaml", "AAA.BBB")
			Printf("%+v\n", p)

			So(err, ShouldBeNil)
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")
		})

		Convey("Configure", func() {
			cfg := Config{}
			err := Configure(&cfg)
			So(err, ShouldBeNil)

			Printf("%+v\n", cfg)

			p := cfg.P1
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")

			p = cfg.P2
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")

			p = cfg.P3
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")

			p = cfg.P4
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")
		})
	})
}

func TestBindConfig(t *testing.T) {
	type Person struct {
		Name   string
		Age    int
		OpenID int `mapstructure:"open_id"`
		EMail  string
	}

	type Person2 struct {
		BindConfig
		Person `mapstructure:",squash"`
	}

	type PersonMap map[string]interface{}

	type bundle struct {
		P1 Person         `confx:"name=test/person.json; read=file; binding"`
		P2 Person2        `confx:"name=test/person.yaml;format=yaml"`
		P3 *Person2       `confx:"name=test/person_path.json;path=AAA.BBB; read=file"`
		P4 Person         `confx:"name=test/person_path.yaml;format=yaml;path=AAA.BBB; binding"`
		P5 func() *Person `confx:"name=test/person_path.yaml;format=yaml;path=AAA.BBB; binding"`
		P6 func() Person2 `confx:"name=test/person.yaml;format=yaml"`

		P7 map[string]interface{} `confx:"name=test/person_path.yaml;format=yaml;path=AAA.BBB; binding"`
		P8 PersonMap              `confx:"name=test/person_path.yaml;format=yaml;path=AAA.BBB; binding"`
		P9 *PersonMap             `confx:"name=test/person_path.yaml;format=yaml;path=AAA.BBB; binding"`

		P10 func() map[string]interface{} `confx:"name=test/person.json; read=file; binding"`

		P11 []string `confx:"name=test/persons2.json;format=json;read=file;binding"`
	}

	type bundle2 struct {
		P5 func() *Person `confx:"name=Bob1;format=yaml;path=AAA.BBB; binding; read=Test"`
		P6 func() Person2 `confx:"name=Bob2;format=yaml; read=Test"`

		P7 func() map[string]interface{} `confx:"name=Bob1;format=yaml;path=AAA.BBB; binding; read=Test"`
		P8 func() *PersonMap             `confx:"name=Bob1;format=yaml;path=AAA.BBB; binding; read=Test"`
		P9 func() PersonMap              `confx:"name=Bob1;format=yaml;path=AAA.BBB; binding; read=Test"`
	}

	Convey("绑定配置", t, func() {
		container := bindx.NewContainer()
		container.AddBinder(binder)
		Convey("绑定函数", func() {
			cfg := bundle{}
			err := container.AddCombo(&cfg)
			So(err, ShouldBeNil)

			err = container.Bind(context.TODO())
			So(err, ShouldBeNil)

			Printf("%+v\n", cfg)

			p := cfg.P1
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")

			p2 := cfg.P2
			So(p2.Name, ShouldEqual, "Bob")
			So(p2.Age, ShouldEqual, 23)
			So(p2.OpenID, ShouldEqual, 99999)
			So(p2.EMail, ShouldEqual, "xxxx@qq.com")

			p3 := cfg.P3
			So(p3.Name, ShouldEqual, "Bob")
			So(p3.Age, ShouldEqual, 23)
			So(p3.OpenID, ShouldEqual, 99999)
			So(p3.EMail, ShouldEqual, "xxxx@qq.com")

			p = cfg.P4
			So(p.Name, ShouldEqual, "Bob")
			So(p.Age, ShouldEqual, 23)
			So(p.OpenID, ShouldEqual, 99999)
			So(p.EMail, ShouldEqual, "xxxx@qq.com")

			p5 := cfg.P5()
			So(p5.Name, ShouldEqual, "Bob")
			So(p5.Age, ShouldEqual, 23)
			So(p5.OpenID, ShouldEqual, 99999)
			So(p5.EMail, ShouldEqual, "xxxx@qq.com")

			p6 := cfg.P6()
			So(p6.Name, ShouldEqual, "Bob")
			So(p6.Age, ShouldEqual, 23)
			So(p6.OpenID, ShouldEqual, 99999)
			So(p6.EMail, ShouldEqual, "xxxx@qq.com")

			m := cfg.P7
			testx.PrintJSONPretty(m)
			So(m["name"], ShouldEqual, "Bob")
			So(m["age"], ShouldEqual, 23)
			So(m["open_id"], ShouldEqual, 99999)
			So(m["email"], ShouldEqual, "xxxx@qq.com")

			m = cfg.P8
			So(m["name"], ShouldEqual, "Bob")
			So(m["age"], ShouldEqual, 23)
			So(m["open_id"], ShouldEqual, 99999)
			So(m["email"], ShouldEqual, "xxxx@qq.com")

			m = *cfg.P9
			So(m["name"], ShouldEqual, "Bob")
			So(m["age"], ShouldEqual, 23)
			So(m["open_id"], ShouldEqual, 99999)
			So(m["email"], ShouldEqual, "xxxx@qq.com")

			pm := cfg.P10()
			fmt.Println("P10")
			testx.PrintJSONPretty(pm)

			strs := cfg.P11
			So(len(strs), ShouldEqual, 3)
			testx.PrintJSONPretty(strs)
		})

		Convey("配置更新", func() {
			watch, ch := NewMapChannelWatch()
			RegisterWatch("Test", watch)

			bytes1, err := ReadFile("test/person_path.yaml")
			So(err, ShouldBeNil)

			bytes2, err := ReadFile("test/person.yaml")
			So(err, ShouldBeNil)

			bytes12, err := ReadFile("test/person_path2.yaml")
			So(err, ShouldBeNil)

			bytes22, err := ReadFile("test/person2.yaml")
			So(err, ShouldBeNil)

			ch <- map[string][]byte{
				"Bob1": bytes1,
				"Bob2": bytes2,
			}

			cfg := bundle2{}
			err = container.AddCombo(&cfg)
			So(err, ShouldBeNil)

			err = container.Bind(context.TODO())
			So(err, ShouldBeNil)

			p5 := cfg.P5()
			So(p5.Name, ShouldEqual, "Bob")
			So(p5.Age, ShouldEqual, 23)
			So(p5.OpenID, ShouldEqual, 99999)
			So(p5.EMail, ShouldEqual, "xxxx@qq.com")

			p6 := cfg.P6()
			So(p6.Name, ShouldEqual, "Bob")
			So(p6.Age, ShouldEqual, 23)
			So(p6.OpenID, ShouldEqual, 99999)
			So(p6.EMail, ShouldEqual, "xxxx@qq.com")

			m := cfg.P7()
			So(m["name"], ShouldEqual, "Bob")
			So(m["age"], ShouldEqual, 23)
			So(m["open_id"], ShouldEqual, 99999)
			So(m["email"], ShouldEqual, "xxxx@qq.com")

			m = *cfg.P8()
			So(m["name"], ShouldEqual, "Bob")
			So(m["age"], ShouldEqual, 23)
			So(m["open_id"], ShouldEqual, 99999)
			So(m["email"], ShouldEqual, "xxxx@qq.com")

			m = cfg.P9()
			So(m["name"], ShouldEqual, "Bob")
			So(m["age"], ShouldEqual, 23)
			So(m["open_id"], ShouldEqual, 99999)
			So(m["email"], ShouldEqual, "xxxx@qq.com")

			ch <- map[string][]byte{
				"Bob1": bytes12,
				"Bob2": bytes2,
			}

			time.Sleep(time.Millisecond * 10)

			p5 = cfg.P5()
			So(p5.Name, ShouldEqual, "Tom")
			So(p5.Age, ShouldEqual, 25)
			So(p5.OpenID, ShouldEqual, 99998)
			So(p5.EMail, ShouldEqual, "xxxx1@qq.com")

			m = cfg.P7()
			So(m["name"], ShouldEqual, "Tom")
			So(m["age"], ShouldEqual, 25)
			So(m["open_id"], ShouldEqual, 99998)
			So(m["email"], ShouldEqual, "xxxx1@qq.com")

			m = *cfg.P8()
			So(m["name"], ShouldEqual, "Tom")
			So(m["age"], ShouldEqual, 25)
			So(m["open_id"], ShouldEqual, 99998)
			So(m["email"], ShouldEqual, "xxxx1@qq.com")

			m = cfg.P9()
			So(m["name"], ShouldEqual, "Tom")
			So(m["age"], ShouldEqual, 25)
			So(m["open_id"], ShouldEqual, 99998)
			So(m["email"], ShouldEqual, "xxxx1@qq.com")

			p6 = cfg.P6()
			So(p6.Name, ShouldEqual, "Bob")
			So(p6.Age, ShouldEqual, 23)
			So(p6.OpenID, ShouldEqual, 99999)
			So(p6.EMail, ShouldEqual, "xxxx@qq.com")

			ch <- map[string][]byte{
				"Bob1": bytes1,
				"Bob2": bytes22,
			}

			time.Sleep(time.Millisecond * 10)

			p5 = cfg.P5()
			So(p5.Name, ShouldEqual, "Bob")
			So(p5.Age, ShouldEqual, 23)
			So(p5.OpenID, ShouldEqual, 99999)
			So(p5.EMail, ShouldEqual, "xxxx@qq.com")

			m = cfg.P7()
			So(m["name"], ShouldEqual, "Bob")
			So(m["age"], ShouldEqual, 23)
			So(m["open_id"], ShouldEqual, 99999)
			So(m["email"], ShouldEqual, "xxxx@qq.com")

			m = *cfg.P8()
			So(m["name"], ShouldEqual, "Bob")
			So(m["age"], ShouldEqual, 23)
			So(m["open_id"], ShouldEqual, 99999)
			So(m["email"], ShouldEqual, "xxxx@qq.com")

			m = cfg.P9()
			So(m["name"], ShouldEqual, "Bob")
			So(m["age"], ShouldEqual, 23)
			So(m["open_id"], ShouldEqual, 99999)
			So(m["email"], ShouldEqual, "xxxx@qq.com")

			p6 = cfg.P6()
			So(p6.Name, ShouldEqual, "Tom")
			So(p6.Age, ShouldEqual, 25)
			So(p6.OpenID, ShouldEqual, 99998)
			So(p6.EMail, ShouldEqual, "xxxx1@qq.com")

			close(ch)
		})

	})
}

func TestCustom(t *testing.T) {
	AddTypeHook(func(tf reflect.Type, tt reflect.Type, data interface{}) (interface{}, error) {
		if tf == reflect.TypeOf(map[string]interface{}{}) && tt == reflect.TypeOf(new(HelloAware)).Elem() {
			res := Person{}
			err := mapstructure.Decode(data, &res)
			if err != nil {
				return nil, err
			}

			return &res, nil
		}
		return data, nil
	})

	Convey("接口读取", t, func() {
		Convey("单个读取", func() {
			hello := new(HelloAware)
			err := ReadFile.Unmarshal("test/person.json", hello, JSON)
			So(err, ShouldBeNil)

			testx.PrintJSONPretty(hello)
		})
		Convey("列表读取", func() {
			type bundle struct {
				Hellos []HelloAware `mapstructure:"persons"`
			}
			hellos := new(bundle)
			err := ReadFile.UnmarshalJSON("test/persons.json", hellos)
			So(err, ShouldBeNil)

			testx.PrintJSONPretty(hellos)
		})
	})

}
