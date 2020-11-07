package bot

import (
	"fmt"
	str "strings"

	"../config"
	"../features"
	dg "github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *dg.Session

func Start() {
	goBot, err := dg.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err)
	}

	BotID = u.ID

	goBot.AddHandler(messageCreate)

	goBot.AddHandler(ready)

	goBot.AddHandler(test)

	goBot.AddHandler(manageChannels)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Bot is Running")
}

func ready(s *dg.Session, event *dg.Ready) {
	//set the playing status
	s.UpdateStatus(0, "Bot in Development")
	//since := 1

	/*
		status := dg.UpdateStatusData{
			//IdleSince: &since,
			Game: &dg.Game{
							Name: "Pydroid & Termux",
							Type: 3,
							Details: "Making this bot that you are watching in go",
							TimeStamps: dg.TimeStamps{
													StartTimestamp: 1602918742000,
												},
						},
			AFK: false,
			Status: "test",
		}
		err := s.UpdateStatusComplex(status)
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	//s.UpdateListeningStatus("Mr.Makra")

}

func messageCreate(s *dg.Session, m *dg.MessageCreate) {

	if str.HasPrefix(m.Content, config.BotPrefix) {

		if m.Author.ID == BotID || m.Author.Bot {
			return
		}
		cmd := command(m.Content)
		switch cmd {
		case "mute":
			features.Mute(s, m)
			break
		case "ping":
			features.Ping(s, m)
			break
		case "remind":
			features.Remind(s, m)
			break
		case "warn":
			features.Warn(s, m)
			break
		default:
			s.ChannelMessageSend(m.ChannelID, "**WHAT !?** :thinking::thinking:")
		}
	}

	if m.Content == "hi bot" || m.Content == "Hi bot" {
		s.ChannelMessageSend(m.ChannelID, "```Hello Hackers```")
	} /*
		if m.ChannelID == "765909517879476224" {
			fmt.Println(m.Content)
		}*/
}

func test(s *dg.Session, m *dg.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}
	if m.Content != "test" {
		return
	}

	textPerm := &dg.PermissionOverwrite{
		ID:   "772504744341798913",
		Type: "role",
		Deny: 0x800 | 0x40,
	}
	data := &dg.ChannelEdit{
		Position:             9,
		PermissionOverwrites: []*dg.PermissionOverwrite{textPerm},
	}

	chnns, _ := s.GuildChannels(m.GuildID)
	for i := range chnns {
		if chnns[i].Type == 0 && chnns[i].ID == "765909517879476224" {
			temp, err := s.ChannelEditComplex("765909517879476224", data)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println(temp.Name)
		}
	}
	s.ChannelMessageSend(m.ChannelID, "done")
	//msg, _ := s.ChannelMessageSend(m.ChannelID, "not edited")
	//s.ChannelMessageEdit(m.ChannelID, msg.ID, "edited")
	//s.MessageReactionAdd("765909517879476224","768750535464714271","232720527448342530")

}

func manageChannels(s *dg.Session, chans *dg.ChannelCreate) {

	features.AddMuteRole(s, chans)

}

/*
func validID(s *dg.Session, m *dg.MessageCreate, id string ) bool  {

	if m.Author.ID == s.State.User.ID {
		return false
	}
	id = str.Trim(id, "<>&!@#")
	_, err:= s.GuildMember(m.GuildID, m.Content)

	if err == nil {
		return true
	}
	return false
}*/

func command(s string) string {
	cmd := make([]rune, 0, 10)
	for i, val := range s {
		if i == 0 {
			continue
		}
		if val != ' ' {
			cmd = append(cmd, val)
		} else {
			break
		}
	}
	return string(cmd)
}
