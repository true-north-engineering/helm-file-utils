package reader

import (
	"bufio"
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"golang.org/x/term"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
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
		if i != 0 && i <= len(result) && name != "" {
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

	fs := memfs.New()
	repository, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:      pathToClone,
		Progress: os.Stdout,
		Auth:     authMethod,
	})

	if err != nil {
		log.Println(err)
	}
	headReference, err := repository.Head()
	if err != nil {
		log.Println(err)
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
			log.Println(err)
		}
	}

	//parse absolute path properly
	pathToGet := regexMap["PathToGet"]
	info, err := fs.Stat(pathToGet)

	if err != nil {
		return InputValue{}, err
	}

	// read flat
	if info.IsDir() {
		dir, err := fs.ReadDir(pathToGet)
		if err != nil {
			return InputValue{}, err
		}
		inputValue := InputValue{Kind: InputKindDir, Value: make(map[string][]byte)}
		for _, fileInfo := range dir {
			if fileInfo.IsDir() {
				continue
			}
			filePath := filepath.Join(pathToGet, fileInfo.Name())
			file, err := fs.Open(filePath)
			if err != nil {
				return InputValue{}, err
			}
			data, err := ioutil.ReadAll(file)
			if err != nil {
				return InputValue{}, err
			}
			inputValue.Value[filePath] = data
		}
		return inputValue, nil
	}

	// read file
	file, err := fs.Open(pathToGet)
	if err != nil {
		return InputValue{}, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return InputValue{}, err
	}
	inputValue := InputValue{Kind: InputKindFile, Value: make(map[string][]byte)}
	inputValue.Value[InputKindFile] = data
	return inputValue, nil
}

// getPassword Returns provided password if exists
// Order for fetching password is as it follows:
//		1. Check if credentials are provided in URI using the [username[:password]@] syntax
//		2. Look for environment variable named FUTL_GIT_PASSWORD
//		3. Look for environment variable named FUTL_CI, if exists prompt user to enter password
func getGitPassword() string {
	if regexMap["Password"] != "" {
		return regexMap["Password"]
	} else if env, ok := os.LookupEnv("FUTL_GIT_PASSWORD"); ok == true {
		return env
	}
	_, ok := os.LookupEnv("FUTL_CI")
	if !ok {
		return ""
	}
	fmt.Print("enter password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(bytePassword))
}

// getUsername Returns provided username if exists
// Order for fetching username is as it follows:
//		1. Check if credentials are provided in URI using the [username[:password]@] syntax
//		2. Look for environment variable named FUTL_GIT_USER
//		3. Look for environment variable named FUTL_CI, if exists prompt user to enter username
func getGitUsername() string {
	if regexMap["Username"] != "" {
		return regexMap["Username"]
	} else if env, ok := os.LookupEnv("FUTL_GIT_USER"); ok == true {
		return env
	}
	_, ok := os.LookupEnv("FUTL_CI")
	if !ok {
		return ""
	}
	fmt.Print("enter username: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(username)
}
