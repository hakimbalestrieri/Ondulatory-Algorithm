package Config

var ServerPorts = map[int]int{0: 3100, 1: 3101, 2: 3102, 3: 3103, 4: 3104, 5: 3105, 6: 3106, 7: 3107, 8: 3108}
var ServerNumber = len(ServerPorts)

type Node struct {
	Server int
	Edges  []*Edge
}
type Edge struct {
	//Pas besoin de stocker de "From, car le From est toujours le numéro du Node courant (Server int)
	To int
}

type Graph struct {
	nodes []*Node
	edges []*Edge
}

var GraphSchema = &Graph{
	nodes: []*Node{
		{
			Server: 0,
			Edges: []*Edge{
				{
					To: 1,
				},
				{
					To: 3,
				},
				{
					To: 4,
				},
			},
		},
		{
			Server: 1,
			Edges: []*Edge{
				{
					To: 0,
				},
				{
					To: 3,
				},
				{
					To: 2,
				},
			},
		},
		{
			Server: 2,
			Edges: []*Edge{
				{
					To: 1,
				},
				{
					To: 3,
				},
				{
					To: 6,
				},
				{
					To: 7,
				},
			},
		},
		{
			Server: 3,
			Edges: []*Edge{
				{
					To: 0,
				},
				{
					To: 1,
				},
				{
					To: 2,
				},
				{
					To: 4,
				},
				{
					To: 6,
				},
			},
		},
		{
			Server: 4,
			Edges: []*Edge{
				{
					To: 0,
				},
				{
					To: 3,
				},
				{
					To: 5,
				},
				{
					To: 6,
				},
			},
		},
		{
			Server: 5,
			Edges: []*Edge{
				{
					To: 4,
				},
				{
					To: 6,
				},
			},
		},
		{
			Server: 6,
			Edges: []*Edge{
				{
					To: 2,
				},
				{
					To: 3,
				},
				{
					To: 4,
				},
				{
					To: 5,
				},
			},
		},
		{
			Server: 7,
			Edges: []*Edge{
				{
					To: 2,
				},
				{
					To: 8,
				},
			},
		},
		{
			Server: 8,
			Edges: []*Edge{
				{
					To: 7,
				},
			},
		},
	},
}
