package gitcollector

// GitHandler consists of behavior to be executed by git handler's implementation
type GitHandler interface {
	FetchGitPRSummary() (summary map[string]int, err error)
	MailToAdmin(summary map[string]int) (err error)
}
