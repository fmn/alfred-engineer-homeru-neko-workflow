package main

import (
	"flag"
	"strings"

	"github.com/pascalw/go-alfred"
)

var images = map[string][]string{
	"01": {"lgtm"},
	"02": {"kirei"},
	"03": {"hello world"},
	"04": {"sasuga"},
	"05": {"shinchoku doudesuka", "How is the progress"},
	"06": {"shinchoku ng"},
	"07": {"teiji"},
	"08": {"vr kitaku", "go home"},
	"09": {"shigoto hayai"},
	"10": {"cheer up"},
	"11": {"kami", "god"},
	"12": {"sugoi", "amazing"},
	"13": {"arigato", "thank you"},
	"14": {"tasukarimasu"},
	"15": {"ryokai"},
	"16": {"onegaishimasu"},
	"17": {"yarushikanai"},
	"18": {"osushi"},
	"19": {"nerumade ga kyou"},
	"20": {"yusho"},
	"21": {"naki", "cry"},
	"22": {"heart"},
	"23": {"owata"},
	"24": {"moshikashite"},
	"25": {"chien", "delay"},
	"26": {"jitaku sagyo", "home work"},
	"27": {"shiyo desu"},
	"28": {"ima hanashikaketemo daijyoubu?"},
	"29": {"otsukaresama"},
	"30": {"yasumou", "rest"},
	"31": {"oyasumi", "good night"},
	"32": {"minomushi"},
	"33": {"iine", "good"},
	"34": {"sakunomi"},
	"35": {"release"},
	"36": {"naruhodo", "wakaran"},
	"37": {"shokai"},
	"38": {"good engineer"},
	"39": {"5oku"},
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
