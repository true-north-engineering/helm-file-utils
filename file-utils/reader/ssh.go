package reader

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	SshPrefix      = "ssh"
	DefaultPort    = "22"
	DefaultPubFile = "id_ed25519"
)

var splitScp []string
var splitUserAndPassword []string

var suitable_answers = []string{"bogus password", "real password"}
var pwIdx = 0

func ReadSsh(sshPath string) (InputValue, error) {

	//trim given input
	//e.g. ssh://user@localhost to user@localhost
	scp := strings.TrimPrefix(sshPath, SshPrefix+"://")

	//split given input
	//e.g. userPassword@localhostPort to [userPassword, localhostPort]
	splitScp = strings.Split(scp, "@")

	//error if input is incompatible
	if len(splitScp) > 2 || len(scp) < 1 {
		return InputValue{}, fmt.Errorf("error parsing ssh : check given values")
	}

	//split given userPassword if exists
	//e.g. userPassword to [user, password]
	if len(splitScp) == 2 {
		splitUserAndPassword = strings.Split(splitScp[0], ":")
	}

	//split given localhostPort
	//e.g. localhostPortFile to [localhost, portFile]
	splitHostAndPort := strings.Split(splitScp[len(splitScp)-1], ":")

	var hostname, port, path string
	var firstSlash int

	hostname = splitHostAndPort[0]

	firstSlash = strings.Index(splitHostAndPort[1], "/")

	port = func() string {
		if firstSlash == 0 {
			return DefaultPort
		} else {
			return splitHostAndPort[1][:firstSlash]
		}
	}()

	path = splitHostAndPort[1][firstSlash:]
	// var keyErr *knownhosts.KeyError

	config := &ssh.ClientConfig{
		User: getUsername(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(checkForPublicKeys()),
			ssh.PasswordCallback(getPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		/*ssh.HostKeyCallback(func(host string, remote net.Addr, pubKey ssh.PublicKey) error {
			kh := checkKnownHosts()
			hErr := kh(host, remote, pubKey)
			// Reference: https://blog.golang.org/go1.13-errors
			// To understand what errors.As is.
			if errors.As(hErr, &keyErr) && len(keyErr.Want) > 0 {
				// Reference: https://www.godoc.org/golang.org/x/crypto/ssh/knownhosts#KeyError
				// if keyErr.Want slice is empty then host is unknown, if keyErr.Want is not empty
				// and if host is known then there is key mismatch the connection is then rejected.
				log.Printf("WARNING: %v is not a key of %s, either a MiTM attack or %s has reconfigured the host pub key.", pubKey, host, host)
				return keyErr
			} else if errors.As(hErr, &keyErr) && len(keyErr.Want) == 0 {
				// host key not found in known_hosts then give a warning and continue to connect.
				log.Printf("WARNING: %s is not trusted, adding this key: %q to known_hosts file.", host, pubKey)
				return addHostKey(host, remote, pubKey)
			}
			log.Printf("Pub key exists for %s.", host)
			return nil
		}),*/
	}

	conn, err := ssh.Dial("tcp", hostname+":"+port, config)

	if err != nil {
		return InputValue{}, err
	}

	session, err := conn.NewSession()

	defer session.Close()

	if err != nil {
		return InputValue{}, err
	}

	result := InputValue{Kind: InputKindFile, Value: make(map[string][]byte)}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	result.Value[InputKindFile] = file

	return result, nil
}

func Challenge(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	for n, q := range questions {
		fmt.Printf("Got question: %s\n", q)
		answers[n] = suitable_answers[pwIdx]
	}
	pwIdx++

	return answers, nil
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

func promptUser(user string, instruction string, questions []string, echos []bool) ssh.KeyboardInteractiveChallenge {

	return nil
}

func getUsername() string {
	var userInput, user string
	if env, ok := os.LookupEnv("FUTL_SSH_USER"); ok == true {
		return env
	} else {
		if len(splitScp) == 1 {
			fmt.Print("enter username: ")
			fmt.Scanln(&userInput)
			return userInput
		}
		user = splitUserAndPassword[0]
	}
	return user
}

func getPassword() (string, error) {
	var userInput, password string

	if env, ok := os.LookupEnv("FUTL_SSH_PASSWORD"); ok == true {
		return env, nil
	} else {
		if len(splitUserAndPassword) != 2 {
			fmt.Print("enter password: ")
			fmt.Scanln(&userInput)
			return userInput, nil
		}
		password = splitUserAndPassword[1]
	}
	return password, nil
}

func addHostKey(host string, remote net.Addr, pubKey ssh.PublicKey) error {
	// add host key if host is not found in known_hosts, error object is return, if nil then connection proceeds,
	// if not nil then connection stops.
	khFilePath := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")

	f, fErr := os.OpenFile(khFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if fErr != nil {
		return fErr
	}
	defer f.Close()

	knownHosts := knownhosts.Normalize(remote.String())
	_, fileErr := f.WriteString(knownhosts.Line([]string{knownHosts}, pubKey))
	return fileErr
}

func createKnownHosts() {
	f, fErr := os.OpenFile(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"), os.O_CREATE, 0600)
	if fErr != nil {
		log.Fatal(fErr)
	}
	f.Close()
}

func checkKnownHosts() ssh.HostKeyCallback {
	createKnownHosts()
	kh, e := knownhosts.New(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	errCallBack(e)
	return kh
}

func errCallBack(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
