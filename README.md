## How to build and run this thing:

###.Directories structure
You should have a $GOPATH env set. Then in your $GOPATH/ directory create following structure:

```
$GOPATH/
    /src
        /non-trivial-go-backend
            ...
            git repo content here
            ...
```

###.Running backend

```
go build
```
and then 
```
./name-of-the-package
```
That's it! :) Now you can go to your browser and type:
```
localhost:8000/test
```


###.Database configuration

Run a mongo server on localhost
```
mongo
```

Connect to database
```
mongod
```

Create collection **opinions** in **lbpl** database
```
use lbpl
db.createCollection("opinions")
```

Insert example data
```
db.opinions.insertMany( [{author: "Andriej", text: "Wujowe, nie polecam", rating: -5}, {author: "Wokulski", text: "Kupiłem Izabeli, teraz w końcu na pewno zamoczę!", rating: 7.5}, {author: "Hektor", text: "Genialne! Dałem Helenie i bez gadania spierdoliła ze mną do Troi!", rating: 10}])
```

##.Manual tests of REST endpoints using curl
###GET
```curl -XGET http://localhost:8000/endpoint?p=v```

###POST
```curl -XPOST -H 'Content-Type: application/json' -d '{"atr": "value"}' http://localhost:8000/endpoint```

###DELETE
```curl -XDELETE http://localhost:8000/endpoint```
