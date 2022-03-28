package reader

import (
	"bufio"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"golang.org/x/term"
	"net/url"
	"os"
	"regexp"
	"strings"
	"syscall"
)

const (
	GitHttpsPrefix = "git_https"
)

var regexMap = make(map[string]string)

func ReadGitHttps(gitPath string) (InputValue, error) {

	regex := regexp.MustCompile("(((?P<Username>[^:]+)(:(?P<Password>[^:]+))?)@)?((?P<PathToClone>[[:ascii:]]*.[[:ascii:]]*)[[:blank:]](?P<PathToGet>[[:ascii:]]*))")

	result := regex.FindStringSubmatch(strings.TrimPrefix(gitPath, GitHttpsPrefix+"://"))

	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			regexMap[name] = result[i]
		}
	}

	if strings.Contains(regexMap["PathToGet"], "#") {
		regexMap["Branch"] = regexMap["PathToGet"][strings.LastIndex(regexMap["PathToGet"], "#")+1:]
		regexMap["PathToGet"] = regexMap["PathToGet"][:strings.LastIndex(regexMap["PathToGet"], "#")]

	} else {
		regexMap["Branch"] = ""
	}

	//path where files are temporarily store
	//e.g. /tmp/helm-file-utils
	pathToLocalTmpDir := "/tmp/" + regexMap["PathToClone"][strings.Index(regexMap["PathToClone"], "/")+1:strings.LastIndex(regexMap["PathToClone"], "/")+1]
	defer os.RemoveAll(pathToLocalTmpDir)

	//path that is cloned via https
	//e.g. https://github.com/true-north-engineering/helm-file-utils
	pathToClone, _ := url.QueryUnescape("https://" + regexMap["PathToClone"])

	authMethod := &http.BasicAuth{
		Username: getGitUsername(),
		Password: getGitPassword(),
	}

	repository, err := git.PlainClone(pathToLocalTmpDir, false, &git.CloneOptions{
		URL:      pathToClone,
		Progress: os.Stdout,
		Auth:     authMethod,
	})

	if err != nil {
		fmt.Println(err)
	}
	headReference, err := repository.Head()
	if err != nil {
		fmt.Println(err)
	}
	_ = strings.TrimPrefix(string(headReference.Name()), "refs/heads/")

	if regexMap["Branch"] != "" {
		w, _ := repository.Worktree()

		err := repository.Fetch(&git.FetchOptions{
			RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
			Auth:     authMethod,
		})

		err = w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", regexMap["Branch"])),
			Force:  true,
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	//parse absolute path properly
	absolutePath := pathToLocalTmpDir + regexMap["PathToGet"]

	fileInfo, err := os.Stat(absolutePath)

	if err != nil {
		return InputValue{}, err
	}

	//once task is finished, delete temporarily created directory

	//if given path is directory, delegate it to reader dir
	if fileInfo.IsDir() {
		return ReadDir(absolutePath)
	}

	return ReadFile(absolutePath)

}

//getPassword Returns provided password if exists, else prompts user for password
func getGitPassword() string {
	if regexMap["Password"] != "" {
		return regexMap["Password"]
	} else if env, ok := os.LookupEnv("FUTL_GIT_PASSWORD"); ok == true {
		return env
	}
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(bytePassword))
}

func getGitUsername() string {
	if regexMap["Username"] != "" {
		return regexMap["Username"]
	} else if env, ok := os.LookupEnv("FUTL_GIT_USER"); ok == true {
		return env
	}
	fmt.Print("enter username: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(username)
}
