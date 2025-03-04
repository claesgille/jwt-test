package model

type Price struct {
	Id          string
	Name        string
	Description string
	MainImage   string
}

type Draw struct {
	Id     string
	Status string
	Name   string
	Prices []Price
}

var Draws = []Draw{
	{Id: "1", Status: "Closed", Name: "Dragning 2022",
		Prices: []Price{
			{Id: "1", Name: "Röd Vas", Description: "En Tjusig röd vas i äkta hammarplast.", MainImage: "/assets/vase.jpg"},
			{Id: "2", Name: "En räv på isen", Description: "Klassisk naturstudie.", MainImage: "/assets/fox.jpg"},
		},
	},
	{Id: "2", Status: "Closed", Name: "Dragning 2023",
		Prices: []Price{
			{Id: "3", Name: "Klassisk målning", Description: "Skulle kunna vara målad av en av de gamla mästarna.", MainImage: "/assets/painting1.jpg"},
			{Id: "4", Name: "Skallepär", Description: "En bissart dekorerad skalle i sydamerikansk stil.", MainImage: "/assets/skull.jpg"},
		},
	},
	{Id: "3", Status: "Open", Name: "Dragning 2024",
		Prices: []Price{
			{Id: "5", Name: "Persiphone", Description: "En klassisk skulptur. För tankarna till de gamla grekerna.", MainImage: "/assets/sculpture1.jpg"},
			{Id: "6", Name: "Kvinna i färg", Description: "Färgsprakane fotokonst som skulle pryda vilken vägg som helst.", MainImage: "/assets/girl.jpg"},
		},
	},
}
