package forge

import (
	"fmt"
	"time"

	. "gopkg.in/check.v1"
)

type BucketSuite struct{}

var _ = Suite(&BucketSuite{})

func (s *BucketSuite) TestCreateBucket(c *C) {

	client, err := New()
	c.Assert(err, IsNil)

	err = client.Authenticate([]string{"bucket:create"})
	c.Assert(err, IsNil)

	name := fmt.Sprintf("test-%d", time.Now().UnixNano())
	policy := "transient"
	bucket := Bucket{
		Name:   name,
		Policy: policy,
	}

	err = client.CreateBucket(&bucket)
	c.Assert(err, IsNil)

	c.Check(bucket.Name, Equals, name)
	c.Check(bucket.Policy, Equals, policy)
	c.Check(len(bucket.Owner) > 0, Equals, true)
	c.Check(bucket.CreatedDate > 0, Equals, true)
}
