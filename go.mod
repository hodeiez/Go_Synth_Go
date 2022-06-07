module hodei/gosynthgo

go 1.18

require (
	github.com/go-audio/audio v1.0.0
	github.com/go-audio/generator v0.0.0-20191129013639-fe5438877d8c
	github.com/gordonklaus/portaudio v0.0.0-20220320131553-cc649ad523c1
)

require github.com/AllenDang/giu v0.6.2

//replace github.com/go-audio/generator => github.com/hodeiez/generator v1.0.0 //I use my fork

replace github.com/go-audio/generator => C:\Users\hodei\Documents\GOLANG\goaudio-generator\generator //I use my fork github.com/hodeiez/generator

require (
	github.com/AllenDang/go-findfont v0.0.0-20200702051237-9f180485aeb8 // indirect
	github.com/AllenDang/imgui-go v1.12.1-0.20220322114136-499bbf6a42ad // indirect
	github.com/faiface/mainthread v0.0.0-20171120011319-8b78f0a41ae3 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220516021902-eb3e265c7661 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/sahilm/fuzzy v0.1.0 // indirect
	golang.org/x/image v0.0.0-20220601225756-64ec528b34cd // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	gopkg.in/eapache/queue.v1 v1.1.0 // indirect
)
