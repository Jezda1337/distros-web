package distro

type Link struct {
	Label string `json:"label"`
	Icon  string `json:"icon"`
	URL   string `json:"url"`
}

type Information struct {
	InstallSize      string `json:"install_size"`
	DownloadSize     string `json:"download_size"`
	Architecture     string `json:"architecture"`
	TotalInstalls    int    `json:"total_installs"`
	LatestVersion    string `json:"latest_version"`
	ReleaseDate      string `json:"release_date"`
	IsRollingRelease bool   `json:"is_rolling_release"`
}

type Metadata struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Distro struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Logo        string      `json:"logo"`
	Images      []string    `json:"images"`
	Description string      `json:"description"`
	Information Information `json:"information"`
	Tags        []string    `json:"tags"`
	Categories  []string    `json:"categories"`
	Links       struct {
		Website       Link `json:"website"`
		IssueTracker  Link `json:"issue_tracker"`
		Forum         Link `json:"forum"`
		Documentation Link `json:"documentation"`
	} `json:"links"`
	Metadata Metadata `json:"metadata"`
}
