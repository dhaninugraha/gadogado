package cherrypicker

type CherryPickerDetail struct {
	Tags        []string
	GetChildren bool
}

func Tags(tags ...string) func(*CherryPickerDetail) {
	return func(c *CherryPickerDetail) {
		if len(tags) > 0 {
			for _, tag := range tags {
				c.Tags = append(c.Tags, tag)
			}
		}
	}
}

func GetChildren() func(*CherryPickerDetail) {
	return func(c *CherryPickerDetail) {
		c.GetChildren = true
	}
}
