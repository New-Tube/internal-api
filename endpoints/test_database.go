package endpoints

import (
	"internal-api/db"
	db_models "internal-api/db/models"
)

func resetDB() error {
	conn, err := db.GetDBConnection()

	if err != nil {
		return err
	}

	result := conn.Exec("DROP TABLE comments, media_sources, reactions, users, videos")

	return result.Error
}

func fillUsers() error {
	conn, err := db.GetDBConnection()

	if err != nil {
		return err
	}

	users := []*db_models.User{
		{
			Name:         "Tanek",
			Surname:      "Richards",
			Nickname:     "nunc",
			PasswordHash: 3984,
		},
		{
			Name:         "Nayda",
			Surname:      "Talley",
			Nickname:     "Mauris",
			PasswordHash: 7778,
		},
		{
			Name:         "Hamish",
			Surname:      "Conner",
			Nickname:     "luctus",
			PasswordHash: 3836,
		},
		{
			Name:         "Dylan",
			Surname:      "Daniel",
			Nickname:     "Quisque",
			PasswordHash: 1954,
		},
		{
			Name:         "Alyssa",
			Surname:      "Walsh",
			Nickname:     "augue",
			PasswordHash: 2311,
		},
	}

	return conn.Create(users).Error
}

func fillVideos() error {
	conn, err := db.GetDBConnection()

	if err != nil {
		return err
	}

	videos := []*db_models.Video{

		{
			UserID:      4,
			Title:       "diam.",
			Description: "Aenean sed pede nec ante blandit viverra. Donec tempus, lorem fringilla ornare placerat, orci lacus vestibulum lorem, sit amet ultricies sem magna nec quam.",
			Privacy:     2,
			Link:        "https://guardian.co.uk/sub",
		},
		{
			UserID:      3,
			Title:       "natoque penatibus et magnis dis parturient montes, nascetur",
			Description: "a purus. Duis",
			Privacy:     1,
			Link:        "https://ebay.com/fr",
		},
		{
			UserID:      3,
			Title:       "dignissim. Maecenas ornare egestas ligula.",
			Description: "dis parturient montes, nascetur ridiculus mus. Aenean eget magna. Suspendisse tristique neque venenatis lacus. Etiam bibendum fermentum metus. Aenean sed pede nec ante blandit viverra. Donec tempus, lorem fringilla ornare placerat, orci lacus vestibulum lorem, sit amet ultricies sem magna nec quam. Curabitur vel lectus. Cum",
			Privacy:     2,
			Link:        "https://bbc.co.uk/en-us",
		},
		{
			UserID:      2,
			Title:       "lobortis",
			Description: "non arcu. Vivamus sit amet risus. Donec egestas. Aliquam nec enim. Nunc ut erat. Sed nunc est, mollis non, cursus non, egestas a, dui. Cras pellentesque. Sed dictum. Proin eget",
			Privacy:     1,
			Link:        "http://whatsapp.com/site",
		},
		{
			UserID:      1,
			Title:       "et, lacinia vitae,",
			Description: "Praesent eu dui. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aenean eget magna. Suspendisse tristique neque venenatis lacus. Etiam bibendum fermentum metus. Aenean sed pede nec ante blandit viverra. Donec tempus, lorem fringilla ornare placerat, orci lacus",
			Privacy:     3,
			Link:        "https://reddit.com/en-ca",
		},
		{
			UserID:      4,
			Title:       "Fusce aliquam,",
			Description: "vitae purus",
			Privacy:     2,
			Link:        "https://youtube.com/user/110",
		},
		{
			UserID:      5,
			Title:       "Donec egestas. Duis ac arcu. Nunc mauris. Morbi non",
			Description: "et netus et malesuada fames ac turpis egestas. Fusce",
			Privacy:     1,
			Link:        "https://twitter.com/sub",
		},
		{
			UserID:      4,
			Title:       "augue. Sed molestie. Sed id risus quis diam luctus",
			Description: "eget, dictum placerat, augue. Sed molestie. Sed id risus quis diam luctus lobortis. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos. Mauris ut quam vel sapien imperdiet ornare. In faucibus. Morbi vehicula. Pellentesque tincidunt tempus risus. Donec egestas. Duis ac arcu. Nunc mauris. Morbi non",
			Privacy:     2,
			Link:        "https://whatsapp.com/fr",
		},
		{
			UserID:      1,
			Title:       "mi. Aliquam gravida mauris ut mi.",
			Description: "augue scelerisque mollis. Phasellus libero mauris, aliquam eu, accumsan sed, facilisis vitae, orci. Phasellus dapibus quam quis diam. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Fusce aliquet magna a neque. Nullam ut nisi a odio semper cursus. Integer mollis. Integer tincidunt aliquam arcu. Aliquam",
			Privacy:     2,
			Link:        "https://wikipedia.org/sub",
		},
		{
			UserID:      2,
			Title:       "tempus risus.",
			Description: "magna a tortor. Nunc commodo auctor velit. Aliquam nisl. Nulla eu neque pellentesque massa lobortis ultrices. Vivamus rhoncus. Donec est. Nunc ullamcorper, velit in aliquet",
			Privacy:     1,
			Link:        "https://netflix.com/en-ca",
		},
	}

	return conn.Create(videos).Error
}

