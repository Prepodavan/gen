module github.com/clipperhouse/gen

go 1.13

replace (
	github.com/clipperhouse/linkedlist => ./submodules/linkedlist
	github.com/clipperhouse/ring => ./submodules/ring
	github.com/clipperhouse/setwriter => ./submodules/set
	github.com/clipperhouse/slice => ./submodules/slice
	github.com/clipperhouse/stringer => ./submodules/stringer
	github.com/clipperhouse/typewriter => ./submodules/typewriter
)

require (
	github.com/clipperhouse/linkedlist v0.0.0-00010101000000-000000000000
	github.com/clipperhouse/ring v0.0.0-00010101000000-000000000000
	github.com/clipperhouse/setwriter v0.0.0-20200110222355-ab7a980082f0
	github.com/clipperhouse/slice v0.0.0-20200107170738-a74fc3888fd9
	github.com/clipperhouse/stringer v0.0.0-20200107165315-e8ef8175ba3b
	github.com/clipperhouse/typewriter v0.0.0-20200107164453-d21420026310
	github.com/fsnotify/fsnotify v1.4.7
)
