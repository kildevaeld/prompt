package form

import (
	"github.com/kildevaeld/go-widgets"
	tm "github.com/kildevaeld/prompt/terminal"
)

type PaginatedList struct {
	Theme *tm.Theme
	Name  string
	widgets.PaginatedList
	Message  string
	Paginate func(page int) []string
}

func (c *PaginatedList) Run() {

	if c.Theme == nil {
		c.Theme = tm.DefaultTheme
	}

	c.Config = themeToWidgetConfig(c.Theme)

	label := c.Message
	if label == "" {
		label = c.Name
	}
	c.PaginatedList.Message = label
	c.PaginatedList.Paginate = c.Paginate
	c.PaginatedList.Run()
}

func (c *PaginatedList) GetValue() interface{} {
	return c.Value
}

func (c *PaginatedList) GetName() string {
	return c.Name
}

func (c *PaginatedList) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
