[![Build Status](https://travis-ci.org/kitech/go-toxcore.svg?branch=master)](https://travis-ci.org/kitech/go-toxcore)

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


Contributing
------------
1. Fork it
2. Create your feature branch (``git checkout -b my-new-feature``)
3. Commit your changes (``git commit -am 'Add some feature'``)
4. Push to the branch (``git push origin my-new-feature``)
5. Create new Pull Request

