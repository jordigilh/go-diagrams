package linux

import "github.com/jordigilh/go-diagrams/diagram"

type unixNode struct {
	path string
	opts []diagram.NodeOption
}

var Unix = &unixNode{
	opts: diagram.OptionSet{diagram.Provider("linux"), diagram.NodeShape("none")},
	path: "assets/linux",
}

func (u *unixNode) Systemd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("../assets/linux/systemd.png")}, u.opts, opts)
	return diagram.NewNode(nopts...)
}

func (u *unixNode) Podman(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("../assets/linux/podman.png")}, u.opts, opts)
	return diagram.NewNode(nopts...)
}

func (u *unixNode) Buffer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("../assets/linux/buffer.png")}, u.opts, opts)
	return diagram.NewNode(nopts...)
}
