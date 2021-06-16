# yaml2json

Converts YAML input to JSON output.

e.g.:

    $ cat sample/foobar.yaml
    a: apple
    b: bat
    c:
    - foo
    - bar

    $ go run cmd/yaml2json.go < sample/foobar.yaml
    {
        "a": "apple",
        "b": "bat",
        "c": [
            "foo",
            "bar"
        ]
    }
