package advertisers

import "strconv"

var moviesAdvertiserName = "movies"

var moviesAdvertiser = advertiser{
	name: moviesAdvertiserName,
	bidConfs: []bidConf{
		{
			adURL:         "http://4.bp.blogspot.com/-VZPczy5xTjE/UUDNVhKBCGI/AAAAAAAAC78/8P05IBZbQdE/s1600/croods-poster.jpg",
			shape:         "tall",
			adDescription: "Tall movie ad for children and teenagers.",
			rules: []rule{
				{
					param: LikesParameter,
					condition: func(likes string) bool {
						return likes == moviesAdvertiserName
					},
					impact: 3,
				},
				{
					param: DislikesParameter,
					condition: func(dislikes string) bool {
						return dislikes == moviesAdvertiserName
					},
					impact: -4,
				},
				{
					param: AgeParameter,
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
					param: LikesParameter,
					condition: func(likes string) bool {
						return likes == moviesAdvertiserName
					},
					impact: 4,
				},
				{
					param: DislikesParameter,
					condition: func(dislikes string) bool {
						return dislikes == moviesAdvertiserName
					},
					impact: -5,
				},
				{
					param: AgeParameter,
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
					param: LikesParameter,
					condition: func(likes string) bool {
						return likes == moviesAdvertiserName
					},
					impact: 3,
				},
				{
					param: DislikesParameter,
					condition: func(dislikes string) bool {
						return dislikes == moviesAdvertiserName
					},
					impact: -4,
				},
				{
					param: AgeParameter,
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
					param: LikesParameter,
					condition: func(likes string) bool {
						return likes == moviesAdvertiserName
					},
					impact: 4,
				},
				{
					param: DislikesParameter,
					condition: func(dislikes string) bool {
						return dislikes == moviesAdvertiserName
					},
					impact: -5,
				},
				{
					param: AgeParameter,
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
