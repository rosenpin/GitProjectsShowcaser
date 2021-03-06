package sorters

import (
	"sort"

	"gitlab.com/rosenpin/git-project-showcaser/models"
)

type stars struct{}

func newStars() Sorter {
	return &stars{}
}

func (stars *stars) Sort(projects models.Projects) models.Projects {
	if projects == nil {
		panic("can't sort nil projects object")
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Stars > projects[j].Stars
	})

	return projects
}

func init() {
	sortersCreators[models.StarsSort] = newStars
}
