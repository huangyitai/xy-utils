package confx

import "time"

// TimeDuration 用于配置中获取Duration的场景，直接从string格式的duration转换为Duration
type TimeDuration time.Duration

// MarshalJSON ...
func (t TimeDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(t))
}

// UnmarshalJSON ...
func (t *TimeDuration) UnmarshalJSON(bytes []byte) error {
	str := ""
	err := json.Unmarshal(bytes, &str)
	if err != nil {
		return err
	}
	dur, err := time.ParseDuration(str)
	if err != nil {
		return err
	}
	*t = TimeDuration(dur)
	return nil
}

// Dur 获取标准库time.Duration
func (t TimeDuration) Dur() time.Duration {
	return time.Duration(t)
}
