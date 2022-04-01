package systems

import "GoRpg/src/interfaces"

var SCENEMANAGER *SceneManager

type SceneManager struct {
	allScenese   []interfaces.IScene
	CurrentScene interfaces.IScene
}

func InitSceneManager() {
	if SCENEMANAGER == nil {
		g := &SceneManager{}
		g.allScenese = make([]interfaces.IScene, 0)
		SCENEMANAGER = g
	}
}

func (s *SceneManager) setScene() {
	l := len(s.allScenese)
	s.CurrentScene = s.allScenese[l-1]

	s.CurrentScene.Init()
}

func (s *SceneManager) Push(v interfaces.IScene) {
	s.allScenese = append(s.allScenese, v)
	s.setScene()
}

func (s *SceneManager) Pop() {
	if len(s.allScenese) <= 0 {
		s.allScenese = make([]interfaces.IScene, 0)
		s.CurrentScene = nil
		return
	}
	l := len(s.allScenese)
	s.allScenese = s.allScenese[:l-1]
	s.CurrentScene = nil
	s.setScene()
}
