package models

// import "time"

type Music struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ArtistID    string `json:"artist_id" gorm:"references:ID"`
	Artist      Artist `json:"artist" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AlbumID     string `json:"album_id" gorm:"references:ID"`
	Album       Album  `json:"album" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TrackNumber int16  `json:"tracknumber"`
	Genre       string `json:"genre"`
	Year        string `json:"year"`
	UrlLink     string `json:"url"`
	Likes       int64  `json:"likes" gorm:"default:0"`
}

type ReceivedMusicData struct {
	Title       string `json:"title"`
	ArtistID    string `json:"artist_id"`
	AlbumID     string `json:"album_id"`
	TrackNumber int16  `json:"tracknumber"`
	Genre       string `json:"genre"`
	Year        string `json:"year"`
	UrlLink     string `json:"url"`
}

type Album struct {
	ID        string `json:"id" gorm:"primarykey"`
	Name      string `json:"name" gorm:"references:ID"`
	ArtistID  string `json:"artist_id"`
	Artist    Artist `json:"artist" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Thumbnail string `json:"thumbnail"`
	// ReleaseDate time.Time
}

type ReceivedAlbum struct {
	Name      string
	ArtistID  string
	Thumbnail string
}

type Artist struct {
	ID        string `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

type LikeData struct {
	ID    string `json:"id"`
	Likes int64  `json:"likes"`
}
