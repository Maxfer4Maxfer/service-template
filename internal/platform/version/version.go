package version

// Version holds everything about a current build.
type Version struct {
	buildTime    string
	buildBranch  string
	buildCommit  string
	buildSummary string
}

// New returns the new instance of a Version struct.
func New(buildTime, buildBranch, buildCommit, buildSummary string) *Version {
	return &Version{
		buildTime:    buildTime,
		buildBranch:  buildBranch,
		buildCommit:  buildCommit,
		buildSummary: buildSummary,
	}
}

// BuildTime is n output of "date +"%Y.%m.%d-%T.%Z"".
func (ver *Version) BuildTime() string {
	return ver.buildTime
}

// BuildBranch is an output of "git symbolic-ref  --short HEAD".
func (ver *Version) BuildBranch() string {
	return ver.buildBranch
}

// BuildCommit is an output of "git rev-parse HEAD".
func (ver *Version) BuildCommit() string {
	return ver.buildCommit
}

// BuildSummary is an output of "git describe --tags --dirty --always".
func (ver *Version) BuildSummary() string {
	return ver.buildSummary
}
