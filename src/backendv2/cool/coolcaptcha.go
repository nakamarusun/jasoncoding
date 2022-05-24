package cool

import (
	"io"
	"os/exec"
	"strings"
)

// Stores the question and the answer
type QuestionAnswer struct {
	Question string
	Answer   string
}

type Result struct {
	Close    func() error     // Very important to defer
	Reader   io.ReadCloser    // Read the file here
	Question []QuestionAnswer // Question answer pack
}

// Generates a new captcha file and contains the answer
func GenCaptcha() (Result, error) {
	command := "-background lightblue -fill blue -pointsize 72 label:Anthony -quality 100% jpg:-"
	commandSplit := strings.Split(command, " ")

	// Execute the command and get the stdoutpipe
	cmd := exec.Command("convert", commandSplit...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return Result{}, err
	}
	if err := cmd.Start(); err != nil {
		return Result{}, err
	}

	return Result{
		Close:  cmd.Wait,
		Reader: stdout,
	}, nil
}

func CheckCaptcha() bool {
	return true
}

// Initializes the coolcaptcha configs
func Init(options CoolOptions) {
	cfg = options
}

