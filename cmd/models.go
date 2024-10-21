package cmd

type PyPIResponse struct {
	Info     PackageInfo              `json:"info"`
	Releases map[string][]ReleaseInfo `json:"releases"`
}

type PackageInfo struct {
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Summary      string   `json:"summary"`
	HomePage     string   `json:"home_page"`
	License      string   `json:"license"`
	RequiresDist []string `json:"requires_dist"`
}

type ReleaseInfo struct {
	Filename   string `json:"filename"`
	URL        string `json:"url"`
	Size       int64  `json:"size"`
	UploadTime string `json:"upload_time"`
}
