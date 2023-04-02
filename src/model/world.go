package model

type Area struct {
	Id               int      `yaml:"id"`
	Name             string   `yaml:"name"`
	Color            int      `yaml:"color"`
	Size             string   `yaml:"size"`
	ConstructionRate string   `yaml:"constructionRate"`
	Building         []string `yaml:"building"`
}
type Road struct {
	name  string
	color int
	*MetaData
}

type Construction struct {
	name  string
	color int
	*MetaData
	Grids []*Grid
}

type MetaData struct {
	objects []string
}

type IMetaGrid interface {
	getObjects() []string
}

type Grid struct {
	Area         *Area
	Construction *Construction
	Road         *Road
	X            int
	Y            int
}

func (g *Grid) SetConstruction(construction *Construction) {
	g.Construction = construction
	construction.Grids = append(construction.Grids, g)
}

func (g *Road) SetRoad(road *Road) {

}

func (m MetaData) getObjects() []string {
	return m.objects
}

var _ IMetaGrid = (*MetaData)(nil)

func NewGrid(area *Area, x int, y int) *Grid {
	grid := Grid{Area: area, X: x, Y: y}
	return &grid
}

func NewConstruction(area *Area) *Construction {
	return &Construction{
		name:     "Construction_" + area.Name,
		color:    0,
		MetaData: nil,
	}
}

func NewRoad(area *Area, area2 *Area) *Road {
	return &Road{
		name:     "Road_" + area.Name + "_" + area2.Name,
		color:    0,
		MetaData: nil,
	}
}
