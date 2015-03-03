package version

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	GitDescriptionRegexp = "v?(?P<major>\\d+)(.(?P<minor>\\d+)(.(?P<patch>\\d+))?)?(-(?P<pre>.*))?-(?P<add>[\\d]+)-g(?P<sha>.+)"
	GitNoTagsFound       = "fatal: No names found, cannot describe anything."
)

var (
	PreReleaseVersion = Version{Full: "0.0.0-prerelease", PreRelease: "prerelease"}
)

type Version struct {
	Full                 string
	Major                int
	Minor                int
	Patch                int
	PreRelease           string
	GitAdditionalCommits int
	GitShortSha          string
}

func ParseGitDescription(desc string) Version {
	if desc == GitNoTagsFound {
		return PreReleaseVersion
	}

	r := regexp.MustCompile(GitDescriptionRegexp)
	matches := r.FindStringSubmatch(desc)

	if matches == nil {
		return Version{}
	}

	lookup := make(map[string]string)
	groups := r.SubexpNames()

	for i, group := range groups {
		if group == "" {
			continue
		}
		lookup[group] = matches[i]
	}

	major, _ := strconv.Atoi(lookup["major"])
	minor, _ := strconv.Atoi(lookup["minor"])
	patch, _ := strconv.Atoi(lookup["patch"])
	additional, _ := strconv.Atoi(lookup["add"])

	return Version{
		Full:                 desc,
		Major:                major,
		Minor:                minor,
		Patch:                patch,
		PreRelease:           lookup["pre"],
		GitAdditionalCommits: additional,
		GitShortSha:          lookup["sha"],
	}
}

func (v Version) String() string {
	return v.Full
}

func (v Version) Desc() string {
	return fmt.Sprintf("major=%v, minor=%v, patch=%v, prerelease=%v, git-additional=%v, git-sha=%v",
		v.Major, v.Minor, v.Patch, v.PreRelease, v.GitAdditionalCommits, v.GitShortSha)
}
