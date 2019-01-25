package advertisers

import "strconv"

var workingAdvertiser = advertiser{
	name: "working",
	bidConfs: []bidConf{
		{
			adURL:         "http://www.baltimoretherapyspot.com/wp-content/uploads/2013/07/Orton-Gollingham-to-be-emailed.jpg",
			shape:         "tall",
			adDescription: "Tall working ad for children",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "working"
					},
					impact: 5,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "working"
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
						return ageInt < 10
					},
					impact: 4,
				},
			},
		},
		{
			adURL:         "https://marketplace.canva.com/MACU_t9hYmk/2/0/thumbnail_large/canva-orange%2C-blue-and-cream-summer-tutor-flyer-MACU_t9hYmk.jpg",
			shape:         "tall",
			adDescription: "Tall working ad for college students",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "working"
					},
					impact: 3,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "working"
					},
					impact: -2,
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
					impact: 6,
				},
			},
		},
		{
			adURL:         "http://4.bp.blogspot.com/--B7p3I8Uqr4/Vkw9loHYECI/AAAAAAAAAFw/_hr5Iy3zsMw/s1600/male%2Bwork.jpg",
			shape:         "tall",
			adDescription: "Tall working ad for adults",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "working"
					},
					impact: 5,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "working"
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
						return ageInt > 25
					},
					impact: 2,
				},
			},
		},
		{
			adURL:         "https://www.jeffco.edu/sites/default/files/asc_banner.jpg",
			shape:         "wide",
			adDescription: "Wide working ad for college students",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "working"
					},
					impact: 3,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "working"
					},
					impact: -2,
				},
				{
					param: "a",
					condition: func(age string) bool {
						ageInt, err := strconv.Atoi(age)
						if err != nil {
							return false
						}
						return (ageInt > 17 && ageInt <= 25)
					},
					impact: 6,
				},
			},
		},
		{
			adURL:         "https://fightingeveryday.files.wordpress.com/2014/11/hacking-academy.jpg",
			shape:         "wide",
			adDescription: "Wide working ad for adults",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "working"
					},
					impact: 5,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "working"
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
						return ageInt > 25
					},
					impact: 2,
				},
			},
		},
	},
}
