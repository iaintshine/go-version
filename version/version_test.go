package version

import (
	"testing"

	. "github.com/mailgun/vulcand/Godeps/_workspace/src/gopkg.in/check.v1"
)

func TestVersion(t *testing.T) { TestingT(t) }

type VersionSuite struct {
}

var _ = Suite(&VersionSuite{})

func (s *VersionSuite) TestVersionParsing(c *C) {
	cases := []struct {
		d string
		v Version
	}{
		{"fatal: No names found, cannot describe anything.", Version{"0.0.0-prerelease", 0, 0, 0, "prerelease", 0, ""}},
		{"v1.0.0-beta.2-10-g3bc96e1", Version{"v1.0.0-beta.2-10-g3bc96e1", 1, 0, 0, "beta.2", 10, "3bc96e1"}},
	}

	for _, test := range cases {
		c.Assert(ParseGitDescription(test.d), DeepEquals, test.v)
	}
}
