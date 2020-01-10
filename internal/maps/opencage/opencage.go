/*
This package encapsulates the OpenCage Geocoder API.

Additional information can be found in our Developer Guide:

https://github.com/photoprism/photoprism/wiki
*/
package opencage

import (
	"github.com/photoprism/photoprism-places/internal/event"
)

const ProviderName = "opencage"

var ApiKey = ""
var log = event.Log

var opencageTypes = map[string]string{
	"aerodrome":           "airport",
	"apron":               "airport",
	"control_tower":       "airport",
	"control_center":      "airport",
	"gate":                "airport",
	"helipad":             "airport",
	"navigationaid":       "airport",
	"runway":              "airport",
	"taxilane":            "airport",
	"taxiway":             "airport",
	"terminal":            "airport",
	"bay":                 "bay",
	"peninsula":           "peninsula",
	"cape":                "cape",
	"wood":                "forest",
	"grassland":           "grassland",
	"beach":               "beach",
	"dune":                "dune",
	"water":               "water",
	"wetland":             "wetland",
	"glacier":             "glacier",
	"strait":              "seashore",
	"coastline":           "seashore",
	"reef":                "reef",
	"geyser":              "geyser",
	"peak":                "mountain",
	"hill":                "hill",
	"volcano":             "volcano",
	"valley":              "valley",
	"ridge":               "mountain",
	"cliff":               "cliff",
	"saddle":              "mountain",
	"isthmus":             "seashore",
	"sinkhole":            "sinkhole",
	"natural":             "nature",
	"sea":                 "ocean",
	"ocean":               "ocean",
	"gallery":             "gallery",
	"museum":              "museum",
	"alpine_hut":          "alpine hut",
	"aquarium":            "aquarium",
	"artwork":             "exhibition",
	"camp_pitch":          "camping",
	"camp_site":           "camping",
	"caravan_site":        "camping",
	"hotel":               "hotel",
	"hostel":              "hotel",
	"motel":               "hotel",
	"information":         "visitor center",
	"picnic_site":         "hiking",
	"theme_park":          "theme park",
	"viewpoint":           "viewpoint",
	"wilderness_hut":      "hiking",
	"zoo":                 "zoo",
	"shop":                "shop",
	"butcher":             "butcher",
	"department_store":    "department store",
	"supermarket":         "supermarket",
	"mall":                "mall",
	"boutique":            "boutique",
	"fashion":             "boutique",
	"fashion_accessories": "boutique",
	"clothes":             "boutique",
	"fabric":              "boutique",
	"leather":             "boutique",
	"baby_goods":          "boutique",
	"bag":                 "boutique",
	"books":               "bookstore",
	"fire_station":        "fire station",
	"bar":                 "bar",
	"biergarten":          "biergarten",
	"cafe":                "cafe",
	"internet_cafe":       "cafe",
	"ice_cream":           "cafe",
	"bistro":              "restaurant",
	"restaurant":          "restaurant",
	"fast_food":           "restaurant",
	"food_court":          "restaurant",
	"pub":                 "pub",
	"college":             "university",
	"university":          "university",
	"kindergarten":        "kindergarten",
	"language_school":     "school",
	"driving_school":      "school",
	"music_school":        "school",
	"school":              "school",
	"car_rental":          "car",
	"ferry_terminal":      "harbor",
	"parking":             "parking",
	"parking_entrance":    "parking",
	"parking_space":       "parking",
	"community_centre":    "community center",
	"bank":                "bank",
	"clinic":              "hospital",
	"hospital":            "hospital",
	"pharmacy":            "pharmacy",
	"arts_centre":         "exhibition",
	"casino":              "casino",
	"cinema":              "cinema",
	"gambling":            "casino",
	"planetarium":         "planetarium",
	"nightclub":           "nightclub",
	"theatre":             "theatre",
	"embassy":             "embassy",
	"grave_yard":          "cemetery",
	"cemetery":            "cemetery",
	"marketplace":         "marketplace",
	"monastery":           "monastery",
	"police":              "police",
	"prison":              "prison",
	"public_bath":         "swimming",
	"shelter":             "shelter",
	"aircraft":            "aircraft",
	"castle":              "castle",
	"castle_wall":         "castle",
	"church":              "church",
	"farm":                "farm",
	"memorial":            "memorial",
	"ship":                "ship",
	"tank":                "tank",
	"tower":               "tower",
	"wreck":               "ship",
	"houseboat":           "ship",
	"office":              "office",
	"warehouse":           "warehouse",
	"cathedral":           "cathedral",
	"chapel":              "chapel",
	"mosque":              "mosque",
	"shrine":              "shrine",
	"synagogue":           "synagogue",
	"temple":              "temple",
	"train_station":       "train station",
	"cowshed":             "farm",
	"greenhouse":          "greenhouse",
	"stable":              "farm",
	"farm_auxiliary":      "farm",
	"barn":                "farm",
	"sty":                 "farm",
	"stadium":             "stadium",
	"hangar":              "hangar",
	"water_tower":         "tower",
	"transformer_tower":   "tower",
	"bunker":              "bunker",
	"bridge":              "bridge",
	"garden":              "botanical garden",
	"adult_gaming_centre": "casino",
	"amusement_arcade":    "casino",
	"beach_resort":        "beach",
	"dog_park":            "dog park",
	"escape_game":         "escape game",
	"firepit":             "camping",
	"golf_course":         "golf",
	"miniature_golf":      "golf",
	"hackerspace":         "hackerspace",
	"marina":              "marina",
	"nature_reserve":      "nature reserve",
	"park":                "park",
	"picnic_table":        "outdoor",
	"pitch":               "sports",
	"sports_centre":       "sports",
	"swimming_area":       "swimming",
	"swimming_pool":       "swimming",
	"water_park":          "water park",
	"wildlife_hide":       "wildlife",
}
