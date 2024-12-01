package auxiliary

import (
	"context"
)

func GetUserInfo(c context.Context) int64 {
	return c.Value("userId").(int64)
}

func SetUserInfo(c *context.Context, userId int64) {
	tc := context.WithValue(*c, "userId", userId)
	*c = tc
}
