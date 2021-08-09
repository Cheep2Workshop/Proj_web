if not exist gen-js mkdir gen-js
protoc *.proto --js_out=./gen-js
pause