	"io/ioutil"
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
		return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content))
	diffRecord := diffMatchPatch.DiffMain(highlight.Code(diffSection.FileName, diff1[1:]), highlight.Code(diffSection.FileName, diff2[1:]), true)
func (diff *Diff) LoadComments(issue *models.Issue, currentUser *models.User) error {
func ParsePatch(maxLines, maxLineCharacters, maxFiles int, reader io.Reader) (*Diff, error) {
			return diff, fmt.Errorf("Invalid first file line: %s", line)
		// TODO: Handle skipping first n files
		if len(diff.Files) >= maxFiles {
			_, err := io.Copy(ioutil.Discard, reader)
				return diff, fmt.Errorf("Copy: %v", err)
						return diff, fmt.Errorf("Unable to ReadLine: %v", err)
				err = fmt.Errorf("Unable to ReadLine: %v", err)
			err = fmt.Errorf("Unable to ReadLine: %v", err)
			if curFileLinesCount >= maxLines {
					err = fmt.Errorf("Unable to ReadLine: %v", err)
			curSection = &DiffSection{}
			if curFileLinesCount >= maxLines {
				err = fmt.Errorf("Unexpected line in hunk: %s", string(lineBytes))
			if curFileLinesCount >= maxLines {
				curSection = &DiffSection{}
			if curFileLinesCount >= maxLines {
				curSection = &DiffSection{}
			if curFileLinesCount >= maxLines {
				curSection = &DiffSection{}
			err = fmt.Errorf("Unexpected line in hunk: %s", string(lineBytes))
					err = fmt.Errorf("Unable to ReadLine: %v", err)
				count, err := models.Count(m)
// GetDiffRangeWithWhitespaceBehavior builds a Diff between two commits of a repository.
func GetDiffRangeWithWhitespaceBehavior(gitRepo *git.Repository, beforeCommitID, afterCommitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	commit, err := gitRepo.GetCommit(afterCommitID)
	ctx, cancel := context.WithTimeout(git.DefaultContext, time.Duration(setting.Git.Timeout.Default)*time.Second)
	defer cancel()
	var cmd *exec.Cmd
	if (len(beforeCommitID) == 0 || beforeCommitID == git.EmptySHA) && commit.ParentCount() == 0 {
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		if len(whitespaceBehavior) != 0 {
			diffArgs = append(diffArgs, whitespaceBehavior)
		diffArgs = append(diffArgs, afterCommitID)
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
		actualBeforeCommitID := beforeCommitID
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		if len(whitespaceBehavior) != 0 {
			diffArgs = append(diffArgs, whitespaceBehavior)
		diffArgs = append(diffArgs, afterCommitID)
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
		beforeCommitID = actualBeforeCommitID
		return nil, fmt.Errorf("StdoutPipe: %v", err)
		return nil, fmt.Errorf("Start: %v", err)
	pid := process.GetManager().Add(fmt.Sprintf("GetDiffRange [repo_path: %s]", repoPath), cancel)
	defer process.GetManager().Remove(pid)
	diff, err := ParsePatch(maxLines, maxLineCharacters, maxFiles, stdout)
	if err != nil {
		return nil, fmt.Errorf("ParsePatch: %v", err)
		tailSection := diffFile.GetTailSection(gitRepo, beforeCommitID, afterCommitID)
		return nil, fmt.Errorf("Wait: %v", err)
	shortstatArgs := []string{beforeCommitID + "..." + afterCommitID}
	if len(beforeCommitID) == 0 || beforeCommitID == git.EmptySHA {
		shortstatArgs = []string{git.EmptyTreeSHA, afterCommitID}
		shortstatArgs = []string{beforeCommitID, afterCommitID}
// GetDiffCommitWithWhitespaceBehavior builds a Diff representing the given commitID.
// The whitespaceBehavior is either an empty string or a git flag
func GetDiffCommitWithWhitespaceBehavior(gitRepo *git.Repository, commitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(gitRepo, "", commitID, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior)
}

		setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(c.Patch))