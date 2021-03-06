package forge

import (
	"fmt"
	"os"
	"time"

	. "gopkg.in/check.v1"
)

type ObjectSuite struct{}

var _ = Suite(&ObjectSuite{})

func (s *ObjectSuite) TestCreateObject(c *C) {

	client, err := New()
	c.Assert(err, IsNil)

	err = client.Authenticate([]string{"bucket:create", "data:create", "data:write"})
	c.Assert(err, IsNil)

	name := fmt.Sprintf("test-%d", time.Now().UnixNano())
	policy := "transient"
	bucket := Bucket{
		Name:   name,
		Policy: policy,
	}

	err = client.CreateBucket(&bucket)
	c.Assert(err, IsNil)

	filepath := "test/test.txt"

	fi, err := os.Stat(filepath)
	c.Assert(err, IsNil)

	// get the size
	contentSize := fi.Size()
	contentType := "text/plain; charset=UTF-8"

	obj := Object{
		BucketName:  name,
		Name:        "test.txt",
		ContentType: contentType,
		ContentSize: int(contentSize),
	}

	file, err := os.Open(filepath)
	c.Assert(err, IsNil)
	err = client.ObjectUpload(&obj, file)
	c.Assert(err, IsNil)
	c.Check(len(obj.Location) > 0, Equals, true)
	c.Check(len(obj.Id) > 0, Equals, true)
}
