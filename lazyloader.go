package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LazyLoader struct {
	textures map[string]rl.Texture2D
}

func NewLazyLoader() *LazyLoader {
	self := &LazyLoader{}

	self.textures = make(map[string]rl.Texture2D)

	return self
}

func (self *LazyLoader) Load(name string) rl.Texture2D {
	_, ok := self.textures[name]
	if !ok {
		self.textures[name] = rl.LoadTexture(name)
	}

	return self.textures[name]
}

func (self *LazyLoader) Unload(name string) {
	texture, ok := self.textures[name]
	if ok {
		rl.UnloadTexture(texture)
		delete(self.textures, name)
	}
}

func (self *LazyLoader) UnloadAll() {
	for name := range self.textures {
		rl.UnloadTexture(self.textures[name])
		delete(self.textures, name)
	}
}
