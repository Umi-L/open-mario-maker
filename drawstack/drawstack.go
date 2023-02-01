package drawstack

import (
	"github.com/hajimehoshi/ebiten/v2"
	"sort"
)

type DrawCall func(image *ebiten.Image)

type DrawStack map[int][]DrawCall

func (drawStack *DrawStack) Draw(screen *ebiten.Image) {

	stack := *drawStack

	//get keys in order
	keys := make([]int, len(stack))

	i := 0
	for k := range stack {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	//make Draw calls
	for j := range stack {
		calls := stack[keys[j]]
		for _, call := range calls {
			call(screen)
		}
	}

	//clear stack
	*drawStack = make(map[int][]DrawCall)
}

func (drawStack *DrawStack) Add(call DrawCall, zIndex int) {
	stack := *drawStack

	stack[zIndex] = append(stack[zIndex], call)
}
