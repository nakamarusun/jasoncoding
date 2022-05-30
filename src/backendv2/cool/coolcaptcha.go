package cool

import (
	"fmt"
	"io"
	"math/rand"
	"mime"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
	"jasoncoding.com/backendv2/utils"
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

type Coords struct {
	X int
	Y int
}

var noiser = "( -size %dx%d xc:black -seed %d -attenuate 0.35 +noise random -channel green -separate +channel -virtual-pixel background -blur 0x1 -auto-level -negate -wave 5x40 ) -compose Multiply -composite"

// Generates a new captcha file and contains the answer
func GenCaptcha(wrongNum int, answerNum int) (Result, error) {

	command := []string{
		"-size",
		strconv.Itoa(cfg.W) + "x" + strconv.Itoa(cfg.H),
		"xc:white",
		"-quality",
		cfg.Quality,
		"-gravity",
		"West",
	}

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

	// Draw random lines
	lineCount := rand.Intn(10) + 5
	for i := 0; i < lineCount; i++ {
		a1 := rand.Intn(cfg.W + cfg.H)
		a2 := rand.Intn(cfg.W + cfg.H)
		l1 := strconv.Itoa(utils.Tern(a1 < cfg.W, a1, cfg.W)) + "," + strconv.Itoa(utils.Tern(a1 > cfg.W, a1-cfg.W, 0))
		l2 := strconv.Itoa(utils.Tern(a2 < cfg.H, 0, a2-cfg.H)) + "," + strconv.Itoa(utils.Tern(a2 < cfg.H, a2, cfg.H))
		command = append(command,
			"-strokewidth",
			strconv.Itoa(rand.Intn(4)+2),
			"-stroke",
			cfg.Colors[rand.Intn(len(cfg.Colors))],
			"-draw",
			fmt.Sprintf("line %s %s", l1, l2),
		)
	}

	command = append(command,
		"-stroke",
		"None",
	)

	words := make([]string, 0, wrongNum+answerNum)
	coords := make([]Coords, 0, wrongNum+answerNum)
	j := 0
	for i := 0; i < wrongNum+answerNum; i++ {
		for {
			if word := cfg.WordList[rand.Intn(len(cfg.WordList))]; !slices.Contains(words, word) {
				// Get words and insert the selected index to the challenges list
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

				// Makes sure that words don't overlap
				var x int
				var y int
				for {
					x = rand.Intn(cfg.W * 6 / 10)
					y = rand.Intn(cfg.H/4) - cfg.H/8
					found := false
					for _, coord := range coords {
						if utils.PowInts(x-coord.X, 2)+utils.PowInts(y-coord.Y, 2) < utils.PowInts(cfg.ColRange, 2) {
							found = true
							break
						}
					}
					if !found {
						break
					}
				}
				coords = append(coords, Coords{x, y})

				// Whether we are using any fonts or nah
				if usefont {
					command = append(command,
						"-font",
						cfg.FontList[rand.Intn(len(cfg.FontList))],
					)
				}

				// Rotation
				rot := rand.Intn(80) - 40

				// Generate word command
				command = append(command,
					"-fill",
					color,
					"-pointsize",
					strconv.Itoa(rand.Intn(35)+55),
					"-annotate",
					fmt.Sprintf("%dx%d+%d+%d", rot, rot, x, y),
					word,
				)
				break
			}
		}
	}

	// Execute the command and get the stdoutpipe
	command = append(command,
		"-wave",
		strconv.Itoa(rand.Intn(5)+3)+"x100", // Waves
	)
	// Add the noisy overlay
	command = append(command, strings.Split(fmt.Sprintf(noiser, cfg.W, cfg.H+50, rand.Int()), " ")...)
	// Print resulting image to stdout
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

// Initializes the coolcaptcha configs
func Init(options CoolOptions) {
	cfg = options
	rand.Seed(time.Now().Unix())
}
