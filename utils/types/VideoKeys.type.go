package types

const VideoKeyLength = 32

type VideoKeys []FileKey

func (vk *VideoKeys) String() string {

	str := ""
	for i, v := range *vk {
		str += v.Id + "," + v.Value
		if len(*vk) != i+1 {
			str += ","
		}
	}
	return str
}
