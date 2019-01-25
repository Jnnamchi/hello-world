package advertisers

import "strconv"

var cookingAdvertiser = advertiser{
	name: "cooking",
	bidConfs: []bidConf{
		{
			adURL:         "https://www.advertgallery.com/wp-content/uploads/2017/12/skywalk-fest-healthy-holiday-cooking-ad-the-hindu-chennai-19-12-2017.jpg",
			shape:         "tall",
			adDescription: "Tall cooking ad for individual under 13 years old.",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "cooking"
					},
					impact: 3,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "cooking"
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
					impact: 2,
				},
			},
		},
		{
			adURL:         "https://i.pinimg.com/originals/a1/04/0f/a1040fdef5374ad68f67e37236196f04.jpg",
			shape:         "tall",
			adDescription: "Tall cooking ad for middle aged adults.",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "cooking"
					},
					impact: 5,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "cooking"
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
			adURL:         "https://shopping.timminspress.com/imgs/media.images/300/300.jpg",
			shape:         "tall",
			adDescription: "Tall cooking ad for senior citizens.",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "cooking"
					},
					impact: 3,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "cooking"
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
						return ageInt > 65
					},
					impact: 4,
				},
			},
		},
		{
			adURL:         "https://thumbs.dreamstime.com/z/assorted-raw-organic-vegetables-ingredients-healthy-cooking-white-wooden-background-top-view-banner-vegetarian-diet-61687830.jpg",
			shape:         "wide",
			adDescription: "Wide cooking ad for adults and children.",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "cooking"
					},
					impact: 3,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "cooking"
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
						return ageInt < 65
					},
					impact: 4,
				},
			},
		},
		{
			adURL:         "https://www.penrithcity.nsw.gov.au/uploadedImages/Penrith_City_Council/Pages/Community_and_Library/Seniors/Reimagine_Ageing_Cooking_650x300.jpg",
			shape:         "wide",
			adDescription: "Wide cooking ad for senior citizens.",
			rules: []rule{
				{
					param: "l",
					condition: func(likes string) bool {
						return likes == "cooking"
					},
					impact: 3,
				},
				{
					param: "d",
					condition: func(dislikes string) bool {
						return dislikes == "cooking"
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
						return ageInt > 65
					},
					impact: 4,
				},
			},
		},
	},
}
