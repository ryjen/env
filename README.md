[![GoDoc](https://godoc.org/github.com/ryjen/env?status.svg)](https://godoc.org/github.com/ryjen/env)
[![Build Status](https://travis-ci.org/ryjen/env.svg?branch=master)](https://travis-ci.org/ryjen/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryjen/env)](https://goreportcard.com/report/github.com/ryjen/env)
[![Coverage Status](https://coveralls.io/repos/github/ryjen/env/badge.png?branch=master)](https://coveralls.io/github/ryjen/env?branch=master)

# env-context

a context aware environment library. 

## env

the environment is the variables defined by the OS.  Such as in a file or in a shell.

## context

the context is where environment variables are defined.  For example, *development* or *production*.

a single environment variable can determine the context.

By default, `ENV_CONTEXT` is used, but can be specified.

## files

the default environment files are (in order of overriding):

1. .env
2. .env.local 
3. .env.context
4. .env.context.local

## command

The utility command is similar to `envsubst` or `env`
```
Usage of env-context:
  -context string
    	specify the context as a string or name of a environment variable
  -dir string
    	the directory to load environment files from
  -env-file string
    	specify environment file name
  -local-ext string
    	specify local extension
  -overload
    	overload existing OS environment variables
```

# credits and inspiration

* [react-scripts](https://create-react-app.dev/docs/adding-custom-environment-variables/#adding-development-environment-variables-in-env)
* [env for go](https://github.com/golobby/env)


## Contributors

* [@ryjen](https://github.com/ryjen)

## License

Released under the [MIT License](http://opensource.org/licenses/mit-license.php).