package common

func (app AppData) MergeStatus(status, resolution, statusp, resolutionp string) (string,string) {
	var statusOrder = []string{
		"Closed",
		"In Progress",
		"Verified",
		"New",
	}
	// TODO: Compute this with a script from the config.
	
	for _, test := range statusOrder {
		if status == test {
			return status, resolution
		}
		if statusp == test {
			return statusp, resolutionp
		}
	}
	// This *should* be unreachable, but things change
	return statusp, resolutionp
}
