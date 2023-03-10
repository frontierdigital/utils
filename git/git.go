package git

type Git interface {
	Add(path string) error
	Checkout(branchName string, create bool) error
	CloneOverHttp(url string, username string, password string) error
	Commit(message string) (string, error)
	GetRepositoryPath() string
	Push(force bool) error
}
