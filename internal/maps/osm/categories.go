package osm

var osmCategories = map[string]string{
	"aeroway=*":                "airport",
	"natural=bay":              "bay",
	"natural=peninsula":        "peninsula",
	"natural=cape":             "cape",
	"natural=wood":             "forest",
	"natural=grassland":        "grassland",
	"*=beach":                  "beach",
	"*=dune":                   "dune",
	"*=water":                  "water",
	"*=wetland":                "wetland",
	"*=glacier":                "glacier",
	"*=strait":                 "seashore",
	"*=coastline":              "seashore",
	"*=reef":                   "reef",
	"*=geyser":                 "geyser",
	"natural=peak":             "mountain",
	"natural=hill":             "hill",
	"natural=volcano":          "volcano",
	"natural=valley":           "valley",
	"natural=ridge":            "mountain",
	"natural=cliff":            "cliff",
	"natural=saddle":           "mountain",
	"natural=isthmus":          "seashore",
	"natural=sinkhole":         "sinkhole",
	"natural=*":                "nature",
	"place=sea":                "ocean",
	"*=ocean":                  "ocean",
	"*=gallery":                "gallery",
	"*=museum":                 "museum",
	"*=alpine_hut":             "alpine hut",
	"*=aquarium":               "aquarium",
	"*=artwork":                "exhibition",
	"*=camp_pitch":             "camping",
	"*=camp_site":              "camping",
	"*=caravan_site":           "camping",
	"*=hotel":                  "hotel",
	"*=hostel":                 "hotel",
	"*=motel":                  "hotel",
	"tourism=information":      "visitor center",
	"*=picnic_site":            "hiking",
	"*=theme_park":             "theme park",
	"*=viewpoint":              "viewpoint",
	"*=wilderness_hut":         "hiking",
	"*=zoo":                    "zoo",
	"shop=*":                   "shop",
	"shop=butcher":             "butcher",
	"shop=department_store":    "department store",
	"*=supermarket":            "supermarket",
	"*=mall":                   "mall",
	"*=boutique":               "boutique",
	"*=fashion":                "boutique",
	"*=fashion_accessories":    "boutique",
	"*=clothes":                "boutique",
	"*=fabric":                 "boutique",
	"shop=leather":             "boutique",
	"shop=baby_goods":          "boutique",
	"shop=bag":                 "boutique",
	"shop=books":               "bookstore",
	"*=fire_station":           "fire station",
	"amenity=bar":              "bar",
	"amenity=biergarten":       "biergarten",
	"amenity=cafe":             "cafe",
	"amenity=internet_cafe":    "cafe",
	"amenity=ice_cream":        "cafe",
	"amenity=bistro":           "restaurant",
	"amenity=restaurant":       "restaurant",
	"amenity=fast_food":        "restaurant",
	"amenity=food_court":       "restaurant",
	"amenity=pub":              "pub",
	"amenity=college":          "university",
	"amenity=university":       "university",
	"amenity=kindergarten":     "kindergarten",
	"amenity=language_school":  "school",
	"amenity=driving_school":   "school",
	"amenity=music_school":     "school",
	"amenity=school":           "school",
	"amenity=car_rental":       "car",
	"amenity=ferry_terminal":   "harbor",
	"amenity=parking":          "parking",
	"amenity=parking_entrance": "parking",
	"amenity=parking_space":    "parking",
	"*=bank":                   "bank",
	"*=clinic":                 "hospital",
	"*=hospital":               "hospital",
	"*=pharmacy":               "pharmacy",
	"*=arts_centre":            "exhibition",
	"*=casino":                 "casino",
	"*=cinema":                 "cinema",
	"*=gambling":               "casino",
	"*=planetarium":            "planetarium",
	"*=nightclub":              "nightclub",
	"*=theatre":                "theatre",
	"*=embassy":                "embassy",
	"*=grave_yard":             "cemetery",
	"*=cemetery":               "cemetery",
	"*=marketplace":            "marketplace",
	"*=monastery":              "monastery",
	"*=police":                 "police",
	"*=prison":                 "prison",
	"*=public_bath":            "swimming",
	"*=shelter":                "shelter",
	"*=aircraft":               "aircraft",
	"*=castle":                 "castle",
	"*=castle_wall":            "castle",
	"*=church":                 "church",
	"*=farm":                   "farm",
	"*=memorial":               "memorial",
	"*=ship":                   "ship",
	"*=tank":                   "tank",
	"*=tower":                  "tower",
	"*=wreck":                  "ship",
	"*=houseboat":              "ship",
	"*=office":                 "office",
	"*=warehouse":              "warehouse",
	"*=cathedral":              "cathedral",
	"*=chapel":                 "chapel",
	"*=mosque":                 "mosque",
	"*=shrine":                 "shrine",
	"*=synagogue":              "synagogue",
	"*=temple":                 "temple",
	"*=train_station":          "train station",
	"*=cowshed":                "farm",
	"*=greenhouse":             "greenhouse",
	"*=stable":                 "farm",
	"*=farm_auxiliary":         "farm",
	"*=barn":                   "farm",
	"*=sty":                    "farm",
	"*=stadium":                "stadium",
	"*=hangar":                 "hangar",
	"*=parking":                "parking",
	"*=water_tower":            "tower",
	"*=transformer_tower":      "tower",
	"*=bunker":                 "bunker",
	"*=bridge":                 "bridge",
	"*=garden":                 "botanical garden",
	"*=adult_gaming_centre":    "casino",
	"*=amusement_arcade":       "casino",
	"*=beach_resort":           "beach",
	"*=dog_park":               "dog park",
	"*=escape_game":            "escape game",
	"*=firepit":                "camping",
	"*=golf_course":            "golf",
	"*=miniature_golf":         "golf",
	"*=hackerspace":            "hackerspace",
	"*=marina":                 "marina",
	"*=nature_reserve":         "nature reserve",
	"*=park":                   "park",
	"*=picnic_table":           "outdoor",
	"*=pitch":                  "sports",
	"*=sports_centre":          "sports",
	"*=swimming_area":          "swimming",
	"*=swimming_pool":          "swimming",
	"*=water_park":             "water park",
	"*=wildlife_hide":          "wildlife",
}
