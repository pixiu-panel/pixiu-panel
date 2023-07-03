package types

import (
	"database/sql/driver"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"strings"
	"time"
)

// 默认时间格式
const dateFormat = "2006-01-02"

// Date 自定义时间类型
type Date time.Time

// Scan implements the Scanner interface.
func (dt *Date) Scan(value any) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", value.(time.Time).String())
	*dt = Date(tTime)
	return nil
}

// Value implements the driver Valuer interface.
func (dt Date) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if dt.String() == "0001-01-01" {
		return nil, nil
	}
	return []byte(dt.Format(dateFormat)), nil
}

// 用于 fmt.Println 和后续验证场景
func (dt Date) String() string {
	return dt.Format(dateFormat)
}

// Format 格式化
func (dt Date) Format(fm string) string {
	return time.Time(dt).Format(fm)
}

// After 时间比较
func (dt *Date) After(now time.Time) bool {
	return time.Time(*dt).After(now)
}

// Before 时间比较
func (dt *Date) Before(now time.Time) bool {
	return time.Time(*dt).Before(now)
}

// IBefore 时间比较
func (dt *Date) IBefore(now Date) bool {
	return dt.Before(time.Time(now))
}

// SubTime 对比
func (dt Date) SubTime(t time.Time) time.Duration {
	return dt.ToTime().Sub(t)
}

// Sub 对比
func (dt Date) Sub(t Date) time.Duration {
	return dt.ToTime().Sub(t.ToTime())
}

// ToTime 转换为golang的时间类型
func (dt Date) ToTime() time.Time {
	return time.Time(dt)
}

// IsNil 是否为空值
func (dt Date) IsNil() bool {
	return dt.Format(dateFormat) == "0001-01-01"
}

// Unix 实现Unix函数
func (dt Date) Unix() int64 {
	return dt.ToTime().Unix()
}

// ========  序列化 JSON ========

// MarshalJSON 时间到字符串
func (dt Date) MarshalJSON() ([]byte, error) {
	// 过滤掉空数据
	if dt.IsNil() {
		return []byte("\"\""), nil
	}
	output := fmt.Sprintf(`"%s"`, dt.Format(dateFormat))
	return []byte(output), nil
}

// UnmarshalJSON 字符串到时间
func (dt *Date) UnmarshalJSON(b []byte) error {
	if len(b) == 2 {
		*dt = Date{}
		return nil
	}
	// 解析指定的格式
	var now time.Time
	var err error
	if strings.HasPrefix(string(b), "\"") {
		now, err = time.ParseInLocation(`"`+dateFormat+`"`, string(b), time.Local)
	} else {
		now, err = time.ParseInLocation(dateFormat, string(b), time.Local)
	}
	*dt = Date(now)
	return err
}

// ========  序列化 RPCX ========

// EncodeMsgpack
// @description: 序列化(由于Msgpack不支持时区，所以这里序列化成字符串)
// @receiver dt
// @param enc
// @return error
func (dt Date) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.Encode(dt.Format("2006-01-02"))
}

// DecodeMsgpack
// @description: 反序列化(由于Msgpack不支持时区，所以这里从字符串反序列化)
// @receiver dt
// @param dec
// @return error
func (dt *Date) DecodeMsgpack(dec *msgpack.Decoder) error {
	var dtStr string
	err := dec.Decode(&dtStr)
	if err != nil {
		return err
	}

	var tm time.Time
	tm, err = time.ParseInLocation("2006-01-02", dtStr, time.Local)
	if err != nil {
		return err
	}
	*dt = Date(tm)
	return nil
}
