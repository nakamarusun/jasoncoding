package cool

import (
	"io"
	"math/rand"
	"mime"
	"os/exec"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
)

// Stores the question and the answer
type QuestionAnswer struct {
	Question string
	Answer   string
}

type Result struct {
	Close     func() error     // Very important to defer
	Reader    io.ReadCloser    // Read the file here
	Challenge []QuestionAnswer // Question answer pack
	Format    string           // Image format
}

// Generates a new captcha file and contains the answer
func GenCaptcha(wrongNum int, answerNum int) (Result, error) {

	command := []string{
		"-size",
		strconv.Itoa(cfg.W) + "x" + strconv.Itoa(cfg.H),
		"xc:white",
		"-quality",
		cfg.Quality}

	challenges := make([]QuestionAnswer, 0, answerNum)

	// Get the right answer indices
	ansIdx := make([]int, 0, answerNum)
	for i := 0; i < answerNum; i++ {
		for {
			if idx := rand.Intn(answerNum + wrongNum); !slices.Contains(ansIdx, idx) {
				ansIdx = append(ansIdx, idx)
				break
			}
		}
	}
	slices.Sort(ansIdx)

	// Generates the command and the answers
	words := make([]string, 0, wrongNum+answerNum)
	j := 0
	for i := 0; i < wrongNum+answerNum; i++ {
		for {
			if word := cfg.WordList[rand.Intn(len(cfg.WordList))]; !slices.Contains(words, word) {
				words = append(words, word)
				color := cfg.Colors[rand.Intn(len(cfg.Colors))]
				if ansIdx[j] == i {
					// This is the answer
					challenges = append(challenges, QuestionAnswer{
						Question: word,
						Answer:   color,
					})
					if j < len(ansIdx)-1 {
						j++
					}
				}
				// Generate command
				command = append(command,
					"-fill",
					color,
					"-pointsize",
					strconv.Itoa(rand.Intn(50)+50),
					"-annotate",
					"0x0+0+60",
					word,
				)
				break
			}
		}
	}

	// Execute the command and get the stdoutpipe
	command = append(command, cfg.ImgFormat+":-")
	cmd := exec.Command("convert", command...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return Result{}, err
	}
	if err := cmd.Start(); err != nil {
		return Result{}, err
	}

	return Result{
		Close:     cmd.Wait,
		Reader:    stdout,
		Challenge: challenges,
		Format:    mime.TypeByExtension("." + cfg.ImgFormat),
	}, nil
}

func CheckCaptcha() bool {
	return true
}

// Initializes the coolcaptcha configs
func Init(options CoolOptions) {
	cfg = options
	rand.Seed(time.Now().Unix())
}
