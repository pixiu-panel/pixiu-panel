package types

import (
	"database/sql/driver"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"strings"
	"time"
)

// 默认时间格式
const dateTimeFormat = "2006-01-02 15:04:05.000"

// 可能包含的时间格式
var formatMap = map[string]string{
	"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
	"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
	"yyyy-mm-dd hh":       "2006-01-02 15:04",
	"yyyy-mm-dd":          "2006-01-02",
	"yyyy-mm":             "2006-01",
	"mm-dd":               "01-02",
	"dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",
	"yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
	"yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
	"yyyy/mm/dd hh":       "2006/01/02 15",
	"yyyy/mm/dd":          "2006/01/02",
	"yyyy/mm":             "2006/01",
	"mm/dd":               "01/02",
	"dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",
	"yyyy":                "2006",
	"mm":                  "01",
	"hh:mm:ss":            "15:04:05",
	"mm:ss":               "04:05",
}

// DateTime 自定义时间类型
type DateTime time.Time

// Scan implements the Scanner interface.
func (dt *DateTime) Scan(value interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", value.(time.Time).String())
	*dt = DateTime(tTime)
	return nil
}

// Value implements the driver Valuer interface.
func (dt DateTime) Value() (dv driver.Value, err error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if dt.String() == "0001-01-01 00:00:00.000" {
		return nil, nil
	}
	dv, err = []byte(dt.Format(dateTimeFormat)), nil
	return
}

// 用于 fmt.Println 和后续验证场景
func (dt DateTime) String() string {
	return dt.Format(dateTimeFormat)
}

// Format 格式化
func (dt *DateTime) Format(fm string) string {
	return time.Time(*dt).Format(fm)
}

// AutoParse 假装是个自动解析时间的函数
func (dt DateTime) AutoParse(timeStr string) (t time.Time, err error) {
	// 循环匹配预设的时间格式
	for _, v := range formatMap {
		// 尝试解析，没报错就是解析成功了
		t, err = time.ParseInLocation(v, timeStr, time.Local)
		if err == nil {
			// 错误为空，表示匹配上了
			return
		}
	}
	return
}

// After 时间比较
func (dt *DateTime) After(now time.Time) bool {
	return time.Time(*dt).After(now)
}

// Before 时间比较
func (dt *DateTime) Before(now time.Time) bool {
	return time.Time(*dt).Before(now)
}

// IBefore 时间比较
func (dt *DateTime) IBefore(now DateTime) bool {
	return dt.Before(time.Time(now))
}

// SubTime 对比
func (dt DateTime) SubTime(t time.Time) time.Duration {
	return dt.ToTime().Sub(t)
}

// Sub 对比
func (dt DateTime) Sub(t DateTime) time.Duration {
	return dt.ToTime().Sub(t.ToTime())
}

// ToTime 转换为golang的时间类型
func (dt DateTime) ToTime() time.Time {
	return time.Time(dt)
}

// IsNil 是否为空值
func (dt DateTime) IsNil() bool {
	return dt.Format(dateTimeFormat) == "0001-01-01 00:00:00.000"
}

// Unix 实现Unix函数
func (dt DateTime) Unix() int64 {
	return dt.ToTime().Unix()
}

// EndOfCentury 获取本世纪最后时间
func (dt DateTime) EndOfCentury() DateTime {
	yearEnd := time.Now().Local().Year()/100*100 + 99
	return DateTime(time.Date(yearEnd, 12, 31, 23, 59, 59, 999999999, time.Local))
}

// ========  序列化 JSON ========

// MarshalJSON 时间到字符串
func (dt DateTime) MarshalJSON() ([]byte, error) {
	// 过滤掉空数据
	if dt.IsNil() {
		return []byte("\"\""), nil
	}
	output := fmt.Sprintf(`"%s"`, dt.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// UnmarshalJSON 字符串到时间
func (dt *DateTime) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 2 {
		*dt = DateTime{}
		return
	}
	// 解析指定的格式
	var now time.Time
	if strings.HasPrefix(string(b), "\"") {
		now, err = dt.AutoParse(string(b)[1 : len(b)-1])
	} else {
		now, err = dt.AutoParse(string(b))
	}
	if err != nil {
		return
	}
	*dt = DateTime(now)
	return
}

// ========  序列化 RPCX ========

// EncodeMsgpack
// @description: 序列化(由于Msgpack不支持时区，所以这里序列化成字符串)
// @receiver dt
// @param enc
// @return error
func (dt DateTime) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.Encode(dt.Format("2006-01-02 15:04:05"))
}

// DecodeMsgpack
// @description: 反序列化(由于Msgpack不支持时区，所以这里从字符串反序列化)
// @receiver dt
// @param dec
// @return error
func (dt *DateTime) DecodeMsgpack(dec *msgpack.Decoder) error {
	var dtStr string
	err := dec.Decode(&dtStr)
	if err != nil {
		return err
	}

	var tm time.Time
	tm, err = dt.AutoParse(dtStr)
	if err != nil {
		return err
	}
	*dt = DateTime(tm)
	return nil
}
