package advertisers

import "strconv"

var sportsAdvertiser = advertiser{
	name: "sports",
	bidConfs: []bidConf{
		{
			adURL:         "https://i.pinimg.com/originals/2b/1e/15/2b1e159f172936e3470fc41d09bf7969.jpg",
			shape:         "tall",
			adDescription: "Tall sports summer camp ad for children and teenagers",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "sports"
					},
					impact: 5,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "sports"
					},
					impact: -3,
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
					impact: 6,
				},
			},
		},
		{
			adURL:         "https://d3nuqriibqh3vw.cloudfront.net/lucozade-sport-rugby-ad-3_aotw.jpgg",
			shape:         "tall",
			adDescription: "Tall sports ad for adults",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "sports"
					},
					impact: 5,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "sports"
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
						return (ageInt > 17 && ageInt <= 65)
					},
					impact: 2,
				},
			},
		},
		{
			adURL:         "http://npfirstumc.org/wp-content/uploads/2012/02/upward-banner.jpg",
			shape:         "wide",
			adDescription: "Wide sports ad for children and teenagers",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "sports"
					},
					impact: 5,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "sports"
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
						return ageInt < 13
					},
					impact: 5,
				},
			},
		},
		{
			adURL:         "https://eurasiaprospective.files.wordpress.com/2017/10/1-jjfcyxskqc9x5cuqnfunka.jpg",
			shape:         "wide",
			adDescription: "Wide sports ad for adults",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "sports"
					},
					impact: 4,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "sports"
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
