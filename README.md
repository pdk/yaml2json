# yaml2json

Converts YAML input to JSON output.

Install:

    go get github.com/pdk/yaml2json@latest

usage:

    $ cat sample/foobar.yaml
    a: apple
    b: bat
    c:
    - foo
    - bar

    $ yaml2json < sample/foobar.yaml
    {
        "a": "apple",
        "b": "bat",
        "c": [
            "foo",
            "bar"
        ]
    }
