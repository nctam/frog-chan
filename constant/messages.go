package constant

import "kaeru.chan/voz/message"

var (
	FonfonMessages = map[string][]message.TargetMessage{
		"đĩ,bitch": {
			{
				Message:    "đồ đĩ",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
			},
		},
		"ktv,bu cu, buscu": {
			{
				Message:    "buông bỏ đặc cầu đi trú",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
			},
		},
		"póng,póng chan,pông nan": {
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
	}

	PongMessages = map[string][]message.TargetMessage{
		"chửi": {
			{
				Message:    "đồ đĩ",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
				HasRef:     true,
			},
			{
				Message:    "ngộ ha, tự nhiên chửi người ta",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
			},
		},
		"join fast,đâm,theo": {
			{
				Message:    "Có tiền không mà theo?",
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
		},
	}

	CommunityMessages = map[string][]message.TargetMessage{
		"bd,đĩ đực,gei lord,gay,gei": {
			{
				Message:    "kingtomb",
				Emoji:      Tsk,
				ReactEmoji: Aragorn,
				Url:        "",
				HasRef:     true,
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
			{
				Message:    "Anh nghèo em cũng chẳng hơn anh...LK Nghèo Trường Vũ",
				Emoji:      TooSad,
				ReactEmoji: Ree,
				Url:        "",
				HasRef:     true,
			},
		},
	}
)
