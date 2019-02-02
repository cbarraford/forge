package forge

import (
	"time"

	. "gopkg.in/check.v1"
)

type JWTSuite struct{}

var _ = Suite(&JWTSuite{})

func (s *JWTSuite) TestSetExpiration(c *C) {
	jwt := JWT{
		ExpiresIn: 500,
	}

	jwt.SetExpiration()
	c.Check(jwt.ExpiresAt.UnixNano() > time.Now().UnixNano(), Equals, true)
}

func (s *JWTSuite) TestIsExpired(c *C) {
	jwt := JWT{
		ExpiresIn: 500,
	}

	jwt.SetExpiration()
	c.Check(jwt.IsExpired(), Equals, false)

	jwt = JWT{
		ExpiresIn: -500,
	}

	jwt.SetExpiration()
	c.Check(jwt.IsExpired(), Equals, true)

}
