package voice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	soundsFolder = "sounds"
)

func searchVoiceChannel(discord *discordgo.Session, user string) (voiceChannelID string) {
	for _, g := range discord.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == user {
				return v.ChannelID
			}
		}
	}
	return ""
}

func PlaySound(discord *discordgo.Session, guildID, userID, filename string) {
	voiceChannelID := searchVoiceChannel(discord, userID)
	// Connect to voice channel.
	// NOTE: Setting mute to false, deaf to true.
	voice, err := discord.ChannelVoiceJoin(guildID, voiceChannelID, false, true)
	if err != nil {
		log.Fatalf("FAILED TO CONNECT TO VOICE CHANNEL: %v", err)
	}
	voice.LogLevel = discordgo.LogWarning

	// Hacky loop to prevent sending on a nil channel.
	// TODO: Find a better way.
	for voice.Ready == false {
		runtime.Gosched()
	}

	path := fmt.Sprintf("%s/%s", soundsFolder, filename)
	log.Printf("Searching audio file: %v", path)

	_, err = os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to find sound file: %v, error: '%v'", filename, err)
	}

	discord.UpdateGameStatus(0, filename)
	playAudioFile(voice, path)

	voice.Disconnect()
	// Close connections
	voice.Close()
}

// playAudioFile will play the given filename to the already connected
// Discord voice server/channel.  voice websocket and udp socket
// must already be setup before this will work.
func playAudioFile(v *discordgo.VoiceConnection, filepath string) {
	log.Printf("Playing sound file: %v", filepath)

	// Send "speaking" packet over the voice websocket
	err := v.Speaking(true)
	if err != nil {
		log.Fatal("Failed setting speaking", err)
	}

	// Send not "speaking" packet over the websocket when we finish
	defer v.Speaking(false)

	opts := dca.StdEncodeOptions
	opts.RawOutput = true
	opts.Bitrate = 120

	encodeSession, err := dca.EncodeFile(filepath, opts)
	if err != nil {
		log.Fatal("Failed creating an encoding session: ", err)
	}

	done := make(chan error)
	stream := dca.NewStream(encodeSession, v, done)

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case err := <-done:
			if err != nil && err != io.EOF {
				log.Fatal("An error occured", err)
			}

			// Clean up incase something happened and ffmpeg is still running
			encodeSession.Truncate()
			return
		case <-ticker.C:
			stats := encodeSession.Stats()
			playbackPosition := stream.PlaybackPosition()

			log.Printf("Playback: %10s, Transcode Stats: Time: %5s, Size: %5dkB, Bitrate: %6.2fkB, Speed: %5.1fx\r", playbackPosition, stats.Duration.String(), stats.Size, stats.Bitrate, stats.Speed)
		}
	}
}
