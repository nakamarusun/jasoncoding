package cool

import (
	"fmt"
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
	Choices   []string         // Choices to question
}

// Generates a new captcha file and contains the answer
func GenCaptcha(wrongNum int, answerNum int) (Result, error) {

	command := []string{
		"-size",
		strconv.Itoa(cfg.W) + "x" + strconv.Itoa(cfg.H),
		"xc:white",
		"-quality",
		cfg.Quality,
		"-gravity",
		"West"}

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

	usefont := len(cfg.FontList) >= 0

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

				rot := rand.Intn(80) - 40
				x := rand.Intn(cfg.W * 6 / 10)
				y := rand.Intn(cfg.H/4) - cfg.H/8

				if usefont {
					command = append(command,
						"-font",
						cfg.FontList[rand.Intn(len(cfg.FontList))],
					)
				}

				// Generate command
				command = append(command,
					"-fill",
					color,
					"-pointsize",
					strconv.Itoa(rand.Intn(50)+40),
					"-annotate",
					fmt.Sprintf("%dx%d+%d+%d", rot, rot, x, y),
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
		Choices:   cfg.Colors,
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
