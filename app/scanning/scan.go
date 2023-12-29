package scanning

import (
	"fmt"
	"image"

	"github.com/airtonix/pixelclicker/app/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// A scanner is an object that scans a sequence of frames
// each describing a matric of pixelsx.
// It provides a method to scan a provided image and reports
// a degree of condifence that the image contains a match.
type IScanner interface {
	AddPattern(pattern image.Image) error
	Scan() error
}

type ScanResult struct {
	Confidence int
}

type Scanner struct {
	pattern   image.Image
	history   *core.Stack[ScanResult]
	threshold int
}

type ScannerOptFunc func(*Scanner)

func WithThreshold(threshold int) ScannerOptFunc {
	return func(s *Scanner) {
		s.threshold = threshold
	}
}

func NewScanner(optionsFn ...ScannerOptFunc) *Scanner {
	scanner := &Scanner{
		history: core.NewStack[ScanResult](10),
	}
	for _, opt := range optionsFn {
		opt(scanner)
	}
	return scanner
}

// AddPattern adds a pattern to the scanner
// A pattern is a matrix of pixels that the scanner
// will attempt to match against an image
func (s *Scanner) AddPattern(pattern image.Image) error {
	s.pattern = pattern
	return nil
}

func (s *Scanner) HasPattern() bool {
	return s.pattern != nil
}

func (s *Scanner) GetPatternSize() string {
	if s.pattern == nil {
		return "0x0"
	}
	// calculate the byte size of the pattern
	bounds := s.pattern.Bounds().Size()
	size := bounds.X * bounds.Y

	return fmt.Sprintf("%d", size)
}

func (s *Scanner) CapturePattern() error {
	captured := CaptureScreen()
	s.AddPattern(captured)
	return nil
}

func (s *Scanner) GetHistory() *core.Stack[ScanResult] {
	return s.history
}

// Scan scans the provided image and returns a degree of confidence
func (s *Scanner) Scan() error {
	// captured := CaptureScreen()

	// if captured == nil {
	// 	log.Printf("captured is nil")
	// 	return nil
	// }

	// log.Printf("captured: %v", captured)

	// loop over patterns
	// look at image, how much does it match the pattern
	// distance, err := CompareImages(
	// 	&captured,
	// 	&s.pattern,
	// )

	// if err != nil {
	// 	return err
	// }

	// record confidence
	s.history.Push(ScanResult{
		Confidence: 1,
	})

	return nil
}

// draws a line chart for each scan down the side of the screen
func (s *Scanner) ReportScan(
	screen *ebiten.Image,
	screenWidth,
	screenHeight float32,
) error {
	if s.history.IsEmpty() {
		return nil
	}

	// draw a line chart
	// for each scan
	// down the side of the screen
	// with the confidence level
	// of the scan

	data := s.history.Pop()

	ebitenutil.DebugPrintAt(
		screen,
		"Confidence: "+string(data.Confidence),
		int(screenWidth-100),
		0,
	)

	return nil
}
