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
                TagUsers:   []string{MikeyID},
            },
            {
                Message:    "Có người bảo em rằng",
                Emoji:      SexyPepe,
                ReactEmoji: Aragorn,
                Url:        "https://cdn.discordapp.com/attachments/880833106297381014/1164216566393622608/image.png?ex=6542682b&is=652ff32b&hm=3542d2fb4277756f3fed4de82f20715bd21f4e77bb79c8c4b229e70d8d65b68a&",
                HasRef:     true,
                TagUsers:   []string{PiChanID},
            },
            {
                Message:    "Buscu time.",
                Emoji:      SexyPepe,
                ReactEmoji: Aragorn,
                Url:        "https://cdn.discordapp.com/attachments/872829278872567890/1281490049434914857/Screenshot_2024-09-06_at_12.43.35_PM.png?ex=66df33dc&is=66dde25c&hm=83bcaec5540bf698e756bcb16f724f04f688582a1bf621c9adb88457f7697304&",
                HasRef:     true,
                TagUsers:   []string{FonChanID},
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
            {
                Message:    "pông nan cc, mua cục sạc đi",
                Emoji:      Tsk,
                ReactEmoji: Tsk,
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
        "nwng,nứng": {
            {
                Message:  "Thư kí nứng đến đây!",
                Emoji:    Nwng1,
                Url:      "https://cdn.discordapp.com/attachments/1048640658795155476/1338393436549681262/2BA5B04F-221F-421B-B3FC-9CD8859E0E08.png?ex=67aaeb85&is=67a99a05&hm=25726decd5122902e15086c163e1db08dff58a9991305207a13b1b8878ca20ad&",
                HasRef:   false,
                TagUsers: []string{CaDuChanID},
            },
            {
                Message:  "Ra chưa?",
                Emoji:    SexyPepe,
                Url:      "",
                HasRef:   false,
                TagUsers: []string{},
            },
            {
                Message:  "Muốn thành kẻ hiepdam!",
                Emoji:    SexyPepe,
                Url:      "https://cdn.discordapp.com/attachments/1048640658795155476/1336593591740534845/image.png?ex=67a50808&is=67a3b688&hm=9eaae88cfca506c821eba21dd140f8a12442710b0b045c73e1beb74ef164773a&",
                HasRef:   false,
                TagUsers: []string{CaDuChanID},
            },
            {
                Message:    "Ai muốn nứng bước ra đây coi!",
                ReactEmoji: Nwng1,
                Url:        "https://media.discordapp.net/attachments/1048640658795155476/1338697210786746388/20250211_091914.png?ex=67ae00ae&is=67acaf2e&hm=4fa0a1ed4b9915eddce22111db9659e982313def4c94519798d726ab8331f2ac&=&format=webp&quality=lossless&width=738&height=359",
                HasRef:     true,
                TagUsers:   []string{PiChanID},
            },
            {
                Message:    "Mất dái rồi",
                ReactEmoji: Nwng1,
                Url:        "https://cdn.discordapp.com/attachments/1048640658795155476/1335876517900849212/image.png?ex=67b19575&is=67b043f5&hm=91240650fa233f84429fdcca4a7cab3868db82c8028f6952bf168f2335d20052&",
                HasRef:     true,
                TagUsers:   []string{PiChanID},
            },
            {
                Message:    "Biết đụ trên thùng xe không?",
                ReactEmoji: Nwng1,
                Url:        "https://cdn.discordapp.com/attachments/1048640658795155476/1337284647171784745/image.png?ex=67b16ee1&is=67b01d61&hm=f75657d916950337bed78923a3bb108b51c0bf4d66ca414c75c432707a62a984&",
                HasRef:     true,
                TagUsers:   []string{},
            },
        },
        "tnt": {
            {
                Message:    "vinh ko phải tnt",
                ReactEmoji: ThietLuon,
                Url:        "https://media.discordapp.net/attachments/1048640658795155476/1335213961792196679/image0.jpg?ex=67b1cf67&is=67b07de7&hm=9978f51580024ebbffdd6bc47f0ebca7a7fbc9bb8c4a0eba5529f57823db83c1&=&format=webp&width=857&height=785",
                HasRef:     true,
                TagUsers:   []string{FonChanID},
            },
        },
        "sạc": {
            {
                Message:    "sạc sạc cl, mua pông nan đi",
                Emoji:      Tsk,
                ReactEmoji: WibuSpank,
                Url:        "",
                HasRef:     true,
            },
        },
    }
    MsgReplyTagged = []string{
        "Không hay ho gì mà tag",
        "Thích bị tag không?",
        "Làm gì mà tag?",
    }
)
