## go-toxcore
The golang bindings for libtoxcore 

### Installation

    go get github.com/kitech/go-toxcore


### Examples

    import "github.com/kitech/go-toxcore"

    // use custom options
    opt := tox.NewToxOptions()
    t := tox.NewTox(opt)
    av := tox.NewToxAv(t)
    
    // use default options
    t := tox.NewTox(nil)
    av := tox.NewToxAv(t)

