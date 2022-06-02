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

var noiser = "( -size %dx%d xc:black -seed %d -attenuate 0.3 +noise random -channel green -separate +channel -virtual-pixel background -blur 0x1 -auto-level -negate -wave 5x40 ) -compose Multiply -composite"

// Line padder
var pad = -16
var textMaxIterate = 15

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
		l1 := strconv.Itoa(utils.Tern(a1 < cfg.W, a1, cfg.W+pad)) + "," + strconv.Itoa(utils.Tern(a1 > cfg.W, a1-cfg.W, -pad))
		l2 := strconv.Itoa(utils.Tern(a2 < cfg.H, -pad, a2-cfg.H)) + "," + strconv.Itoa(utils.Tern(a2 < cfg.H, a2, cfg.H+pad))
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
	colorLeft := make([]string, 0, wrongNum+answerNum)
	copy(colorLeft, cfg.Colors)

	// Index to keep track of font answers
	curAnswerI := 0
	for i := 0; i < wrongNum+answerNum; i++ {
		var word string
		for {
			if word = cfg.WordList[rand.Intn(len(cfg.WordList))]; !slices.Contains(words, word) {
				// Get words and insert the selected index to the challenges list
				words = append(words, word)
				break
			}
		}
		
		// Get distinct colors
		var color string
		if len(colorLeft) > 0 {
			colI := rand.Intn(len(colorLeft))
			color = colorLeft[colI]
			colorLeft = utils.RemoveIndex(colorLeft, colI)
		} else {
			color = cfg.Colors[rand.Intn(len(cfg.Colors))]
		}

		if ansIdx[curAnswerI] == i {
			// This is the answer
			challenges = append(challenges, QuestionAnswer{
				Question: word,
				Answer:   color,
			})
			if curAnswerI < len(ansIdx)-1 {
				curAnswerI++
			}
		}

		// Makes sure that words don't overlap
		var x int
		var yFromMid int // Because we are using gravity west, this starts from mid
		for i := 0; i < textMaxIterate; i += 1 {
			x = rand.Intn(cfg.W * 6 / 10)
			yFromMid = rand.Intn(cfg.H*3/5) - cfg.H*3/10
			found := false
			for _, coord := range coords {
				if utils.PowInts(x-coord.X, 2)+utils.PowInts(yFromMid-coord.Y, 2) < utils.PowInts(cfg.ColRange, 2) {
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
		coords = append(coords, Coords{x, yFromMid})

		// Whether we are using any fonts or nah
		if usefont {
			command = append(command,
				"-font",
				cfg.FontList[rand.Intn(len(cfg.FontList))],
			)
		}

		// Rotation
		rot := rand.Intn(30) - 15

		// Generate word command
		command = append(command,
			"-fill",
			color,
			"-pointsize",
			strconv.Itoa(rand.Intn(35)+55),
			"-annotate",
			fmt.Sprintf("%dx%d+%d+%d", rot, rot, x, yFromMid),
			word,
		)
	}

	// Add waves
	command = append(command,
		"-wave",
		strconv.Itoa(rand.Intn(3)+3)+"x100", // Waves
	)
	// Add the noisy overlay
	command = append(command, strings.Split(fmt.Sprintf(noiser, cfg.W, cfg.H+50, rand.Int()), " ")...)
	// Crop
	command = append(command,
		"-gravity",
		"Center",
		"-crop",
		fmt.Sprintf("%dx%d+0+0", cfg.W, cfg.H),
		"+repage",
	)
	// Print resulting image to stdout
	command = append(command, cfg.ImgFormat+":-")

	// Execute the command and get the stdoutpipe
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
