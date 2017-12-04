package main

import (
	"flag"
	"strings"

	"github.com/pascalw/go-alfred"
)

var images = map[string][]string{
	"01": {"lgtm"},
	"02": {"code ga kirei!"},
	"03": {"hello world"},
	"04": {"sasuga desu"},
	"05": {"shinchoku doudesuka"},
	"06": {"shinchoku dame desu"},
	"07": {"teiji agari"},
	"08": {"vr kitaku", "go home"},
	"09": {"shigoto hayai"},
	"10": {"genki dashite", "cheer up"},
	"11": {"kami", "god"},
	"12": {"sugoi", "amazing"},
	"13": {"arigatougozaimasu", "thank you"},
	"14": {"tasukarimasu"},
	"15": {"ryokai desu"},
	"16": {"onegaishimasu"},
	"17": {"yarushikanai"},
	"18": {"osushi dozo"},
	"19": {"nerumade ga kyou"},
	"20": {"yusho shita"},
	"21": {"naki", "cry"},
	"22": {"heart"},
	"23": {"owata"},
	"24": {"moshikashite"},
	"25": {"chien de okuremasu"},
	"26": {"jitaku sagyo shimasu", "home work"},
	"27": {"shiyou desu"},
	"28": {"ima hanashikaketemo daijyoubu kana"},
	"29": {"otsukaresama"},
	"30": {"yasumou", "rest"},
	"31": {"oyasumi", "good night"},
	"32": {"minomushi"},
	"33": {"iine", "good"},
	"34": {"sakunomi shiyou"},
	"35": {"release shimashita"},
	"36": {"naruhodo", "wakaran"},
	"37": {"yoi engineer shokai shite"},
	"38": {"yoi engineer da"},
	"39": {"nenpou 5oku"},
	"40": {"soredewa yoi engineer life wo", "good engineer life"},
}

var (
	keyword = flag.String("k", "", "help message for k option")
)

func main() {

	flag.Parse()

	response := alfred.NewResponse()

	for k, values := range images {
		for _, v := range values {
			if strings.Contains(v, *keyword) {
				response.AddItem(&alfred.AlfredResponseItem{
					Valid:    true,
					Arg:      k,
					Uid:      k,
					Title:    strings.Join(values, ", "),
					Icon:     "images/" + k + ".png",
					Subtitle: "Enter to copy url to clipboard. (Cmd + Enter to copy as markdown.)",
				})

				break
			}
		}
	}

	response.Print()
}
