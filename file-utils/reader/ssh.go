package reader

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

const (
	SshPrefix      = "ssh"
	DefaultPort    = "22"
	DefaultPubFile = "id_ed25519"
)

var regexMapSsh = make(map[string]string)

// ReadSsh Reader protocol that allows user to read content via ssh.
func ReadSsh(sshPath string) (InputValue, error) {

	regex := regexp.MustCompile("(((?P<Username>[^:]+)(:(?P<Password>[^:]+))?)@)?(?P<Hostname>[[:ascii:]]*):(?P<Port>\\d*)\\/((?P<Path>[[:ascii:]]*))")

	regexResult := regex.FindStringSubmatch(strings.TrimPrefix(sshPath, GitHttpsPrefix+"://"))

	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			regexMapSsh[name] = regexResult[i]
		}
	}

	port := func() string {
		if regexMapSsh["Port"] == "" {
			return DefaultPort
		} else {
			return regexMapSsh["Port"]
		}
	}()

	config := &ssh.ClientConfig{
		User: getSshUsername(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(checkForPublicKeys()),
			ssh.PasswordCallback(getSshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", regexMapSsh["Hostname"]+":"+port, config)

	if err != nil {
		return InputValue{}, err
	}

	session, err := conn.NewSession()

	defer session.Close()

	if err != nil {
		return InputValue{}, err
	}

	result := InputValue{Kind: InputKindFile, Value: make(map[string][]byte)}
	file, err := ioutil.ReadFile(regexMapSsh["Path"])
	if err != nil {
		return InputValue{}, err
	}
	result.Value[InputKindFile] = file

	return result, nil
}

func checkForPublicKeys() ssh.Signer {
	//get default HOME environment variable
	//e.g. default is /home/user
	homeEnv, _ := os.LookupEnv("HOME")

	//get first pub file located in /$HOME/.ssh/
	//e.g. default is /home/user
	pubFile := getPubFile(homeEnv)

	key, err := ioutil.ReadFile(homeEnv + "/.ssh/" + pubFile)

	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)

	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	return signer
}

// getPubFile Reads
func getPubFile(homeEnv string) string {
	cmd, _ := exec.Command("ls", homeEnv+"/.ssh/", "-A1").Output()
	lines := strings.Split(string(cmd), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "id_") {
			return line
		}
	}
	return DefaultPubFile
}

// getSshUsername Returns provided username if exists
// Order for fetching username is as it follows:
//		1. Check if credentials are provided in URI using the [username[:password]@] syntax
//		2. Look for environment variable named FUTL_SSH_USER
//		3. Look for environment variable named FUTL_CI, if exists prompt user to enter username
func getSshUsername() string {
	if regexMap["Username"] != "" {
		return regexMap["Username"]
	} else if env, ok := os.LookupEnv("FUTL_SSH_USER"); ok == true {
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

// getSshPassword Returns provided password if exists
// Order for fetching password is as it follows:
//		1. Check if credentials are provided in URI using the [username[:password]@] syntax
//		2. Look for environment variable named FUTL_SSH_PASSWORD
//		3. Look for environment variable named FUTL_CI, if exists prompt user to enter password
func getSshPassword() (string, error) {

	if regexMapSsh["Password"] != "" {
		return regexMapSsh["Password"], nil
	} else if env, ok := os.LookupEnv("FUTL_SSH_PASSWORD"); ok == true {
		return env, nil
	}
	_, ok := os.LookupEnv("FUTL_CI")
	if !ok {
		return "", nil
	}
	fmt.Print("enter password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", nil
	}
	return strings.TrimSpace(string(bytePassword)), nil
}
