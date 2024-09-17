# json-clean

A stupid simple binary that removes characters from JSON that are formatted to be logged within a JSON log. 
I have set Spring Boot to log in JSON format, so the output of a logged JSON payload looks something like:
```
{\n  \"field1\" : \"value1\",\n  \"field2\" : 123.5\n}
```
If we pipe this to `json-clean` it will remove all backslash (`\`) characters, whitespace (` `), along with any characters that do
not belong outside of double quotes (in this case the letter `n`).
```sh
$ echo '{\n  \"field1\" : \"value1\",\n  \"field2\" : 123.5\n}' | jc
{"field1":"value1","field2":123.5}
```
We can then pipe this to [jq](https://jqlang.github.io/jq/) for a nice output.
```json5
$ echo '{\n  \"field1\" : \"value1\",\n  \"field2\" : 123.5\n}' | jc | jq
{
  "field1": "value1",
  "field2": 123.5
}
```
