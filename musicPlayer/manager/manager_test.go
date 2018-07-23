package manager

import "testing"

func TestOpts(t *testing.T) {
	mm:= NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManger Failed , not empty")
	}

	m0 := &Music{
		"1",
		"MyHeart Will Go On",
		"Cellion Dion",
		"http://qbox.me/24501234",
		"MP3",
	}

	mm.Add(m0)

	if mm.Len() != 1{
		t.Error("NewMusicManger Failed , not empty")
	}

	m := mm.Find(m0.Name)

	if m!= nil{
		t.Error("MusicManger Find Failed ")
	}


	if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManger Find Failed 111")
	}

	m, err := mm.Get(0)
	if m!= nil{
		t.Error("MusicManger Get Failed ", err)
	}

	m = mm.Remove(0)


	if m == nil || mm.Len() != 0 {
		t.Error("MusicManger Remove Failed ")
	}


}
