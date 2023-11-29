package helper

import "os"

func CheckAndCreateAppDirs() error {

	var err error

	current_dir := GetCurrentDir()

	_, err = os.Stat(current_dir + "/ffmpeg")

	if err != nil {

		err = CreateDirectory(current_dir + "/ffmpeg")

		if err != nil {

			return err
		}
	}

	_, err = os.Stat(current_dir + "/packager")

	if err != nil {

		err = CreateDirectory(current_dir + "/packager")

		if err != nil {

			return err
		}
	}

	_, err = os.Stat(current_dir + "/storage")

	if err != nil {

		err = CreateDirectory(current_dir + "/storage")

		if err != nil {

			return err
		}
	}

	_, err = os.Stat(current_dir + "/storage/stage1")

	if err != nil {

		err = CreateDirectory(current_dir + "/storage/stage1")

		if err != nil {

			return err
		}
	}

	_, err = os.Stat(current_dir + "/storage/stage2")

	if err != nil {

		err = CreateDirectory(current_dir + "/storage/stage2")

		if err != nil {

			return err
		}
	}

	_, err = os.Stat(current_dir + "/storage/stage3")

	if err != nil {

		err = CreateDirectory(current_dir + "/storage/stage3")

		if err != nil {

			return err
		}
	}

	_, err = os.Stat(current_dir + "/storage/stage4")

	if err != nil {

		err = CreateDirectory(current_dir + "/storage/stage4")

		if err != nil {

			return err
		}
	}

	return nil
}
