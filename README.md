# gonc
Golang Network CLI Client (Trainig project)


gonc http

    gonc http -verb head "https://example.com/"

    gonc http -verb get "https://example.com/"

    gonc http post "https://example.com/" \
        -type 'application/json' \
        -body '{"auth":"secret"}'

    gonc http post "https://example.com/" \
        -type 'application/json' \
        -body-file './some-file.json'
    