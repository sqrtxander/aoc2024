package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func genGradient(start, end tcell.Color, count int) []tcell.Color {
	result := make([]tcell.Color, count)
	result[0] = start
	result[count-1] = end
	r1, g1, b1 := start.RGB()
	r2, g2, b2 := end.RGB()

	r := float32(r1)
	g := float32(g1)
	b := float32(b1)

	dr := float32(r2-r1) / float32(count-1)
	dg := float32(g2-g1) / float32(count-1)
	db := float32(b2-b1) / float32(count-1)

	for i := 1; i < count-1; i++ {
		r += dr
		g += dg
		b += db
		result[i] = tcell.NewRGBColor(int32(r), int32(g), int32(b))
	}

	return result
}

type stat struct {
	stat  string
	style tcell.Style
}

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	uniqueNums := 10
	numBGCols := genGradient(
		tcell.NewHexColor(0xff0000),
		tcell.NewHexColor(0x00ff00),
		uniqueNums,
	)
	numStyles := make([]tcell.Style, uniqueNums)
	for i := 0; i < uniqueNums; i++ {
		numStyles[i] = tcell.StyleDefault.Background(numBGCols[i]).Foreground(tcell.ColorWhite)
	}

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	ox, oy := 1, 1

	eventChan := make(chan tcell.Event)
	go func() {
		for {
			eventChan <- s.PollEvent()
		}
	}()

	gridStr := getContents()
	gridLines := strings.Split(gridStr, "\n")
	resultChan := make(chan updateState)
	go solve(gridStr, &resultChan)

	drawText(s, ox, oy, ox+len(gridLines[0])+1, oy+len(gridLines)+1, defStyle, gridStr)

	headers := []string{
		"      |  Perimeter  |  Area  |  Price  |",
		"      +=============+========+=========+",
	}
	internal := "      |             |        |         |"
	footers := []string{
		"      +=============+========+=========+",
		"TOTAL |                                |",
		"      +================================+",
	}

	statX := ox + len(gridLines[0]) + ox
	footY := oy + len(gridLines) - len(footers) - 1

	drawText(s, statX, oy, statX+len(headers[0])+1, oy+len(headers)+1, defStyle, strings.Join(headers, "\n"))
	for y := oy + len(headers); y < footY; y++ {
		drawText(s, statX, y, statX+len(internal)+1, y, defStyle, internal)
	}

	drawText(s, statX, footY, statX+len(footers[0])+1, footY+len(footers)+1, defStyle, strings.Join(footers, "\n"))
	s.Show()
	total := 0
	prevScore := 0
	statQueue := utils.Queue[stat]{}
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	for {
		select {
		case ev := <-eventChan:
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					return
				}
			}
		case <-ticker.C:
			select {
			case state := <-resultChan:
				s.SetContent(ox+state.x, oy+state.y, rune(gridLines[state.y][state.x]), nil, state.style)
				if state.new {
					statQueue = statQueue.Push(stat{
						stat:  fmt.Sprintf("  % 9d  |  % 4d  |  % 5d  ", state.perim, state.area, state.perim*state.area),
						style: state.style,
					})
				} else {
					statQueue[len(statQueue)-1] = stat{
						stat:  fmt.Sprintf("  % 9d  |  % 4d  |  % 5d  ", state.perim, state.area, state.perim*state.area),
						style: state.style,
					}
					total -= prevScore
				}
				prevScore = state.perim * state.area
				total += state.perim * state.area
				if len(statQueue) > len(gridLines)-len(headers)-len(footers)-1 {
					statQueue, _ = statQueue.Pop()
				}
				y := footY - 1
				for i := len(statQueue) - 1; i >= 0; i-- {
					stat := statQueue[i]
					s.SetContent(statX+len("TOTAL "), y, '|', nil, defStyle)
					x := statX + len("TOTAL ") + 1
					for _, str := range strings.Split(stat.stat, "|") {
						drawText(s, x, y, x+len(str)+1, y, stat.style, str)
						x += len(str)
						s.SetContent(x, y, '|', nil, defStyle)
						x++
					}
					y--
				}
				totalStr := fmt.Sprintf("TOTAL |  % 28d  |", total)
				drawText(s, statX, footY+1, statX+len(totalStr), footY+1, defStyle, totalStr)
			default:
				break
			}
			s.Show()
		}
	}
}
