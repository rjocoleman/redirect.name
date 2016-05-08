# [redirect.name](http://redirect.name/) [![Build status][ci-image]][ci-url]

Please refer to [redirect.name](http://redirect.name/) for documentation.

[ci-image]: https://img.shields.io/circleci/project/holic/redirect.name/master.svg?style=flat
[ci-url]: https://circleci.com/gh/holic/redirect.name


**this is a fork**

See [redirect.name](http://redirect.name) or [http://github.com/holic/redirect.name](http://github.com/holic/redirect.name) for the canonical information.

### Differences from upstream

* `DEFAULT_RECORD` environment variable enables a default fallback mode. e.g. if `DEFAULT_RECORD=default` then you look up `foo.example.com` and an explicit `_redirect.foo` doesn't exist then `_redirect.default` will be used instead.

* Makefile for easier manual deployment. 
