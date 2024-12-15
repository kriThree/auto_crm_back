package auxiliary

import (
	"context"
)

func GetUserInfo(c context.Context) (int64, string) {
	if c.Value("userId") == nil || c.Value("role") == nil {
		return 0, ""
	}
	return c.Value("userId").(int64) , c.Value("role").(string)
}

func SetUserInfo(c *context.Context, userId int64, role string) {
	tc := context.WithValue(*c, "userId", userId)
	tc = context.WithValue(tc, "role", role)
	*c = tc
}
