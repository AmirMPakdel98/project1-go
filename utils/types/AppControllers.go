package types

type CreateFilesReq struct {
	File_size uint32 `json:"file_size"`
	File_type int    `json:"file_type"`
	File_ext  string `json:"file_ext"`
}

type CreateFilesResData struct {
	Object_id  uint32 `json:"object_id"`
	Upload_url string `json:"upload_url"`
}

type GetFilesResData struct {
	File_status int       `json:"file_status"`
	File_url    string    `json:"file_url"`
	File_keys   []FileKey `json:"file_keys"`
}
