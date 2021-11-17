package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/bwmarrin/discordgo"
	"mvdan.cc/xurls/v2"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "!pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content[0:7] == "!censor" {

		var filename string = ""
		xurlsStrict := xurls.Strict()
		output := xurlsStrict.FindAllString(m.Content, -1)

		// download as png
		if output[0][len(output[0])-4:] == ".png" {
			filename = "img_" + strconv.Itoa(rand.Intn(1000)) + "_" + m.Author.ID + ".png"
			DownloadFile(filename, output[0])
		}

		// read file for editing
		img, err := imgio.Open(filename)

		result := blur.Gaussian(img, 256.0)

		if err := imgio.Save("censored_"+filename, result, imgio.PNGEncoder()); err != nil {
			fmt.Println(err)
			return
		}

		// read file for sending
		file, err := os.Open("censored_" + filename)
		if err != nil {
			fmt.Println("Fucked")
		}
		s.ChannelFileSend(m.ChannelID, "censored.png", file)
		os.Remove(filename)
		os.Remove("censored_" + filename)
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
