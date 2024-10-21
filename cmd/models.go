package cmd

type PyPIResponse struct {
	Info     PackageInfo              `json:"info"`
	Releases map[string][]ReleaseInfo `json:"releases"`
}

type PackageInfo struct {
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Summary      string   `json:"summary"`
	Description  string   `json:"description"`
	Author       string   `json:"author"`
	AuthorEmail  string   `json:"author_email"`
	License      string   `json:"license"`
	HomePage     string   `json:"home_page"`
	ProjectURL   string   `json:"project_url"`
	RequiresDist []string `json:"requires_dist"`
}

type ReleaseInfo struct {
	Filename   string `json:"filename"`
	URL        string `json:"url"`
	Size       int64  `json:"size"`
	UploadTime string `json:"upload_time"`
}
