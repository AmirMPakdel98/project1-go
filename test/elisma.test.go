package test

/*
import (
	"fmt"
	"log"
	"os/exec"
)

func segmentedWithKeys() {

	packager := ""

	keys := []struct {
		id    string
		value string
	}{
		{id: "f3c5e0361e6654b28f8049c778b23946", value: "a4631a153a443df9eed0593043db7519"},
		{id: "6d76f25cb17f5e16b8eaef6bbf582d8e", value: "cb541084c99731aef4fff74500c12ead"},
		{id: "5d76f25cb17f5e16b8eaef6bbf582d8e", value: "ab541084c99731aef4fff74500c12ead"},
	}

	file_id := fmt.Sprintf("%d", file.Id)

	video480_input := "./../storage/stage2/" + file_id + "_480.mp4"
	video720_input := "./../storage/stage2/" + file_id + "_720.mp4"
	audio_init := "./../storage/stage3/" + file_id + "/" + file_id + "_audio_init.mp4"
	audio_segtem := "./../storage/stage3/" + file_id + "/" + file_id + "_audio_$Number$.m4s"
	video480_init := "./../storage/stage3/" + file_id + "/" + file_id + "_480p_init.mp4"
	video480_segtem := "./../storage/stage3/" + file_id + "/" + file_id + "_480p_$Number$.m4s"
	video720_init := "./../storage/stage3/" + file_id + "/" + file_id + "_720p_init.mp4"
	video720_segtem := "./../storage/stage3/" + file_id + "/" + file_id + "_720p_$Number$.m4s"
	mpd_out := "./../storage/stage3/" + file_id + "/" + file_id + ".mpd"
	segment_duration := "8"

	cmd := exec.Command(packager,
		fmt.Sprintf("in=%s,stream=audio,init_segment=%s,segment_template=%s,drm_label=AUDIO", video720_input, audio_init, audio_segtem),
		fmt.Sprintf("in=%s,stream=video,init_segment=%s,segment_template=%s,drm_label=480p", video480_input, video480_init, video480_segtem),
		fmt.Sprintf("in=%s,stream=video,init_segment=%s,segment_template=%s,drm_label=720p", video720_input, video720_init, video720_segtem),
		"--keys",
		fmt.Sprintf("label=AUDIO:key_id=%s:key=%s,", keys[0].id, keys[0].value)+
			fmt.Sprintf("label=480p:key_id=%s:key=%s,", keys[1].id, keys[1].value)+
			fmt.Sprintf("label=720p:key_id=%s:key=%s", keys[2].id, keys[2].value),
		"--pssh",
		"000000317073736800000000EDEF8BA979D64ACEA3C827DCD51D21ED00000011220F7465737420636F6E74656E74206964",
		"--enable_raw_key_encryption",
		"--segment_duration",
		segment_duration,
		"--mpd_output",
		mpd_out,
	)

	log.Printf("%v", cmd)
}

func normal() {

	packager := ""

	file_id := fmt.Sprintf("%d", file.Id)

	video720_input := "./../storage/stage2/" + file_id + "_720.mp4"
	audio_init := "./../storage/stage3/" + file_id + "/" + file_id + "_audio_init.mp4"
	video720_init := "./../storage/stage3/" + file_id + "/" + file_id + "_720p_init.mp4"
	mpd_out := "./../storage/stage3/" + file_id + "/" + file_id + ".mpd"

	cmd := exec.Command(packager,
		fmt.Sprintf("in=%s,stream=audio,output=%s,drm_label=AUDIO", video720_input, audio_init),
		fmt.Sprintf("in=%s,stream=video,output=%s,drm_label=720p", video720_input, video720_init),
		"--pssh",
		"000000317073736800000000EDEF8BA979D64ACEA3C827DCD51D21ED00000011220F7465737420636F6E74656E74206964",
		"--mpd_output",
		mpd_out,
	)

	log.Printf("%v", cmd)
}
*/

/*
./packager-osx-x64 in=input.mp4,stream=audio,init_segment=output2_au_init.mp4,segment_template=output2_au_$Number$.m4s in=input.mp4,stream=video,init_segment=output_720p_init.mp4,segment_template=output2_720p_$Number$.m4s --segment_duration 4 --mpd_output output_1.mpd
*/
