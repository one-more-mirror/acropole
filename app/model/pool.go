package model

import "github.com/bwmarrin/discordgo"

type Vote struct {
	Time *discordgo.Timestamp
	User *discordgo.User
	Yes  bool
}

type Poll struct {
	Id      int
	Votes   []*Vote
	Creator *discordgo.User
}

func createPool(pool *Poll) {

}

func getPools() ([]*Poll, error) {
	pools := []*Poll{
		{
			Votes: []*Vote{
				{
					Time: nil,
					User: nil,
					Yes:  true,
				},
				{
					Time: nil,
					User: nil,
					Yes:  false,
				},
			},
			Creator: nil,
			Id:      1,
		},
	}

	return pools, nil
}
