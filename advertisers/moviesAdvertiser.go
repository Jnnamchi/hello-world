package advertisers

import "strconv"

var moviesAdvertiser = advertiser{
	name: "movies",
	bidConfs: []bidConf{
		{
			adURL:         "http://4.bp.blogspot.com/-VZPczy5xTjE/UUDNVhKBCGI/AAAAAAAAC78/8P05IBZbQdE/s1600/croods-poster.jpg",
			shape:         "tall",
			adDescription: "Tall movie ad for children and teenagers.",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "movies"
					},
					impact: 3,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "movies"
					},
					impact: -4,
				},
				{
					param: "a",
					condition: func(age string) bool {
						ageInt, err := strconv.Atoi(age)
						if err != nil {
							return false
						}
						return ageInt < 13
					},
					impact: 3,
				},
			},
		},
		{
			adURL:         "https://media-assets-05.thedrum.com/cache/images/thedrum-prod/s3-black-panther-imax-poster-1080815--default--826.jpeg",
			shape:         "tall",
			adDescription: "Tall movie ad for adults and seniors",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "movies"
					},
					impact: 4,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "movies"
					},
					impact: -5,
				},
				{
					param: "a",
					condition: func(age string) bool {
						ageInt, err := strconv.Atoi(age)
						if err != nil {
							return false
						}
						return ageInt > 17
					},
					impact: 3,
				},
			},
		},
		{
			adURL:         "https://www.joblo.com/assets/images/oldsite/posters/images/full/2009-monsters_vs_aliens-7.jpg",
			shape:         "wide",
			adDescription: "Wide movie ad for children and teenagers.",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "movies"
					},
					impact: 3,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "movies"
					},
					impact: -4,
				},
				{
					param: "a",
					condition: func(age string) bool {
						ageInt, err := strconv.Atoi(age)
						if err != nil {
							return false
						}
						return ageInt < 13
					},
					impact: 3,
				},
			},
		},
		{
			adURL:         "http://cdn.collider.com/wp-content/uploads/dark-knight-rises-movie-poster-banner-catwoman.jpg",
			shape:         "wide",
			adDescription: "Wide movie ad for adults and seniors",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "movies"
					},
					impact: 4,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "movies"
					},
					impact: -5,
				},
				{
					param: "a",
					condition: func(age string) bool {
						ageInt, err := strconv.Atoi(age)
						if err != nil {
							return false
						}
						return ageInt > 17
					},
					impact: 3,
				},
			},
		},
	},
}
