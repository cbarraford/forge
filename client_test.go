package forge

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestPackage(t *testing.T) { TestingT(t) }

type ClientSuite struct{}

var _ = Suite(&ClientSuite{})

func (s *ClientSuite) TestNew(c *C) {
	_, err := New()
	c.Assert(err, IsNil)

	client, err := NewWithCreds("foo", "bar")
	c.Assert(err, IsNil)
	c.Check(client.clientId, Equals, "foo")
	c.Check(client.clientSecret, Equals, "bar")
}

func (s *ClientSuite) TestPath(c *C) {
	client := Client{
		baseURL: "foo",
	}

	c.Check(client.Path("/bar"), Equals, "foo/bar")
}

func (s *ClientSuite) TestAuthenticate(c *C) {
	client, err := New()
	c.Assert(err, IsNil)

	err = client.Authenticate([]string{"data:read"})
	c.Assert(err, IsNil)

	c.Check(client.jwt.IsExpired(), Equals, false)
	c.Check(client.jwt.ExpiresIn > 0, Equals, true)
}
