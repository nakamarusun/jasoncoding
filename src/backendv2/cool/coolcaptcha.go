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

// Add noise overlay
var noiser = "( -size %dx%d xc:black -seed %d -attenuate 0.3 +noise random -channel green -separate +channel -virtual-pixel background -blur 0x1 -auto-level -negate -wave 5x40 ) -compose Multiply -composite"

// Line padder
var pad = -16

// Maximum iterations for text trying to find a place that does not collide with each other
var textMaxIterate = 5

// Ranges (min, max)
var linesRange = [2]int{5, 15}
var strokeRange = [2]int{2, 6}
var xRange [2]int
var yHalfRange [2]int
var rotateRange = [2]int{-15, 15}
var fontSizeRange = [2]int{55, 90}
var wavelengthRange = [2]int{5, 10}

// Generates a list of selected words with indices of what should be the right answers.
func generateWords(wrongNum int, answerNum int) []int {

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
	return ansIdx
}

func drawRandomLines(command []string) []string {
	// Draw random lines
	lineCount := utils.RandomRangeArray(linesRange)
	for i := 0; i < lineCount; i++ {
		a1 := rand.Intn(cfg.W + cfg.H)
		a2 := rand.Intn(cfg.W + cfg.H)
		l1 := strconv.Itoa(utils.Tern(a1 < cfg.W, a1, cfg.W+pad)) + "," + strconv.Itoa(utils.Tern(a1 > cfg.W, a1-cfg.W, -pad))
		l2 := strconv.Itoa(utils.Tern(a2 < cfg.H, -pad, a2-cfg.H)) + "," + strconv.Itoa(utils.Tern(a2 < cfg.H, a2, cfg.H+pad))
		command = append(command,
			"-strokewidth",
			strconv.Itoa(utils.RandomRangeArray(strokeRange)),
			"-stroke",
			cfg.Colors[rand.Intn(len(cfg.Colors))],
			"-draw",
			fmt.Sprintf("line %s %s", l1, l2),
		)
	}

	// Reset stroke
	command = append(command,
		"-stroke",
		"None",
	)
	return command
}

func drawAndGenerateChalenges(command []string, textAmount int, correctAnswerIndex []int) ([]string, []QuestionAnswer) {
	// Sort it
	slices.Sort(correctAnswerIndex)

	// Whether to use font or not
	usefont := len(cfg.FontList) >= 0

	// Store challenges here
	challenges := make([]QuestionAnswer, 0)

	// Words used
	words := make([]string, 0, textAmount)
	// Coords list
	coords := make([]Coords, 0, textAmount)
	// Colors left that we can use
	colorLeft := make([]string, textAmount)

	// Copy the color slice so we have a reference
	copy(colorLeft, cfg.Colors)

	// Index to keep track of font answers
	curAnswerI := 0

	// Loop to create challenges
	for i := 0; i < textAmount; i++ {
		// Get distict word
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

		// If the current index is the right answer, append to challenge
		if correctAnswerIndex[curAnswerI] == i {
			// This is the answer
			challenges = append(challenges, QuestionAnswer{
				Question: word,
				Answer:   color,
			})
			if curAnswerI < len(correctAnswerIndex)-1 {
				curAnswerI++
			}
		}

		// Creates the words at certain coordinates and stores the location.
		// Location is stored so we can create words that does not collide with previous words.
		var x int
		var yFromMid int // Because we are using gravity west, this starts from mid
		for i := 0; i < textMaxIterate; i += 1 {
			x = utils.RandomRangeArray(xRange)
			yFromMid = utils.RandomRangeArray(yHalfRange)
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
		rot := utils.RandomRangeArray(rotateRange)

		// Generate word command
		command = append(command,
			"-fill",
			color,
			"-pointsize",
			strconv.Itoa(utils.RandomRangeArray(fontSizeRange)),
			"-annotate",
			fmt.Sprintf("%dx%d+%d+%d", rot, rot, x, yFromMid),
			word,
		)
	}

	return command, challenges
}

// Generates a new captcha file and contains the answer
func GenCaptcha(wrongNum int, answerNum int) (Result, error) {

	// Initial command
	command := []string{
		"-size",
		strconv.Itoa(cfg.W) + "x" + strconv.Itoa(cfg.H),
		"xc:white",
		"-quality",
		cfg.Quality,
		"-gravity",
		"West",
	}

	ansIdx := generateWords(wrongNum, answerNum)
	command = drawRandomLines(command)
	var challenges []QuestionAnswer
	command, challenges = drawAndGenerateChalenges(command, wrongNum+answerNum, ansIdx)

	// Add waves
	command = append(command,
		"-wave",
		strconv.Itoa(utils.RandomRangeArray(wavelengthRange))+"x100", // Waves
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
	xRange = [2]int{0, cfg.W * 6 / 10}
	yHalfRange = [2]int{-cfg.H * 3 / 10, cfg.H * 3 / 10}
}
