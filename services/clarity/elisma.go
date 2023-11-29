package clarity

import (
	fileModel "c-vod/models/fileModel"
	"c-vod/utils/globals"
	"c-vod/utils/helper"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

/*
Elisma will handle the video stream formatting
*/
type Elisma struct{}

func (el *Elisma) convertToStreamableFormat(file *fileModel.File) error {

	// + elisma will convert 480p and 720p video
	var packager string

	switch runtime.GOOS {
	case "darwin":
		packager = "./packager-osx-x64"
	case "win32":
		packager = "./packager-win-x64.exe"
	case "linux":
		packager = "./packager-linux-x64"
	}

	_keys, err := helper.GenerateVideoKeys()
	keys := *_keys

	if err != nil {
		//TODO: rollback
		fmt.Println("error in calling GenerateVideoKeys() :", err.Error())
		return err
	}

	// + elisma will call edborn to insert vide keys record in db
	err = Edborn.insertVideoKeys(_keys, file.Id)

	if err != nil {
		//TODO: rollback
		fmt.Println("error on inserting videoKeys record:", err.Error())
		return err
	}

	/*
		keys = []struct {
			id    string
			value string
		}{
			{id: "f3c5e0361e6654b28f8049c778b23946", value: "a4631a153a443df9eed0593043db7519"},
			{id: "abba271e8bcf552bbd2e86a434a9a5d9", value: "69eaa802a6763af979e8d1940fb88392"},
			{id: "6d76f25cb17f5e16b8eaef6bbf582d8e", value: "cb541084c99731aef4fff74500c12ead"},
		}
	*/

	file_id := fmt.Sprintf("%d", file.Id)

	video480_input := "./../storage/stage2/" + file_id + "_480.mp4"
	video720_input := "./../storage/stage2/" + file_id + "_720.mp4"
	audio_init := "./../storage/stage3/" + file_id + "/__0000_ZGFzaF9zdHJlYW1fc2VnbWVudF9maWxl_" + file_id + "_audio_init.mp4"
	audio_segtem := "./../storage/stage3/" + file_id + "/__0000_ZGFzaF9zdHJlYW1fc2VnbWVudF9maWxl_" + file_id + "_audio_$Number$.m4s"
	video480_init := "./../storage/stage3/" + file_id + "/__0000_ZGFzaF9zdHJlYW1fc2VnbWVudF9maWxl_" + file_id + "_video_sd_init.mp4"
	video480_segtem := "./../storage/stage3/" + file_id + "/__0000_ZGFzaF9zdHJlYW1fc2VnbWVudF9maWxl_" + file_id + "_video_sd_$Number$.m4s"
	video720_init := "./../storage/stage3/" + file_id + "/__0000_ZGFzaF9zdHJlYW1fc2VnbWVudF9maWxl_" + file_id + "_video_hd_init.mp4"
	video720_segtem := "./../storage/stage3/" + file_id + "/__0000_ZGFzaF9zdHJlYW1fc2VnbWVudF9maWxl_" + file_id + "_video_hd_$Number$.m4s"
	mpd_out := "./../storage/stage3/" + file_id + "/__0000_ZGFzaF9zdHJlYW1fc2VnbWVudF9maWxl_" + file_id + "_dash_stream_manifest.mpd"
	segment_duration := "10"

	cmd := exec.Command(packager,
		fmt.Sprintf("in=%s,stream=audio,init_segment=%s,segment_template=%s,drm_label=AUDIO", video720_input, audio_init, audio_segtem),
		fmt.Sprintf("in=%s,stream=video,init_segment=%s,segment_template=%s,drm_label=SD", video480_input, video480_init, video480_segtem),
		fmt.Sprintf("in=%s,stream=video,init_segment=%s,segment_template=%s,drm_label=HD", video720_input, video720_init, video720_segtem),
		"--keys",
		fmt.Sprintf("label=AUDIO:key_id=%s:key=%s,", keys[0].Id, keys[0].Value)+
			fmt.Sprintf("label=SD:key_id=%s:key=%s,", keys[1].Id, keys[1].Value)+
			fmt.Sprintf("label=HD:key_id=%s:key=%s", keys[2].Id, keys[2].Value),
		"--enable_raw_key_encryption",
		"--segment_duration",
		segment_duration,
		"--generate_static_live_mpd",
		"--mpd_output",
		mpd_out,
	)

	cmd.Dir = filepath.Join(helper.GetCurrentDir(), "packager")

	if globals.App.Config.Log_enabled == "true" {
		// showing packager output in terminal
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	// + elisma will store output in a new folder with dbfile id in stage3 folder
	err = cmd.Run()

	if err != nil {
		//TODO: rollback
		fmt.Println("packager cmd error :", err.Error())
		return err
	}

	// + elisam will call trauma to delete stage2 files

	rm_err := Trauma.deleteNormalizedVideoSource(file)

	if rm_err != nil {
		//TODO: rollback
		fmt.Println("error on removing file from stage2 :", rm_err.Error())
		return err
	}

	// + elisam will call edborn to update status to step3

	err = Edborn.updateFileStatus(file, fileModel.READY_TO_STORE)

	if err != nil {
		//TODO: rollback
		fmt.Println("error on updating file status to step3 :", err.Error())
		return err
	}

	return nil
}
