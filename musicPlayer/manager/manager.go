package manager

import "github.com/pkg/errors"

// 定义音乐的结构体
type Music struct {
	Id string
	Name string
	Artist string
	Source string
	Type string

}

type MusicManager struct {
	musics []Music
}

func NewMusicManager() *MusicManager{
	return &MusicManager{make([]Music, 0)}
}

func (m *MusicManager) Len() int  {
	return len(m.musics)
}

func (m *MusicManager) Get (index int) (music *Music, err error){
	if index <0 || index >= len(m.musics){
		return nil, errors.New("Index out of range.")
	}

	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) (music *Music)  {
	if len(m.musics) == 0{
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Add (music *Music)  {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *Music  {
	if index<0 || index>=len(m.musics){
		return nil
	}

	removedMusic := &m.musics[index]

	if index < len(m.musics)-1 {
		// 删除中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		// 删除仅有的一个元素
		m.musics = make([]Music, 0)
	} else {
		// 删除最后的一个元素
		m.musics = m.musics[:index-1]
	}

	return removedMusic

}
