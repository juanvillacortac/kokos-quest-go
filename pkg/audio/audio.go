package audio

import (
	"fmt"
	"strings"

	"kokos_quest/assets"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

var (
	suffixes = []string{".mp3", ".ogg", ".wav"}

	audioContext *audio.Context

	bgmPlayers = map[string]*audio.Player{}
	sePlayers  = map[string]*audio.Player{}

	mute = false

	bgmKeys []string
	seKeys  []string
)

func bgm(slice map[string]*audio.Player) []string {
	keys := make([]string, 0, len(slice))
	for key := range slice {
		keys = append(keys, key)
	}
	return keys
}

func IsMute() bool {
	return mute
}

func SetMute(enabled bool) {
	mute = enabled
}

func init() {
	bgmKeys = assets.LoadKeys("music", suffixes)
	seKeys = assets.LoadKeys("sounds", suffixes)
}

func init() {
	const sampleRate = 44100
	var err error
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		panic(err)
	}
}

func Load() error {
	soundDirs := []string{
		"music",
		"sounds",
	}
	for _, dir := range soundDirs {
		filenames := make([]string, 0)
		switch dir {
		case "music":
			filenames = bgmKeys
		case "sounds":
			filenames = seKeys
		}
		for _, n := range filenames {
			b, err := assets.Asset(dir + "/" + n)
			if err != nil {
				return err
			}
			f := audio.BytesReadSeekCloser(b)
			var stream, decoded audio.ReadSeekCloser
			var length int64
			switch {
			case strings.HasSuffix(n, ".mp3"):
				decoded, err = mp3.Decode(audioContext, f)
				if err != nil {
					return err
				}
				length = decoded.(*mp3.Stream).Length()
				n = strings.TrimSuffix(n, ".mp3")
			case strings.HasSuffix(n, ".ogg"):
				decoded, err = vorbis.Decode(audioContext, f)
				if err != nil {
					return err
				}
				length = decoded.(*vorbis.Stream).Length()
				n = strings.TrimSuffix(n, ".ogg")
			case strings.HasSuffix(n, ".wav"):
				decoded, err = wav.Decode(audioContext, f)
				if err != nil {
					return err
				}
				length = decoded.(*wav.Stream).Length()
				n = strings.TrimSuffix(n, ".wav")
			default:
				panic("invalid file name")
			}
			switch dir {
			case "music":
				stream = audio.NewInfiniteLoop(decoded, length)
			case "sounds":
				stream = decoded
			}
			p, err := audio.NewPlayer(audioContext, stream)
			if err != nil {
				return err
			}
			switch dir {
			case "music":
				bgmPlayers[n] = p
			case "sounds":
				sePlayers[n] = p
			}
		}
	}

	return nil
}

func Finalize() error {
	soundDirs := []string{
		"music",
		"sounds",
	}
	for _, dir := range soundDirs {
		var soundPlayers map[string]*audio.Player
		switch dir {
		case "music":
			soundPlayers = bgmPlayers
		case "sounds":
			soundPlayers = sePlayers
		}
		for _, p := range soundPlayers {
			if err := p.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}

func SetBGMVolume(volume float64) {
	if mute {
		return
	}
	for _, p := range bgmPlayers {
		if !p.IsPlaying() {
			continue
		}
		p.SetVolume(volume)
		return
	}
}

func PauseBGM() {
	if mute {
		return
	}
	for _, p := range bgmPlayers {
		p.Pause()
	}
}

func ResumeBGM(bgm string) {
	if mute {
		return
	}
	PauseBGM()
	p := bgmPlayers[bgm]
	if p == nil {
		err := fmt.Errorf(`BGM "%s" doesn't found on memory`, bgm)
		panic(err)
	}
	p.SetVolume(1)
	p.Play()
}

func PlayBGM(bgm string) error {
	if mute {
		return nil
	}
	PauseBGM()
	p := bgmPlayers[bgm]
	if p == nil {
		err := fmt.Errorf(`BGM "%s" doesn't found on memory`, bgm)
		panic(err)
	}
	p.SetVolume(1)
	if err := p.Rewind(); err != nil {
		return err
	}
	p.Play()
	return nil
}

func PlaySE(se string) {
	if mute {
		return
	}
	p := sePlayers[se]
	if p == nil {
		err := fmt.Errorf(`Sound Effect "%s" doesn't found on memory`, se)
		panic(err)
	}
	p.Rewind()
	p.Play()
}