func fillComments() error {
	conn, err := db.GetDBConnection()

	if err != nil {
		return err
	}

	comments := []*db_models.Comment{
		{
			UserID:  4,
			VideoID: 5,
			Text:    "vitae purus",
		},
		{
			UserID:  5,
			VideoID: 0,
			Text:    "et netus et malesuada fames ac turpis egestas. Fusce",
		},
		{
			UserID:  4,
			VideoID: 3,
			Text:    "eget, dictum placerat, augue. Sed molestie. Sed id risus quis diam luctus lobortis. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos. Mauris ut quam vel sapien imperdiet ornare. In faucibus. Morbi vehicula. Pellentesque tincidunt tempus risus. Donec egestas. Duis ac arcu. Nunc mauris. Morbi non",
		},
		{
			UserID:  1,
			VideoID: 5,
			Text:    "augue scelerisque mollis. Phasellus libero mauris, aliquam eu, accumsan sed, facilisis vitae, orci. Phasellus dapibus quam quis diam. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Fusce aliquet magna a neque. Nullam ut nisi a odio semper cursus. Integer mollis. Integer tincidunt aliquam arcu. Aliquam",
		},
		{
			UserID:  2,
			VideoID: 0,
			Text:    "magna a tortor. Nunc commodo auctor velit. Aliquam nisl. Nulla eu neque pellentesque massa lobortis ultrices. Vivamus rhoncus. Donec est. Nunc ullamcorper, velit in aliquet",
		},
		{
			UserID:  3,
			VideoID: 1,
			Text:    "dignissim. Maecenas ornare egestas ligula. Nullam feugiat placerat velit. Quisque varius. Nam porttitor scelerisque neque. Nullam nisl. Maecenas malesuada fringilla est.",
		},
		{
			UserID:    3,
			CommentID: 7,
			Text:      "Nunc mauris sapien, cursus in, hendrerit consectetuer, cursus et, magna. Praesent interdum ligula eu enim. Etiam imperdiet dictum magna. Ut tincidunt orci quis lectus. Nullam suscipit, est ac facilisis facilisis, magna tellus faucibus leo, in lobortis tellus",
		},
		{
			UserID:    4,
			CommentID: 3,
			Text:      "vehicula. Pellentesque tincidunt tempus risus. Donec egestas. Duis ac arcu. Nunc mauris. Morbi non sapien molestie orci tincidunt adipiscing.",
		},
		{
			UserID:    1,
			CommentID: 12,
			Text:      "Duis at lacus. Quisque purus sapien, gravida non, sollicitudin a, malesuada id, erat. Etiam vestibulum massa rutrum magna. Cras convallis convallis dolor. Quisque tincidunt pede ac urna. Ut tincidunt vehicula risus. Nulla eget metus eu erat semper rutrum. Fusce dolor",
		},
		{
			UserID:    5,
			CommentID: 12,
			Text:      "fringilla euismod enim. Etiam gravida molestie arcu. Sed eu nibh vulputate mauris",
		},
		{
			UserID:    3,
			CommentID: 3,
			Text:      "libero mauris, aliquam eu, accumsan sed, facilisis vitae, orci. Phasellus dapibus quam quis diam. Pellentesque habitant morbi tristique senectus et netus et",
		},
		{
			UserID:    3,
			CommentID: 13,
			Text:      "Cras sed leo. Cras vehicula aliquet libero. Integer in magna. Phasellus dolor elit, pellentesque a, facilisis non, bibendum sed, est. Nunc laoreet lectus quis massa. Mauris vestibulum, neque sed dictum eleifend, nunc risus varius",
		},
		{
			UserID:    3,
			CommentID: 12,
			Text:      "ipsum porta elit, a feugiat tellus lorem eu metus. In lorem. Donec elementum, lorem ut aliquam iaculis, lacus pede sagittis augue, eu tempor erat neque",
		},
		{
			UserID:    1,
			CommentID: 13,
			Text:      "cursus et, magna. Praesent interdum ligula eu enim. Etiam imperdiet dictum magna. Ut tincidunt orci quis lectus. Nullam suscipit, est ac facilisis facilisis, magna tellus faucibus leo, in lobortis tellus justo sit amet nulla. Donec non justo. Proin non massa non ante bibendum ullamcorper. Duis cursus, diam at pretium",
		},
		{
			UserID:    3,
			CommentID: 13,
			Text:      "feugiat non, lobortis quis, pede. Suspendisse dui. Fusce diam nunc, ullamcorper eu, euismod ac, fermentum vel, mauris. Integer sem elit, pharetra ut, pharetra sed, hendrerit a, arcu. Sed et libero. Proin mi. Aliquam gravida mauris ut mi. Duis risus odio, auctor vitae, aliquet nec, imperdiet nec, leo. Morbi",
		},
		{
			UserID:    2,
			CommentID: 15,
			Text:      "leo. Cras vehicula aliquet libero.",
		},
		{
			UserID:  2,
			VideoID: 4,
			Text:    "nec, imperdiet nec, leo. Morbi neque tellus, imperdiet non, vestibulum nec, euismod in, dolor. Fusce feugiat. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aliquam auctor, velit eget laoreet posuere, enim nisl elementum purus,",
		},
		{
			UserID:  3,
			VideoID: 3,
			Text:    "ipsum leo elementum sem, vitae aliquam eros turpis non enim. Mauris quis turpis vitae purus gravida sagittis. Duis gravida. Praesent eu nulla at sem molestie sodales. Mauris blandit enim consequat purus. Maecenas libero est, congue a, aliquet vel, vulputate eu,",
		},
		{
			UserID:  3,
			VideoID: 8,
			Text:    "eu, placerat eget, venenatis a, magna. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Etiam laoreet, libero et tristique pellentesque, tellus sem mollis dui, in sodales elit erat vitae risus. Duis a mi fringilla mi lacinia mattis. Integer eu lacus. Quisque imperdiet, erat nonummy",
		},
		{
			UserID:  4,
			VideoID: 9,
			Text:    "Donec tempor, est ac",
		},
	}

	return conn.Create(comments).Error
}
