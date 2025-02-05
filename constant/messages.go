package constant

import "kaeru.chan/voz/message"

var (
	CommunityMessages = map[string][]message.Template{
		"bê đê,pê đê,buê đuê": {
			{
				Message:    "đồng ý, thật kingtomb",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
				HasRef:     true,
			},
			{
				Message:    "Có người bảo em rằng",
				Emoji:      SexyPepe,
				ReactEmoji: Aragorn,
				Url:        "https://media.discordapp.net/attachments/880833106297381014/1161861154839208007/image.png?ex=6539d685&is=65276185&hm=0efcc32bfdf9c1b4cbc511340d620598c28566a2900656def07dfc2b05d04efc&=&width=1520&height=312",
				HasRef:     true,
				TagUsers:   []string{"596658100841480192"},
			},
			{
				Message:    "Có người bảo em rằng",
				Emoji:      SexyPepe,
				ReactEmoji: Aragorn,
				Url:        "https://cdn.discordapp.com/attachments/880833106297381014/1164216566393622608/image.png?ex=6542682b&is=652ff32b&hm=3542d2fb4277756f3fed4de82f20715bd21f4e77bb79c8c4b229e70d8d65b68a&",
				HasRef:     true,
				TagUsers:   []string{"632196353354629160"},
			},
			{
				Message:    "Buscu time.",
				Emoji:      SexyPepe,
				ReactEmoji: Aragorn,
				Url:        "https://cdn.discordapp.com/attachments/872829278872567890/1281490049434914857/Screenshot_2024-09-06_at_12.43.35_PM.png?ex=66df33dc&is=66dde25c&hm=83bcaec5540bf698e756bcb16f724f04f688582a1bf621c9adb88457f7697304&",
				HasRef:     true,
				TagUsers:   []string{"367309546823221250"},
			},
		},
		"nghèo": {
			{
				Message:    "Nghèo thì đứng sang 1 bên?",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
				HasRef:     true,
			},
		},
		"bô": {
			{
				Message:    "Bô bô cl, úp hoài",
				Emoji:      Tsk,
				ReactEmoji: What,
				Url:        "",
				HasRef:     true,
			},
			{
				Message:    "Bô này có vẻ thơm",
				Emoji:      SexyPika,
				ReactEmoji: Aragorn,
				Url:        "",
				TagUsers:   []string{EchChanID},
			},
			{
				Message:    "Bô bô cc",
				Emoji:      WibuSpank,
				ReactEmoji: Tsk,
				Url:        "",
			},
		},
		"chửi": {
			{
				ReactEmoji: WibuSpank,
				Url:        "https://cdn.discordapp.com/attachments/872829278872567890/1162264345925931028/image.png?ex=653b4e05&is=6528d905&hm=32f1575374d725bd3b6a4bb70cab4b2ff0b5c988e93f29bb16885a12dcabf72b&",
			},
			//{
			//	Message:    "ngộ ha, tự nhiên chửi người ta",
			//	Emoji:      Tsk,
			//	ReactEmoji: WibuSpank,
			//	Url:        "",
			//	HasRef:     true,
			//},
		},
		"long xiên,póng chan,pông nan": {
			{
				Message:    "Hãy để trú yên, trú cần thời gian đi úp bô",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
			},
			{
				Message:    "Bận làm pông nan rồi, tag cl",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
			},
		},
		"join fast": {
			{
				Message:    "Có cho tiền không mà kêu theo?",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
				HasRef:     true,
			},
			{
				Message:    "Ai nhà nghèo thì đứng sang 1 bên",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
			},
			{
				Message:    "Anh theo em theo",
				Emoji:      SexyPepe,
				ReactEmoji: Aragorn,
				Url:        "",
				HasRef:     true,
			},
		},
		"nwng, nứng": {
			{
				Message:    "Thư kí nứng đến đây!",
				Emoji:      SexyPepe,
				ReactEmoji: SexyPika,
				Url:        "",
				HasRef:     true,
				TagUsers:   []string{"376386392311201793"},
			},
		},
	}
	MsgReplyTagged = []string{
		"Không hay ho gì mà tag",
		"Thích bị tag không?",
		"Làm gì mà tag?",
	}
)
