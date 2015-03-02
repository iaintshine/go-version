package version

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	GitDescriptionRegexp = "v(\\d+).(\\d+).(\\d+)-(.*)-([\\d]+)-g(.+)"
)

type Version struct {
	Id                   string
	Major                int
	Minor                int
	Patch                int
	PreRelease           string
	GitAdditionalCommits int
	GitShortSha          string
}

func ParseGitDescription(desc string) Version {
	r := regexp.MustCompile(GitDescriptionRegexp)
	matches := r.FindStringSubmatch(desc)

	if matches == nil {
		return Version{}
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])
	additional, _ := strconv.Atoi(matches[5])

	return Version{
		Id:                   desc,
		Major:                major,
		Minor:                minor,
		Patch:                patch,
		PreRelease:           matches[4],
		GitAdditionalCommits: additional,
		GitShortSha:          matches[6],
	}
}

func (v Version) String() string {
	return fmt.Sprintf("major=%v, minor=%v, patch=%v, prerelease=%v, additional=%v, sha=%v",
		v.Major, v.Minor, v.Patch, v.PreRelease, v.GitAdditionalCommits, v.GitShortSha)
}
