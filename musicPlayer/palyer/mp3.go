package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat int
	progress int
}

func (p *MP3Player) Play(source string){
	fmt.Println("Playing music", source)
	p.progress = 0

	for p.progress <100{
		time.Sleep(100*time.Millisecond)
		fmt.Print('.')
		p.progress +=10
	}

	fmt.Println("\nFinished Player", source)
}

type WAVPlayer struct {
	stat int
	progress int
}

func (p *WAVPlayer) Play(source string){
	fmt.Println("Playing music", source)
	p.progress = 0

	for p.progress <100{
		time.Sleep(100*time.Millisecond)
		fmt.Print('.')
		p.progress +=10
	}

	fmt.Println("\nFinished Player", source)
}