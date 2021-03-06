// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names_test

import (
	gc "launchpad.net/gocheck"

	"github.com/juju/names"
)

type environSuite struct{}

var _ = gc.Suite(&environSuite{})

var parseEnvironTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{{
	tag: "",
	err: names.InvalidTagError("", ""),
}, {
	tag:      "environment-dave",
	expected: names.NewEnvironTag("dave"),
}, {
	tag: "dave",
	err: names.InvalidTagError("dave", ""),
	//}, {
	// TODO(dfc) passes, but should not
	//	tag: "environment-",
	//	err: names.InvalidTagError("environment", ""),
}, {
	tag: "service-dave",
	err: names.InvalidTagError("service-dave", names.EnvironTagKind),
}}

func (s *environSuite) TestParseEnvironTag(c *gc.C) {
	for i, t := range parseEnvironTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseEnvironTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}
